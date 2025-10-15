package store

import (
	"database/sql"
	"errors"
	"todo-app-backend/internal/model"
)

type TodoStore struct {
	DB *sql.DB
}

func (t *TodoStore) GetTodosByUserID(userID int64) ([]*model.Todo, error) {
    rows, err := t.DB.Query(
        `SELECT id, user_id, task, completed, created_at 
         FROM todos WHERE user_id = ? ORDER BY created_at DESC`,
        userID,
    )
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var todos []*model.Todo
    for rows.Next() {
        var todo model.Todo
        err := rows.Scan(
            &todo.ID,
            &todo.UserID,
            &todo.Task,
            &todo.Completed,
            &todo.CreatedAt,
        )
        if err != nil {
            return nil, err
        }
        todos = append(todos, &todo)
    }

    return todos, rows.Err()
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

func (t *TodoStore) UpdateTodoStatus(userID, todoID int64, completed bool) error {
    result, err := t.DB.Exec(
        `UPDATE todos SET completed = ? WHERE id = ? AND user_id = ?`,
        completed, todoID, userID,
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
