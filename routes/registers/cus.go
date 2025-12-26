package registerRoutes

import (
	registerHandlers "algebra-isosofts-api/handlers/registers"

	"github.com/gin-gonic/gin"
)

func CUSRoutes(rg *gin.RouterGroup) {
	var cusHandler registerHandlers.CUSHandler
	rg.GET("/all", cusHandler.GetAll) // query: status
	rg.POST("/one", cusHandler.Create)
	rg.GET("/one/:id", cusHandler.Get)
	rg.PUT("/one/:id", cusHandler.Update)
	rg.PUT("/all/archive", cusHandler.Archive)
	rg.PUT("/all/unarchive", cusHandler.Unarchive)
	rg.PUT("/all/delete", cusHandler.Delete)
	rg.PUT("/all/undelete", cusHandler.Undelete)
}
