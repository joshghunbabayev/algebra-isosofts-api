package registerComponentRoutes

import (
	registerComponentHandlers "algebra-isosofts-api/handlers/registers/components"

	"github.com/gin-gonic/gin"
)

func VendorFeedbackRoutes(rg *gin.RouterGroup) {
	var vendorFeedbackHandler registerComponentHandlers.VendorFeedbackHandler
	rg.GET("/all", vendorFeedbackHandler.GetAll) // query: status, registerId
	rg.POST("/one", vendorFeedbackHandler.Create)
	rg.PUT("/one/:id", vendorFeedbackHandler.Update)
	rg.PUT("/all/delete", vendorFeedbackHandler.Delete)
	rg.PUT("/all/undelete", vendorFeedbackHandler.Undelete)
}
