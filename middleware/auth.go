package middleware

import (
	"net/http"
	"strings"

	"github.com/clapat-bb/memo/logger"
	"github.com/clapat-bb/memo/util"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Pls access after login",
			})
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := util.ParseToken(tokenStr)
		if err != nil {
			logger.Log.Warnf("token parse error: %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expire TOKEN"})
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Next()
	}
}
