package registerRoutes

import (
	registerHandlers "algebra-isosofts-api/handlers/registers"

	"github.com/gin-gonic/gin"
)

func AOPRoutes(rg *gin.RouterGroup) {
	var aopHandler registerHandlers.AOPHandler
	rg.GET("/all", aopHandler.GetAll) // query: status
	rg.POST("/one", aopHandler.Create)
	rg.PUT("/one/:id", aopHandler.Update)
	rg.PUT("/all/archive", aopHandler.Archive)
	rg.PUT("/all/unarchive", aopHandler.Unarchive)
	rg.PUT("/all/delete", aopHandler.Delete)
	rg.PUT("/all/undelete", aopHandler.Undelete)
}
