package main

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)

	// Where the audit events live
	root := os.Args[1]
	if root == "" {
		panic("Must provide a directory for audit events")
	}

	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Only want to publish files
		if info.IsDir() {
			return nil
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		li := datadog.NewHTTPLogItem()
		li.SetMessage(string(data))
		li.SetDdsource("auditevents")
		// TODO: Set this based on test/live in bucket name
		li.SetDdtags("env:prod")

		// TODO: Queue up as many as possible
		body := []datadog.HTTPLogItem{*li}

		ctx := datadog.NewDefaultContext(context.Background())
		_, _, err = apiClient.LogsApi.SubmitLog(ctx, body)
		if err != nil {
			return err
		}

		// fmt.Println(string(data))
		fmt.Print(".")

		return nil
	})

	fmt.Println("")

	if err != nil {
		panic(err)
	}
}
