package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/willamesSantoos/api-todo-list/internal/models"
	"github.com/willamesSantoos/api-todo-list/internal/repository"
)

type reponseCreate struct {
    Message string `json:"message"`
}

func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Contente-Type", "application/json")

	var todo models.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)

	if err != nil {
		log.Println("Erro ao fazer decode do json: v%", err)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := repository.Insert(todo)

	var resp reponseCreate

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		resp = reponseCreate {
			Message: fmt.Sprintf("Oops! Something went wrong, Err: %v", err),
		}
	} else {
		w.WriteHeader(http.StatusCreated)

		resp =  reponseCreate {
			Message: fmt.Sprintf("Registration created successfully! Id: %d", id),
		}
	}	
	
	json.NewEncoder(w).Encode(resp)
}
