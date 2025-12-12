package registerRoutes

import (
	registerHandlers "algebra-isosofts-api/handlers/registers"

	"github.com/gin-gonic/gin"
)

func EIARoutes(rg *gin.RouterGroup) {
	var eiaHandler registerHandlers.EIAHandler
	rg.GET("/all", eiaHandler.GetAll) // query: status
	rg.POST("/one", eiaHandler.Create)
	rg.PUT("/one/:id", eiaHandler.Update)
	rg.PUT("/all/archive", eiaHandler.Archive)
	rg.PUT("/all/unarchive", eiaHandler.Unarchive)
	rg.PUT("/all/delete", eiaHandler.Delete)
	rg.PUT("/all/undelete", eiaHandler.Undelete)
}
