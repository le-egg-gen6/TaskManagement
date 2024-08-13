package middleware

import (
	"github.com/gin-gonic/gin"
	"go-ecommerce/utils/tokenutil"
	"net/http"
	"strings"
)

func JWTAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{})
			c.Abort()
		}
		authToken := t[1]
		authorized, err := tokenutil.IsAuthorized(authToken, secret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{})
			c.Abort()
			return
		}
		if authorized {
			userID, err := tokenutil.ExtractIDFromToken(authToken, secret)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{})
				c.Abort()
				return
			}
			c.Set("x-user-id", userID)
			c.Next()
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{})
		c.Abort()
		return
	}
}
