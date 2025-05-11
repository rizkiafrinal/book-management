package main

import (
    "book-management/config"
    "book-management/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    db := config.ConnectDB()
    defer db.Close()

    r := gin.Default()
    routes.SetupRoutes(r, db)
    r.Run()
}