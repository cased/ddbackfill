package main

import (
	"fmt"
	"os"

	"github.com/cased/ddbackfill/pkg/backfill"
)

func main() {
	// Where the audit events live
	if len(os.Args) == 1 {
		fmt.Println("Must provide a directory for audit events")
		os.Exit(1)
	}

	b := backfill.New(os.Args[1])
	defer b.Close()

	if err := b.Backfill(); err != nil {
		panic(err)
	}
}
