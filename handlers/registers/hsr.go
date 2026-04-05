package registerHandlers

import (
	"algebra-isosofts-api/middlewares"
	registerModels "algebra-isosofts-api/models/registers"
	registerTypes "algebra-isosofts-api/types/registers"
	tableComponentTypes "algebra-isosofts-api/types/tableComponents"

	"github.com/gin-gonic/gin"
)

type HSRHandler struct {
}

func (*HSRHandler) GetAll(c *gin.Context) {
	account, _ := c.MustGet("account").(middlewares.RemoteAccount)
	status := c.Query("status")

	if status == "" {
		status = "active"
	}

	var hsrModel registerModels.HSRModel

	hsrs, err := hsrModel.GetAll(map[string]interface{}{
		"dbStatus":  status,
		"companyId": account.CompanyId,
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

	var hsrModel registerModels.HSRModel

	account, _ := c.MustGet("account").(middlewares.RemoteAccount)

	hsrModel.Create(registerTypes.HSR{
		Id:        hsrModel.GenerateUniqueId(),
		CompanyId: account.CompanyId,
		No:        hsrModel.GenerateUniqueNo(),
		Process: tableComponentTypes.DropDownListItem{
			Id: body.Process,
		},
		Hazard: tableComponentTypes.DropDownListItem{
			Id: body.Hazard,
		},
		Risk: tableComponentTypes.DropDownListItem{
			Id: body.Risk,
		},
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

func (*HSRHandler) Update(c *gin.Context) {
	account, _ := c.MustGet("account").(middlewares.RemoteAccount)
	Id := c.Param("id")

	var hsrModel registerModels.HSRModel

	currentHsr, _ := hsrModel.GetById(Id)

	if currentHsr.IsEmpty() || currentHsr.CompanyId != account.CompanyId {
		c.IndentedJSON(404, gin.H{})
		return
	}

	var body struct {
		Process                string `json:"process"`
		Hazard                 string `json:"hazard"`
		Risk                   string `json:"risk"`
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

	hsrModel.Update(Id, map[string]interface{}{
		"process":                body.Process,
		"hazard":                 body.Hazard,
		"risk":                   body.Risk,
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

func (*HSRHandler) Archive(c *gin.Context) {
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

	var hsrModel registerModels.HSRModel

	for _, Id := range body.Ids {
		currentHsr, _ := hsrModel.GetById(Id)
		if currentHsr.IsEmpty() || currentHsr.CompanyId != account.CompanyId {
			c.IndentedJSON(404, gin.H{})
			return
		}

		if currentHsr.DbStatus != "active" {
			continue
		}

		hsrModel.Update(Id, map[string]interface{}{
			"dbStatus":     "archived",
			"dbLastStatus": currentHsr.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*HSRHandler) Unarchive(c *gin.Context) {
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

	var hsrModel registerModels.HSRModel

	for _, Id := range body.Ids {
		currentHsr, _ := hsrModel.GetById(Id)
		if currentHsr.IsEmpty() || currentHsr.CompanyId != account.CompanyId {
			c.IndentedJSON(404, gin.H{})
			return
		}

		if currentHsr.DbStatus != "archived" {
			continue
		}

		hsrModel.Update(Id, map[string]interface{}{
			"dbStatus":     "active",
			"dbLastStatus": currentHsr.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*HSRHandler) Delete(c *gin.Context) {
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

	var hsrModel registerModels.HSRModel

	for _, Id := range body.Ids {
		currentHsr, _ := hsrModel.GetById(Id)
		if currentHsr.IsEmpty() || currentHsr.CompanyId != account.CompanyId {
			c.IndentedJSON(404, gin.H{})
			return
		}

		if currentHsr.DbStatus == "deleted" {
			continue
		}

		hsrModel.Update(Id, map[string]interface{}{
			"dbStatus":     "deleted",
			"dbLastStatus": currentHsr.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*HSRHandler) Undelete(c *gin.Context) {
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

	var hsrModel registerModels.HSRModel

	for _, Id := range body.Ids {
		currentHsr, _ := hsrModel.GetById(Id)
		if currentHsr.IsEmpty() || currentHsr.CompanyId != account.CompanyId {
			c.IndentedJSON(404, gin.H{})
			return
		}

		if currentHsr.DbStatus != "deleted" {
			continue
		}

		hsrModel.Update(Id, map[string]interface{}{
			"dbStatus":     currentHsr.DbLastStatus,
			"dbLastStatus": currentHsr.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}
