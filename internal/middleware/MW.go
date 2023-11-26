package middleware

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)
//middleware that checks the username and if his name is "admin" then writes "admin" to the console
func RoleCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		val := c.Request.Header.Get("User-Role")
		val = strings.ToLower(val)

		if strings.Contains(val, "admin") {
			log.Println("red button user detected")
		}

		c.Next()

	}
}
