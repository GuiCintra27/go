package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/GuiCintra27/go/projects/api_postgresql_project/models"
	"github.com/go-chi/chi/v5"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		log.Printf("Error parsing todo ID: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	rows, err := models.Delete(int64(id))

	if err != nil {
		log.Printf("Error deleting todo: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Printf("Error updating todo: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}