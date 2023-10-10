package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/willamesSantoos/api-todo-list/internal/models"
	"github.com/willamesSantoos/api-todo-list/internal/repository"
)

type reponseUpdate struct {
    Message string `json:"message"`
}

func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Println("Erro ao fazer parser do Id: v%", err)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var todo models.Todo

	err = json.NewDecoder(r.Body).Decode(&todo)

	if err != nil {
		log.Println("Erro ao fazer decode do json: v%", err)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := repository.Update(int64(id), todo)	

	if err != nil {
		log.Println("Erro ao fazer decode do json: v%", err)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Println("Error: foram atualizados %d registros", rows)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	resp := reponseUpdate {
		Message: "Registro atualizado com sucesso",
	}

	json.NewEncoder(w).Encode(resp)
}
