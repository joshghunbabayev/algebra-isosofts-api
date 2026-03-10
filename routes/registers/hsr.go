package registerRoutes

import (
	registerHandlers "algebra-isosofts-api/handlers/registers"
	"algebra-isosofts-api/middlewares"

	"github.com/gin-gonic/gin"
)

func HSRRoutes(rg *gin.RouterGroup) {
	var hsrHandler registerHandlers.HSRHandler
	rg.Use(middlewares.AccessMiddleware())

	rg.GET("/all", hsrHandler.GetAll) // query: status
	rg.POST("/one", hsrHandler.Create)
	rg.PUT("/one/:id", hsrHandler.Update)
	rg.PUT("/all/archive", hsrHandler.Archive)
	rg.PUT("/all/unarchive", hsrHandler.Unarchive)
	rg.PUT("/all/delete", hsrHandler.Delete)
	rg.PUT("/all/undelete", hsrHandler.Undelete)
}
