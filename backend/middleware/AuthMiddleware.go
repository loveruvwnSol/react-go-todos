package middleware

import (
	"app/app/handler"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
		claims, err := handler.ParseJWT(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		c.Set("user_id", claims.UserID)
		c.Next()
	}
}
