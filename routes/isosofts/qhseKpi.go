package isosoftsRoutes

import (
	isosoftsHandlers "algebra-isosofts-api/handlers/isosofts"

	"github.com/gin-gonic/gin"
)

func KPIRoutes(rg *gin.RouterGroup) {
	var kpiHandler isosoftsHandlers.KPIHandler
	rg.GET("/duplicate-defaults", kpiHandler.DuplicateDefaults) // query: companyId
}
