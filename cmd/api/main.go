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
	
	// Apply CORS middleware globally
	r.Use(middleware.CORSMiddleware)

	// Public routes
	r.HandleFunc("/register", userHandler.Register).Methods("POST", "OPTIONS")
	r.HandleFunc("/login", userHandler.Login).Methods("POST", "OPTIONS")

	// Protected routes
	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.AuthMiddleware(cfg.JWTSecretKey))

	api.HandleFunc("/todos", todoHandler.GetTodos).Methods("GET", "OPTIONS")
	api.HandleFunc("/todos", todoHandler.CreateTodo).Methods("POST", "OPTIONS")
	api.HandleFunc("/todos/{id}", todoHandler.UpdateTodo).Methods("PUT", "OPTIONS")
	api.HandleFunc("/todos/{id}", todoHandler.DeleteTodo).Methods("DELETE", "OPTIONS")

	log.Printf("🚀 Server running on port: %s", cfg.APIPort)

	if err := http.ListenAndServe(":"+cfg.APIPort, r); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}