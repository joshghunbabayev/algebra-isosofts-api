package dashboardRoutes

import (
	"algebra-isosofts-api/middlewares"

	"github.com/gin-gonic/gin"
)

func MainRoutes(rg *gin.RouterGroup) {
	rg.Use(middlewares.AuthMiddleware())
	rg.Use(middlewares.AccessMiddleware())
	QhseKPIRoutes(rg.Group("/qhseKpi"))
	ActionLogRoutes(rg.Group("/actionLog"))
}
