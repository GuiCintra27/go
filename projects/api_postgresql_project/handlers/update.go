package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/GuiCintra27/go/projects/api_postgresql_project/models"
	"github.com/go-chi/chi/v5"
)

func Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		log.Printf("Error parsing todo ID: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var todo models.Todo

	err = json.NewDecoder(r.Body).Decode(&todo)

	if err != nil {
		log.Printf("Error decoding todo: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	rows, err := models.Update(int64(id), todo)

	if err != nil {
		log.Printf("Error updating todo: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Printf("Error updating todo: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}