package registerRoutes

import (
	registerHandlers "algebra-isosofts-api/handlers/registers"
	"algebra-isosofts-api/middlewares"

	"github.com/gin-gonic/gin"
)

func BRRoutes(rg *gin.RouterGroup) {
	var brHandler registerHandlers.BRHandler
	rg.Use(middlewares.AccessMiddleware())

	rg.GET("/all", brHandler.GetAll) // query: status
	rg.POST("/one", brHandler.Create)
	rg.PUT("/one/:id", brHandler.Update)
	rg.PUT("/all/archive", brHandler.Archive)
	rg.PUT("/all/unarchive", brHandler.Unarchive)
	rg.PUT("/all/delete", brHandler.Delete)
	rg.PUT("/all/undelete", brHandler.Undelete)
}
