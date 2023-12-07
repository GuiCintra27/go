package main

import (
	"context"
	"fmt"
	"time"
)

func main()  {
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func ()  {
		time.Sleep(time.Second*3)
		cancel()
	}()

	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	select {
		case <- ctx.Done():
			fmt.Println("Canceled")
		case <- time.After(time.Second * 4):
			fmt.Println("Booked")
	}
}