package middlewares

import (
    "net/http"
    "os"

    "github.com/gin-gonic/gin"
)

func BasicAuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        user, pass, hasAuth := c.Request.BasicAuth()
        if !hasAuth || user != os.Getenv("BASIC_AUTH_USER") || pass != os.Getenv("BASIC_AUTH_PASS") {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            return
        }
        c.Next()
    }
}