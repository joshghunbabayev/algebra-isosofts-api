package registerRoutes

import (
	registerHandlers "algebra-isosofts-api/handlers/registers"

	"github.com/gin-gonic/gin"
)

func EARoutes(rg *gin.RouterGroup) {
	var eaHandler registerHandlers.EAHandler
	rg.GET("/all", eaHandler.GetAll) // query: status
	rg.POST("/one", eaHandler.Create)
	rg.PUT("/one/:id", eaHandler.Update)
	rg.PUT("/all/archive", eaHandler.Archive)
	rg.PUT("/all/unarchive", eaHandler.Unarchive)
	rg.PUT("/all/delete", eaHandler.Delete)
	rg.PUT("/all/undelete", eaHandler.Undelete)
}
