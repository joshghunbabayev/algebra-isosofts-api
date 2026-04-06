package registerHandlers

import (
	"algebra-isosofts-api/middlewares"
	registerModels "algebra-isosofts-api/models/registers"
	registerTypes "algebra-isosofts-api/types/registers"
	tableComponentTypes "algebra-isosofts-api/types/tableComponents"

	"github.com/gin-gonic/gin"
)

type EAIHandler struct {
}

func (*EAIHandler) GetAll(c *gin.Context) {
	account, _ := c.MustGet("account").(middlewares.RemoteAccount)
	status := c.Query("status")

	if status == "" {
		status = "active"
	}

	var eaiModel registerModels.EAIModel

	eais, err := eaiModel.GetAll(map[string]interface{}{
		"dbStatus":  status,
		"companyId": account.CompanyId,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, eais)
}

func (*EAIHandler) Create(c *gin.Context) {
	var body struct {
		Process           string `json:"process"`
		Aspect            string `json:"aspect"`
		Impact            string `json:"impact"`
		AffectedReceptors string `json:"affectedReceptors"`
		ECM               string `json:"ecm"`
		IDOSProbability   int8   `json:"idosProbability"`
		IDOSSeverity      int8   `json:"idosSeverity"`
		IDOSDuration      int8   `json:"idosDuration"`
		IDOSScale         int8   `json:"idosScale"`
		ACM               string `json:"acm"`
		RDOSProbability   int8   `json:"rdosProbability"`
		RDOSSeverity      int8   `json:"rdosSeverity"`
		RDOSDuration      int8   `json:"rdosDuration"`
		RDOSScale         int8   `json:"rdosScale"`
		Comment           string `json:"comment"`
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

	var eaiModel registerModels.EAIModel

	account, _ := c.MustGet("account").(middlewares.RemoteAccount)

	eaiModel.Create(registerTypes.EAI{
		Id:        eaiModel.GenerateUniqueId(),
		CompanyId: account.CompanyId,
		No:        eaiModel.GenerateUniqueNo(account.CompanyId),
		Process: tableComponentTypes.DropDownListItem{
			Id: body.Process,
		},
		Aspect: tableComponentTypes.DropDownListItem{
			Id: body.Aspect,
		},
		Impact: body.Impact,
		AffectedReceptors: tableComponentTypes.DropDownListItem{
			Id: body.AffectedReceptors,
		},
		ECM:             body.ECM,
		IDOSProbability: body.IDOSProbability,
		IDOSSeverity:    body.IDOSSeverity,
		IDOSDuration:    body.IDOSDuration,
		IDOSScale:       body.IDOSScale,
		ACM:             body.ACM,
		RDOSProbability: body.RDOSProbability,
		RDOSSeverity:    body.RDOSSeverity,
		RDOSDuration:    body.RDOSDuration,
		RDOSScale:       body.RDOSScale,
		Comment:         body.Comment,
		DbStatus:        "active",
		DbLastStatus:    "active",
	})

	c.IndentedJSON(201, gin.H{})
}

func (*EAIHandler) Update(c *gin.Context) {
	account, _ := c.MustGet("account").(middlewares.RemoteAccount)
	Id := c.Param("id")

	var eaiModel registerModels.EAIModel

	currentEai, _ := eaiModel.GetById(Id)

	if currentEai.IsEmpty() || currentEai.CompanyId != account.CompanyId {
		c.IndentedJSON(404, gin.H{})
		return
	}

	var body struct {
		Process           string `json:"process"`
		Aspect            string `json:"aspect"`
		Impact            string `json:"impact"`
		AffectedReceptors string `json:"affectedReceptors"`
		ECM               string `json:"ecm"`
		IDOSProbability   int8   `json:"idosProbability"`
		IDOSSeverity      int8   `json:"idosSeverity"`
		IDOSDuration      int8   `json:"idosDuration"`
		IDOSScale         int8   `json:"idosScale"`
		ACM               string `json:"acm"`
		RDOSProbability   int8   `json:"rdosProbability"`
		RDOSSeverity      int8   `json:"rdosSeverity"`
		RDOSDuration      int8   `json:"rdosDuration"`
		RDOSScale         int8   `json:"rdosScale"`
		Comment           string `json:"comment"`
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

	eaiModel.Update(Id, map[string]interface{}{
		"process":           body.Process,
		"aspect":            body.Aspect,
		"impact":            body.Impact,
		"affectedReceptors": body.AffectedReceptors,
		"ecm":               body.ECM,
		"idosProbability":   body.IDOSProbability,
		"idosSeverity":      body.IDOSSeverity,
		"idosDuration":      body.IDOSDuration,
		"idosScale":         body.IDOSScale,
		"acm":               body.ACM,
		"rdosProbability":   body.RDOSProbability,
		"rdosSeverity":      body.RDOSSeverity,
		"rdosDuration":      body.RDOSDuration,
		"rdosScale":         body.RDOSScale,
		"comment":           body.Comment,
	})

	c.JSON(200, gin.H{})
}

func (*EAIHandler) Archive(c *gin.Context) {
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

	var eaiModel registerModels.EAIModel

	for _, Id := range body.Ids {
		currentEai, _ := eaiModel.GetById(Id)
		if currentEai.IsEmpty() || currentEai.CompanyId != account.CompanyId {
			continue
		}

		if currentEai.DbStatus != "active" {
			continue
		}

		eaiModel.Update(Id, map[string]interface{}{
			"dbStatus":     "archived",
			"dbLastStatus": currentEai.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*EAIHandler) Unarchive(c *gin.Context) {
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

	var eaiModel registerModels.EAIModel

	for _, Id := range body.Ids {
		currentEai, _ := eaiModel.GetById(Id)
		if currentEai.IsEmpty() || currentEai.CompanyId != account.CompanyId {
			continue
		}

		if currentEai.DbStatus != "archived" {
			continue
		}

		eaiModel.Update(Id, map[string]interface{}{
			"dbStatus":     "active",
			"dbLastStatus": currentEai.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*EAIHandler) Delete(c *gin.Context) {
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

	var eaiModel registerModels.EAIModel

	for _, Id := range body.Ids {
		currentEai, _ := eaiModel.GetById(Id)
		if currentEai.IsEmpty() || currentEai.CompanyId != account.CompanyId {
			continue
		}

		if currentEai.DbStatus == "deleted" {
			continue
		}

		eaiModel.Update(Id, map[string]interface{}{
			"dbStatus":     "deleted",
			"dbLastStatus": currentEai.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*EAIHandler) Undelete(c *gin.Context) {
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

	var eaiModel registerModels.EAIModel

	for _, Id := range body.Ids {
		currentEai, _ := eaiModel.GetById(Id)
		if currentEai.IsEmpty() || currentEai.CompanyId != account.CompanyId {
			continue
		}

		if currentEai.DbStatus != "deleted" {
			continue
		}

		eaiModel.Update(Id, map[string]interface{}{
			"dbStatus":     currentEai.DbLastStatus,
			"dbLastStatus": currentEai.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}
