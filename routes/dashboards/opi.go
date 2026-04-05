package dashboardRoutes

import (
	dashboardHandlers "algebra-isosofts-api/handlers/dashboards"

	"github.com/gin-gonic/gin"
)

func OPIRoutes(rg *gin.RouterGroup) {
	var opiHandler dashboardHandlers.OPIHandler
	rg.GET("/all", opiHandler.GetAll) // query: status
	rg.POST("/one", opiHandler.Create)
	rg.PUT("/one/:id", opiHandler.Update)
	rg.PUT("/all/archive", opiHandler.Archive)
	rg.PUT("/all/unarchive", opiHandler.Unarchive)
	rg.PUT("/all/delete", opiHandler.Delete)
	rg.PUT("/all/undelete", opiHandler.Undelete)
}
