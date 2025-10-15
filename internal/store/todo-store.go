package store

import (
	"database/sql"
)

type TodoStore struct {
	DB	*sql.DB
}

func (t *TodoStore) CreateTodo(todo string) error {
	_, err := t.DB.Exec(`INSERT INTO todos ()`)
}