package registerHandlers

import (
	"algebra-isosofts-api/middlewares"
	registerModels "algebra-isosofts-api/models/registers"
	registerTypes "algebra-isosofts-api/types/registers"
	tableComponentTypes "algebra-isosofts-api/types/tableComponents"

	"github.com/gin-gonic/gin"
)

type LEGHandler struct {
}

func (*LEGHandler) GetAll(c *gin.Context) {
	account, _ := c.MustGet("account").(middlewares.RemoteAccount)
	status := c.Query("status")

	if status == "" {
		status = "active"
	}

	var legModel registerModels.LEGModel

	legs, err := legModel.GetAll(map[string]interface{}{
		"dbStatus":  status,
		"companyId": account.CompanyId,
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
		AffectedPositions      string `json:"affectedPositions"`
		ECM                    string `json:"ecm"`
		InitialRiskSeverity    int8   `json:"initialRiskSeverity"`
		InitialRiskLikelihood  int8   `json:"initialRiskLikelihood"`
		ACM                    string `json:"acm"`
		ResidualRiskSeverity   int8   `json:"residualRiskSeverity"`
		ResidualRiskLikelihood int8   `json:"residualRiskLikelihood"`
		Comment                string `json:"comment"`
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

	account, _ := c.MustGet("account").(middlewares.RemoteAccount)

	legModel.Create(registerTypes.LEG{
		Id:        legModel.GenerateUniqueId(),
		CompanyId: account.CompanyId,
		No:        legModel.GenerateUniqueNo(),
		Process: tableComponentTypes.DropDownListItem{
			Id: body.Process,
		},
		Legislation:     body.Legislation,
		Section:         body.Section,
		Requirement:     body.Requirement,
		RiskOfViolation: body.RiskOfViolation,
		AffectedPositions: tableComponentTypes.DropDownListItem{
			Id: body.AffectedPositions,
		},
		ECM:                    body.ECM,
		InitialRiskSeverity:    body.InitialRiskSeverity,
		InitialRiskLikelihood:  body.InitialRiskLikelihood,
		ACM:                    body.ACM,
		ResidualRiskSeverity:   body.ResidualRiskSeverity,
		ResidualRiskLikelihood: body.ResidualRiskLikelihood,
		Comment:                body.Comment,
		DbStatus:               "active",
		DbLastStatus:           "active",
	})

	c.IndentedJSON(201, gin.H{})
}

func (*LEGHandler) Update(c *gin.Context) {
	account, _ := c.MustGet("account").(middlewares.RemoteAccount)
	Id := c.Param("id")

	var legModel registerModels.LEGModel

	currentLeg, _ := legModel.GetById(Id)

	if currentLeg.IsEmpty() || currentLeg.CompanyId != account.CompanyId {
		c.IndentedJSON(404, gin.H{})
		return
	}

	var body struct {
		Process                string `json:"process"`
		Legislation            string `json:"legislation"`
		Section                string `json:"section"`
		Requirement            string `json:"requirement"`
		RiskOfViolation        string `json:"riskOfViolation"`
		AffectedPositions      string `json:"affectedPositions"`
		ECM                    string `json:"ecm"`
		InitialRiskSeverity    int8   `json:"initialRiskSeverity"`
		InitialRiskLikelihood  int8   `json:"initialRiskLikelihood"`
		ACM                    string `json:"acm"`
		ResidualRiskSeverity   int8   `json:"residualRiskSeverity"`
		ResidualRiskLikelihood int8   `json:"residualRiskLikelihood"`
		Comment                string `json:"comment"`
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
		"affectedPositions":      body.AffectedPositions,
		"ecm":                    body.ECM,
		"initialRiskSeverity":    body.InitialRiskSeverity,
		"initialRiskLikelihood":  body.InitialRiskLikelihood,
		"acm":                    body.ACM,
		"residualRiskSeverity":   body.ResidualRiskSeverity,
		"residualRiskLikelihood": body.ResidualRiskLikelihood,
		"comment":                body.Comment,
	})

	c.JSON(200, gin.H{})
}

func (*LEGHandler) Archive(c *gin.Context) {
	account, _ := c.MustGet("account").(middlewares.RemoteAccount)
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
		currentLeg, _ := legModel.GetById(Id)
		if currentLeg.IsEmpty() || currentLeg.CompanyId != account.CompanyId {
			continue
		}

		if currentLeg.DbStatus != "active" {
			continue
		}

		legModel.Update(Id, map[string]interface{}{
			"dbStatus":     "archived",
			"dbLastStatus": currentLeg.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*LEGHandler) Unarchive(c *gin.Context) {
	account, _ := c.MustGet("account").(middlewares.RemoteAccount)
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
		currentLeg, _ := legModel.GetById(Id)
		if currentLeg.IsEmpty() || currentLeg.CompanyId != account.CompanyId {
			continue
		}

		if currentLeg.DbStatus != "archived" {
			continue
		}

		legModel.Update(Id, map[string]interface{}{
			"dbStatus":     "active",
			"dbLastStatus": currentLeg.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*LEGHandler) Delete(c *gin.Context) {
	account, _ := c.MustGet("account").(middlewares.RemoteAccount)
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
		currentLeg, _ := legModel.GetById(Id)
		if currentLeg.IsEmpty() || currentLeg.CompanyId != account.CompanyId {
			continue
		}

		if currentLeg.DbStatus == "deleted" {
			continue
		}

		legModel.Update(Id, map[string]interface{}{
			"dbStatus":     "deleted",
			"dbLastStatus": currentLeg.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*LEGHandler) Undelete(c *gin.Context) {
	account, _ := c.MustGet("account").(middlewares.RemoteAccount)
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
		currentLeg, _ := legModel.GetById(Id)
		if currentLeg.IsEmpty() || currentLeg.CompanyId != account.CompanyId {
			continue
		}

		if currentLeg.DbStatus != "deleted" {
			continue
		}

		legModel.Update(Id, map[string]interface{}{
			"dbStatus":     currentLeg.DbLastStatus,
			"dbLastStatus": currentLeg.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}
