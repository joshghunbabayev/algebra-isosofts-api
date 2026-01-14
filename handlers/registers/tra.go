package registerHandlers

import (
	registerModels "algebra-isosofts-api/models/registers"
	registerTypes "algebra-isosofts-api/types/registers"

	"github.com/gin-gonic/gin"
)

type TRAHandler struct {
}

func (*TRAHandler) GetAll(c *gin.Context) {
	status := c.Query("status")

	if status == "" {
		status = "active"
	}

	var traModel registerModels.TRAModel

	tras, err := traModel.GetAll(map[string]interface{}{
		"dbStatus": status,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, tras)
}

func (*TRAHandler) Create(c *gin.Context) {
	var body struct {
		EmployeeName     string `json:"employeeName"`
		Position         string `json:"position"`
		CLName           string `json:"clname"`
		TCLID            string `json:"nvcd"`
		CLNumber         string `json:"clnumber"`
		NCD              string `json:"ncd"`
		CompetencyStatus int8   `json:"competencyStatus"`
		Effectiveness    string `json:"effectiveness"`
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

	var traModel registerModels.TRAModel

	traModel.Create(registerTypes.TRA{
		Id:               traModel.GenerateUniqueId(),
		No:               traModel.GenerateUniqueNo(),
		EmployeeName:     body.EmployeeName,
		Position:         body.Position,
		CLName:           body.CLName,
		TCLID:            body.TCLID,
		CLNumber:         body.CLNumber,
		NCD:              body.NCD,
		CompetencyStatus: body.CompetencyStatus,
		Effectiveness:    body.Effectiveness,
		DbStatus:         "active",
		DbLastStatus:     "active",
	})

	c.IndentedJSON(201, gin.H{})
}

func (*TRAHandler) Update(c *gin.Context) {
	Id := c.Param("id")

	var traModel registerModels.TRAModel

	currentTra, _ := traModel.GetById(Id)

	if currentTra.IsEmpty() {
		c.IndentedJSON(404, gin.H{})
		return
	}

	var body struct {
		EmployeeName     string `json:"employeeName"`
		Position         string `json:"position"`
		CLName           string `json:"clname"`
		TCLID            string `json:"nvcd"`
		CLNumber         string `json:"clnumber"`
		NCD              string `json:"ncd"`
		CompetencyStatus int8   `json:"competencyStatus"`
		Effectiveness    string `json:"effectiveness"`
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

	traModel.Update(Id, map[string]interface{}{
		"employeeName":     body.EmployeeName,
		"position":         body.Position,
		"clname":           body.CLName,
		"nvcd":             body.TCLID,
		"clnumber":         body.CLNumber,
		"ncd":              body.NCD,
		"competencyStatus": body.CompetencyStatus,
		"effectiveness":    body.Effectiveness,
	})

	c.JSON(200, gin.H{})
}

func (*TRAHandler) Archive(c *gin.Context) {
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

	var traModel registerModels.TRAModel

	for _, Id := range body.Ids {
		currentTra, _ := traModel.GetById(Id)
		if currentTra.IsEmpty() {
			continue
		}

		if currentTra.DbStatus != "active" {
			continue
		}

		traModel.Update(Id, map[string]interface{}{
			"dbStatus":     "archived",
			"dbLastStatus": currentTra.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*TRAHandler) Unarchive(c *gin.Context) {
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

	var traModel registerModels.TRAModel

	for _, Id := range body.Ids {
		currentTra, _ := traModel.GetById(Id)
		if currentTra.IsEmpty() {
			continue
		}

		if currentTra.DbStatus != "archived" {
			continue
		}

		traModel.Update(Id, map[string]interface{}{
			"dbStatus":     "active",
			"dbLastStatus": currentTra.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*TRAHandler) Delete(c *gin.Context) {
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

	var traModel registerModels.TRAModel

	for _, Id := range body.Ids {
		currentTra, _ := traModel.GetById(Id)
		if currentTra.IsEmpty() {
			continue
		}

		if currentTra.DbStatus == "deleted" {
			continue
		}

		traModel.Update(Id, map[string]interface{}{
			"dbStatus":     "deleted",
			"dbLastStatus": currentTra.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*TRAHandler) Undelete(c *gin.Context) {
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

	var traModel registerModels.TRAModel

	for _, Id := range body.Ids {
		currentTra, _ := traModel.GetById(Id)
		if currentTra.IsEmpty() {
			continue
		}

		if currentTra.DbStatus != "deleted" {
			continue
		}

		traModel.Update(Id, map[string]interface{}{
			"dbStatus":     currentTra.DbLastStatus,
			"dbLastStatus": currentTra.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}
