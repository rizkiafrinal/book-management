package main

import (
    "github.com/gin-gonic/gin"
    "log"
    "os"
    "bookstore/config"
    "bookstore/handler"
    "bookstore/middleware"
    "bookstore/repository"
    "bookstore/service"
    _ "github.com/lib/pq"
)

func main() {
    db, err := config.ConnectDB()
    if err != nil {
        log.Fatal("Database connection failed:", err)
    }

    categoryRepo := repository.NewCategoryRepository(db)
    bookRepo := repository.NewBookRepository(db)
    categoryService := service.NewCategoryService(categoryRepo)
    bookService := service.NewBookService(bookRepo)
    categoryHandler := handler.NewCategoryHandler(categoryService)
    bookHandler := handler.NewBookHandler(bookService)

    r := gin.Default()
    r.Use(middleware.BasicAuth())

    api := r.Group("/api")
    {
        cat := api.Group("/categories")
        {
            cat.GET("", categoryHandler.GetAll)
            cat.POST("", categoryHandler.Create)
            cat.GET("/:id", categoryHandler.GetByID)
            cat.DELETE("/:id", categoryHandler.Delete)
            cat.GET("/:id/books", categoryHandler.GetBooksByCategory)
        }

        book := api.Group("/books")
        {
            book.GET("", bookHandler.GetAll)
            book.POST("", bookHandler.Create)
            book.GET("/:id", bookHandler.GetByID)
            book.DELETE("/:id", bookHandler.Delete)
        }
    }

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    r.Run(":" + port)
}