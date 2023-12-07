package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/GuiCintra27/go/api_postgresql_project/models"
)

func List(w http.ResponseWriter, r *http.Request) {
	todos, err := models.GetAll()
	if err != nil {
		log.Printf("Error getting todos: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}