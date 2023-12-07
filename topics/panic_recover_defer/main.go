package main

import "fmt"

func divide(x, y int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered")
			fmt.Printf("Error: %v\n", r)
		}
	}()

	return x / y
}

func main()  {
	fmt.Printf("\n%d\n", divide(10, 0))
}