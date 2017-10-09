package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, result chan<- int) {
	for job := range jobs {
		fmt.Println("worker", id, "started job", job)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", job)
		result <- job * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	result := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, result)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}

	close(jobs)

	for a := 1; a <= 5; a++ {
		<-result
	}
}
