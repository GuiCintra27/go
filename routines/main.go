package main

import (
	"fmt"
	"time"
)

func counter(typ string) {

	for i := 0; i < 10; i++ {
		fmt.Println(typ, i)
		time.Sleep(time.Second)
	}
}

func main() {

	go counter("without go routines:")
	go counter("goroutine:")

	time.Sleep(time.Second * 10)
}