package main

import (
	"fmt"
	"time"
)

func worker (id int, jobs <-chan int) {

	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
	}
}

func main() {
	jobs := make(chan int)

	go worker(1, jobs)
	go worker(2, jobs)

	for w := 0; w < 10; w++ {
		jobs <- w
	}
}