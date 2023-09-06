package main

import (
	"context"
	"os"

	"dagger.io/dagger"
)

func main() {
	ctx := context.Background()
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		panic(err)
	}

	// creating an empty container, using alpine:3 as base and export
	// the image to current directory as a gzip file
	client.Container().From("alpine:3").Export(ctx, "./sample.tar.gz")
}
