package dashboardRoutes

import (
	dashboardHandlers "algebra-isosofts-api/handlers/dashboards"

	"github.com/gin-gonic/gin"
)

func KPIRoutes(rg *gin.RouterGroup) {
	var kpiHandler dashboardHandlers.KPIHandler
	rg.GET("/all", kpiHandler.GetAll) // query: status
}
