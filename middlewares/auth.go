package middlewares

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func BasicAuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        user, pass, hasAuth := c.Request.BasicAuth()
        if !hasAuth || user != "admin" || pass != "secret" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            return
        }
        c.Next()
    }
}