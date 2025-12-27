package logRoutes

import (
	logHandlers "algebra-isosofts-api/handlers/logs"

	"github.com/gin-gonic/gin"
)

func ActionRoutes(rg *gin.RouterGroup) {
	var actionHandler logHandlers.ActionHandler
	rg.GET("/all", actionHandler.GetAll) // query: status
}
