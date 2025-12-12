package registerRoutes

import (
	registerComponentRoutes "algebra-isosofts-api/routes/registers/components"

	"github.com/gin-gonic/gin"
)

func MainRoutes(rg *gin.RouterGroup) {
	BRRoutes(rg.Group("/br"))
	HSRRoutes(rg.Group("/hsr"))
	LEGRoutes(rg.Group("/leg"))
	EIARoutes(rg.Group("/eia"))
	EIRoutes(rg.Group("/ei"))
	TRARoutes(rg.Group("/tra"))
	EARoutes(rg.Group("/ea"))
	MRMRoutes(rg.Group("/mrm"))
	registerComponentRoutes.MainComponentRoutes(rg.Group("/component"))
}
