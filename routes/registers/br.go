package registerRoutes

import (
	registerHandlers "algebra-isosofts-api/handlers/registers"

	"github.com/gin-gonic/gin"
)

func BrRoutes(rg *gin.RouterGroup) {
	var brHandler registerHandlers.BrHandler
	rg.GET("/all", brHandler.GetAll)
	rg.POST("/one", brHandler.Create)
	rg.PUT("/one/:id", brHandler.Update)
	rg.PUT("/all/archive", brHandler.Archive)
	rg.PUT("/all/unarchive", brHandler.Unarchive)
	rg.PUT("/all/delete", brHandler.Delete)
	rg.PUT("/all/undelete", brHandler.Undelete)
}
