package main

import (
    "book-management/config"
    "book-management/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    config.LoadEnv()
    db := config.ConnectDB()
    defer db.Close()

    r := gin.Default()
    routes.SetupRoutes(r, db)
    r.Run()
}