package registerHandlers

import (
	registerModels "algebra-isosofts-api/models/registers"
	registerTypes "algebra-isosofts-api/types/registers"
	tableComponentTypes "algebra-isosofts-api/types/tableComponents"

	"github.com/gin-gonic/gin"
)

type FBHandler struct {
}

func (*FBHandler) GetAll(c *gin.Context) {
	status := c.Query("status")

	if status == "" {
		status = "active"
	}

	var fbModel registerModels.FBModel

	fbs, err := fbModel.GetAll(map[string]interface{}{
		"dbStatus": status,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, fbs)
}

func (*FBHandler) Create(c *gin.Context) {
	var body struct {
		JobNumber         string `json:"jobNumber"`
		JobStartDate      string `json:"jobStartDate"`
		JobCompletionDate string `json:"jobCompletionDate"`
		Scope             string `json:"scope"`
		CustomerId        string `json:"customerId"`
		TypeOfFinding     string `json:"typeOfFinding"`
		QGS               int8   `json:"qgs"`
		Communication     int8   `json:"communication"`
		OTD               int8   `json:"otd"`
		Documentation     int8   `json:"documentation"`
		HS                int8   `json:"hs"`
		Environment       int8   `json:"environment"`
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

	var fbModel registerModels.FBModel

	fbModel.Create(registerTypes.FB{
		Id:                fbModel.GenerateUniqueId(),
		No:                fbModel.GenerateUniqueNo(),
		JobNumber:         body.JobNumber,
		JobStartDate:      body.JobStartDate,
		JobCompletionDate: body.JobCompletionDate,
		Scope: tableComponentTypes.DropDownListItem{
			Id: body.Scope,
		},
		CustomerId: body.CustomerId,
		TypeOfFinding: tableComponentTypes.DropDownListItem{
			Id: body.TypeOfFinding,
		},
		QGS:           body.QGS,
		Communication: body.Communication,
		OTD:           body.OTD,
		Documentation: body.Documentation,
		HS:            body.HS,
		Environment:   body.Environment,
		DbStatus:      "active",
		DbLastStatus:  "active",
	})

	c.IndentedJSON(201, gin.H{})
}

func (*FBHandler) Update(c *gin.Context) {
	Id := c.Param("id")

	var fbModel registerModels.FBModel

	currentFb, _ := fbModel.GetById(Id)

	if currentFb.IsEmpty() {
		c.IndentedJSON(404, gin.H{})
		return
	}

	var body struct {
		JobNumber         string `json:"jobNumber"`
		JobStartDate      string `json:"jobStartDate"`
		JobCompletionDate string `json:"jobCompletionDate"`
		Scope             string `json:"scope"`
		CustomerId        string `json:"customerId"`
		TypeOfFinding     string `json:"typeOfFinding"`
		QGS               int8   `json:"qgs"`
		Communication     int8   `json:"communication"`
		OTD               int8   `json:"otd"`
		Documentation     int8   `json:"documentation"`
		HS                int8   `json:"hs"`
		Environment       int8   `json:"environment"`
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

	fbModel.Update(Id, map[string]interface{}{
		"jobNumber":         body.JobNumber,
		"jobStartDate":      body.JobStartDate,
		"jobCompletionDate": body.JobCompletionDate,
		"scope":             body.Scope,
		"customerId":        body.CustomerId,
		"typeOfFinding":     body.TypeOfFinding,
		"qgs":               body.QGS,
		"communication":     body.Communication,
		"otd":               body.OTD,
		"documentation":     body.Documentation,
		"hs":                body.HS,
		"environment":       body.Environment,
	})

	c.JSON(200, gin.H{})
}

func (*FBHandler) Archive(c *gin.Context) {
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

	var fbModel registerModels.FBModel

	for _, Id := range body.Ids {
		currentFb, _ := fbModel.GetById(Id)
		if currentFb.IsEmpty() {
			continue
		}

		if currentFb.DbStatus != "active" {
			continue
		}

		fbModel.Update(Id, map[string]interface{}{
			"dbStatus":     "archived",
			"dbLastStatus": currentFb.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*FBHandler) Unarchive(c *gin.Context) {
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

	var fbModel registerModels.FBModel

	for _, Id := range body.Ids {
		currentFb, _ := fbModel.GetById(Id)
		if currentFb.IsEmpty() {
			continue
		}

		if currentFb.DbStatus != "archived" {
			continue
		}

		fbModel.Update(Id, map[string]interface{}{
			"dbStatus":     "active",
			"dbLastStatus": currentFb.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*FBHandler) Delete(c *gin.Context) {
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

	var fbModel registerModels.FBModel

	for _, Id := range body.Ids {
		currentFb, _ := fbModel.GetById(Id)
		if currentFb.IsEmpty() {
			continue
		}

		if currentFb.DbStatus == "deleted" {
			continue
		}

		fbModel.Update(Id, map[string]interface{}{
			"dbStatus":     "deleted",
			"dbLastStatus": currentFb.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*FBHandler) Undelete(c *gin.Context) {
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

	var fbModel registerModels.FBModel

	for _, Id := range body.Ids {
		currentFb, _ := fbModel.GetById(Id)
		if currentFb.IsEmpty() {
			continue
		}

		if currentFb.DbStatus != "deleted" {
			continue
		}

		fbModel.Update(Id, map[string]interface{}{
			"dbStatus":     currentFb.DbLastStatus,
			"dbLastStatus": currentFb.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}
