package domain

import (
	"log"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Base
	Name string `json:"name" gorm:"type:varchar(255);not null"`
	Email string `json:"email" gorm:"type:varchar(255);not null;unique_index"`
	Password string `json:"-" gorm:"type:varchar(255);not null"`
	Token string `json:"token" gorm:"type:varchar(255);not null;unique_index"`
}

func NewUser() *User {
	return &User{}
}

func (u *User) Prepare() error {
	password, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Fatalf("Error hashing password: %v", err)
		return err
	}

	u.Password = string(password)
	u.Token = uuid.New().String()

	err = u.Validate()

	if err != nil {
		log.Fatalf("Error validating user: %v", err)
		return err
	}

	return nil
}

func (u *User) Validate() error {
	return nil
}