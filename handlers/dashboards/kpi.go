package dashboardHandlers

import (
	HelperFuncs "algebra-isosofts-api/helper/funcs"
	registerModels "algebra-isosofts-api/models/registers"
	dashboardTypes "algebra-isosofts-api/types/dashboards"
	registerComponentTypes "algebra-isosofts-api/types/registers/components"

	"github.com/gin-gonic/gin"
)

type KPIHandler struct {
}

func (*KPIHandler) GetAll(c *gin.Context) {
	status := c.Query("status")

	if status == "" {
		status = "active"
	}

	var actions []registerComponentTypes.Action

	var brModel registerModels.BRModel
	var hsrModel registerModels.HSRModel
	var legModel registerModels.LEGModel
	var eaiModel registerModels.EAIModel
	var eiModel registerModels.EIModel
	var traModel registerModels.TRAModel
	var docModel registerModels.DOCModel
	var venModel registerModels.VENModel
	var cusModel registerModels.CUSModel
	var eaModel registerModels.EAModel
	var mocModel registerModels.MOCModel
	var finModel registerModels.FINModel
	var mrmModel registerModels.MRMModel
	var aopModel registerModels.AOPModel

	brs, err := brModel.GetAll(map[string]interface{}{
		"dbStatus": status,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	hsrs, err := hsrModel.GetAll(map[string]interface{}{
		"dbStatus": status,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	legs, err := legModel.GetAll(map[string]interface{}{
		"dbStatus": status,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	eais, err := eaiModel.GetAll(map[string]interface{}{
		"dbStatus": status,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	eis, err := eiModel.GetAll(map[string]interface{}{
		"dbStatus": status,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	tras, err := traModel.GetAll(map[string]interface{}{
		"dbStatus": status,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	docs, err := docModel.GetAll(map[string]interface{}{
		"dbStatus": status,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	vens, err := venModel.GetAll(map[string]interface{}{
		"dbStatus": status,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	cuss, err := cusModel.GetAll(map[string]interface{}{
		"dbStatus": status,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	eas, err := eaModel.GetAll(map[string]interface{}{
		"dbStatus": status,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	mocs, err := mocModel.GetAll(map[string]interface{}{
		"dbStatus": status,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	fins, err := finModel.GetAll(map[string]interface{}{
		"dbStatus": status,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	mrms, err := mrmModel.GetAll(map[string]interface{}{
		"dbStatus": status,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	aops, err := aopModel.GetAll(map[string]interface{}{
		"dbStatus": status,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	for _, br := range brs {
		for _, action := range br.Actions {
			actions = append(actions, action)
		}
	}

	for _, hsr := range hsrs {
		for _, action := range hsr.Actions {
			actions = append(actions, action)
		}
	}

	for _, leg := range legs {
		for _, action := range leg.Actions {
			actions = append(actions, action)
		}
	}

	for _, eai := range eais {
		for _, action := range eai.Actions {
			actions = append(actions, action)
		}
	}

	for _, ei := range eis {
		for _, action := range ei.Actions {
			actions = append(actions, action)
		}
	}

	for _, tra := range tras {
		for _, action := range tra.Actions {
			actions = append(actions, action)
		}
	}

	for _, doc := range docs {
		for _, action := range doc.Actions {
			actions = append(actions, action)
		}
	}

	for _, ven := range vens {
		for _, action := range ven.Actions {
			actions = append(actions, action)
		}
	}

	for _, cus := range cuss {
		for _, action := range cus.Actions {
			actions = append(actions, action)
		}
	}

	for _, ea := range eas {
		for _, action := range ea.Actions {
			actions = append(actions, action)
		}
	}

	for _, moc := range mocs {
		for _, action := range moc.Actions {
			actions = append(actions, action)
		}
	}

	for _, fin := range fins {
		for _, action := range fin.Actions {
			actions = append(actions, action)
		}
	}

	for _, mrm := range mrms {
		for _, action := range mrm.Actions {
			actions = append(actions, action)
		}
	}

	for _, aop := range aops {
		for _, action := range aop.Actions {
			actions = append(actions, action)
		}
	}

	var kpi dashboardTypes.KPI

	kpi.TotalActions = int64(len(actions))
	if len(actions) == 0 {
		c.IndentedJSON(200, kpi)
		return
	}

	// init maps
	kpi.Status.Counts = make(map[string]int64)
	kpi.VerificationStatus = make(map[string]int64)
	kpi.Confirmation = make(map[string]int64)

	var statusSum int

	var jan, feb, mar, apr, may, jun, jul, aug, sep, oct, nov, dec int

	for _, action := range actions {

		// ===== STATUS =====
		if !action.Status.IsEmpty() {
			kpi.Status.Counts[action.Status.Value]++
			statusSum += HelperFuncs.ParseInt(action.Status.Value)
		}

		// ===== VERIFICATION STATUS =====
		if !action.VerificationStatus.IsEmpty() {
			kpi.VerificationStatus[action.VerificationStatus.Value]++
		}

		// ===== CONFIRMATION =====
		if !action.Confirmation.IsEmpty() {
			kpi.Confirmation[action.Confirmation.Value]++
		}

		// ===== MONTHLY =====
		jan += HelperFuncs.ParseInt(action.January.Value)
		feb += HelperFuncs.ParseInt(action.February.Value)
		mar += HelperFuncs.ParseInt(action.March.Value)
		apr += HelperFuncs.ParseInt(action.April.Value)
		may += HelperFuncs.ParseInt(action.May.Value)
		jun += HelperFuncs.ParseInt(action.June.Value)
		jul += HelperFuncs.ParseInt(action.July.Value)
		aug += HelperFuncs.ParseInt(action.August.Value)
		sep += HelperFuncs.ParseInt(action.September.Value)
		oct += HelperFuncs.ParseInt(action.October.Value)
		nov += HelperFuncs.ParseInt(action.November.Value)
		dec += HelperFuncs.ParseInt(action.December.Value)
	}

	count := float64(len(actions))

	// ===== AVERAGES =====
	kpi.Status.Average = float64(statusSum) / count

	kpi.MonthlyProgress.January = float64(jan) / count
	kpi.MonthlyProgress.February = float64(feb) / count
	kpi.MonthlyProgress.March = float64(mar) / count
	kpi.MonthlyProgress.April = float64(apr) / count
	kpi.MonthlyProgress.May = float64(may) / count
	kpi.MonthlyProgress.June = float64(jun) / count
	kpi.MonthlyProgress.July = float64(jul) / count
	kpi.MonthlyProgress.August = float64(aug) / count
	kpi.MonthlyProgress.September = float64(sep) / count
	kpi.MonthlyProgress.October = float64(oct) / count
	kpi.MonthlyProgress.November = float64(nov) / count
	kpi.MonthlyProgress.December = float64(dec) / count

	c.IndentedJSON(200, kpi)
}
