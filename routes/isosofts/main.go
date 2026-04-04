package isosoftsRoutes

import (
	"github.com/gin-gonic/gin"
)

func MainRoutes(rg *gin.RouterGroup) {
	QhseKPIRoutes(rg.Group("/qhseKpi"))
}
