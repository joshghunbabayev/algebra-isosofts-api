package registerComponentRoutes

import (
	"github.com/gin-gonic/gin"
)

func MainComponentRoutes(rg *gin.RouterGroup) {
	ActionRoutes(rg.Group("/action"))
	VendorFeedbackRoutes(rg.Group("/vendorFeedback"))
}
