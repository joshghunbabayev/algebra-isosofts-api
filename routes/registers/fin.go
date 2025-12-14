package registerRoutes

import (
	registerHandlers "algebra-isosofts-api/handlers/registers"

	"github.com/gin-gonic/gin"
)

func FINRoutes(rg *gin.RouterGroup) {
	var finHandler registerHandlers.FINHandler
	rg.GET("/all", finHandler.GetAll) // query: status
	rg.POST("/one", finHandler.Create)
	rg.PUT("/one/:id", finHandler.Update)
	rg.PUT("/all/archive", finHandler.Archive)
	rg.PUT("/all/unarchive", finHandler.Unarchive)
	rg.PUT("/all/delete", finHandler.Delete)
	rg.PUT("/all/undelete", finHandler.Undelete)
}
