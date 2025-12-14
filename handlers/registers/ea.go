package registerHandlers

import (
	registerModels "algebra-isosofts-api/models/registers"
	registerTypes "algebra-isosofts-api/types/registers"

	"github.com/gin-gonic/gin"
)

type EAHandler struct {
}

func (*EAHandler) GetAll(c *gin.Context) {
	status := c.Query("status")

	if status == "" {
		status = "active"
	}

	var eaModel registerModels.EAModel

	eas, err := eaModel.GetAll(map[string]interface{}{
		"dbStatus": status,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, eas)
}

func (*EAHandler) Create(c *gin.Context) {
	var body struct {
		EmployeeName    string `json:"employeeName"`
		Position        string `json:"position"`
		LineManager     string `json:"lineManager"`
		ESD             string `json:"esd"`
		AppraisalDate   string `json:"appraisalDate"`
		AppraisalType   string `json:"appraisalType"`
		TCA             string `json:"tca"`
		SkillsAppraisal string `json:"skillsAppraisal"`
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

	var eaModel registerModels.EAModel

	eaModel.Create(registerTypes.EA{
		Id:              eaModel.GenerateUniqueId(),
		No:              eaModel.GenerateUniqueNo(),
		EmployeeName:    body.EmployeeName,
		Position:        body.Position,
		LineManager:     body.LineManager,
		ESD:             body.ESD,
		AppraisalDate:   body.AppraisalDate,
		AppraisalType:   body.AppraisalType,
		TCA:             body.TCA,
		SkillsAppraisal: body.SkillsAppraisal,
		DbStatus:        "active",
		DbLastStatus:    "active",
	})

	c.IndentedJSON(201, gin.H{})
}

func (*EAHandler) Update(c *gin.Context) {
	Id := c.Param("id")

	var eaModel registerModels.EAModel

	currentEa, _ := eaModel.GetById(Id)

	if currentEa.IsEmpty() {
		c.IndentedJSON(404, gin.H{})
		return
	}

	var body struct {
		EmployeeName    string `json:"employeeName"`
		Position        string `json:"position"`
		LineManager     string `json:"lineManager"`
		ESD             string `json:"esd"`
		AppraisalDate   string `json:"appraisalDate"`
		AppraisalType   string `json:"appraisalType"`
		TCA             string `json:"tca"`
		SkillsAppraisal string `json:"skillsAppraisal"`
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

	eaModel.Update(Id, map[string]interface{}{
		"employeeName":    body.EmployeeName,
		"position":        body.Position,
		"lineManager":     body.LineManager,
		"esd":             body.ESD,
		"appraisalDate":   body.AppraisalDate,
		"appraisalType":   body.AppraisalType,
		"tca":             body.TCA,
		"skillsAppraisal": body.SkillsAppraisal,
	})

	c.JSON(200, gin.H{})
}

func (*EAHandler) Archive(c *gin.Context) {
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

	var eaModel registerModels.EAModel

	for _, Id := range body.Ids {
		currentEa, _ := eaModel.GetById(Id)
		if currentEa.IsEmpty() {
			continue
		}

		if currentEa.DbStatus != "active" {
			continue
		}

		eaModel.Update(Id, map[string]interface{}{
			"dbStatus":     "archived",
			"dbLastStatus": currentEa.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*EAHandler) Unarchive(c *gin.Context) {
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

	var eaModel registerModels.EAModel

	for _, Id := range body.Ids {
		currentEa, _ := eaModel.GetById(Id)
		if currentEa.IsEmpty() {
			continue
		}

		if currentEa.DbStatus != "archived" {
			continue
		}

		eaModel.Update(Id, map[string]interface{}{
			"dbStatus":     "active",
			"dbLastStatus": currentEa.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*EAHandler) Delete(c *gin.Context) {
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

	var eaModel registerModels.EAModel

	for _, Id := range body.Ids {
		currentEa, _ := eaModel.GetById(Id)
		if currentEa.IsEmpty() {
			continue
		}

		if currentEa.DbStatus == "deleted" {
			continue
		}

		eaModel.Update(Id, map[string]interface{}{
			"dbStatus":     "deleted",
			"dbLastStatus": currentEa.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*EAHandler) Undelete(c *gin.Context) {
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

	var eaModel registerModels.EAModel

	for _, Id := range body.Ids {
		currentEa, _ := eaModel.GetById(Id)
		if currentEa.IsEmpty() {
			continue
		}

		if currentEa.DbStatus != "deleted" {
			continue
		}

		eaModel.Update(Id, map[string]interface{}{
			"dbStatus":     currentEa.DbLastStatus,
			"dbLastStatus": currentEa.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}
