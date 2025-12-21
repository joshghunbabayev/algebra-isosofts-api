package registerHandlers

import (
	registerModels "algebra-isosofts-api/models/registers"
	registerTypes "algebra-isosofts-api/types/registers"
	tableComponentTypes "algebra-isosofts-api/types/tableComponents"

	"github.com/gin-gonic/gin"
)

type CUSHandler struct {
}

func (*CUSHandler) GetAll(c *gin.Context) {
	status := c.Query("status")

	if status == "" {
		status = "active"
	}

	var cusModel registerModels.CUSModel

	cuss, err := cusModel.GetAll(map[string]interface{}{
		"dbStatus": status,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, cuss)
}

func (*CUSHandler) Create(c *gin.Context) {
	var body struct {
		Name             string `json:"name"`
		RegNumber        string `json:"regNumber"`
		Scope1           string `json:"scope1"`
		Scope2           string `json:"scope2"`
		Scope3           string `json:"scope3"`
		RegistrationDate string `json:"registrationDate"`
		ReviewDate       string `json:"reviewDate"`
		Actual           int8   `json:"actual"`
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

	var cusModel registerModels.CUSModel

	cusModel.Create(registerTypes.CUS{
		Id:        cusModel.GenerateUniqueId(),
		No:        cusModel.GenerateUniqueNo(),
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
		Actual:           body.Actual,
		DbStatus:         "active",
		DbLastStatus:     "active",
	})

	c.IndentedJSON(201, gin.H{})
}

func (*CUSHandler) Update(c *gin.Context) {
	Id := c.Param("id")

	var cusModel registerModels.CUSModel

	currentCus, _ := cusModel.GetById(Id)

	if currentCus.IsEmpty() {
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
		Actual           int8   `json:"actual"`
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

	cusModel.Update(Id, map[string]interface{}{
		"name":             body.Name,
		"regNumber":        body.RegNumber,
		"scope1":           body.Scope1,
		"scope2":           body.Scope2,
		"scope3":           body.Scope3,
		"registrationDate": body.RegistrationDate,
		"reviewDate":       body.ReviewDate,
		"actual":           body.Actual,
	})

	c.JSON(200, gin.H{})
}

func (*CUSHandler) Archive(c *gin.Context) {
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

	var cusModel registerModels.CUSModel

	for _, Id := range body.Ids {
		currentCus, _ := cusModel.GetById(Id)
		if currentCus.IsEmpty() {
			continue
		}

		if currentCus.DbStatus != "active" {
			continue
		}

		cusModel.Update(Id, map[string]interface{}{
			"dbStatus":     "archived",
			"dbLastStatus": currentCus.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*CUSHandler) Unarchive(c *gin.Context) {
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

	var cusModel registerModels.CUSModel

	for _, Id := range body.Ids {
		currentCus, _ := cusModel.GetById(Id)
		if currentCus.IsEmpty() {
			continue
		}

		if currentCus.DbStatus != "archived" {
			continue
		}

		cusModel.Update(Id, map[string]interface{}{
			"dbStatus":     "active",
			"dbLastStatus": currentCus.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*CUSHandler) Delete(c *gin.Context) {
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

	var cusModel registerModels.CUSModel

	for _, Id := range body.Ids {
		currentCus, _ := cusModel.GetById(Id)
		if currentCus.IsEmpty() {
			continue
		}

		if currentCus.DbStatus == "deleted" {
			continue
		}

		cusModel.Update(Id, map[string]interface{}{
			"dbStatus":     "deleted",
			"dbLastStatus": currentCus.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*CUSHandler) Undelete(c *gin.Context) {
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

	var cusModel registerModels.CUSModel

	for _, Id := range body.Ids {
		currentCus, _ := cusModel.GetById(Id)
		if currentCus.IsEmpty() {
			continue
		}

		if currentCus.DbStatus != "deleted" {
			continue
		}

		cusModel.Update(Id, map[string]interface{}{
			"dbStatus":     currentCus.DbLastStatus,
			"dbLastStatus": currentCus.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}
