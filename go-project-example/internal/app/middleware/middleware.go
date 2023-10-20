package middleware

import "github.com/gin-gonic/gin"

func AuthMiddileware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// auth

		c.Next()
	}
}
