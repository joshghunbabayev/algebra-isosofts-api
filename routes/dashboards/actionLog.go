package dashboardRoutes

import (
	dashboardHandlers "algebra-isosofts-api/handlers/dashboards"

	"github.com/gin-gonic/gin"
)

func ActionLogRoutes(rg *gin.RouterGroup) {
	var actionHandler dashboardHandlers.ActionLogHandler
	rg.GET("/all", actionHandler.GetAll) // query: status
}
