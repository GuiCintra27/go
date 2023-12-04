package main

import (
	"context"
	"log"

	"github.com/GuiCintra27/go/user_grpc_platform_project/framework/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main()  {
	connection, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

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
		Email: "jdoe@me.com",
		Password: "123456",
	}

	response, err := client.CreateUser(context.Background(), request)

	if err != nil {
		log.Fatalf("Could not create user: %v", err)
	}

	log.Println(response)
}