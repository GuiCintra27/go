package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/GuiCintra27/go/user_grpc_platform_project/framework/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main()  {
	port := flag.Int("port", 0, "Insert the server port")

	flag.Parse()
	log.Printf("Making request on port %d\n", *port)	

	address := fmt.Sprintf("localhost:%d", *port)

	connection, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer connection.Close()

	client := pb.NewUserServiceClient(connection)

	CreateUser(client)
}

func CreateUser(client pb.UserServiceClient) {
	request := &pb.UserRequest{
		Name: "John",
		Email: "jdoec@me.com",
		Password: "123456",
	}

	response, err := client.CreateUser(context.Background(), request)

	if err != nil {
		log.Fatalf("Could not create user: %v", err)
	}

	log.Println(response)
}