package isosoftsRoutes

import (
	isosoftsHandlers "algebra-isosofts-api/handlers/isosofts"

	"github.com/gin-gonic/gin"
)

func DropDownListItemRoutes(rg *gin.RouterGroup) {
	var dropDownListItemHandler isosoftsHandlers.DropDownListItemHandler
	rg.GET("/duplicate-defaults", dropDownListItemHandler.DuplicateDefaults) // query: companyId
}
