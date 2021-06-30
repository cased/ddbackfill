package backfill

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

const (
	// Maximum array size if sending multiple logs in an array: 1000 entries
	BatchItemLimit = 1000

	// Maximum size for a single log: 1MB
	BatchMaxItemLength = 1000000

	// Maximum content size per payload (uncompressed): 5MB
	BatchMaxLength = 5000000
)

type Batch struct {
	Items []datadog.HTTPLogItem

	// Maximum content size per payload (uncompressed): 5MB
	BodySize int

	callback func([]datadog.HTTPLogItem)
}

func NewBatch(callback func([]datadog.HTTPLogItem)) *Batch {
	return &Batch{
		Items:    []datadog.HTTPLogItem{},
		BodySize: 0,
		callback: callback,
	}
}

func (b *Batch) Submit() {
	if len(b.Items) > 0 {
		b.callback(b.Items)
	}

	b.Clear()
}

func (b *Batch) Clear() {
	b.Items = []datadog.HTTPLogItem{}
	b.BodySize = 0
}

func (b *Batch) Append(item *datadog.HTTPLogItem) {
	if b.BodySize >= BatchMaxLength {
		fmt.Printf("Max body size encounted %d of %d limit\n", b.BodySize, BatchMaxItemLength)
		b.Submit()
	}

	c := len(*item.Message)

	// This entry will exceed max item length, we need to submit what we have then
	// submit the one item on its own
	if c > BatchMaxItemLength {
		fmt.Printf("Hit BatchMaxItemLength with %d\n", c)
		b.Submit()
		b.Items = append(b.Items, *item)
		b.BodySize += c
		b.Submit()
		return
	} else if b.BodySize+c > BatchMaxLength {
		fmt.Printf("Hit BatchMaxLength with %d\n", b.BodySize+c)
		// By appending we'd exceed the batch limit, need to submit first
		b.Submit()
	}

	b.Items = append(b.Items, *item)
	b.BodySize += c

	if len(b.Items) == BatchItemLimit {
		fmt.Println("Hit BatchItemLimit")
		b.Submit()
	}
}
