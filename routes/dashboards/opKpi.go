package dashboardRoutes

import (
	dashboardHandlers "algebra-isosofts-api/handlers/dashboards"

	"github.com/gin-gonic/gin"
)

func OpKPIRoutes(rg *gin.RouterGroup) {
	var opKPIHandler dashboardHandlers.OpKPIHandler
	rg.GET("/all", opKPIHandler.GetAll) // query: status
	rg.POST("/one", opKPIHandler.Create)
	rg.PUT("/one/:id", opKPIHandler.Update)
	rg.PUT("/all/archive", opKPIHandler.Archive)
	rg.PUT("/all/unarchive", opKPIHandler.Unarchive)
	rg.PUT("/all/delete", opKPIHandler.Delete)
	rg.PUT("/all/undelete", opKPIHandler.Undelete)
}
