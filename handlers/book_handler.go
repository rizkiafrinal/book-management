package handlers

import (
    "database/sql"
    "net/http"
    "book-management/repositories"
    "github.com/gin-gonic/gin"
)

func RegisterBookRoutes(rg *gin.RouterGroup, db *sql.DB) {
    bookRoute := rg.Group("/books")
    bookRoute.GET("/", func(c *gin.Context) {
        books, err := repositories.GetAllBooks(db)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, books)
    })
}