package middleware

import (
	"api/config"
	"api/structs"

	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

const BEARER_PREFIX = "Bearer "

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

		claims := token.Claims.(jwt.MapClaims)

		userId, _ := strconv.Atoi(claims["iss"].(string))

		if token.Valid {
			authScope := structs.AuthScope{
				UserID: userId,
			}
			c.Set("authScope", authScope)
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "Bad token.", "data": nil})
			return
		}
	}
}
