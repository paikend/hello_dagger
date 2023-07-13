package main

import (
	"context"
	"dagger.io/dagger"
	"os"
)

func main() {
	ctx := context.Background()
	// create a Dagger client
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		panic(err)
	}
	defer client.Close()
}
