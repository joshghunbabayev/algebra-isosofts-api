package registerRoutes

import (
	registerComponentRoutes "algebra-isosofts-api/routes/registers/components"

	"github.com/gin-gonic/gin"
)

func MainRoutes(rg *gin.RouterGroup) {
	BRRoutes(rg.Group("/br"))
	HSRRoutes(rg.Group("/hsr"))
	LEGRoutes(rg.Group("/leg"))
	EAIRoutes(rg.Group("/eai"))
	EIRoutes(rg.Group("/ei"))
	TRARoutes(rg.Group("/tra"))
	DOCRoutes(rg.Group("/doc"))
	EARoutes(rg.Group("/ea"))
	MOCRoutes(rg.Group("/moc"))
	FINRoutes(rg.Group("/fin"))
	MRMRoutes(rg.Group("/mrm"))
	AOPRoutes(rg.Group("/aop"))
	registerComponentRoutes.MainComponentRoutes(rg.Group("/component"))
}
