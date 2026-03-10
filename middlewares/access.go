package middlewares

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func AccessMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "GET" {
			c.Next()
			return
		}

		fullPath := c.Request.URL.Path
		parts := strings.Split(strings.Trim(fullPath, "/"), "/")

		if len(parts) < 3 {
			c.AbortWithStatusJSON(400, gin.H{"error": "Invalid route structure for access check"})
			return
		}
		register := parts[2]

		token := c.Query("token")
		if token == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Token is required for access check"})
			return
		}

		isosoftsUrl := os.Getenv("ISOSOFTS_API_URL") + "/api/algebra/check-access?register=" + register + "&token=" + token

		resp, err := http.Get(isosoftsUrl)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": "Identity service is unreachable"})
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			c.AbortWithStatusJSON(403, gin.H{
				"error":    "You don't have write access for this register",
				"register": register,
			})
			return
		}

		c.Next()
	}
}
