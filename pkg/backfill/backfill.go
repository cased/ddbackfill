package backfill

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"sync"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/cheggaaa/pb/v3"
)

const StateFile = ".ddbackfill"

type Backfill struct {
	// Path is the location which contains audit events.
	Path string

	// LastFile is the last file that was processed.
	LastFile string

	// Uploading ...
	Uploading bool

	// WorkerCount ...
	WorkerCount int

	// AuditEventPaths contains list of ordered audit events
	AuditEventPaths []string

	ddSource string
	ddClient *datadog.APIClient
	ddTags   string

	pb *pb.ProgressBar
}

func New(path string) *Backfill {
	b := &Backfill{
		Path:            path,
		Uploading:       true,
		AuditEventPaths: []string{},
		WorkerCount:     runtime.NumCPU(),
		ddSource:        "auditevents",
	}
	b.init()

	return b
}

func (b *Backfill) Close() error {
	if b.pb != nil {
		b.pb.Finish()
	}

	return os.WriteFile(StateFile, []byte(b.LastFile), 0644)
}

func (b *Backfill) Backfill() error {
	done := make(chan struct{})
	defer close(done)

	paths, errc := b.loadAuditEventPaths(done)

	fmt.Println("outside of b.loadAuditEventPaths")
	c := make(chan bool)

	b.pb = pb.StartNew(len(b.AuditEventPaths))

	var wg sync.WaitGroup

	for i := 0; i < b.WorkerCount; i++ {
		wg.Add(1)

		go func() {
			b.worker(done, paths, c)
			fmt.Println("done with worker")
			wg.Done()
		}()
	}
	go func() {
		// Wait for all workers to consume channels
		wg.Wait()
		close(c)
	}()

	fmt.Printf("Started %d workers\n", b.WorkerCount)

	return <-errc
}

func (b *Backfill) worker(done chan struct{}, paths <-chan string, c chan<- bool) {
	batch := NewBatch(func(body []datadog.HTTPLogItem) {
		fmt.Printf("Batch prepared (%d)\n", len(body))

		ctx := datadog.NewDefaultContext(context.Background())
		b.ddClient.LogsApi.SubmitLog(ctx, body)

		b.pb.Add(len(body))
	})

	for path := range paths {
		if !b.Uploading {
			if b.LastFile != "" && path >= b.LastFile {
				// We hit the last file uploaded, we can continue
				b.Uploading = true
			}

			b.pb.Increment()

			fmt.Println("b.Uploading")
			continue
		}

		data, err := os.ReadFile(path)
		if err != nil {
			continue
		}

		li := datadog.NewHTTPLogItem()
		li.SetMessage(string(data))
		li.SetDdsource(b.ddSource)
		// TODO: Set this based on test/live in bucket name
		li.SetDdtags(b.ddTags)

		batch.Append(li)

		b.LastFile = path

		select {
		case c <- true:
		case <-done:
			batch.Submit()
			return
		}
	}

	batch.Submit()
}

func (b *Backfill) init() {
	_, err := os.Stat(StateFile)
	if err != nil {
		if !errors.Is(err, fs.ErrNotExist) {
			panic(err)
		}
	} else {
		data, err := os.ReadFile(StateFile)
		if err != nil {
			panic(err)
		}
		b.LastFile = string(data)

		// We have a state file and file to start after, don't upload right away
		b.Uploading = false
	}

	// Prepare Datadog tags to be included with each audit event
	if strings.Contains(b.Path, "cased-test") {
		b.ddTags = "env:test"
	} else {
		b.ddTags = "env:prod"
	}

	// Setup the Datadog API Client
	configuration := datadog.NewConfiguration()
	b.ddClient = datadog.NewAPIClient(configuration)

}

func (b *Backfill) loadAuditEventPaths(done <-chan struct{}) (<-chan string, <-chan error) {
	fmt.Println("Loading audit events…")
	paths := make(chan string)
	errc := make(chan error, 1)

	go func() {
		defer close(paths)

		errc <- filepath.Walk(b.Path, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				fmt.Println(err)
				return err
			}

			// Only want to publish files
			if info.IsDir() {
				return nil
			}

			b.AuditEventPaths = append(b.AuditEventPaths, path)

			select {
			case paths <- path:
				fmt.Println(path)
				fmt.Println("194 ended")
			case <-done:
				fmt.Println("197 ended")
				return errors.New("walk canceled")
			}
			return nil
		})

		fmt.Println("ended")
	}()

	fmt.Println("Sorting audit events…")
	fmt.Println(len(b.AuditEventPaths))
	fmt.Println(len(paths))
	sort.Strings(b.AuditEventPaths)

	return paths, errc
}
