package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

func ConnectDB(host, port, user, password, dbname string) (*pgx.Conn, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
