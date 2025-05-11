package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	dbURL := "host=postgres.railway.internal port=5432 user=postgres password=RJynwNElWuDSbyYcJyxpbrkyXVPJraoL dbname=railway sslmode=disable"
	if dbURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error connecting to DB: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("Cannot reach DB: %v", err)
	}
	fmt.Println("Database connected!")
	return db
}
