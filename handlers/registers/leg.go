package registerHandlers

import (
	registerModels "algebra-isosofts-api/models/registers"
	registerTypes "algebra-isosofts-api/types/registers"
	tableComponentTypes "algebra-isosofts-api/types/tableComponents"

	"github.com/gin-gonic/gin"
)

type LEGHandler struct {
}

func (*LEGHandler) GetAll(c *gin.Context) {
	status := c.Query("status")

	if status == "" {
		status = "active"
	}

	var legModel registerModels.LEGModel

	legs, err := legModel.GetAll(map[string]interface{}{
		"dbStatus": status,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, legs)
}

func (*LEGHandler) Create(c *gin.Context) {
	var body struct {
		Process                string `json:"process"`
		Legislation            string `json:"legislation"`
		Section                string `json:"section"`
		Requirement            string `json:"requirement"`
		RiskOfViolation        string `json:"riskOfViolation"`
		InitialRiskSeverity    int8   `json:"initialRiskSeverity"`
		InitialRiskLikelyhood  int8   `json:"initialRiskLikelyhood"`
		ResidualRiskSeverity   int8   `json:"residualRiskSeverity"`
		ResidualRiskLikelyhood int8   `json:"residualRiskLikelyhood"`
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

	var legModel registerModels.LEGModel

	legModel.Create(registerTypes.LEG{
		Id: legModel.GenerateUniqueId(),
		No: legModel.GenerateUniqueNo(),
		Process: tableComponentTypes.DropDownListItem{
			Id: body.Process,
		},
		Legislation:            body.Legislation,
		Section:                body.Section,
		Requirement:            body.Requirement,
		RiskOfViolation:        body.RiskOfViolation,
		InitialRiskSeverity:    body.InitialRiskSeverity,
		InitialRiskLikelyhood:  body.InitialRiskLikelyhood,
		ResidualRiskSeverity:   body.ResidualRiskSeverity,
		ResidualRiskLikelyhood: body.ResidualRiskLikelyhood,
		DbStatus:               "active",
		DbLastStatus:           "active",
	})

	c.IndentedJSON(201, gin.H{})
}

func (*LEGHandler) Update(c *gin.Context) {
	Id := c.Param("id")

	var legModel registerModels.LEGModel

	currentLEG, _ := legModel.GetById(Id)

	if currentLEG.IsEmpty() {
		c.IndentedJSON(404, gin.H{})
		return
	}

	var body struct {
		Process                string `json:"process"`
		Legislation            string `json:"legislation"`
		Section                string `json:"section"`
		Requirement            string `json:"requirement"`
		RiskOfViolation        string `json:"riskOfViolation"`
		InitialRiskSeverity    int8   `json:"initialRiskSeverity"`
		InitialRiskLikelyhood  int8   `json:"initialRiskLikelyhood"`
		ResidualRiskSeverity   int8   `json:"residualRiskSeverity"`
		ResidualRiskLikelyhood int8   `json:"residualRiskLikelyhood"`
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

	legModel.Update(Id, map[string]interface{}{
		"process":                body.Process,
		"legislation":            body.Legislation,
		"section":                body.Section,
		"requirement":            body.Requirement,
		"riskOfViolation":        body.RiskOfViolation,
		"initialRiskSeverity":    body.InitialRiskSeverity,
		"initialRiskLikelyhood":  body.InitialRiskLikelyhood,
		"residualRiskSeverity":   body.ResidualRiskSeverity,
		"residualRiskLikelyhood": body.ResidualRiskLikelyhood,
	})

	c.JSON(200, gin.H{})
}

func (*LEGHandler) Archive(c *gin.Context) {
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

	var legModel registerModels.LEGModel

	for _, Id := range body.Ids {
		currentLEG, _ := legModel.GetById(Id)
		if currentLEG.IsEmpty() {
			continue
		}

		if currentLEG.DbStatus != "active" {
			continue
		}

		legModel.Update(Id, map[string]interface{}{
			"dbStatus":     "archived",
			"dbLastStatus": currentLEG.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*LEGHandler) Unarchive(c *gin.Context) {
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

	var legModel registerModels.LEGModel

	for _, Id := range body.Ids {
		currentLEG, _ := legModel.GetById(Id)
		if currentLEG.IsEmpty() {
			continue
		}

		if currentLEG.DbStatus != "archived" {
			continue
		}

		legModel.Update(Id, map[string]interface{}{
			"dbStatus":     "active",
			"dbLastStatus": currentLEG.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*LEGHandler) Delete(c *gin.Context) {
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

	var legModel registerModels.LEGModel

	for _, Id := range body.Ids {
		currentLEG, _ := legModel.GetById(Id)
		if currentLEG.IsEmpty() {
			continue
		}

		if currentLEG.DbStatus == "deleted" {
			continue
		}

		legModel.Update(Id, map[string]interface{}{
			"dbStatus":     "deleted",
			"dbLastStatus": currentLEG.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*LEGHandler) Undelete(c *gin.Context) {
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

	var legModel registerModels.LEGModel

	for _, Id := range body.Ids {
		currentLEG, _ := legModel.GetById(Id)
		if currentLEG.IsEmpty() {
			continue
		}

		if currentLEG.DbStatus != "deleted" {
			continue
		}

		legModel.Update(Id, map[string]interface{}{
			"dbStatus":     currentLEG.DbLastStatus,
			"dbLastStatus": currentLEG.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}
