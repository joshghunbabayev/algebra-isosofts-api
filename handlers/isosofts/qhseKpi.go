package isosoftsHandlers

import (
	dashboardModels "algebra-isosofts-api/models/dashboards"

	"github.com/gin-gonic/gin"
)

type QhseKPIHandler struct {
}

func (*QhseKPIHandler) DuplicateDefaults(c *gin.Context) {
	companyId := c.Query("companyId")

	if companyId == "" {
		c.IndentedJSON(400, gin.H{})
	}

	var qhseKPIModel dashboardModels.QhseKPIModel
	qhseKPIModel.DuplicateDefaults(companyId)
	c.IndentedJSON(201, gin.H{})
}
