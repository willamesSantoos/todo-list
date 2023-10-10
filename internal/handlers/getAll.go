package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/willamesSantoos/api-todo-list/internal/repository"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	todos, err := repository.GetAll()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(todos)
}
