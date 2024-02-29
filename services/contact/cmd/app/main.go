package main

import (
	"architecture_go/pkg/store/postgres"
	"architecture_go/services/contact/cmd/internal/delivery"
	"architecture_go/services/contact/cmd/internal/repository"
	"architecture_go/services/contact/cmd/internal/usecase"
	"context"
	"fmt"
	"net/http"
)

func main() {
	// Connect to the database
	db, err := postgres.ConnectDB("localhost", "5432", "postgres", "dek@123455", "contact")
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
	defer db.Close(context.Background())

	contactRepo := repository.NewContactRepository(db)
	groupRepo := repository.NewGroupRepository(db)

	contactUsecase := usecase.NewContactUsecase(contactRepo, groupRepo)

	contactHandler := delivery.NewContactHandler(*contactUsecase)

	http.HandleFunc("/contacts", contactHandler.HandleContacts)

	fmt.Println("Server is running on port :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
