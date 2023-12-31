package servers

import (
	"context"
	"log"

	"github.com/GuiCintra27/go/projects/user_grpc_platform_project/application/usecases"
	"github.com/GuiCintra27/go/projects/user_grpc_platform_project/domain"
	"github.com/GuiCintra27/go/projects/user_grpc_platform_project/framework/pb"
)

type UserServer struct {
	User domain.User
	UserUseCase usecases.UserUsecase
}

func NewUserServer() *UserServer {
	return &UserServer{}
}

func (UserServer *UserServer) CreateUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	UserServer.User.Name = req.GetName()
	UserServer.User.Email = req.GetEmail()
	UserServer.User.Password = req.GetPassword()

	user, err := UserServer.UserUseCase.CreateUser(&UserServer.User)

	if err != nil {
		log.Fatalf("Error during RPC Create User: %v", err)
		return nil, err
	}

	return &pb.UserResponse{
		Token: user.Token,
	}, nil 
}