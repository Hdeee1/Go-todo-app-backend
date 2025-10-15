package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo-app-backend/internal/middleware"
	"todo-app-backend/internal/model"

	"github.com/gorilla/mux"
)

type TodoHandler struct {
	Store	*store.TodoStore
}

func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(int64)
	if !ok {
		http.Error(w, "user not found", http.StatusInternalServerError)
		return
	}


	var data struct {
		Task	string `json:"task"`
	}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "failed to process json input", http.StatusBadRequest)
		return
	}

	if data.Task == "" {
		http.Error(w, "data can't be empty", http.StatusBadRequest)
		return
	}

	todo :=  model.Todo {
		UserID: userID,
		Task: data.Task,
	}

	if err := h.Store.CreateTodo(&todo); err != nil {
		http.Error(w, "failed to create to do", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

func (h *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	userID, ok :=  r.Context().Value(middleware.UserIDKey).(int64)
	if !ok {
		http.Error(w, "user not found", http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "missing todo id", http.StatusBadRequest)
		return
	}

	todoID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid todo id", http.StatusBadRequest)
		return
	}

	if err := h.Store.DeleteTodo(userID, todoID); err != nil {
		http.Error(w, "failed to delete todo", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}