package registerHandlers

import (
	registerModels "algebra-isosofts-api/models/registers"
	tableComponentModels "algebra-isosofts-api/models/tableComponents"
	registerTypes "algebra-isosofts-api/types/registers"
	tableComponentTypes "algebra-isosofts-api/types/tableComponents"

	"github.com/gin-gonic/gin"
)

type DOCHandler struct {
}

func (*DOCHandler) GetAll(c *gin.Context) {
	status := c.Query("status")

	if status == "" {
		status = "active"
	}

	var docModel registerModels.DOCModel

	docs, err := docModel.GetAll(map[string]interface{}{
		"dbStatus": status,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, docs)
}

func (*DOCHandler) Create(c *gin.Context) {
	var body struct {
		Name              string `json:"name"`
		Origin            string `json:"origin"`
		Number            string `json:"number"`
		DepntFunctionName string `json:"depntFunctionName"`
		Type              string `json:"type"`
		SerialNumber      string `json:"serialNumber"`
		RevNumber         string `json:"revNumber"`
		Issuer            string `json:"issuer"`
		Approver          string `json:"approver"`
		IssueDate         string `json:"issueDate"`
		NextReviewDate    string `json:"nextReviewDate"`
		Actual            int8   `json:"actual"`
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

	var docModel registerModels.DOCModel
	var dropDownListItemModel tableComponentModels.DropDownListItemModel
	var number string

	originDDLI, _ := dropDownListItemModel.GetById(body.Origin)
	if originDDLI.Value == "Internal" {
		depntFunctionNameDDLI, _ := dropDownListItemModel.GetById(body.DepntFunctionName)
		typeDDLI, _ := dropDownListItemModel.GetById(body.Type)
		number = docModel.GenerateUniqueNumber(depntFunctionNameDDLI.ShortValue + "/" + typeDDLI.ShortValue)
	} else if originDDLI.Value == "External" {
		number = body.Number
	}

	docModel.Create(registerTypes.DOC{
		Id:   docModel.GenerateUniqueId(),
		No:   docModel.GenerateUniqueNo(),
		Name: body.Name,
		Origin: tableComponentTypes.DropDownListItem{
			Id: body.Origin,
		},
		Number: number,
		DepntFunctionName: tableComponentTypes.DropDownListItem{
			Id: body.DepntFunctionName,
		},
		Type: tableComponentTypes.DropDownListItem{
			Id: body.Type,
		},
		SerialNumber:   body.SerialNumber,
		RevNumber:      body.RevNumber,
		Issuer:         body.Issuer,
		Approver:       body.Approver,
		IssueDate:      body.IssueDate,
		NextReviewDate: body.NextReviewDate,
		Actual:         body.Actual,
		DbStatus:       "active",
		DbLastStatus:   "active",
	})

	c.IndentedJSON(201, gin.H{})
}

func (*DOCHandler) Update(c *gin.Context) {
	Id := c.Param("id")

	var docModel registerModels.DOCModel

	currentDoc, _ := docModel.GetById(Id)

	if currentDoc.IsEmpty() {
		c.IndentedJSON(404, gin.H{})
		return
	}

	var body struct {
		Name              string `json:"name"`
		Origin            string `json:"origin"`
		Number            string `json:"number"`
		DepntFunctionName string `json:"depntFunctionName"`
		Type              string `json:"type"`
		SerialNumber      string `json:"serialNumber"`
		RevNumber         string `json:"revNumber"`
		Issuer            string `json:"issuer"`
		Approver          string `json:"approver"`
		IssueDate         string `json:"issueDate"`
		NextReviewDate    string `json:"nextReviewDate"`
		Actual            int8   `json:"actual"`
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

	var dropDownListItemModel tableComponentModels.DropDownListItemModel
	var number string

	originDDLI, _ := dropDownListItemModel.GetById(body.Origin)
	if originDDLI.Value == "Internal" {
		depntFunctionNameDDLI, _ := dropDownListItemModel.GetById(body.DepntFunctionName)
		typeDDLI, _ := dropDownListItemModel.GetById(body.Type)
		number = docModel.GenerateUniqueNumber(depntFunctionNameDDLI.ShortValue + "/" + typeDDLI.ShortValue)
	} else if originDDLI.Value == "External" {
		number = body.Number
	}

	docModel.Update(Id, map[string]interface{}{
		"name":              body.Name,
		"origin":            body.Origin,
		"number":            number,
		"depntFunctionName": body.DepntFunctionName,
		"type":              body.Type,
		"serialNumber":      body.SerialNumber,
		"revNumber":         body.RevNumber,
		"issuer":            body.Issuer,
		"approver":          body.Approver,
		"issueDate":         body.IssueDate,
		"nextReviewDate":    body.NextReviewDate,
		"actual":            body.Actual,
	})

	c.JSON(200, gin.H{})
}

func (*DOCHandler) Archive(c *gin.Context) {
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

	var docModel registerModels.DOCModel

	for _, Id := range body.Ids {
		currentDoc, _ := docModel.GetById(Id)
		if currentDoc.IsEmpty() {
			continue
		}

		if currentDoc.DbStatus != "active" {
			continue
		}

		docModel.Update(Id, map[string]interface{}{
			"dbStatus":     "archived",
			"dbLastStatus": currentDoc.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*DOCHandler) Unarchive(c *gin.Context) {
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

	var docModel registerModels.DOCModel

	for _, Id := range body.Ids {
		currentDoc, _ := docModel.GetById(Id)
		if currentDoc.IsEmpty() {
			continue
		}

		if currentDoc.DbStatus != "archived" {
			continue
		}

		docModel.Update(Id, map[string]interface{}{
			"dbStatus":     "active",
			"dbLastStatus": currentDoc.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*DOCHandler) Delete(c *gin.Context) {
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

	var docModel registerModels.DOCModel

	for _, Id := range body.Ids {
		currentDoc, _ := docModel.GetById(Id)
		if currentDoc.IsEmpty() {
			continue
		}

		if currentDoc.DbStatus == "deleted" {
			continue
		}

		docModel.Update(Id, map[string]interface{}{
			"dbStatus":     "deleted",
			"dbLastStatus": currentDoc.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*DOCHandler) Undelete(c *gin.Context) {
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

	var docModel registerModels.DOCModel

	for _, Id := range body.Ids {
		currentDoc, _ := docModel.GetById(Id)
		if currentDoc.IsEmpty() {
			continue
		}

		if currentDoc.DbStatus != "deleted" {
			continue
		}

		docModel.Update(Id, map[string]interface{}{
			"dbStatus":     currentDoc.DbLastStatus,
			"dbLastStatus": currentDoc.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}
