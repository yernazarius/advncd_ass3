package main

import (
	"architecture_go/pkg/store/postgres"
	"architecture_go/services/contact/internal/delivery"
	"architecture_go/services/contact/internal/repository"
	"architecture_go/services/contact/internal/usecase"
	"log"
	"net/http"
)

func main() {
	db, err := postgres.DBconnection("postgres", "admin", "localhost", "5432", "clean_architecture")
	if err != nil {
		log.Fatalf("Connection failed: %v", err)
	}

	defer db.Close()
	log.Println("Connected to the database")

	repo := repository.NewContactRepository()
	usecase := usecase.NewContactUseCase(repo)
	delivery := delivery.NewContactDelivery(usecase)

	_ = delivery

	log.Println("Server is starting on port :4000")
	if err := http.ListenAndServe(":4000", nil); err != nil {
		log.Fatalf("The server could not be started: %v", err)
	}

}
