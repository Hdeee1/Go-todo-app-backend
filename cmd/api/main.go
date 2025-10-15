package main

import (
	"log"
	"net/http"
	"todo-app-backend/internal/config"
	"todo-app-backend/internal/handler"
	"todo-app-backend/internal/middleware"
	"todo-app-backend/internal/store"

	"github.com/gorilla/mux"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	dbStore, err := store.NewStore(cfg.DBSource)
	if err != nil {
		log.Fatal(err)
	}

	userStore := &store.UserStore{DB: dbStore.DB}
	userHandler := &handler.UserHandler{Store: userStore, Secret: cfg.JWTSecretKey}

	todoStore := store.TodoStore{DB: dbStore.DB}
	todoHandler := handler.NewTodoHandler(todoStore)

	r := mux.NewRouter()

	r.HandleFunc("/register", userHandler.Register).Methods("POST")
	r.HandleFunc("/login", userHandler.Login).Methods("POST")

	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.AuthMiddleware(cfg.JWTSecretKey))

	api.HandleFunc("/todos", todoHandler.GetTodos).Methods("GET")
	api.HandleFunc("/todos", todoHandler.CreateTodo).Methods("POST")
	api.HandleFunc("/todos/{id}", todoHandler.UpdateTodo).Methods("PUT")
	api.HandleFunc("/todos/{id}", todoHandler.DeleteTodo).Methods("DELETE")

	log.Printf("Server running on port : %s", cfg.APIPort)

	if err := http.ListenAndServe(":"+cfg.APIPort, r); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}