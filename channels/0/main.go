package main

import "fmt"


func main() {

	forever := make(chan string)

	go func() {
		x := true
		for {
			if x {
				continue
			}
		}
	}()

	fmt.Println("aguardando para sempre")
	<- forever
}