package registerRoutes

import (
	registerHandlers "algebra-isosofts-api/handlers/registers"
	"algebra-isosofts-api/middlewares"

	"github.com/gin-gonic/gin"
)

func MOCRoutes(rg *gin.RouterGroup) {
	var mocHandler registerHandlers.MOCHandler
	rg.Use(middlewares.AccessMiddleware())

	rg.GET("/all", mocHandler.GetAll) // query: status
	rg.POST("/one", mocHandler.Create)
	rg.PUT("/one/:id", mocHandler.Update)
	rg.PUT("/all/archive", mocHandler.Archive)
	rg.PUT("/all/unarchive", mocHandler.Unarchive)
	rg.PUT("/all/delete", mocHandler.Delete)
	rg.PUT("/all/undelete", mocHandler.Undelete)
}
