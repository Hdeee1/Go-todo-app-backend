package store

import (
	"database/sql"
	"errors"
	"todo-app-backend/internal/model"
)

type TodoStore struct {
	DB *sql.DB
}

func (t *TodoStore) GetTodosByUserID() {
	
}

func (t *TodoStore) CreateTodo(todo *model.Todo) error {
	result, err := t.DB.Exec(
		`INSERT INTO todos (user_id, task, completed) VALUES (?, ?, ?)`,
		todo.UserID, todo.Task, false,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	todo.ID = id
	return nil
}

func (t *TodoStore) DeleteTodo(userID, todoID int64) error {
	result, err := t.DB.Exec(
		`DELETE FROM todos WHERE id = ? AND user_id = ?`,
		todoID, userID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("todo not found or unauthorized")
	}

	return nil
}
