package config

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

var DB *pgxpool.Pool

func InitDB() {
    dsn := "postgres://postgres:123@localhost:5432/echoAPI"
    var err error
    DB, err = pgxpool.Connect(context.Background(), dsn)
    if err != nil {
        log.Fatalf("Unable to connect to database: %v\n", err)
    }
    log.Println("Connected to database")
}
