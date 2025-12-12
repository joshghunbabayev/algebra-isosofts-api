package registerComponentHandlers

import (
	registerModels "algebra-isosofts-api/models/registers"
	registerComponentModels "algebra-isosofts-api/models/registers/components"
	registerComponentTypes "algebra-isosofts-api/types/registers/components"
	tableComponentTypes "algebra-isosofts-api/types/tableComponents"

	"github.com/gin-gonic/gin"
)

type ActionHandler struct {
}

func (*ActionHandler) GetAll(c *gin.Context) {
	registerId := c.Query("registerId")
	status := c.Query("status")

	if status == "" {
		status = "active"
	}

	var actionModel registerComponentModels.ActionModel
	var actions []registerComponentTypes.Action
	var err error

	if registerId == "" {
		actions, err = actionModel.GetAll(map[string]interface{}{
			"dbStatus": status,
		})
	} else {
		actions, err = actionModel.GetAll(map[string]interface{}{
			"registerId": registerId,
			"dbStatus":   status,
		})
	}

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, actions)
}

func (*ActionHandler) Create(c *gin.Context) {
	var body struct {
		RegisterId         string `json:"registerId"`
		Title              string `json:"title"`
		RaiseDate          string `json:"raiseDate"`
		Resources          int64  `json:"resources"`
		Currency           string `json:"currency"`
		RelativeFunction   string `json:"relativeFunction"`
		Responsible        string `json:"responsible"`
		Deadline           string `json:"deadline"`
		Confirmation       string `json:"confirmation"`
		Status             string `json:"status"`
		CompletionDate     string `json:"completionDate"`
		VerificationStatus string `json:"verificationStatus"`
		Comment            string `json:"comment"`
		January            string `json:"january"`
		February           string `json:"february"`
		March              string `json:"march"`
		April              string `json:"april"`
		May                string `json:"may"`
		June               string `json:"june"`
		July               string `json:"july"`
		August             string `json:"august"`
		September          string `json:"september"`
		October            string `json:"october"`
		November           string `json:"november"`
		December           string `json:"december"`
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

	var brModel registerModels.BRModel

	br, _ := brModel.GetById(body.RegisterId)

	if br.IsEmpty() {
		c.IndentedJSON(404, gin.H{})
		return
	}

	if br.DbStatus != "active" {
		c.IndentedJSON(400, gin.H{})
		return
	}

	var actionModel registerComponentModels.ActionModel

	actionModel.Create(registerComponentTypes.Action{
		Id:         actionModel.GenerateUniqueId(),
		RegisterId: body.RegisterId,
		Title:      body.Title,
		RaiseDate:  body.RaiseDate,
		Resources:  body.Resources,
		Currency:   body.Currency,
		RelativeFunction: tableComponentTypes.DropDownListItem{
			Id: body.RelativeFunction,
		},
		Responsible: tableComponentTypes.DropDownListItem{
			Id: body.Responsible,
		},
		Deadline: body.Deadline,
		Confirmation: tableComponentTypes.DropDownListItem{
			Id: body.Confirmation,
		},
		Status: tableComponentTypes.DropDownListItem{
			Id: body.Status,
		},
		CompletionDate: body.CompletionDate,
		VerificationStatus: tableComponentTypes.DropDownListItem{
			Id: body.VerificationStatus,
		},
		Comment: body.Comment,
		January: tableComponentTypes.DropDownListItem{
			Id: body.January,
		},
		February: tableComponentTypes.DropDownListItem{
			Id: body.February,
		},
		March: tableComponentTypes.DropDownListItem{
			Id: body.March,
		},
		April: tableComponentTypes.DropDownListItem{
			Id: body.April,
		},
		May: tableComponentTypes.DropDownListItem{
			Id: body.May,
		},
		June: tableComponentTypes.DropDownListItem{
			Id: body.June,
		},
		July: tableComponentTypes.DropDownListItem{
			Id: body.July,
		},
		August: tableComponentTypes.DropDownListItem{
			Id: body.August,
		},
		September: tableComponentTypes.DropDownListItem{
			Id: body.September,
		},
		October: tableComponentTypes.DropDownListItem{
			Id: body.October,
		},
		November: tableComponentTypes.DropDownListItem{
			Id: body.November,
		},
		December: tableComponentTypes.DropDownListItem{
			Id: body.December,
		},
		DbStatus:     "active",
		DbLastStatus: "active",
	})

	c.IndentedJSON(201, gin.H{})
}

func (*ActionHandler) Update(c *gin.Context) {
	Id := c.Param("id")

	var actionModel registerComponentModels.ActionModel

	currentAction, _ := actionModel.GetById(Id)

	if currentAction.IsEmpty() {
		c.IndentedJSON(404, gin.H{})
		return
	}

	var body struct {
		Title              string `json:"title"`
		RaiseDate          string `json:"raiseDate"`
		Resources          int64  `json:"resources"`
		Currency           string `json:"currency"`
		RelativeFunction   string `json:"relativeFunction"`
		Responsible        string `json:"responsible"`
		Deadline           string `json:"deadline"`
		Confirmation       string `json:"confirmation"`
		Status             string `json:"status"`
		CompletionDate     string `json:"completionDate"`
		VerificationStatus string `json:"verificationStatus"`
		Comment            string `json:"comment"`
		January            string `json:"january"`
		February           string `json:"february"`
		March              string `json:"march"`
		April              string `json:"april"`
		May                string `json:"may"`
		June               string `json:"june"`
		July               string `json:"july"`
		August             string `json:"august"`
		September          string `json:"september"`
		October            string `json:"october"`
		November           string `json:"november"`
		December           string `json:"december"`
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

	actionModel.Update(Id, map[string]interface{}{
		"title":              body.Title,
		"raiseDate":          body.RaiseDate,
		"resources":          body.Resources,
		"currency":           body.Currency,
		"relativeFunction":   body.RelativeFunction,
		"responsible":        body.Responsible,
		"deadline":           body.Deadline,
		"confirmation":       body.Confirmation,
		"status":             body.Status,
		"completionDate":     body.CompletionDate,
		"verificationStatus": body.VerificationStatus,
		"comment":            body.Comment,
		"january":            body.January,
		"february":           body.February,
		"march":              body.March,
		"april":              body.April,
		"may":                body.May,
		"june":               body.June,
		"july":               body.July,
		"august":             body.August,
		"september":          body.September,
		"october":            body.October,
		"november":           body.November,
		"december":           body.December,
	})

	c.JSON(200, gin.H{})
}

func (*ActionHandler) Delete(c *gin.Context) {
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

	var actionModel registerComponentModels.ActionModel

	for _, Id := range body.Ids {
		currentAction, _ := actionModel.GetById(Id)
		if currentAction.IsEmpty() {
			continue
		}

		if currentAction.DbStatus == "deleted" {
			continue
		}

		actionModel.Update(Id, map[string]interface{}{
			"dbStatus":     "deleted",
			"dbLastStatus": currentAction.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*ActionHandler) Undelete(c *gin.Context) {
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

	var actionModel registerComponentModels.ActionModel

	for _, Id := range body.Ids {
		currentAction, _ := actionModel.GetById(Id)
		if currentAction.IsEmpty() {
			continue
		}

		if currentAction.DbStatus != "deleted" {
			continue
		}

		actionModel.Update(Id, map[string]interface{}{
			"dbStatus":     currentAction.DbLastStatus,
			"dbLastStatus": currentAction.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}
