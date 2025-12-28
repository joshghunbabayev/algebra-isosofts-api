package dashboardHandlers

import (
	registerModels "algebra-isosofts-api/models/registers"
	dashboardTypes "algebra-isosofts-api/types/dashboards"

	"github.com/gin-gonic/gin"
)

type ActionLogHandler struct {
}

func (*ActionLogHandler) GetAll(c *gin.Context) {
	status := c.Query("status")

	if status == "" {
		status = "active"
	}

	var regs []interface{}

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
		if br.Actions != nil {
			regs = append(regs, dashboardTypes.ActionLog{
				Id:      br.Id,
				No:      br.No,
				Actions: br.Actions,
			})
		}
	}

	for _, hsr := range hsrs {
		if hsr.Actions != nil {
			regs = append(regs, dashboardTypes.ActionLog{
				Id:      hsr.Id,
				No:      hsr.No,
				Actions: hsr.Actions,
			})
		}
	}

	for _, leg := range legs {
		if leg.Actions != nil {
			regs = append(regs, dashboardTypes.ActionLog{
				Id:      leg.Id,
				No:      leg.No,
				Actions: leg.Actions,
			})
		}
	}

	for _, eai := range eais {
		if eai.Actions != nil {
			regs = append(regs, dashboardTypes.ActionLog{
				Id:      eai.Id,
				No:      eai.No,
				Actions: eai.Actions,
			})
		}
	}

	for _, ei := range eis {
		if ei.Actions != nil {
			regs = append(regs, dashboardTypes.ActionLog{
				Id:      ei.Id,
				No:      ei.No,
				Actions: ei.Actions,
			})
		}
	}

	for _, tra := range tras {
		if tra.Actions != nil {
			regs = append(regs, dashboardTypes.ActionLog{
				Id:      tra.Id,
				No:      tra.No,
				Actions: tra.Actions,
			})
		}
	}

	for _, doc := range docs {
		if doc.Actions != nil {
			regs = append(regs, dashboardTypes.ActionLog{
				Id:      doc.Id,
				No:      doc.No,
				Actions: doc.Actions,
			})
		}
	}

	for _, ven := range vens {
		if ven.Actions != nil {
			regs = append(regs, dashboardTypes.ActionLog{
				Id:      ven.Id,
				No:      ven.No,
				Actions: ven.Actions,
			})
		}
	}

	for _, cus := range cuss {
		if cus.Actions != nil {
			regs = append(regs, dashboardTypes.ActionLog{
				Id:      cus.Id,
				No:      cus.No,
				Actions: cus.Actions,
			})
		}
	}

	for _, ea := range eas {
		if ea.Actions != nil {
			regs = append(regs, dashboardTypes.ActionLog{
				Id:      ea.Id,
				No:      ea.No,
				Actions: ea.Actions,
			})
		}
	}

	for _, moc := range mocs {
		if moc.Actions != nil {
			regs = append(regs, dashboardTypes.ActionLog{
				Id:      moc.Id,
				No:      moc.No,
				Actions: moc.Actions,
			})
		}
	}

	for _, fin := range fins {
		if fin.Actions != nil {
			regs = append(regs, dashboardTypes.ActionLog{
				Id:      fin.Id,
				No:      fin.No,
				Actions: fin.Actions,
			})
		}
	}

	for _, mrm := range mrms {
		if mrm.Actions != nil {
			regs = append(regs, dashboardTypes.ActionLog{
				Id:      mrm.Id,
				No:      mrm.No,
				Actions: mrm.Actions,
			})
		}
	}

	for _, aop := range aops {
		if aop.Actions != nil {
			regs = append(regs, dashboardTypes.ActionLog{
				Id:      aop.Id,
				No:      aop.No,
				Actions: aop.Actions,
			})
		}
	}

	c.IndentedJSON(200, regs)
}
