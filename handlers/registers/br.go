package registerHandlers

import (
	"algebra-isosofts-api/middlewares"
	registerModels "algebra-isosofts-api/models/registers"
	registerTypes "algebra-isosofts-api/types/registers"
	tableComponentTypes "algebra-isosofts-api/types/tableComponents"

	"github.com/gin-gonic/gin"
)

type BRHandler struct {
}

func (*BRHandler) GetAll(c *gin.Context) {
	account, _ := c.MustGet("account").(middlewares.RemoteAccount)
	status := c.Query("status")

	if status == "" {
		status = "active"
	}

	var brModel registerModels.BRModel

	brs, err := brModel.GetAll(map[string]interface{}{
		"dbStatus":  status,
		"companyId": account.CompanyId,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, brs)
}

func (*BRHandler) Create(c *gin.Context) {
	var body struct {
		Swot                   string `json:"swot"`
		Pestle                 string `json:"pestle"`
		InterestedParty        string `json:"interestedParty"`
		RiskOpportunity        string `json:"riskOpportunity"`
		Objective              string `json:"objective"`
		KPI                    string `json:"kpi"`
		Process                string `json:"process"`
		ECM                    string `json:"ecm"`
		InitialRiskSeverity    int8   `json:"initialRiskSeverity"`
		InitialRiskLikelihood  int8   `json:"initialRiskLikelihood"`
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

	var brModel registerModels.BRModel

	account, _ := c.MustGet("account").(middlewares.RemoteAccount)

	brModel.Create(registerTypes.BR{
		Id:        brModel.GenerateUniqueId(),
		CompanyId: account.CompanyId,
		No:        brModel.GenerateUniqueNo(),
		Swot: tableComponentTypes.DropDownListItem{
			Id: body.Swot,
		},
		Pestle: tableComponentTypes.DropDownListItem{
			Id: body.Pestle,
		},
		InterestedParty: tableComponentTypes.DropDownListItem{
			Id: body.InterestedParty,
		},
		RiskOpportunity: body.RiskOpportunity,
		Objective:       body.Objective,
		KPI:             body.KPI,
		Process: tableComponentTypes.DropDownListItem{
			Id: body.Process,
		},
		ECM:                    body.ECM,
		InitialRiskSeverity:    body.InitialRiskSeverity,
		InitialRiskLikelihood:  body.InitialRiskLikelihood,
		ResidualRiskSeverity:   body.ResidualRiskSeverity,
		ResidualRiskLikelihood: body.ResidualRiskLikelihood,
		DbStatus:               "active",
		DbLastStatus:           "active",
	})

	c.IndentedJSON(201, gin.H{})
}

func (*BRHandler) Update(c *gin.Context) {
	account, _ := c.MustGet("account").(middlewares.RemoteAccount)
	Id := c.Param("id")

	var brModel registerModels.BRModel

	currentBr, _ := brModel.GetById(Id)

	if currentBr.IsEmpty() || currentBr.CompanyId != account.CompanyId {
		c.IndentedJSON(404, gin.H{})
		return
	}

	var body struct {
		Swot                   string `json:"swot"`
		Pestle                 string `json:"pestle"`
		InterestedParty        string `json:"interestedParty"`
		RiskOpportunity        string `json:"riskOpportunity"`
		Objective              string `json:"objective"`
		KPI                    string `json:"kpi"`
		Process                string `json:"process"`
		ECM                    string `json:"ecm"`
		InitialRiskSeverity    int8   `json:"initialRiskSeverity"`
		InitialRiskLikelihood  int8   `json:"initialRiskLikelihood"`
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

	brModel.Update(Id, map[string]interface{}{
		"swot":                   body.Swot,
		"pestle":                 body.Pestle,
		"interestedParty":        body.InterestedParty,
		"riskOpportunity":        body.RiskOpportunity,
		"objective":              body.Objective,
		"kpi":                    body.KPI,
		"process":                body.Process,
		"ecm":                    body.ECM,
		"initialRiskSeverity":    body.InitialRiskSeverity,
		"initialRiskLikelihood":  body.InitialRiskLikelihood,
		"residualRiskSeverity":   body.ResidualRiskSeverity,
		"residualRiskLikelihood": body.ResidualRiskLikelihood,
	})

	c.JSON(200, gin.H{})
}

func (*BRHandler) Archive(c *gin.Context) {
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

	var brModel registerModels.BRModel

	for _, Id := range body.Ids {
		currentBr, _ := brModel.GetById(Id)
		if currentBr.IsEmpty() || currentBr.CompanyId != account.CompanyId {
			continue
		}

		if currentBr.DbStatus != "active" {
			continue
		}

		brModel.Update(Id, map[string]interface{}{
			"dbStatus":     "archived",
			"dbLastStatus": currentBr.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*BRHandler) Unarchive(c *gin.Context) {
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

	var brModel registerModels.BRModel

	for _, Id := range body.Ids {
		currentBr, _ := brModel.GetById(Id)
		if currentBr.IsEmpty() || currentBr.CompanyId != account.CompanyId {
			continue
		}

		if currentBr.DbStatus != "archived" {
			continue
		}

		brModel.Update(Id, map[string]interface{}{
			"dbStatus":     "active",
			"dbLastStatus": currentBr.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*BRHandler) Delete(c *gin.Context) {
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

	var brModel registerModels.BRModel

	for _, Id := range body.Ids {
		currentBr, _ := brModel.GetById(Id)
		if currentBr.IsEmpty() || currentBr.CompanyId != account.CompanyId {
			continue
		}

		if currentBr.DbStatus == "deleted" {
			continue
		}

		brModel.Update(Id, map[string]interface{}{
			"dbStatus":     "deleted",
			"dbLastStatus": currentBr.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*BRHandler) Undelete(c *gin.Context) {
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

	var brModel registerModels.BRModel

	for _, Id := range body.Ids {
		currentBr, _ := brModel.GetById(Id)
		if currentBr.IsEmpty() || currentBr.CompanyId != account.CompanyId {
			continue
		}

		if currentBr.DbStatus != "deleted" {
			continue
		}

		brModel.Update(Id, map[string]interface{}{
			"dbStatus":     currentBr.DbLastStatus,
			"dbLastStatus": currentBr.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}
