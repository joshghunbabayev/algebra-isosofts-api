package logHandlers

import (
	registerModels "algebra-isosofts-api/models/registers"
	logTypes "algebra-isosofts-api/types/logs"

	"github.com/gin-gonic/gin"
)

type ActionHandler struct {
}

func (*ActionHandler) GetAll(c *gin.Context) {
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
		regs = append(regs, logTypes.Action{
			Id:      br.Id,
			No:      br.No,
			Actions: br.Actions,
		})
	}

	for _, hsr := range hsrs {
		regs = append(regs, logTypes.Action{
			Id:      hsr.Id,
			No:      hsr.No,
			Actions: hsr.Actions,
		})
	}

	for _, leg := range legs {
		regs = append(regs, logTypes.Action{
			Id:      leg.Id,
			No:      leg.No,
			Actions: leg.Actions,
		})
	}

	for _, eai := range eais {
		regs = append(regs, logTypes.Action{
			Id:      eai.Id,
			No:      eai.No,
			Actions: eai.Actions,
		})
	}

	for _, ei := range eis {
		regs = append(regs, logTypes.Action{
			Id:      ei.Id,
			No:      ei.No,
			Actions: ei.Actions,
		})
	}

	for _, tra := range tras {
		regs = append(regs, logTypes.Action{
			Id:      tra.Id,
			No:      tra.No,
			Actions: tra.Actions,
		})
	}

	for _, doc := range docs {
		regs = append(regs, logTypes.Action{
			Id:      doc.Id,
			No:      doc.No,
			Actions: doc.Actions,
		})
	}

	for _, ven := range vens {
		regs = append(regs, logTypes.Action{
			Id:      ven.Id,
			No:      ven.No,
			Actions: ven.Actions,
		})
	}

	for _, cus := range cuss {
		regs = append(regs, logTypes.Action{
			Id:      cus.Id,
			No:      cus.No,
			Actions: cus.Actions,
		})
	}

	for _, ea := range eas {
		regs = append(regs, logTypes.Action{
			Id:      ea.Id,
			No:      ea.No,
			Actions: ea.Actions,
		})
	}

	for _, moc := range mocs {
		regs = append(regs, logTypes.Action{
			Id:      moc.Id,
			No:      moc.No,
			Actions: moc.Actions,
		})
	}

	for _, fin := range fins {
		regs = append(regs, logTypes.Action{
			Id:      fin.Id,
			No:      fin.No,
			Actions: fin.Actions,
		})
	}

	for _, mrm := range mrms {
		regs = append(regs, logTypes.Action{
			Id:      mrm.Id,
			No:      mrm.No,
			Actions: mrm.Actions,
		})
	}

	for _, aop := range aops {
		regs = append(regs, logTypes.Action{
			Id:      aop.Id,
			No:      aop.No,
			Actions: aop.Actions,
		})
	}

	c.IndentedJSON(200, regs)
}
