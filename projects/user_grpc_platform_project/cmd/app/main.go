package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/GuiCintra27/go/projects/user_grpc_platform_project/application/repositories"
	"github.com/GuiCintra27/go/projects/user_grpc_platform_project/application/usecases"
	"github.com/GuiCintra27/go/projects/user_grpc_platform_project/framework/pb"
	"github.com/GuiCintra27/go/projects/user_grpc_platform_project/framework/servers"
	"github.com/GuiCintra27/go/projects/user_grpc_platform_project/framework/utils"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var db *gorm.DB

func main()  {
	db = utils.ConnectDB()

	port := flag.Int("port", 0, "Choose the server port")

	flag.Parse()
	log.Printf("Starting server on port %d\n", *port)	

	userServer := setUpUserServer()

	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, userServer)
	reflection.Register(grpcServer)

	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalf("Error creating listener: %v", err)
	}

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func setUpUserServer() *servers.UserServer {
	UserServer := servers.NewUserServer()
	userRepository := repositories.UserRepositoryDb{Db: db}
	UserServer.UserUseCase = usecases.UserUsecase{UserRepository:userRepository}

	return UserServer
}