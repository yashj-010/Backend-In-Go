package main

import (
	"fmt"
	"time"
)

func assignjob(id int, jobs <-chan int, results chan<- int){
	for j := range jobs{
		fmt.Println("Worker",id,"processing job",j)
		time.Sleep(time.Second)
		fmt.Println("Worker",id,"finished job",j)
		results <- j*2
	}
}

func worker(){
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	
	for w:=1;w<=3;w++{
		go assignjob(w,jobs,results)
	}

	time.Sleep(3*time.Second)

	for j:=1;j<=numJobs;j++{
		jobs <- j
	}
	close(jobs)

	for a:=1;a<=numJobs;a++{
		fmt.Println("Result:",<-results)
	}
}