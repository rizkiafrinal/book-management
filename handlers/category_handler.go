package handlers

import (
	"book-management/repositories"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterCategoryRoutes(rg *gin.RouterGroup, db *sql.DB) {
	categoryRoute := rg.Group("/categories")
	categoryRoute.GET("/", func(c *gin.Context) {
		categories, err := repositories.GetAllCategories(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, categories)
	})
}
