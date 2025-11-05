package tableComponentRoutes

import "github.com/gin-gonic/gin"

func MainRoutes(rg *gin.RouterGroup) {
	DropDownListItemRoutes(rg.Group("/dropdownlistitem"))
}
