package registerRoutes

import (
	registerHandlers "algebra-isosofts-api/handlers/registers"

	"github.com/gin-gonic/gin"
)

func VENRoutes(rg *gin.RouterGroup) {
	var venHandler registerHandlers.VENHandler
	rg.GET("/all", venHandler.GetAll) // query: status
	rg.POST("/one", venHandler.Create)
	rg.GET("/one/:id", venHandler.Get)
	rg.PUT("/one/:id", venHandler.Update)
	rg.PUT("/all/archive", venHandler.Archive)
	rg.PUT("/all/unarchive", venHandler.Unarchive)
	rg.PUT("/all/delete", venHandler.Delete)
	rg.PUT("/all/undelete", venHandler.Undelete)
}
