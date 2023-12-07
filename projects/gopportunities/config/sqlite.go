package config

import (
	"os"

	"github.com/GuiCintra27/go/projects/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("Database not found\n")
		logger.Info("Creating sqlite database...")

		err := os.Mkdir("./db", os.ModePerm)
		if err != nil {
			logger.Errorf("Error creating db folder: %v", err)
			return nil, err
		}

		file, err := os.Create(dbPath)
		if err != nil {
			logger.Errorf("Error creating db file: %v", err)
			return nil, err
		}

		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		logger.Errorf("Error initializing sqlite: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})

	if err != nil {
		logger.Errorf("Error migrating sqlite: %v", err)
		return nil, err
	}

	return db, nil
}