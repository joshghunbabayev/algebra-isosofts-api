package registerHandlers

import (
	registerModels "algebra-isosofts-api/models/registers"
	registerTypes "algebra-isosofts-api/types/registers"
	tableComponentTypes "algebra-isosofts-api/types/tableComponents"

	"github.com/gin-gonic/gin"
)

type BrHandler struct {
}

func (*BrHandler) GetAll(c *gin.Context) {
	status := c.Query("status")

	if status == "" {
		status = "active"
	}

	var brModel registerModels.BrModel

	brs, err := brModel.GetAll(map[string]interface{}{
		"dbStatus": status,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, brs)
}

func (*BrHandler) Create(c *gin.Context) {
	var body struct {
		Swot                   string `json:"swot"`
		Pestle                 string `json:"pestle"`
		InterestedParty        string `json:"interestedParty"`
		RiskOpportunity        string `json:"riskOpportunity"`
		Objective              string `json:"objective"`
		KPI                    string `json:"kpi"`
		Process                string `json:"process"`
		ERMEOA                 string `json:"ermeoa"`
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

	var brModel registerModels.BrModel

	brModel.Create(registerTypes.Br{
		Id: brModel.GenerateUniqueId(),
		No: "DEFVALUE!",
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
		ERMEOA:                 body.ERMEOA,
		InitialRiskSeverity:    body.InitialRiskSeverity,
		InitialRiskLikelyhood:  body.InitialRiskLikelyhood,
		ResidualRiskSeverity:   body.ResidualRiskSeverity,
		ResidualRiskLikelyhood: body.ResidualRiskLikelyhood,
		DbStatus:               "active",
		DbLastStatus:           "active",
	})

	c.IndentedJSON(201, gin.H{})
}

func (*BrHandler) Update(c *gin.Context) {
	Id := c.Param("id")

	var brModel registerModels.BrModel

	currentBr, _ := brModel.GetById(Id)

	if currentBr.IsEmpty() {
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
		ERMEOA                 string `json:"ermeoa"`
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

	brModel.Update(Id, map[string]interface{}{
		"swot":                   body.Swot,
		"pestle":                 body.Pestle,
		"interestedParty":        body.InterestedParty,
		"riskOpportunity":        body.RiskOpportunity,
		"objective":              body.Objective,
		"kpi":                    body.KPI,
		"process":                body.Process,
		"ermeoa":                 body.ERMEOA,
		"initialRiskSeverity":    body.InitialRiskSeverity,
		"initialRiskLikelyhood":  body.InitialRiskLikelyhood,
		"residualRiskSeverity":   body.ResidualRiskSeverity,
		"residualRiskLikelyhood": body.ResidualRiskLikelyhood,
	})

	c.JSON(200, gin.H{})
}

func (*BrHandler) Archive(c *gin.Context) {
	Id := c.Param("id")

	var brModel registerModels.BrModel

	currentBr, _ := brModel.GetById(Id)

	if currentBr.IsEmpty() {
		c.IndentedJSON(404, gin.H{})
		return
	}

	// ikinci shert -> eger active dirse

	brModel.Update(Id, map[string]interface{}{
		"dbStatus":     "archived",
		"dbLastStatus": currentBr.DbStatus,
	})

	c.JSON(200, gin.H{})
}

func (*BrHandler) Unarchive(c *gin.Context) {
	Id := c.Param("id")

	var brModel registerModels.BrModel

	currentBr, _ := brModel.GetById(Id)

	if currentBr.IsEmpty() {
		c.IndentedJSON(404, gin.H{})
		return
	}

	// ikinci shert -> eger arxivdirse dirse

	brModel.Update(Id, map[string]interface{}{
		"dbStatus":     "active",
		"dbLastStatus": currentBr.DbStatus,
	})

	c.JSON(200, gin.H{})
}

func (*BrHandler) Delete(c *gin.Context) {
	Id := c.Param("id")

	var brModel registerModels.BrModel

	currentBr, _ := brModel.GetById(Id)

	if currentBr.IsEmpty() {
		c.IndentedJSON(404, gin.H{})
		return
	}

	brModel.Update(Id, map[string]interface{}{
		"dbStatus":     "deleted",
		"dbLastStatus": currentBr.DbStatus,
	})

	c.JSON(200, gin.H{})
}

func (*BrHandler) Undelete(c *gin.Context) {
	Id := c.Param("id")

	var brModel registerModels.BrModel

	currentBr, _ := brModel.GetById(Id)

	if currentBr.IsEmpty() {
		c.IndentedJSON(404, gin.H{})
		return
	}

	// ikinci shert -> eger silinibse dirse

	brModel.Update(Id, map[string]interface{}{
		"dbStatus":     currentBr.DbLastStatus,
		"dbLastStatus": "deleted",
	})

	c.JSON(200, gin.H{})
}
