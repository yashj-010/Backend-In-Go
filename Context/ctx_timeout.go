package main

import (
	"context"
	"fmt"
	"time"
)

func work(ctx context.Context) {
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Work completed")
	case <-ctx.Done():
		fmt.Println("Timeout occurred:", ctx.Err())
	}
}

func ctx_timeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	work(ctx)
}