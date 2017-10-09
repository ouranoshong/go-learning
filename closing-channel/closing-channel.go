package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool, 1)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all job")
				done <- true
				return
			}
		}
	}()

	for i := 0; i <= 3; i++ {
		jobs <- i
		fmt.Println("sent job", i)
	}
	fmt.Println("sent all jobs")
	close(jobs)
	<-done
}
