package dashboardRoutes

import (
	dashboardHandlers "algebra-isosofts-api/handlers/dashboards"

	"github.com/gin-gonic/gin"
)

func QhseKPIRoutes(rg *gin.RouterGroup) {
	var qhseKPIHandler dashboardHandlers.QhseKPIHandler
	rg.GET("", qhseKPIHandler.GetAll)
	rg.PUT("/:id", qhseKPIHandler.Update)
}
