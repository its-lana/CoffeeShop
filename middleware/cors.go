package middleware

import "github.com/gin-gonic/gin"

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set specific origin
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		// c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		// Allow credentials
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		// Allow specific headers
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Strict-Transport-Security, X-Content-Type-Options, X-Frame-Options, X-XSS-Protection, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		// Allow specific methods
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		// Expose specific headers
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, Authorization")

		// Handle preflight OPTIONS request
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
