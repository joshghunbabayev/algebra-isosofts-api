package registerRoutes

import (
	registerHandlers "algebra-isosofts-api/handlers/registers"
	"algebra-isosofts-api/middlewares"

	"github.com/gin-gonic/gin"
)

func TRARoutes(rg *gin.RouterGroup) {
	var traHandler registerHandlers.TRAHandler
	rg.Use(middlewares.AccessMiddleware())

	rg.GET("/all", traHandler.GetAll) // query: status
	rg.POST("/one", traHandler.Create)
	rg.PUT("/one/:id", traHandler.Update)
	rg.PUT("/all/archive", traHandler.Archive)
	rg.PUT("/all/unarchive", traHandler.Unarchive)
	rg.PUT("/all/delete", traHandler.Delete)
	rg.PUT("/all/undelete", traHandler.Undelete)
}
