package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func ConnectDB() *sql.DB {
	dbURL := os.Getenv("DATABASE_URL")
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
