package isosoftsRoutes

import (
	"github.com/gin-gonic/gin"
)

func MainRoutes(rg *gin.RouterGroup) {
	KPIRoutes(rg.Group("/kpi"))
	DropDownListItemRoutes(rg.Group("/dropdownlistitem"))
}
