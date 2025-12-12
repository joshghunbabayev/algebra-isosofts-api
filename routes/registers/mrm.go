package registerRoutes

import (
	registerHandlers "algebra-isosofts-api/handlers/registers"

	"github.com/gin-gonic/gin"
)

func MRMRoutes(rg *gin.RouterGroup) {
	var mrmHandler registerHandlers.MRMHandler
	rg.GET("/all", mrmHandler.GetAll) // query: status
	rg.POST("/one", mrmHandler.Create)
	rg.PUT("/one/:id", mrmHandler.Update)
	rg.PUT("/all/archive", mrmHandler.Archive)
	rg.PUT("/all/unarchive", mrmHandler.Unarchive)
	rg.PUT("/all/delete", mrmHandler.Delete)
	rg.PUT("/all/undelete", mrmHandler.Undelete)
}
