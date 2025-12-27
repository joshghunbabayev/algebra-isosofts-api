package logRoutes

import "github.com/gin-gonic/gin"

func MainRoutes(rg *gin.RouterGroup) {
	ActionRoutes(rg.Group("/action"))
}
