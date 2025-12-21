package registerHandlers

import (
	registerModels "algebra-isosofts-api/models/registers"
	registerTypes "algebra-isosofts-api/types/registers"
	tableComponentTypes "algebra-isosofts-api/types/tableComponents"

	"github.com/gin-gonic/gin"
)

type VENHandler struct {
}

func (*VENHandler) GetAll(c *gin.Context) {
	status := c.Query("status")

	if status == "" {
		status = "active"
	}

	var venModel registerModels.VENModel

	vens, err := venModel.GetAll(map[string]interface{}{
		"dbStatus": status,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, vens)
}

func (*VENHandler) Create(c *gin.Context) {
	var body struct {
		Name             string `json:"name"`
		RegNumber        string `json:"regNumber"`
		Scope1           string `json:"scope1"`
		Scope2           string `json:"scope2"`
		Scope3           string `json:"scope3"`
		RegistrationDate string `json:"registrationDate"`
		ReviewDate       string `json:"reviewDate"`
		Approved         int8   `json:"approved"`
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

	var venModel registerModels.VENModel

	venModel.Create(registerTypes.VEN{
		Id:        venModel.GenerateUniqueId(),
		No:        venModel.GenerateUniqueNo(),
		Name:      body.Name,
		RegNumber: body.RegNumber,
		Scope1: tableComponentTypes.DropDownListItem{
			Id: body.Scope1,
		},
		Scope2: tableComponentTypes.DropDownListItem{
			Id: body.Scope2,
		},
		Scope3: tableComponentTypes.DropDownListItem{
			Id: body.Scope3,
		},
		RegistrationDate: body.RegistrationDate,
		ReviewDate:       body.ReviewDate,
		Approved:         body.Approved,
		DbStatus:         "active",
		DbLastStatus:     "active",
	})

	c.IndentedJSON(201, gin.H{})
}

func (*VENHandler) Update(c *gin.Context) {
	Id := c.Param("id")

	var venModel registerModels.VENModel

	currentVen, _ := venModel.GetById(Id)

	if currentVen.IsEmpty() {
		c.IndentedJSON(404, gin.H{})
		return
	}

	var body struct {
		Name             string `json:"name"`
		RegNumber        string `json:"regNumber"`
		Scope1           string `json:"scope1"`
		Scope2           string `json:"scope2"`
		Scope3           string `json:"scope3"`
		RegistrationDate string `json:"registrationDate"`
		ReviewDate       string `json:"reviewDate"`
		Approved         int8   `json:"approved"`
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

	venModel.Update(Id, map[string]interface{}{
		"name":             body.Name,
		"regNumber":        body.RegNumber,
		"scope1":           body.Scope1,
		"scope2":           body.Scope2,
		"scope3":           body.Scope3,
		"registrationDate": body.RegistrationDate,
		"reviewDate":       body.ReviewDate,
		"approved":         body.Approved,
	})

	c.JSON(200, gin.H{})
}

func (*VENHandler) Archive(c *gin.Context) {
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

	var venModel registerModels.VENModel

	for _, Id := range body.Ids {
		currentVen, _ := venModel.GetById(Id)
		if currentVen.IsEmpty() {
			continue
		}

		if currentVen.DbStatus != "active" {
			continue
		}

		venModel.Update(Id, map[string]interface{}{
			"dbStatus":     "archived",
			"dbLastStatus": currentVen.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*VENHandler) Unarchive(c *gin.Context) {
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

	var venModel registerModels.VENModel

	for _, Id := range body.Ids {
		currentVen, _ := venModel.GetById(Id)
		if currentVen.IsEmpty() {
			continue
		}

		if currentVen.DbStatus != "archived" {
			continue
		}

		venModel.Update(Id, map[string]interface{}{
			"dbStatus":     "active",
			"dbLastStatus": currentVen.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*VENHandler) Delete(c *gin.Context) {
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

	var venModel registerModels.VENModel

	for _, Id := range body.Ids {
		currentVen, _ := venModel.GetById(Id)
		if currentVen.IsEmpty() {
			continue
		}

		if currentVen.DbStatus == "deleted" {
			continue
		}

		venModel.Update(Id, map[string]interface{}{
			"dbStatus":     "deleted",
			"dbLastStatus": currentVen.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*VENHandler) Undelete(c *gin.Context) {
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

	var venModel registerModels.VENModel

	for _, Id := range body.Ids {
		currentVen, _ := venModel.GetById(Id)
		if currentVen.IsEmpty() {
			continue
		}

		if currentVen.DbStatus != "deleted" {
			continue
		}

		venModel.Update(Id, map[string]interface{}{
			"dbStatus":     currentVen.DbLastStatus,
			"dbLastStatus": currentVen.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}
