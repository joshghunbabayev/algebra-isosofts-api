package tableComponentRoutes

import (
	tableComponentHandlers "algebra-isosofts-api/handlers/tableComponents"

	"github.com/gin-gonic/gin"
)

func DropDownListItemRoutes(rg *gin.RouterGroup) {
	var dropDownListItemHandler tableComponentHandlers.DropDownListItemHandler
	rg.GET("", dropDownListItemHandler.GetAll)
}
