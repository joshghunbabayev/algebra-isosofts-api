package registerRoutes

import (
	registerHandlers "algebra-isosofts-api/handlers/registers"

	"github.com/gin-gonic/gin"
)

func BrRoutes(rg *gin.RouterGroup) {
	var brHandler registerHandlers.BrHandler
	rg.GET("/", brHandler.GetAll)
	rg.POST("/", brHandler.Create)
	rg.PUT("/:id", brHandler.Update)
}
