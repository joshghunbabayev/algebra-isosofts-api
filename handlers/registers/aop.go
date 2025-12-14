package registerHandlers

import (
	registerModels "algebra-isosofts-api/models/registers"
	registerTypes "algebra-isosofts-api/types/registers"
	tableComponentTypes "algebra-isosofts-api/types/tableComponents"

	"github.com/gin-gonic/gin"
)

type AOPHandler struct {
}

func (*AOPHandler) GetAll(c *gin.Context) {
	status := c.Query("status")

	if status == "" {
		status = "active"
	}

	var aopModel registerModels.AOPModel

	aops, err := aopModel.GetAll(map[string]interface{}{
		"dbStatus": status,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, aops)
}

func (*AOPHandler) Create(c *gin.Context) {
	var body struct {
		ActivityDescription string `json:"activityDescription"`
		AuditorInspector    string `json:"auditorInspector"`
		AuditeeInspectee    string `json:"auditeeInspectee"`
		ReviewedPremises    string `json:"reviewedPremises"`
		ReviewedProcess     string `json:"reviewedProcess"`
		RTIC                string `json:"rtic"`
		Frequency           string `json:"frequency"`
		AuditDate           string `json:"auditDate"`
		InspectionFrequency string `json:"inspectionFrequency"`
		NextAuditDate       string `json:"nextAuditDate"`
		AuditStatus         string `json:"auditStatus"`
	}

	var errs = make(map[string]interface{})

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if len(errs) > 0 {
		c.JSON(400, gin.H{
			"token": "aaa",
			"errs":  errs,
		})
		return
	}

	var aopModel registerModels.AOPModel

	aopModel.Create(registerTypes.AOP{
		Id: aopModel.GenerateUniqueId(),
		No: aopModel.GenerateUniqueNo(),
		ActivityDescription: tableComponentTypes.DropDownListItem{
			Id: body.ActivityDescription,
		},
		AuditorInspector: body.AuditorInspector,
		AuditeeInspectee: body.AuditeeInspectee,
		ReviewedPremises: tableComponentTypes.DropDownListItem{
			Id: body.ReviewedPremises,
		},
		ReviewedProcess: tableComponentTypes.DropDownListItem{
			Id: body.ReviewedProcess,
		},
		RTIC:      body.RTIC,
		Frequency: body.Frequency,
		AuditDate: body.AuditDate,
		InspectionFrequency: tableComponentTypes.DropDownListItem{
			Id: body.InspectionFrequency,
		},
		NextAuditDate: body.NextAuditDate,
		AuditStatus:   body.AuditStatus,
		DbStatus:      "active",
		DbLastStatus:  "active",
	})

	c.IndentedJSON(201, gin.H{})
}

func (*AOPHandler) Update(c *gin.Context) {
	Id := c.Param("id")

	var aopModel registerModels.AOPModel

	currentAop, _ := aopModel.GetById(Id)

	if currentAop.IsEmpty() {
		c.IndentedJSON(404, gin.H{})
		return
	}

	var body struct {
		ActivityDescription string `json:"activityDescription"`
		AuditorInspector    string `json:"auditorInspector"`
		AuditeeInspectee    string `json:"auditeeInspectee"`
		ReviewedPremises    string `json:"reviewedPremises"`
		ReviewedProcess     string `json:"reviewedProcess"`
		RTIC                string `json:"rtic"`
		Frequency           string `json:"frequency"`
		AuditDate           string `json:"auditDate"`
		InspectionFrequency string `json:"inspectionFrequency"`
		NextAuditDate       string `json:"nextAuditDate"`
		AuditStatus         string `json:"auditStatus"`
	}

	var errs = make(map[string]interface{})

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if len(errs) > 0 {
		c.JSON(400, gin.H{
			"token": "aaa",
			"errs":  errs,
		})
		return
	}

	aopModel.Update(Id, map[string]interface{}{
		"activityDescription": body.ActivityDescription,
		"auditorInspector":    body.AuditorInspector,
		"auditeeInspectee":    body.AuditeeInspectee,
		"reviewedPremises":    body.ReviewedPremises,
		"reviewedProcess":     body.ReviewedProcess,
		"rtic":                body.RTIC,
		"frequency":           body.Frequency,
		"auditDate":           body.AuditDate,
		"inspectionFrequency": body.InspectionFrequency,
		"nextAuditDate":       body.NextAuditDate,
		"auditStatus":         body.AuditStatus,
	})

	c.JSON(200, gin.H{})
}

func (*AOPHandler) Archive(c *gin.Context) {
	var body struct {
		Ids []string `json:"ids"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{})
		return
	}

	if len(body.Ids) == 0 {
		c.JSON(404, gin.H{})
		return
	}

	var aopModel registerModels.AOPModel

	for _, Id := range body.Ids {
		currentAop, _ := aopModel.GetById(Id)
		if currentAop.IsEmpty() {
			continue
		}

		if currentAop.DbStatus != "active" {
			continue
		}

		aopModel.Update(Id, map[string]interface{}{
			"dbStatus":     "archived",
			"dbLastStatus": currentAop.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*AOPHandler) Unarchive(c *gin.Context) {
	var body struct {
		Ids []string `json:"ids"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{})
		return
	}

	if len(body.Ids) == 0 {
		c.JSON(404, gin.H{})
		return
	}

	var aopModel registerModels.AOPModel

	for _, Id := range body.Ids {
		currentAop, _ := aopModel.GetById(Id)
		if currentAop.IsEmpty() {
			continue
		}

		if currentAop.DbStatus != "archived" {
			continue
		}

		aopModel.Update(Id, map[string]interface{}{
			"dbStatus":     "active",
			"dbLastStatus": currentAop.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*AOPHandler) Delete(c *gin.Context) {
	var body struct {
		Ids []string `json:"ids"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{})
		return
	}

	if len(body.Ids) == 0 {
		c.JSON(404, gin.H{})
		return
	}

	var aopModel registerModels.AOPModel

	for _, Id := range body.Ids {
		currentAop, _ := aopModel.GetById(Id)
		if currentAop.IsEmpty() {
			continue
		}

		if currentAop.DbStatus == "deleted" {
			continue
		}

		aopModel.Update(Id, map[string]interface{}{
			"dbStatus":     "deleted",
			"dbLastStatus": currentAop.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*AOPHandler) Undelete(c *gin.Context) {
	var body struct {
		Ids []string `json:"ids"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{})
		return
	}

	if len(body.Ids) == 0 {
		c.JSON(404, gin.H{})
		return
	}

	var aopModel registerModels.AOPModel

	for _, Id := range body.Ids {
		currentAop, _ := aopModel.GetById(Id)
		if currentAop.IsEmpty() {
			continue
		}

		if currentAop.DbStatus != "deleted" {
			continue
		}

		aopModel.Update(Id, map[string]interface{}{
			"dbStatus":     currentAop.DbLastStatus,
			"dbLastStatus": currentAop.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}
