package repositories

import (
	"log"

	"github.com/GuiCintra27/go/user_grpc_platform_project/domain"
	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	Insert(user *domain.User) (*domain.User, error)
}

type UserRepositoryDb struct {
	Db *gorm.DB
}

func (repo UserRepositoryDb) Insert(user *domain.User) (*domain.User, error) {

	err := user.Prepare()

	if err != nil {
		log.Fatalf("Error preparing user: %v", err)
		return nil, err
	}

	err = repo.Db.Create(user).Error

	if err != nil {
		log.Fatalf("Error inserting user: %v", err)
		return nil, err
	}

	return user, nil
}