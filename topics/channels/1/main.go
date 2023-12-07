package main

import (
	"fmt"
	"time"
)
func main() {
	hello := make(chan string)
	go func ()  {
		time.Sleep(time.Second*2)
		hello <- "hello"
	}()

	select{
		case result := <- hello:
			fmt.Println(result)
		default:
			fmt.Println("waiting")
	}

	time.Sleep(time.Second*5)
}