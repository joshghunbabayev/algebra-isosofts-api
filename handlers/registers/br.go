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
	var brModel registerModels.BrModel

	brs, err := brModel.GetAll()

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

	br := registerTypes.Br{
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
	}

	brModel.Create(br)

	c.JSON(201, gin.H{})
}
