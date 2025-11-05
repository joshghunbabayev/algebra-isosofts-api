package routes

import (
	tableComponentRoutes "algebra-isosofts-api/routes/tableComponents"

	"github.com/gin-gonic/gin"
)

func APIRoutes(rg *gin.RouterGroup) {
	tableComponentRoutes.MainRoutes(rg.Group("/tablecomponent"))
}
