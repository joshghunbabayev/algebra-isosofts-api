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
	rg.PUT("/one/:id/archive", brHandler.Archive)
	rg.PUT("/one/:id/unarchive", brHandler.Unarchive)
	rg.PUT("/one/:id/delete", brHandler.Delete)
	rg.PUT("/one/:id/undelete", brHandler.Undelete)
}
