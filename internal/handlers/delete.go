package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/willamesSantoos/api-todo-list/internal/repository"
)

type reponseDelete struct {
	Message string `json:"message"`
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Println("Erro ao fazer parser do Id: v%", err)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := repository.Delete(int64(id))

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Println("Error: foram atualizados %d registros", rows)
	}

	resp := reponseDelete{
		Message: "Registro removido com sucesso",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)

	json.NewEncoder(w).Encode(resp)
}
