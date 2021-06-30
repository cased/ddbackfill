package main

import (
	"fmt"
	"os"

	"github.com/cased/ddbackfill/pkg/backfill"
	"go.uber.org/zap"
)

func main() {
	// Where the audit events live
	if len(os.Args) == 1 {
		fmt.Println("Must provide a directory for audit events")
		os.Exit(1)
	}

	debug := os.Getenv("DEBUG")
	if debug != "" {
		logger, err := zap.NewDevelopment()
		if err != nil {
			panic(err)
		}
		zap.ReplaceGlobals(logger)
	} else {
		logger, err := zap.NewProduction()
		if err != nil {
			panic(err)
		}
		zap.ReplaceGlobals(logger)
	}

	b := backfill.New(os.Args[1])
	defer b.Close()

	if err := b.Backfill(); err != nil {
		panic(err)
	}
}
