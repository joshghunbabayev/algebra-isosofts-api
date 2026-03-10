package registerRoutes

import (
	registerHandlers "algebra-isosofts-api/handlers/registers"
	"algebra-isosofts-api/middlewares"

	"github.com/gin-gonic/gin"
)

func LEGRoutes(rg *gin.RouterGroup) {
	var legHandler registerHandlers.LEGHandler
	rg.Use(middlewares.AccessMiddleware())

	rg.GET("/all", legHandler.GetAll) // query: status
	rg.POST("/one", legHandler.Create)
	rg.PUT("/one/:id", legHandler.Update)
	rg.PUT("/all/archive", legHandler.Archive)
	rg.PUT("/all/unarchive", legHandler.Unarchive)
	rg.PUT("/all/delete", legHandler.Delete)
	rg.PUT("/all/undelete", legHandler.Undelete)
}
