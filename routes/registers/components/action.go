package registerComponentRoutes

import (
	registerComponentHandlers "algebra-isosofts-api/handlers/registers/components"

	"github.com/gin-gonic/gin"
)

func ActionRoutes(rg *gin.RouterGroup) {
	var actionHandler registerComponentHandlers.ActionHandler
	rg.GET("/all", actionHandler.GetAll) // query: status, registerId
	rg.POST("/one", actionHandler.Create)
	rg.PUT("/one/:id", actionHandler.Update)
	rg.PUT("/all/delete", actionHandler.Delete)
	rg.PUT("/all/undelete", actionHandler.Undelete)
}
