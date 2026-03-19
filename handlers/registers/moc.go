package registerHandlers

import (
	"algebra-isosofts-api/middlewares"
	registerModels "algebra-isosofts-api/models/registers"
	registerTypes "algebra-isosofts-api/types/registers"
	tableComponentTypes "algebra-isosofts-api/types/tableComponents"

	"github.com/gin-gonic/gin"
)

type MOCHandler struct {
}

func (*MOCHandler) GetAll(c *gin.Context) {
	account, _ := c.MustGet("account").(middlewares.RemoteAccount)
	status := c.Query("status")

	if status == "" {
		status = "active"
	}

	var mocModel registerModels.MOCModel

	mocs, err := mocModel.GetAll(map[string]interface{}{
		"dbStatus":  status,
		"companyId": account.CompanyId,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, mocs)
}

func (*MOCHandler) Create(c *gin.Context) {
	account, _ := c.MustGet("account").(middlewares.RemoteAccount)
	var body struct {
		Issuer                 string `json:"issuer"`
		ReasonOfChange         string `json:"reasonOfChange"`
		Process                string `json:"process"`
		Action                 string `json:"action"`
		Risks                  string `json:"risks"`
		InitialRiskSeverity    int8   `json:"initialRiskSeverity"`
		InitialRiskLikelihood  int8   `json:"initialRiskLikelihood"`
		ACM                    string `json:"acm"`
		ResidualRiskSeverity   int8   `json:"residualRiskSeverity"`
		ResidualRiskLikelihood int8   `json:"residualRiskLikelihood"`
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

	var mocModel registerModels.MOCModel

	mocModel.Create(registerTypes.MOC{
		Id:             mocModel.GenerateUniqueId(),
		CompanyId:      account.CompanyId,
		No:             mocModel.GenerateUniqueNo(),
		Issuer:         body.Issuer,
		ReasonOfChange: body.ReasonOfChange,
		Process: tableComponentTypes.DropDownListItem{
			Id: body.Process,
		},
		Action:                 body.Action,
		Risks:                  body.Risks,
		InitialRiskSeverity:    body.InitialRiskSeverity,
		InitialRiskLikelihood:  body.InitialRiskLikelihood,
		ACM:                    body.ACM,
		ResidualRiskSeverity:   body.ResidualRiskSeverity,
		ResidualRiskLikelihood: body.ResidualRiskLikelihood,
		DbStatus:               "active",
		DbLastStatus:           "active",
	})

	c.IndentedJSON(201, gin.H{})
}

func (*MOCHandler) Update(c *gin.Context) {
	account, _ := c.MustGet("account").(middlewares.RemoteAccount)
	Id := c.Param("id")

	var mocModel registerModels.MOCModel

	currentMoc, _ := mocModel.GetById(Id)

	if currentMoc.IsEmpty() || currentMoc.CompanyId != account.CompanyId {
		c.IndentedJSON(404, gin.H{})
		return
	}

	var body struct {
		Issuer                 string `json:"issuer"`
		ReasonOfChange         string `json:"reasonOfChange"`
		Process                string `json:"process"`
		Action                 string `json:"action"`
		Risks                  string `json:"risks"`
		InitialRiskSeverity    int8   `json:"initialRiskSeverity"`
		InitialRiskLikelihood  int8   `json:"initialRiskLikelihood"`
		ACM                    string `json:"acm"`
		ResidualRiskSeverity   int8   `json:"residualRiskSeverity"`
		ResidualRiskLikelihood int8   `json:"residualRiskLikelihood"`
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

	mocModel.Update(Id, map[string]interface{}{
		"issuer":                 body.Issuer,
		"reasonOfChange":         body.ReasonOfChange,
		"process":                body.Process,
		"action":                 body.Action,
		"risks":                  body.Risks,
		"initialRiskSeverity":    body.InitialRiskSeverity,
		"initialRiskLikelihood":  body.InitialRiskLikelihood,
		"acm":                    body.ACM,
		"residualRiskSeverity":   body.ResidualRiskSeverity,
		"residualRiskLikelihood": body.ResidualRiskLikelihood,
	})

	c.JSON(200, gin.H{})
}

func (*MOCHandler) Archive(c *gin.Context) {
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

	var mocModel registerModels.MOCModel

	for _, Id := range body.Ids {
		currentMoc, _ := mocModel.GetById(Id)
		if currentMoc.IsEmpty() || currentMoc.CompanyId != account.CompanyId {
			continue
		}

		if currentMoc.DbStatus != "active" {
			continue
		}

		mocModel.Update(Id, map[string]interface{}{
			"dbStatus":     "archived",
			"dbLastStatus": currentMoc.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*MOCHandler) Unarchive(c *gin.Context) {
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

	var mocModel registerModels.MOCModel

	for _, Id := range body.Ids {
		currentMoc, _ := mocModel.GetById(Id)
		if currentMoc.IsEmpty() || currentMoc.CompanyId != account.CompanyId {
			continue
		}

		if currentMoc.DbStatus != "archived" {
			continue
		}

		mocModel.Update(Id, map[string]interface{}{
			"dbStatus":     "active",
			"dbLastStatus": currentMoc.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*MOCHandler) Delete(c *gin.Context) {
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

	var mocModel registerModels.MOCModel

	for _, Id := range body.Ids {
		currentMoc, _ := mocModel.GetById(Id)
		if currentMoc.IsEmpty() || currentMoc.CompanyId != account.CompanyId {
			continue
		}

		if currentMoc.DbStatus == "deleted" {
			continue
		}

		mocModel.Update(Id, map[string]interface{}{
			"dbStatus":     "deleted",
			"dbLastStatus": currentMoc.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*MOCHandler) Undelete(c *gin.Context) {
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

	var mocModel registerModels.MOCModel

	for _, Id := range body.Ids {
		currentMoc, _ := mocModel.GetById(Id)
		if currentMoc.IsEmpty() || currentMoc.CompanyId != account.CompanyId {
			continue
		}

		if currentMoc.DbStatus != "deleted" {
			continue
		}

		mocModel.Update(Id, map[string]interface{}{
			"dbStatus":     currentMoc.DbLastStatus,
			"dbLastStatus": currentMoc.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}
