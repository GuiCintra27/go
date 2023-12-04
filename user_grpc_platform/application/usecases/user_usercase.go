package usecases

import (
	"log"

	"github.com/GuiCintra27/go/user_grpc_platform/application/repositories"
	"github.com/GuiCintra27/go/user_grpc_platform/domain"
)

type UserUsecase struct {
	UserRepository repositories.UserRepository
}

func (u UserUsecase) CreateUser(user *domain.User) (*domain.User, error) {
		user, err := u.UserRepository.Insert(user)

		if err != nil {
			log.Fatalf("Error creating user: %v", err)
			return user, err
		}

		return user, nil
}