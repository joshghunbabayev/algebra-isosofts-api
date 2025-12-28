package routes

import (
	dashboardRoutes "algebra-isosofts-api/routes/dashboards"
	registerRoutes "algebra-isosofts-api/routes/registers"
	tableComponentRoutes "algebra-isosofts-api/routes/tableComponents"

	"github.com/gin-gonic/gin"
)

func APIRoutes(rg *gin.RouterGroup) {
	dashboardRoutes.MainRoutes(rg.Group("/dashboard"))
	registerRoutes.MainRoutes(rg.Group("/register"))
	tableComponentRoutes.MainRoutes(rg.Group("/tablecomponent"))
}
