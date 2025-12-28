package dashboardRoutes

import "github.com/gin-gonic/gin"

func MainRoutes(rg *gin.RouterGroup) {
	ActionLogRoutes(rg.Group("/actionLog"))
}
