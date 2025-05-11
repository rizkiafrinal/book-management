package main

import (
    "book-management-fixed/config"
    "book-management-fixed/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    db := config.ConnectDB()
    defer db.Close()

    r := gin.Default()
    routes.SetupRoutes(r, db)
    r.Run()
}