package main

import (
	"context"
	"dagger.io/dagger"
	"fmt"
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

	src := client.Host().Directory("./src")

	// run test code
	test := client.Container().From("golang:1.20-alpine"). // Build an gradle image with gradle and bash installed
		WithDirectory("/src", src).WithWorkdir("/src"). // mount source directory to /src
		WithExec([]string{"go", "test"})
	out, err := test.Stdout(ctx)
	if err != nil {
		fmt.Printf("Error test: %s", err)
		os.Exit(1)
	}
	fmt.Println(out)

	// execute code
	execute := client.Container().From("golang:1.20-alpine").
		WithDirectory("/src", src).WithWorkdir("/src"). // mount source directory to /src
		WithExec([]string{"go", "run", "main.go"})
	out, err = execute.Stdout(ctx)
	if err != nil {
		fmt.Printf("Error running code: %s", err)
		os.Exit(1)
	}
	fmt.Println(out)
}
