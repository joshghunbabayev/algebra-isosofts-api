package dashboardRoutes

import "github.com/gin-gonic/gin"

func MainRoutes(rg *gin.RouterGroup) {
	KPIRoutes(rg.Group("/kpi"))
	ActionLogRoutes(rg.Group("/actionLog"))
}
