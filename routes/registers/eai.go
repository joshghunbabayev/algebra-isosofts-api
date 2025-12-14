package registerRoutes

import (
	registerHandlers "algebra-isosofts-api/handlers/registers"

	"github.com/gin-gonic/gin"
)

func EAIRoutes(rg *gin.RouterGroup) {
	var eaiHandler registerHandlers.EAIHandler
	rg.GET("/all", eaiHandler.GetAll) // query: status
	rg.POST("/one", eaiHandler.Create)
	rg.PUT("/one/:id", eaiHandler.Update)
	rg.PUT("/all/archive", eaiHandler.Archive)
	rg.PUT("/all/unarchive", eaiHandler.Unarchive)
	rg.PUT("/all/delete", eaiHandler.Delete)
	rg.PUT("/all/undelete", eaiHandler.Undelete)
}
