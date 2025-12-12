package registerHandlers

import (
	registerModels "algebra-isosofts-api/models/registers"
	registerTypes "algebra-isosofts-api/types/registers"
	tableComponentTypes "algebra-isosofts-api/types/tableComponents"

	"github.com/gin-gonic/gin"
)

type HSRHandler struct {
}

func (*HSRHandler) GetAll(c *gin.Context) {
	status := c.Query("status")

	if status == "" {
		status = "active"
	}

	var hsrModel registerModels.HSRModel

	hsrs, err := hsrModel.GetAll(map[string]interface{}{
		"dbStatus": status,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, hsrs)
}

func (*HSRHandler) Create(c *gin.Context) {
	var body struct {
		Process                string `json:"process"`
		Hazard                 string `json:"hazard"`
		Risk                   string `json:"risk"`
		ERMA                   string `json:"erma"`
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

	var hsrModel registerModels.HSRModel

	hsrModel.Create(registerTypes.HSR{
		Id: hsrModel.GenerateUniqueId(),
		No: hsrModel.GenerateUniqueNo(),
		Process: tableComponentTypes.DropDownListItem{
			Id: body.Process,
		},
		Hazard: tableComponentTypes.DropDownListItem{
			Id: body.Hazard,
		},
		Risk: tableComponentTypes.DropDownListItem{
			Id: body.Risk,
		},
		ERMA:                   body.ERMA,
		InitialRiskSeverity:    body.InitialRiskSeverity,
		InitialRiskLikelyhood:  body.InitialRiskLikelyhood,
		ResidualRiskSeverity:   body.ResidualRiskSeverity,
		ResidualRiskLikelyhood: body.ResidualRiskLikelyhood,
		DbStatus:               "active",
		DbLastStatus:           "active",
	})

	c.IndentedJSON(201, gin.H{})
}

func (*HSRHandler) Update(c *gin.Context) {
	Id := c.Param("id")

	var hsrModel registerModels.HSRModel

	currentHSR, _ := hsrModel.GetById(Id)

	if currentHSR.IsEmpty() {
		c.IndentedJSON(404, gin.H{})
		return
	}

	var body struct {
		Process                tableComponentTypes.DropDownListItem `json:"process"`
		Hazard                 tableComponentTypes.DropDownListItem `json:"hazard"`
		Risk                   tableComponentTypes.DropDownListItem `json:"risk"`
		ERMA                   string                               `json:"erma"`
		InitialRiskSeverity    int8                                 `json:"initialRiskSeverity"`
		InitialRiskLikelyhood  int8                                 `json:"initialRiskLikelyhood"`
		ResidualRiskSeverity   int8                                 `json:"residualRiskSeverity"`
		ResidualRiskLikelyhood int8                                 `json:"residualRiskLikelyhood"`
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

	hsrModel.Update(Id, map[string]interface{}{
		"process":                body.Process,
		"hazard":                 body.Hazard,
		"risk":                   body.Risk,
		"erma":                   body.ERMA,
		"initialRiskSeverity":    body.InitialRiskSeverity,
		"initialRiskLikelyhood":  body.InitialRiskLikelyhood,
		"residualRiskSeverity":   body.ResidualRiskSeverity,
		"residualRiskLikelyhood": body.ResidualRiskLikelyhood,
	})

	c.JSON(200, gin.H{})
}

func (*HSRHandler) Archive(c *gin.Context) {
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

	var hsrModel registerModels.HSRModel

	for _, Id := range body.Ids {
		currentHSR, _ := hsrModel.GetById(Id)
		if currentHSR.IsEmpty() {
			continue
		}

		if currentHSR.DbStatus != "active" {
			continue
		}

		hsrModel.Update(Id, map[string]interface{}{
			"dbStatus":     "archived",
			"dbLastStatus": currentHSR.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*HSRHandler) Unarchive(c *gin.Context) {
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

	var hsrModel registerModels.HSRModel

	for _, Id := range body.Ids {
		currentHSR, _ := hsrModel.GetById(Id)
		if currentHSR.IsEmpty() {
			continue
		}

		if currentHSR.DbStatus != "archived" {
			continue
		}

		hsrModel.Update(Id, map[string]interface{}{
			"dbStatus":     "active",
			"dbLastStatus": currentHSR.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*HSRHandler) Delete(c *gin.Context) {
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

	var hsrModel registerModels.HSRModel

	for _, Id := range body.Ids {
		currentHSR, _ := hsrModel.GetById(Id)
		if currentHSR.IsEmpty() {
			continue
		}

		if currentHSR.DbStatus == "deleted" {
			continue
		}

		hsrModel.Update(Id, map[string]interface{}{
			"dbStatus":     "deleted",
			"dbLastStatus": currentHSR.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*HSRHandler) Undelete(c *gin.Context) {
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

	var hsrModel registerModels.HSRModel

	for _, Id := range body.Ids {
		currentHSR, _ := hsrModel.GetById(Id)
		if currentHSR.IsEmpty() {
			continue
		}

		if currentHSR.DbStatus != "deleted" {
			continue
		}

		hsrModel.Update(Id, map[string]interface{}{
			"dbStatus":     currentHSR.DbLastStatus,
			"dbLastStatus": currentHSR.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}
