package repository

import (
	"github.com/willamesSantoos/api-todo-list/internal/db"
	"github.com/willamesSantoos/api-todo-list/internal/models"
)

func Get(id int64) (todo models.Todo, err error) {
	conn, err := db.OpenConnection()
	if err!= nil {
        return 
    }
	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM todos WHERE id = $1`, id)

	err = row.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Done)
	return 
}