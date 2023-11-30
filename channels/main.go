package main

import "fmt"

func main() {
	hello := make(chan string)

	go func() {
		hello <- "hello"
	}()

	result := <-hello
	fmt.Println(result)
}