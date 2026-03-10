package registerRoutes

import (
	registerHandlers "algebra-isosofts-api/handlers/registers"
	"algebra-isosofts-api/middlewares"

	"github.com/gin-gonic/gin"
)

func MRMRoutes(rg *gin.RouterGroup) {
	var mrmHandler registerHandlers.MRMHandler
	rg.Use(middlewares.AccessMiddleware())

	rg.GET("/all", mrmHandler.GetAll) // query: status
	rg.POST("/one", mrmHandler.Create)
	rg.PUT("/one/:id", mrmHandler.Update)
	rg.PUT("/all/archive", mrmHandler.Archive)
	rg.PUT("/all/unarchive", mrmHandler.Unarchive)
	rg.PUT("/all/delete", mrmHandler.Delete)
	rg.PUT("/all/undelete", mrmHandler.Undelete)
}
