package dashboardHandlers

import (
	"algebra-isosofts-api/middlewares"
	dashboardModels "algebra-isosofts-api/models/dashboards"
	registerModels "algebra-isosofts-api/models/registers"
	registerComponentModels "algebra-isosofts-api/models/registers/components"
	"algebra-isosofts-api/modules"

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

		case 1: // Number of Not inspected Inventory/Equipment
			var eiModel registerModels.EIModel

			eis, _ := eiModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": account.CompanyId,
			})

			for _, ei := range eis {
				if !modules.IsDateBigger(ei.NVCD) {
					calculatedValue++
				}
			}

		case 2: // Number of Overdue Trainings
			var traModel registerModels.TRAModel

			tras, _ := traModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": account.CompanyId,
			})

			for _, tra := range tras {
				if !modules.IsDateBigger(tra.NCD) {
					calculatedValue++
				}
			}

		case 3: // Number of Not reviewed Documents
			var docModel registerModels.DOCModel

			docs, _ := docModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": account.CompanyId,
			})

			for _, doc := range docs {
				if !modules.IsDateBigger(doc.NextReviewDate) {
					calculatedValue++
				}
			}

		case 4: // Number of Not evaluated Vendors
			var venModel registerModels.VENModel

			vens, _ := venModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": account.CompanyId,
			})

			for _, ven := range vens {
				if !modules.IsDateBigger(ven.NRD) {
					calculatedValue++
				}
			}

		case 5: // Number of Not evaluated Customers
			var cusModel registerModels.CUSModel

			cuss, _ := cusModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": account.CompanyId,
			})

			for _, cus := range cuss {
				if !modules.IsDateBigger(cus.ReviewDate) {
					calculatedValue++
				}
			}

		case 6: // Number of Not evaluated Employees
			var eaModel registerModels.EAModel

			eas, _ := eaModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": account.CompanyId,
			})

			for _, ea := range eas {
				if !modules.IsDateBigger(ea.NextAppraisalDate) {
					calculatedValue++
				}
			}

		case 7: // Number of Open Findings
			var finModel registerModels.FINModel

			fins, _ := finModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": account.CompanyId,
			})

			for _, fin := range fins {
				if fin.FindingStatus.Value == "Open" {
					calculatedValue++
				}
			}

		case 8: // Number of Not completed A&O Activities
			var aopModel registerModels.AOPModel

			aops, _ := aopModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": account.CompanyId,
			})

			for _, aop := range aops {
				if !modules.IsDateBigger(aop.NextAoaDate) {
					calculatedValue++
				}
			}

		case 9: // Number of Overdue Actions
			var actionModel registerComponentModels.ActionModel

			actions, _ := actionModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": account.CompanyId,
			})

			for _, action := range actions {
				if action.VerificationStatus.Value == "Delayed" {
					calculatedValue++
				}
			}

		case 10: // Number of Residual High Business Risks / Opportunity Level
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

		case 11: // Number of Residual High H&S Risks
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

		case 12: // Number of Legal Requirment with Residual High Risk
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

		case 13: // Number of E&A Aspects with High Residual Significance Level
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

		case 14: // Equipment Safety Rate %
			var eiModel registerModels.EIModel

			var num int64 = 0

			eis, _ := eiModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": account.CompanyId,
			})

			for _, ei := range eis {
				if modules.IsDateBigger(ei.NVCD) {
					num++
				}
			}

			if len(eis) > 0 {
				calculatedValue = int64(float64(num) / float64(len(eis)) * 100)
			} else {
				calculatedValue = 0
			}

		case 15: // Training Validity Rate %
			var traModel registerModels.TRAModel

			var num int64 = 0

			tras, _ := traModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": account.CompanyId,
			})

			for _, tra := range tras {
				if modules.IsDateBigger(tra.NCD) {
					num++
				}
			}

			if len(tras) > 0 {
				calculatedValue = int64(float64(num) / float64(len(tras)) * 100)
			} else {
				calculatedValue = 0
			}

		case 16: // Documents Review Rate %
			var docModel registerModels.DOCModel

			var num int64 = 0

			docs, _ := docModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": account.CompanyId,
			})

			for _, doc := range docs {
				if modules.IsDateBigger(doc.NextReviewDate) {
					num++
				}
			}

			if len(docs) > 0 {
				calculatedValue = int64(float64(num) / float64(len(docs)) * 100)
			} else {
				calculatedValue = 0
			}

		case 17: // Vendors Evaluation Rate %
			var venModel registerModels.VENModel

			var num int64 = 0

			vens, _ := venModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": account.CompanyId,
			})

			for _, ven := range vens {
				if modules.IsDateBigger(ven.NRD) {
					num++
				}
			}

			if len(vens) > 0 {
				calculatedValue = int64(float64(num) / float64(len(vens)) * 100)
			} else {
				calculatedValue = 0
			}

		case 18: // Average Vendors Satisfuction Score
			var venModel registerModels.VENModel

			var num int64 = 0

			vens, _ := venModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": account.CompanyId,
			})

			if len(vens) > 0 {
				for _, ven := range vens {
					num += int64(ven.EvaluationDone + ven.QGS + ven.Communication + ven.OTD + ven.Documentation + ven.HS + ven.Environment)
				}

				calculatedValue = int64(float64(num) / float64(len(vens)))
			} else {
				calculatedValue = 0
			}

		case 19: // Customer Feedback Evaluation Rate %
			var cusModel registerModels.CUSModel

			var num int64 = 0

			cuss, _ := cusModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": account.CompanyId,
			})

			for _, cus := range cuss {
				if modules.IsDateBigger(cus.ReviewDate) {
					num++
				}
			}

			if len(cuss) > 0 {
				calculatedValue = int64(float64(num) / float64(len(cuss)) * 100)
			} else {
				calculatedValue = 0
			}

		case 20: // Average Customer Satisfuction Score
			var cusModel registerModels.CUSModel

			var num int64 = 0

			cuss, _ := cusModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": account.CompanyId,
			})

			if len(cuss) > 0 {
				for _, cus := range cuss {
					num += int64(cus.EvaluationDone + cus.QGS + cus.Communication + cus.OTD + cus.Documentation + cus.HS + cus.Environment)
				}

				calculatedValue = int64(float64(num) / float64(len(cuss)))
			} else {
				calculatedValue = 0
			}

		case 21: // Number of Jobs
			var fbModel registerModels.FBModel

			fbs, _ := fbModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": account.CompanyId,
			})

			calculatedValue = int64(len(fbs))

		case 22: // Employee Performance Appraisal Status Rate %
			var eaModel registerModels.EAModel

			var num int64 = 0

			eas, _ := eaModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": account.CompanyId,
			})

			for _, ea := range eas {
				if modules.IsDateBigger(ea.NextAppraisalDate) {
					num++
				}
			}

			if len(eas) > 0 {
				calculatedValue = int64(float64(num) / float64(len(eas)) * 100)
			} else {
				calculatedValue = 0
			}

		case 23: // Average Employee Skills Appraisal Score
			var eaModel registerModels.EAModel

			var num int64 = 0

			eas, _ := eaModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": account.CompanyId,
			})

			if len(eas) > 0 {
				for _, ea := range eas {
					num += int64(ea.JobQuality + ea.LeadershipSkills + ea.ManagementSkills + ea.BehavioralSkills + ea.EffectivenessOfTrainings)
				}

				calculatedValue = int64(float64(num) / float64(len(eas)))
			} else {
				calculatedValue = 0
			}

		case 24: // Number of Residual High MoC Risks
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

		case 25: // Findings Closure Rate
			var finModel registerModels.FINModel

			var num int64 = 0

			fins, _ := finModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": account.CompanyId,
			})

			for _, fin := range fins {
				if fin.FindingStatus.Value == "Closed" {
					num++
				}
			}

			if len(fins) > 0 {
				calculatedValue = int64(float64(num) / float64(len(fins)) * 100)
			} else {
				calculatedValue = 0
			}

		case 26: // Number of Non-Conformancies
			var finModel registerModels.FINModel

			fins, _ := finModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": account.CompanyId,
			})

			for _, fin := range fins {
				if fin.CategoryOfFinding.Value == "Non conformance" {
					calculatedValue++
				}
			}

		case 27: // Number of Opportunities for Improvement
			var finModel registerModels.FINModel

			fins, _ := finModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": account.CompanyId,
			})

			for _, fin := range fins {
				if fin.CategoryOfFinding.Value == "Opportunity for Improvement" {
					calculatedValue++
				}
			}

		case 28: // Number of Internal Complaints
			var finModel registerModels.FINModel

			fins, _ := finModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": account.CompanyId,
			})

			for _, fin := range fins {
				if fin.CategoryOfFinding.Value == "Internal complaint" {
					calculatedValue++
				}
			}

		case 29: // Number of External Complaints
			var finModel registerModels.FINModel

			fins, _ := finModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": account.CompanyId,
			})

			for _, fin := range fins {
				if fin.CategoryOfFinding.Value == "External Complaint" {
					calculatedValue++
				}
			}

		case 30: // Number of Good Practices
			var finModel registerModels.FINModel

			fins, _ := finModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": account.CompanyId,
			})

			for _, fin := range fins {
				if fin.CategoryOfFinding.Value == "Good practice" {
					calculatedValue++
				}
			}

		case 31: // Number of Near-Misses
			var finModel registerModels.FINModel

			fins, _ := finModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": account.CompanyId,
			})

			for _, fin := range fins {
				if fin.CategoryOfFinding.Value == "Near miss" {
					calculatedValue++
				}
			}

		case 32: // Number of Incidents
			var finModel registerModels.FINModel

			fins, _ := finModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": account.CompanyId,
			})

			for _, fin := range fins {
				if fin.CategoryOfFinding.Value == "Incident" {
					calculatedValue++
				}
			}

		case 33: // Rate of Incidents %
			var finModel registerModels.FINModel

			var num int64 = 0

			fins, _ := finModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": account.CompanyId,
			})

			for _, fin := range fins {
				if fin.CategoryOfFinding.Value == "Incident" {
					num++
				}
			}

			if len(fins) > 0 {
				calculatedValue = int64(float64(num) / float64(len(fins)) * 100)
			} else {
				calculatedValue = 0
			}

		case 34: // Number of Accident
			var finModel registerModels.FINModel

			fins, _ := finModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": account.CompanyId,
			})

			for _, fin := range fins {
				if fin.CategoryOfFinding.Value == "Accident" {
					calculatedValue++
				}
			}

		case 35: // Number of Unsafe Actions
			var finModel registerModels.FINModel

			fins, _ := finModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": account.CompanyId,
			})

			for _, fin := range fins {
				if fin.CategoryOfFinding.Value == "Unsafe actions" {
					calculatedValue++
				}
			}

		case 36: // Number of Unsafe Conditions
			var finModel registerModels.FINModel

			fins, _ := finModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": account.CompanyId,
			})

			for _, fin := range fins {
				if fin.CategoryOfFinding.Value == "Unsafe condition" {
					calculatedValue++
				}
			}

		case 37: // Number of Environmental Incidents
			var finModel registerModels.FINModel

			fins, _ := finModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": account.CompanyId,
			})

			for _, fin := range fins {
				if fin.CategoryOfFinding.Value == "Environmental Incident" {
					calculatedValue++
				}
			}

		case 38: // Assurance & Oversight Plan Implementation Rate %
			var aopModel registerModels.AOPModel

			var num int64 = 0

			aops, _ := aopModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": account.CompanyId,
			})

			for _, aop := range aops {
				if modules.IsDateBigger(aop.NextAoaDate) {
					num++
				}
			}

			if len(aops) > 0 {
				calculatedValue = int64(float64(num) / float64(len(aops)) * 100)
			} else {
				calculatedValue = 0
			}

		case 39: // Actions Closure Rate
			var actionModel registerComponentModels.ActionModel

			var num int64 = 0

			actions, _ := actionModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": account.CompanyId,
			})

			for _, action := range actions {
				if action.Status.Value == "100" {
					num++
				}
			}

			if len(actions) > 0 {
				calculatedValue = int64(float64(num) / float64(len(actions)) * 100)
			} else {
				calculatedValue = 0
			}

		default:
			calculatedValue = 0
		}

		kpis[i].ActualKPI = calculatedValue
		kpis[i].April = calculatedValue
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
