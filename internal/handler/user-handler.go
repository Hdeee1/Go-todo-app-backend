package handler

import (
	"encoding/json"
	"net/http"
	"todo-app-backend/internal/auth"
	"todo-app-backend/internal/store"

	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct  {
	Username	string	`json:"username"`
	Password 	string	`json:"password"`
}

type UserHandler struct {
	Store	*store.UserStore
	Secret	string
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username	string	`json:"username"`
		Password	string	`json:"password"`
	} 

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "failed to hash password", http.StatusInternalServerError)
		return
	}

	err = h.Store.CreateUser(req.Username, string(hashedPassword))
	if err != nil {
		http.Error(w, "failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "user registered"})
}


func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username		string	`json:"username"`
		Password		string	`json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	user, err := h.Store.GetUserByUsername(req.Username)
	if err != nil {
		http.Error(w, "invalid username or password", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(req.Password))
	if err !=  nil {
		http.Error(w, "invalid username or password", http.StatusUnauthorized)
		return
	}

	token, err := auth.GenerateToken(h.Secret, user.ID)
	if err != nil {
		http.Error(w, "failed to generate token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
