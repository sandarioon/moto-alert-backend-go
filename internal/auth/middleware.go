package auth

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sandarioon/moto-alert-backend-go/internal/errors"
)

type Claims struct {
	UserId int `json:"user_id"`
	jwt.RegisteredClaims
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		// I do not use most common "Bearer ${jwt}"" or "Token ${jwt}" implementation
		// That is not a mistake
		// I need backwards compatibility with old clients ¯\_(ツ)_/¯
		if authHeader == "" {
			errors.NewErrorResponse(c, http.StatusUnauthorized, "authorization header missing")
			c.Abort()
			return
		}

		claims := &Claims{}

		token, err := jwt.ParseWithClaims(authHeader, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			errors.NewErrorResponse(c, http.StatusUnauthorized, "invalid or expired token")
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserId)
		c.Next()
	}
}
