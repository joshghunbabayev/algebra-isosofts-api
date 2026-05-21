package tableComponentRoutes

import (
	tableComponentHandlers "algebra-isosofts-api/handlers/tableComponents"
	"algebra-isosofts-api/middlewares"

	"github.com/gin-gonic/gin"
)

func DropDownListItemRoutes(rg *gin.RouterGroup) {
	var dropDownListItemHandler tableComponentHandlers.DropDownListItemHandler
	rg.Use(middlewares.AccessMiddleware())

	rg.GET("", dropDownListItemHandler.GetAll)
	rg.POST("", dropDownListItemHandler.Create)
	rg.PUT("/:id", dropDownListItemHandler.Update)
	rg.DELETE("/:id", dropDownListItemHandler.Delete)
}
