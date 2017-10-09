package main

import "fmt"
import "time"

func work(done chan bool) {

	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {
	done := make(chan bool, 1)

	go work(done)

	<-done
}
