package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/willamesSantoos/api-todo-list/internal/configs"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	stringConnection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Password, conf.Database)

	conn, err := sql.Open("postgres", stringConnection)

	if err != nil {
		panic(err)
	}

	err = conn.Ping()

	return conn, err
}
