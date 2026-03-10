package isosoftsHandlers

import (
	dashboardModels "algebra-isosofts-api/models/dashboards"

	"github.com/gin-gonic/gin"
)

type KPIHandler struct {
}

func (*KPIHandler) DuplicateDefaults(c *gin.Context) {
	companyId := c.Query("companyId")

	if companyId == "" {
		c.IndentedJSON(400, gin.H{})
	}

	var kpiModel dashboardModels.KPIModel
	kpiModel.DuplicateDefaults(companyId)
	c.IndentedJSON(201, gin.H{})
}
