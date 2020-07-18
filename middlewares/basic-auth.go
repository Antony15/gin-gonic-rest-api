// Middlewares package for proper http error statuses
package middlewares

import "github.com/gin-gonic/gin"

// Function for basic authentication
func BasicAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"admin": "admin123",
	})
}
