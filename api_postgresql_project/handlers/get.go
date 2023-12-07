package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/GuiCintra27/go/api_postgresql_project/models"
	"github.com/go-chi/chi/v5"
)

func Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		log.Printf("Error parsing todo ID: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	todo, err := models.Get(int64(id))

	if err != nil {
		log.Printf("Error updating todo: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}