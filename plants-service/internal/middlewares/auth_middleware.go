package middlewares

import (
	"fmt"
	"net/http"
	"plant-care-app/plants-service/config"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cfg := config.GetInstance()

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "JWT not exist"})
			c.Abort()
			return
		}

		contents := strings.Split(authHeader, " ")

		if len(contents) != 2 || contents[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Format JWT wrong"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(contents[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				c.Abort()
				return nil, fmt.Errorf("unexpected method: %s", token.Header["alg"])
			}

			return []byte(cfg.GetJWTSercretKey()), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": err.Error()})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Can't get claims"})
			c.Abort()
			return
		}

		c.Set("user_id", claims["sub"])
		c.Next()
	}
}
