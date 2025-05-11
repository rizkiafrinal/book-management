package config

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
    dbURL := "postgres://postgres:RJynwNElWuDSbyYcJyxpbrkyXVPJraoL@postgres.railway.internal:5432/railway"
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