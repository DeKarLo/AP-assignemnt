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
	db, err := postgres.ConnectDB("localhost", "5432", "postgres", "dek@123455", "contact")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Connected to the database")
	defer db.Close(context.Background())

	contactRepo := repository.NewContactRepository()
	groupRepo := repository.NewGroupRepository()
	contactUsecase := usecase.NewContactUsecase(contactRepo, groupRepo)
	contactHandler := delivery.NewContactHandler(contactUsecase)

	http.HandleFunc("/contacts", contactHandler.HandleContacts)

	fmt.Println("Сервер запущен на порту :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Ошибка при запуске сервера: %s\n", err)
	}
}
