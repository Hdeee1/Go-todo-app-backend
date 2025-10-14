package main

import (
	"log"
	"net/http"
	"todo-app-backend/internal/config"
	"todo-app-backend/internal/handler"
	"todo-app-backend/internal/store"
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

	http.HandleFunc("/register", userHandler.Register)
	http.HandleFunc("/login", userHandler.Login)

	log.Println("Server running on port :", cfg.APIPort)
	http.ListenAndServe(":"+cfg.APIPort, nil)
}

