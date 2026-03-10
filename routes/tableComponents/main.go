package tableComponentRoutes

import (
	"algebra-isosofts-api/middlewares"

	"github.com/gin-gonic/gin"
)

func MainRoutes(rg *gin.RouterGroup) {
	rg.Use(middlewares.AuthMiddleware())
	DropDownListItemRoutes(rg.Group("/dropdownlistitem"))
}
