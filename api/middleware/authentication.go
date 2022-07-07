package middleware

import (
	"api/config"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

const BEARER_PREFIX = "Bearer "

type AuthScope struct {
	userId string
}

func Protected() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if len(authHeader) == 0 {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "error", "message": "No authorization header passed.", "data": nil})
			return
		}

		tokenString := authHeader[len(BEARER_PREFIX):]

		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(config.Config("AUTH_TOKEN_SECRET")), nil
		})

		fmt.Println("Token:", token)
		fmt.Println("Token Valid:", token.Valid)

		claims := token.Claims.(jwt.MapClaims)

		fmt.Println("Claims", claims["exp"])

		if token.Valid {
			authScope := AuthScope{
				userId: claims["iss"].(string),
			}
			c.Set("authScope", authScope)
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "Bad token.", "data": nil})
			return
		}
	}
}
