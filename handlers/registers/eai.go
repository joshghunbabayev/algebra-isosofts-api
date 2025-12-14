package registerHandlers

import (
	registerModels "algebra-isosofts-api/models/registers"
	registerTypes "algebra-isosofts-api/types/registers"
	tableComponentTypes "algebra-isosofts-api/types/tableComponents"

	"github.com/gin-gonic/gin"
)

type EAIHandler struct {
}

func (*EAIHandler) GetAll(c *gin.Context) {
	status := c.Query("status")

	if status == "" {
		status = "active"
	}

	var eaiModel registerModels.EAIModel

	eais, err := eaiModel.GetAll(map[string]interface{}{
		"dbStatus": status,
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
		ExistingControls  string `json:"existingControls"`
		IDOSProbability   int8   `json:"idosProbability"`
		IDOSSeverity      int8   `json:"idosSeverity"`
		IDOSDuration      int8   `json:"idosDuration"`
		IDOSScale         int8   `json:"idosScale"`
		RDOSProbability   int8   `json:"rdosProbability"`
		RDOSSeverity      int8   `json:"rdosSeverity"`
		RDOSDuration      int8   `json:"rdosDuration"`
		RDOSScale         int8   `json:"rdosScale"`
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

	eaiModel.Create(registerTypes.EAI{
		Id: eaiModel.GenerateUniqueId(),
		No: eaiModel.GenerateUniqueNo(),
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
		ExistingControls: body.ExistingControls,
		IDOSProbability:  body.IDOSProbability,
		IDOSSeverity:     body.IDOSSeverity,
		IDOSDuration:     body.IDOSDuration,
		IDOSScale:        body.IDOSScale,
		RDOSProbability:  body.RDOSProbability,
		RDOSSeverity:     body.RDOSSeverity,
		RDOSDuration:     body.RDOSDuration,
		RDOSScale:        body.RDOSScale,
		DbStatus:         "active",
		DbLastStatus:     "active",
	})

	c.IndentedJSON(201, gin.H{})
}

func (*EAIHandler) Update(c *gin.Context) {
	Id := c.Param("id")

	var eaiModel registerModels.EAIModel

	currentEai, _ := eaiModel.GetById(Id)

	if currentEai.IsEmpty() {
		c.IndentedJSON(404, gin.H{})
		return
	}

	var body struct {
		Process           string `json:"process"`
		Aspect            string `json:"aspect"`
		Impact            string `json:"impact"`
		AffectedReceptors string `json:"affectedReceptors"`
		ExistingControls  string `json:"existingControls"`
		IDOSProbability   int8   `json:"idosProbability"`
		IDOSSeverity      int8   `json:"idosSeverity"`
		IDOSDuration      int8   `json:"idosDuration"`
		IDOSScale         int8   `json:"idosScale"`
		RDOSProbability   int8   `json:"rdosProbability"`
		RDOSSeverity      int8   `json:"rdosSeverity"`
		RDOSDuration      int8   `json:"rdosDuration"`
		RDOSScale         int8   `json:"rdosScale"`
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
		"existingControls":  body.ExistingControls,
		"idosProbability":   body.IDOSProbability,
		"idosSeverity":      body.IDOSSeverity,
		"idosDuration":      body.IDOSDuration,
		"idosScale":         body.IDOSScale,
		"rdosProbability":   body.RDOSProbability,
		"rdosSeverity":      body.RDOSSeverity,
		"rdosDuration":      body.RDOSDuration,
		"rdosScale":         body.RDOSScale,
	})

	c.JSON(200, gin.H{})
}

func (*EAIHandler) Archive(c *gin.Context) {
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
		if currentEai.IsEmpty() {
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
		if currentEai.IsEmpty() {
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
		if currentEai.IsEmpty() {
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
		if currentEai.IsEmpty() {
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
