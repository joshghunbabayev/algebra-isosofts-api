package registerRoutes

import (
	"algebra-isosofts-api/middlewares"
	registerComponentRoutes "algebra-isosofts-api/routes/registers/components"

	"github.com/gin-gonic/gin"
)

func MainRoutes(rg *gin.RouterGroup) {
	// Bütün sorğular üçün Auth yoxlanılır
	rg.Use(middlewares.AuthMiddleware())

	// Komponentlər üçün AccessMiddleware lazım deyil (adətən oxumaq üçündür)
	registerComponentRoutes.MainComponentRoutes(rg.Group("/component"))

	// İcazə (Access) yoxlanışı tələb olunan registerlər
	rrg := rg.Group("")
	rrg.Use(middlewares.AccessMiddleware())
	{
		// DİQQƏT: rg.Group yox, rrg.Group istifadə olunmalıdır!
		BRRoutes(rrg.Group("/br"))
		HSRRoutes(rrg.Group("/hsr"))
		LEGRoutes(rrg.Group("/leg"))
		EAIRoutes(rrg.Group("/eai"))
		EIRoutes(rrg.Group("/ei"))
		TRARoutes(rrg.Group("/tra"))
		DOCRoutes(rrg.Group("/doc"))
		VENRoutes(rrg.Group("/ven"))
		CUSRoutes(rrg.Group("/cus"))
		FBRoutes(rrg.Group("/fb"))
		EARoutes(rrg.Group("/ea"))
		MOCRoutes(rrg.Group("/moc"))
		FINRoutes(rrg.Group("/fin"))
		MRMRoutes(rrg.Group("/mrm"))
		AOPRoutes(rrg.Group("/aop"))
	}
}
