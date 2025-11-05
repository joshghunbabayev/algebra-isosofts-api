package registerRoutes

import (
	registerHandlers "algebra-isosofts-api/handlers/registers"

	"github.com/gin-gonic/gin"
)

// package tableComponentRoutes

// import (
// 	tableComponentHandlers "algebra-isosofts-api/handlers/tableComponents"

// 	"github.com/gin-gonic/gin"
// )

func BrRoutes(rg *gin.RouterGroup) {
	var brHandler registerHandlers.BrHandler
	rg.GET("/br", brHandler.GetAll)
}
