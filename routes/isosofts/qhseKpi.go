package isosoftsRoutes

import (
	isosoftsHandlers "algebra-isosofts-api/handlers/isosofts"

	"github.com/gin-gonic/gin"
)

func QhseKPIRoutes(rg *gin.RouterGroup) {
	var qhseKPIHandler isosoftsHandlers.QhseKPIHandler
	rg.GET("/duplicate-defaults", qhseKPIHandler.DuplicateDefaults) // query: companyId
}
