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

	img := client.Container().From("alpine:3")

	c1 := img.WithExec([]string{"touch", "file1.txt"}).File("file1.txt")
	c2 := img.WithExec([]string{"touch", "file2.txt"}).File("file2.txt")
	c3 := img.WithExec([]string{"touch", "file3.txt"}).File("file3.txt")

	ok, err := img.
		WithFile("./file1.txt", c1).
		WithFile("./file2.txt", c2).
		WithFile("./file3.txt", c3).
		Export(ctx, "./sample.tar.gz")
	if err != nil {
		panic(err)
	}

	if !ok {
		panic("failed to export image")
	}
}
