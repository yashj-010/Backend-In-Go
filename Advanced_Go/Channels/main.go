package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

type Truck struct{
	id string
	capacity int
}

func processTruck(ctx context.Context, t Truck) error {
	time.Sleep(2*time.Second)
	select{
	case <-ctx.Done():
		return ctx.Err()
	default:
		fmt.Println("Processing truck", t.id)
		return nil
	}
}

func process(ctx context.Context, truck []Truck) error {
	var wg sync.WaitGroup
	errorsChan := make(chan error,len(truck))

	for _,t := range truck{
		wg.Add(1)
		go func(t Truck){
			if err := processTruck(ctx,t); err != nil {
				log.Println(err)
				errorsChan <- err
			}
			wg.Done()
		}(t)
	}
	
	wg.Wait()
	close(errorsChan)

	var errs []error
	for err := range errorsChan{
		log.Println("Error:",err)
		errs = append(errs, err)
	}

	if len(errs) > 0{
		return fmt.Errorf("found %d errors",len(errs))
	}

	return nil
}

func main(){
	trucks := []Truck{
		{id: "T1", capacity: 10},
		{id: "T2", capacity: 20},
		{id: "T3", capacity: 30},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	process(ctx, trucks)
}