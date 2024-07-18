package middlewares

import (
	"api_goshop/handleError"
	"api_goshop/helper"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			handleError.HandleError(c, handleError.UnauthorizedError{Message: "Unauthorization"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header"})
			c.Abort()
			return
		}
		tokenStr := parts[1]

		userID, err := helper.ValidateToken(tokenStr)
		if err != nil {
			handleError.HandleError(c, handleError.UnauthorizedError{Message: err.Error()})
			c.Abort()
			return
		}
		c.Set("userID", userID)
		c.Next()
	}
}
