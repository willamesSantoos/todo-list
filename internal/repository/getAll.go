package repository

import (
	"log"

	"github.com/willamesSantoos/api-todo-list/internal/db"
	"github.com/willamesSantoos/api-todo-list/internal/models"
)

func GetAll() (todos []models.Todo, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM todos`)

	if err != nil {
		return
	}

	for rows.Next() {
		var todo models.Todo

		err = rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Done)

		if err != nil {
			log.Println(err)

			continue
		}

		todos = append(todos, todo)
	}

	return
}
