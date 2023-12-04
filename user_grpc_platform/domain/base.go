package domain

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Base struct {
	ID string `json:"id" gorm:"type:uuid;primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" sql:"index"`
}

func (u *Base) BeforeCreate(scope *gorm.Scope) error {
	err := scope.SetColumn("CreatedAt", time.Now())

	if err != nil {
		log.Fatalf("Error setting CreatedAt: %v", err)
		return err
	}

	err = scope.SetColumn("ID", uuid.New().String())

	if err != nil {
		log.Fatalf("Error setting ID: %v", err)
		return err
	}
	
	return nil
}