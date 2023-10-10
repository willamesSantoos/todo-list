package routes

import (
	"github.com/gorilla/mux"
	"github.com/willamesSantoos/api-todo-list/internal/handlers"
)

func Initializer(router *mux.Router) {
	urlBase := "/api/v1/todos"

	router.HandleFunc(urlBase, handlers.GetAll).Methods("GET")
	router.HandleFunc(urlBase+"/:id", handlers.Get).Methods("GET")
	router.HandleFunc(urlBase, handlers.Create).Methods("POST")
	router.HandleFunc(urlBase, handlers.Update).Methods("PUT")
	router.HandleFunc(urlBase, handlers.Delete).Methods("DELETE")
}
