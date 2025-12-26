package registerRoutes

import (
	registerHandlers "algebra-isosofts-api/handlers/registers"

	"github.com/gin-gonic/gin"
)

func FBRoutes(rg *gin.RouterGroup) {
	var fbHandler registerHandlers.FBHandler
	rg.GET("/all", fbHandler.GetAll) // query: status
	rg.POST("/one", fbHandler.Create)
	rg.PUT("/one/:id", fbHandler.Update)
	rg.PUT("/all/archive", fbHandler.Archive)
	rg.PUT("/all/unarchive", fbHandler.Unarchive)
	rg.PUT("/all/delete", fbHandler.Delete)
	rg.PUT("/all/undelete", fbHandler.Undelete)
}
