package registerRoutes

import "github.com/gin-gonic/gin"

func MainRoutes(rg *gin.RouterGroup) {
	BrRoutes(rg.Group("/br"))
}
