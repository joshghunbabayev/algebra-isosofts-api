package dashboardHandlers

import (
	"algebra-isosofts-api/middlewares"
	dashboardModels "algebra-isosofts-api/models/dashboards"
	registerModels "algebra-isosofts-api/models/registers"

	"github.com/gin-gonic/gin"
)

type KPIHandler struct {
}

func (*KPIHandler) GetAll(c *gin.Context) {
	account, _ := c.MustGet("account").(middlewares.RemoteAccount)

	var kpiModel dashboardModels.KPIModel

	kpis, err := kpiModel.GetAll(map[string]interface{}{
		"companyId": account.CompanyId,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": "Məlumatlar gətirilərkən xəta baş verdi"})
		return
	}

	for i := range kpis {

		var calculatedValue int64 = 0

		switch kpis[i].SNo {
		case 1: // Number of Residual High Business Risks / Opportunity Level
			var brModel registerModels.BRModel

			brs, _ := brModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": account.CompanyId,
			})

			for _, br := range brs {
				if br.ResidualRiskSeverity*br.ResidualRiskLikelihood >= 12 {
					calculatedValue++
				}
			}
		case 2: // Number of Residual High H&S Risks
			var hsrModel registerModels.HSRModel

			hsrs, _ := hsrModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": account.CompanyId,
			})

			for _, hsr := range hsrs {
				if hsr.ResidualRiskSeverity*hsr.ResidualRiskLikelihood >= 12 {
					calculatedValue++
				}
			}
		case 3: // Number of Legal Requirment with Residual High Risk
			var legModel registerModels.LEGModel

			legs, _ := legModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": account.CompanyId,
			})

			for _, leg := range legs {
				if leg.ResidualRiskSeverity*leg.ResidualRiskLikelihood >= 12 {
					calculatedValue++
				}
			}
		case 4: // Number of E&A Aspects with High Residual Significance Level
			var eaiModel registerModels.EAIModel

			eais, _ := eaiModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": account.CompanyId,
			})

			for _, eai := range eais {
				if eai.RDOSProbability*eai.RDOSSeverity*eai.RDOSDuration*eai.RDOSScale >= 80 {
					calculatedValue++
				}
			}
		case 5: // Equipment Safety Rate %
			var eiModel registerModels.EIModel

			eis, _ := eiModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": account.CompanyId,
			})

			eisSafeties, _ := eiModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": account.CompanyId,
				"eis":       1,
			})

			if len(eis) > 0 {
				calculatedValue = int64(float64(len(eisSafeties)) / float64(len(eis)) * 100)
			} else {
				calculatedValue = 0
			}
		case 6: // Training Validity Rate %
			var traModel registerModels.TRAModel

			tras, _ := traModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": account.CompanyId,
			})

			trasSafeties, _ := traModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": account.CompanyId,
				"tras":      1,
			})

			if len(tras) > 0 {
				calculatedValue = int64(float64(len(trasSafeties)) / float64(len(tras)) * 100)
			} else {
				calculatedValue = 0
			}
		case 7:
			// calculatedValue = 300
		case 8:
			// calculatedValue = 300
		case 9:
			// calculatedValue = 300
		case 10:
			// calculatedValue = 300
		default:
			calculatedValue = 0
		}

		kpis[i].ActualKPI = calculatedValue
		kpis[i].March = calculatedValue
	}

	c.IndentedJSON(200, kpis)
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
