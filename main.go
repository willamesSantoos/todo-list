package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/willamesSantoos/api-todo-list/internal/configs"
	"github.com/willamesSantoos/api-todo-list/internal/routes"
)

func main() {
	err := configs.Load()

	if err!= nil {
        panic(err)
    }

	router := mux.NewRouter()

	routes.Initializer(router)

	log.Println(fmt.Sprintf("Server running on port %s", configs.GetServerPort()))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), router)) 
}
