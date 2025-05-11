package routes

import (
    "database/sql"
    "book-management-fixed/handlers"
    "book-management-fixed/middlewares"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, db *sql.DB) {
    r.Use(middlewares.BasicAuthMiddleware())

    api := r.Group("/api")
    {
        handlers.RegisterCategoryRoutes(api, db)
        handlers.RegisterBookRoutes(api, db)
    }
}