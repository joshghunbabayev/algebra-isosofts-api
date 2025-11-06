package routes

import (
	registerRoutes "algebra-isosofts-api/routes/registers"
	tableComponentRoutes "algebra-isosofts-api/routes/tableComponents"

	"github.com/gin-gonic/gin"
)

func APIRoutes(rg *gin.RouterGroup) {
	registerRoutes.MainRoutes(rg.Group("/register"))
	tableComponentRoutes.MainRoutes(rg.Group("/tablecomponent"))
}
