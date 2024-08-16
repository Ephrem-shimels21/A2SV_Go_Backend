package infrastructure

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(cxt *gin.Context) {
		authHeader := cxt.GetHeader("Authorization")

		if authHeader == "" {
			cxt.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			cxt.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return []byte(""), nil
		})

		if err != nil || !token.Valid {
			cxt.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			cxt.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			cxt.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token claims"})
			cxt.Abort()
			return
		}

		cxt.Set("userID", claims["userID"])
		cxt.Set("username", claims["username"])
		cxt.Set("role", claims["role"])
		cxt.Next()
	}
}

func AdminOnlyMiddleware() gin.HandlerFunc {
	return func(cxt *gin.Context) {
		role, exists := cxt.Get("role")
		println("Role", role)
		if !exists || role != "admin" {
			cxt.JSON(http.StatusForbidden, gin.H{"error": "admin access required"})
			cxt.Abort()
			return
		}
		cxt.Next()
	}
}
