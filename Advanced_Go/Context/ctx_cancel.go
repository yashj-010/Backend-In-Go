package main

import (
	"context"
	"fmt"
	"time"
)

func run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Stopped:", ctx.Err())
			return
		default:
			fmt.Println("Running...")
			time.Sleep(1 * time.Second)
		}
	}
}

func ctx_cancel() {
	ctx, cancel := context.WithCancel(context.Background())

	go run(ctx)

	time.Sleep(3 * time.Second)
	cancel()

	time.Sleep(1 * time.Second)
}