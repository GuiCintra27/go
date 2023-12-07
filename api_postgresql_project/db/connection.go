package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/GuiCintra27/go/api_postgresql_project/configs"
	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	conf := configs.GetDB()

	sc := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", conf.User, conf.Password, conf.Host, conf.Port, conf.Database)

	conn, err := sql.Open("postgres", sc)

	if err != nil {
		log.Fatalf("Error opening connection: %v", err)
	}

	err = conn.Ping()
	
	if err != nil {
		log.Fatalf("Error pinging connection: %v", err)
	}
	
	return conn, err
}