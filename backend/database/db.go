package database

import (
	"fmt"
	"os"

	"github.com/jackc/pgx"
)

func NewConnection() *pgx.ConnPool {
	pgxPool, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			User:     "admin",
			Password: "admin",
			Database: "db",
		},
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to the database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Connected to database")

	return pgxPool
}
