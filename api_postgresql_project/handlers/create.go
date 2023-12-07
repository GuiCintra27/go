package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/GuiCintra27/go/api_postgresql_project/models"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)

	if err != nil {
		log.Printf("Error creating todo: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = models.Insert(todo)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}