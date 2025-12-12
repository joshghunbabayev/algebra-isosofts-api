package registerHandlers

import (
	registerModels "algebra-isosofts-api/models/registers"
	registerTypes "algebra-isosofts-api/types/registers"
	tableComponentTypes "algebra-isosofts-api/types/tableComponents"

	"github.com/gin-gonic/gin"
)

type EIAHandler struct {
}

func (*EIAHandler) GetAll(c *gin.Context) {
	status := c.Query("status")

	if status == "" {
		status = "active"
	}

	var eiaModel registerModels.EIAModel

	eias, err := eiaModel.GetAll(map[string]interface{}{
		"dbStatus": status,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, eias)
}

func (*EIAHandler) Create(c *gin.Context) {
	var body struct {
		Process          string `json:"process"`
		Aspect           string `json:"aspect"`
		Impact           string `json:"impact"`
		ExistingControls string `json:"existingControls"`
		IDOSProbability  int8   `json:"idosProbability"`
		IDOSSeverity     int8   `json:"idosSeverity"`
		IDOSDuration     int8   `json:"idosDuration"`
		IDOSScale        int8   `json:"idosScale"`
		RDOSProbability  int8   `json:"rdosProbability"`
		RDOSSeverity     int8   `json:"rdosSeverity"`
		RDOSDuration     int8   `json:"rdosDuration"`
		RDOSScale        int8   `json:"rdosScale"`
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

	var eiaModel registerModels.EIAModel

	eiaModel.Create(registerTypes.EIA{
		Id: eiaModel.GenerateUniqueId(),
		No: eiaModel.GenerateUniqueNo(),
		Process: tableComponentTypes.DropDownListItem{
			Id: body.Process,
		},
		Aspect: tableComponentTypes.DropDownListItem{
			Id: body.Aspect,
		},
		Impact:           body.Impact,
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

func (*EIAHandler) Update(c *gin.Context) {
	Id := c.Param("id")

	var eiaModel registerModels.EIAModel

	currentEIA, _ := eiaModel.GetById(Id)

	if currentEIA.IsEmpty() {
		c.IndentedJSON(404, gin.H{})
		return
	}

	var body struct {
		Process          string `json:"process"`
		Aspect           string `json:"aspect"`
		Impact           string `json:"impact"`
		ExistingControls string `json:"existingControls"`
		IDOSProbability  int8   `json:"idosProbability"`
		IDOSSeverity     int8   `json:"idosSeverity"`
		IDOSDuration     int8   `json:"idosDuration"`
		IDOSScale        int8   `json:"idosScale"`
		RDOSProbability  int8   `json:"rdosProbability"`
		RDOSSeverity     int8   `json:"rdosSeverity"`
		RDOSDuration     int8   `json:"rdosDuration"`
		RDOSScale        int8   `json:"rdosScale"`
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

	eiaModel.Update(Id, map[string]interface{}{
		"process":          body.Process,
		"aspect":           body.Aspect,
		"impact":           body.Impact,
		"existingControls": body.ExistingControls,
		"idosProbability":  body.IDOSProbability,
		"idosSeverity":     body.IDOSSeverity,
		"idosDuration":     body.IDOSDuration,
		"idosScale":        body.IDOSScale,
		"rdosProbability":  body.RDOSProbability,
		"rdosSeverity":     body.RDOSSeverity,
		"rdosDuration":     body.RDOSDuration,
		"rdosScale":        body.RDOSScale,
	})

	c.JSON(200, gin.H{})
}

func (*EIAHandler) Archive(c *gin.Context) {
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

	var eiaModel registerModels.EIAModel

	for _, Id := range body.Ids {
		currentEIA, _ := eiaModel.GetById(Id)
		if currentEIA.IsEmpty() {
			continue
		}

		if currentEIA.DbStatus != "active" {
			continue
		}

		eiaModel.Update(Id, map[string]interface{}{
			"dbStatus":     "archived",
			"dbLastStatus": currentEIA.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*EIAHandler) Unarchive(c *gin.Context) {
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

	var eiaModel registerModels.EIAModel

	for _, Id := range body.Ids {
		currentEIA, _ := eiaModel.GetById(Id)
		if currentEIA.IsEmpty() {
			continue
		}

		if currentEIA.DbStatus != "archived" {
			continue
		}

		eiaModel.Update(Id, map[string]interface{}{
			"dbStatus":     "active",
			"dbLastStatus": currentEIA.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*EIAHandler) Delete(c *gin.Context) {
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

	var eiaModel registerModels.EIAModel

	for _, Id := range body.Ids {
		currentEIA, _ := eiaModel.GetById(Id)
		if currentEIA.IsEmpty() {
			continue
		}

		if currentEIA.DbStatus == "deleted" {
			continue
		}

		eiaModel.Update(Id, map[string]interface{}{
			"dbStatus":     "deleted",
			"dbLastStatus": currentEIA.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*EIAHandler) Undelete(c *gin.Context) {
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

	var eiaModel registerModels.EIAModel

	for _, Id := range body.Ids {
		currentEIA, _ := eiaModel.GetById(Id)
		if currentEIA.IsEmpty() {
			continue
		}

		if currentEIA.DbStatus != "deleted" {
			continue
		}

		eiaModel.Update(Id, map[string]interface{}{
			"dbStatus":     currentEIA.DbLastStatus,
			"dbLastStatus": currentEIA.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}
