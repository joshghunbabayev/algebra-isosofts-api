package registerRoutes

import (
	registerComponentRoutes "algebra-isosofts-api/routes/registers/components"

	"github.com/gin-gonic/gin"
)

func MainRoutes(rg *gin.RouterGroup) {
	BrRoutes(rg.Group("/br"))
	registerComponentRoutes.ActionRoutes(rg.Group("/action"))
}
