package middlewares

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type RemoteAccount struct {
	Id            string `json:"id"`
	CompanyId     string `json:"companyId"`
	LineManagerId string `json:"lineManagerId"`
	IsAdmin       int8   `json:"isAdmin"`
	IsActive      int8   `json:"isActive"`
	Name          string `json:"name"`
	Surname       string `json:"surname"`
	Email         string `json:"email"`
	PhoneNumber   string `json:"phoneNumber"`
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")

		if token == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Token is required"})
			return
		}

		isosoftsUrl := os.Getenv("ISOSOFTS_API_URL") + "/api/algebra/self?token=" + token

		resp, err := http.Get(isosoftsUrl)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": "Identity service is unreachable"})
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized by Isosofts"})
			return
		}

		var account RemoteAccount
		if err := json.NewDecoder(resp.Body).Decode(&account); err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": "Failed to parse identity data"})
			return
		}

		c.Set("account", account)
		c.Next()
	}
}
