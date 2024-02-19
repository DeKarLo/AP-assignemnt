package main

import (
	"architecture_go/pkg/store/postgres"
	"context"
	"fmt"
)

func main() {
	db, err := postgres.ConnectDB("localhost", "5432", "postgres", "dek@123455", "contact")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Connected to the database")
	defer db.Close(context.Background())
}
