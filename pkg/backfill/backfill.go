package backfill

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/cased/consumers/pkg/auditevents"
	"github.com/cheggaaa/pb/v3"
	"go.uber.org/zap"
)

const StateFile = ".ddbackfill"

type Backfill struct {
	// Path is the location which contains audit events.
	Path string

	// LastFile is the last file that was processed.
	LastFile string

	// WorkerCount ...
	WorkerCount int

	// AuditEventPaths contains list of ordered audit events
	AuditEventPaths []string

	ddSource  string
	ddClient  *datadog.APIClient
	ddTags    string
	errorChan chan error
	pb        *pb.ProgressBar
}

func New(path string) *Backfill {
	b := &Backfill{
		Path:            path,
		AuditEventPaths: []string{},
		WorkerCount:     runtime.NumCPU(),
		ddSource:        "auditevents",
		errorChan:       make(chan error),
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
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	b.pb = pb.StartNew(len(b.AuditEventPaths))
	pathsChan := make(chan string)

	for i := 0; i < b.WorkerCount; i++ {
		wg.Add(1)

		go b.worker(ctx, pathsChan, i, &wg)
	}

	zap.L().Debug("started workers", zap.Int("count", b.WorkerCount))

	for _, path := range b.AuditEventPaths {
		pathsChan <- path
	}
	cancel()

	wg.Wait()

	return nil
}

func (b *Backfill) worker(ctx context.Context, pathsChan <-chan string, worker int, wg *sync.WaitGroup) {
	defer wg.Done()

	batch := NewBatch(func(body []datadog.HTTPLogItem) {
		ctx := datadog.NewDefaultContext(context.Background())
		_, resp, err := b.ddClient.LogsApi.SubmitLog(ctx, body)
		b.pb.Add(len(body))
		if err != nil {
			zap.L().Error("could not submit logs", zap.Error(err))
		}

		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			zap.L().Debug("successfully submitted logs", zap.Int("count", len(body)))
		} else {
			zap.L().Error("could not submit logs", zap.Int("count", len(body)))
		}
	})
	defer batch.Submit()

	for {
		select {
		case <-ctx.Done():
			zap.L().Debug("Shutting down worker", zap.Int("worker", worker))
			return
		case path, ok := <-pathsChan:
			// Channel has been closed.
			if !ok {
				zap.L().Debug("Shutting down worker", zap.Int("worker", worker))
				return
			}

			if b.LastFile != "" && path <= b.LastFile {
				// We hit the last file uploaded, we can continue
				b.pb.Increment()
				continue
			}

			data, err := os.ReadFile(path)
			if err != nil {
				zap.L().Warn("could not read file", zap.Error(err))
				continue
			}

			// Prepare log item
			li := datadog.NewHTTPLogItem()
			li.SetDdsource(b.ddSource)
			li.SetDdtags(b.ddTags)

			aep, err := auditevents.NewAuditEventPayload(data)
			if err != nil {
				panic(err)
			}

			// Log events can be submitted up to 18h in the past and 2h in the future.
			diff := 18 * time.Hour
			then := time.Now().Add(-diff)
			if aep.DotCased.PublishedAt.Before(then) {
				aep.DotCased.PublishedAt = time.Now()
			}

			d, err := json.Marshal(aep)
			if err != nil {
				panic(err)
			}
			li.SetMessage(string(d))

			// Add to batch
			batch.Append(li)

			b.LastFile = path
		}
	}
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

	zap.L().Debug("initialized backfill", zap.String("tags", b.ddTags), zap.String("last_file", b.LastFile))

	b.loadAuditEventPaths()
}

func (b *Backfill) loadAuditEventPaths() {
	walkStart := time.Now()
	filepath.Walk(b.Path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}

		// Only want to publish files
		if info.IsDir() {
			return nil
		}

		b.AuditEventPaths = append(b.AuditEventPaths, path)

		return nil
	})
	zap.L().Debug("loaded audit events", zap.Duration("took", time.Since(walkStart)))

	sortStart := time.Now()
	sort.Strings(b.AuditEventPaths)
	zap.L().Debug("sorted audit events", zap.Duration("took", time.Since(sortStart)))
}
