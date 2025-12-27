package routes

import (
	logRoutes "algebra-isosofts-api/routes/logs"
	registerRoutes "algebra-isosofts-api/routes/registers"
	tableComponentRoutes "algebra-isosofts-api/routes/tableComponents"

	"github.com/gin-gonic/gin"
)

func APIRoutes(rg *gin.RouterGroup) {
	registerRoutes.MainRoutes(rg.Group("/register"))
	logRoutes.MainRoutes(rg.Group("/log"))
	tableComponentRoutes.MainRoutes(rg.Group("/tablecomponent"))
}
