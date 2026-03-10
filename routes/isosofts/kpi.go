package isosoftsRoutes

import (
	isosoftsHandlers "algebra-isosofts-api/handlers/isosofts"

	"github.com/gin-gonic/gin"
)

func KPIRoutes(rg *gin.RouterGroup) {
	var isosoftsHandler isosoftsHandlers.KPIHandler
	rg.GET("/duplicate-defaults", isosoftsHandler.DuplicateDefaults) // query: companyId
}
