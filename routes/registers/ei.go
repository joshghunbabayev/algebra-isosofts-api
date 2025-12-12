package registerRoutes

import (
	registerHandlers "algebra-isosofts-api/handlers/registers"

	"github.com/gin-gonic/gin"
)

func EIRoutes(rg *gin.RouterGroup) {
	var eiHandler registerHandlers.EIHandler
	rg.GET("/all", eiHandler.GetAll) // query: status
	rg.POST("/one", eiHandler.Create)
	rg.PUT("/one/:id", eiHandler.Update)
	rg.PUT("/all/archive", eiHandler.Archive)
	rg.PUT("/all/unarchive", eiHandler.Unarchive)
	rg.PUT("/all/delete", eiHandler.Delete)
	rg.PUT("/all/undelete", eiHandler.Undelete)
}
