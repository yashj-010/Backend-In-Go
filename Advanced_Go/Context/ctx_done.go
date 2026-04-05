package main

import (
	"context"
	"fmt"
	"time"
)

func process(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Stopped:", ctx.Err())
			return
		default:
			fmt.Println("Working...")
			time.Sleep(1 * time.Second)
		}
	}
}

func ctx_done() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	process(ctx)
}