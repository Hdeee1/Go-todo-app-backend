package store

import (
	"database/sql"
	"errors"
	"todo-app-backend/internal/model"
)

type UserStore struct {
	DB	*sql.DB
}

func (s *UserStore) CreateUser(username, hashedPassword string) error {
	_, err := s.DB.Exec(`INSERT INTO users (username, hashed_password) VALUES (?, ?) `,username, hashedPassword)
	return err
}

func (s *UserStore) GetUserByUsername(username string) (*model.User, error) {
	var u model.User
	err := s.DB.QueryRow(`SELECT id, username, hashed_password, created_at FROM users WHERE username = ?`, username).Scan(&u.ID, &u.Username, &u.HashedPassword, &u.CreatedAt)
	
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}

	return &u, err
}
