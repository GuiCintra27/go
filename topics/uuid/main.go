package main

import (
	"fmt"

	"github.com/google/uuid"
)

//studying how to working with go.mod :)

func main()  {
	uuid := uuid.New()
	fmt.Println(uuid)
}