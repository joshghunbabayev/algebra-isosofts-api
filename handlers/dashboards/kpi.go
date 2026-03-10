package dashboardHandlers

import (
	"algebra-isosofts-api/middlewares"
	dashboardModels "algebra-isosofts-api/models/dashboards"

	"github.com/gin-gonic/gin"
)

type KPIHandler struct {
}

func (*KPIHandler) GetAll(c *gin.Context) {
	c.IndentedJSON(201, gin.H{})
}

func (*KPIHandler) Update(c *gin.Context) {
	account, _ := c.MustGet("account").(middlewares.RemoteAccount)
	id := c.Param("id")

	var body struct {
		Function     string `json:"function"`
		LYKPI        int64  `json:"lykpi"`
		AnnualTarget int64  `json:"annualTarget"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var kpiModel dashboardModels.KPIModel
	currentKpi, err := kpiModel.GetById(id)

	if err != nil || currentKpi.IsEmpty() || currentKpi.CompanyId != account.CompanyId {
		c.JSON(404, gin.H{"error": "KPI not found or access denied"})
		return
	}

	err = kpiModel.Update(id, map[string]interface{}{
		"function":     body.Function,
		"lykpi":        body.LYKPI,
		"annualTarget": body.AnnualTarget,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to update KPI"})
		return
	}

	c.JSON(200, gin.H{"message": "KPI updated successfully"})
}
