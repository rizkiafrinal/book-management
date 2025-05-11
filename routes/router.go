package routes

import (
	"book-management/handlers"
	"book-management/middlewares"
	"database/sql"

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
