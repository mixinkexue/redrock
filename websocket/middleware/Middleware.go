package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cookie() gin.HandlerFunc {
	return func(c *gin.Context) {

		_, err := c.Request.Cookie("username")
		if err == nil {
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "没有权限",
			})
			c.Abort()
		}
	}
}