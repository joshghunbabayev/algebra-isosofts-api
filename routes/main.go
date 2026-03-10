package routes

import (
	dashboardRoutes "algebra-isosofts-api/routes/dashboards"
	isosoftsRoutes "algebra-isosofts-api/routes/isosofts"
	registerRoutes "algebra-isosofts-api/routes/registers"
	tableComponentRoutes "algebra-isosofts-api/routes/tableComponents"

	"github.com/gin-gonic/gin"
)

func APIRoutes(rg *gin.RouterGroup) {
	isosoftsRoutes.MainRoutes(rg.Group("/isosofts"))
	dashboardRoutes.MainRoutes(rg.Group("/dashboard"))
	registerRoutes.MainRoutes(rg.Group("/register"))
	tableComponentRoutes.MainRoutes(rg.Group("/tablecomponent"))
}
