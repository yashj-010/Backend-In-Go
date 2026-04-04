package main

import (
	"context"
	"fmt"
	"time"
)

func fetchdata(ctx context.Context){
	select{
	case <-time.After(3*time.Second):
		fmt.Println("Data3 fetched Successfully")
	case <-ctx.Done():
		fmt.Println("Timeout occurred:", ctx.Err())
	}
}

func ctx_api(){
	ctx, cancel := context.WithTimeout(context.Background(),2*time.Second)
	defer cancel()
	fetchdata(ctx)
}