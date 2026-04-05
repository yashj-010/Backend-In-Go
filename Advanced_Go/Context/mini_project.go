package main

import (
	"context"
	"fmt"
	"time"
)

func task(ctx context.Context) {
	for i := 1; ; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("Task stopped:", ctx.Err())
			return
		default:
			fmt.Println("Processing step", i)
			time.Sleep(1 * time.Second)
		}
	}
}

func mini_project() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go task(ctx)

	time.Sleep(4 * time.Second)
}