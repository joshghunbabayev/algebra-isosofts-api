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
				"dbStatus":       "active",
				"companyId":      account.CompanyId,
				"validityStatus": 1,
			})

			if len(tras) > 0 {
				calculatedValue = int64(float64(len(trasSafeties)) / float64(len(tras)) * 100)
			} else {
				calculatedValue = 0
			}

		case 7: // Documents Review Rate %
			var docModel registerModels.DOCModel

			docs, _ := docModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": account.CompanyId,
			})

			docsActuals, _ := docModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": account.CompanyId,
				"actual":    1,
			})

			if len(docs) > 0 {
				calculatedValue = int64(float64(len(docsActuals)) / float64(len(docs)) * 100)
			} else {
				calculatedValue = 0
			}

		case 8: // Vendors Evaluation Rate %
			var venModel registerModels.DOCModel

			vens, _ := venModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": account.CompanyId,
			})

			vensActuals, _ := venModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": account.CompanyId,
				"actual":    1,
			})

			if len(vens) > 0 {
				calculatedValue = int64(float64(len(vensActuals)) / float64(len(vens)) * 100)
			} else {
				calculatedValue = 0
			}

		case 9: // Average Vendors Satisfuction Score
			// calculatedValue = 300
		case 10: // Customer Feedback Evaluation Rate %
			// calculatedValue = 300
		case 11: // Average Customer Satisfuction Score
			// calculatedValue = 300
		case 12: // Number of Jobs
			var fbModel registerModels.FBModel

			fbs, _ := fbModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": account.CompanyId,
			})

			calculatedValue = int64(len(fbs))

		case 13: // Employee Performance Appraisal Status Rate %
			// calculatedValue = 300
		case 14: // Average Employee Skills Appraisal Score
			// calculatedValue = 300
		case 15: // Number of Residual High MoC Risks
			var mocModel registerModels.MOCModel

			mocs, _ := mocModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": account.CompanyId,
			})

			for _, moc := range mocs {
				if moc.ResidualRiskSeverity*moc.ResidualRiskLikelihood >= 12 {
					calculatedValue++
				}
			}

		case 16: // Findings Closure Rate
			// calculatedValue = 300
		case 17: // Number of Non-Conformancies
			// calculatedValue = 300
		case 18: // Number of Opportunities for Improvement
			// calculatedValue = 300
		case 19: // Number of Internal Complaints
			// calculatedValue = 300
		case 20: // Number of External Complaints
			// calculatedValue = 300
		case 21: // Number of Good Practices
			// calculatedValue = 300
		case 22: // Number of Near-Misses
			// calculatedValue = 300
		case 23: // Number of Incidents
			// calculatedValue = 300
		case 24: // Rate of Incidents %
			// calculatedValue = 300
		case 25: // Number of Accident
			// calculatedValue = 300
		case 26: // Number of Unsafe Actions
			// calculatedValue = 300
		case 27: // Number of Unsafe Conditions
			// calculatedValue = 300
		case 28: // Number of Environmental Incidents
			// calculatedValue = 300
		case 29: // Assurance & Oversight Plan Implementation Rate %
			// calculatedValue = 300
		case 30: // Actions Closure Rate
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
