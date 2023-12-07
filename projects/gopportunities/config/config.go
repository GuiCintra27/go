package config

import (
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
	logger *Logger
)

func Init() error {
	return nil
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}