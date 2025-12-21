package registerRoutes

import (
	registerHandlers "algebra-isosofts-api/handlers/registers"

	"github.com/gin-gonic/gin"
)

func DOCRoutes(rg *gin.RouterGroup) {
	var docHandler registerHandlers.DOCHandler
	rg.GET("/all", docHandler.GetAll) // query: status
	rg.POST("/one", docHandler.Create)
	rg.PUT("/one/:id", docHandler.Update)
	rg.PUT("/all/archive", docHandler.Archive)
	rg.PUT("/all/unarchive", docHandler.Unarchive)
	rg.PUT("/all/delete", docHandler.Delete)
	rg.PUT("/all/undelete", docHandler.Undelete)
}
