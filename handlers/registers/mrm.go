package registerHandlers

import (
	registerModels "algebra-isosofts-api/models/registers"
	registerTypes "algebra-isosofts-api/types/registers"
	tableComponentTypes "algebra-isosofts-api/types/tableComponents"

	"github.com/gin-gonic/gin"
)

type MRMHandler struct {
}

func (*MRMHandler) GetAll(c *gin.Context) {
	status := c.Query("status")

	if status == "" {
		status = "active"
	}

	var mrmModel registerModels.MRMModel

	mrms, err := mrmModel.GetAll(map[string]interface{}{
		"dbStatus": status,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, mrms)
}

func (*MRMHandler) Create(c *gin.Context) {
	var body struct {
		RISOS   string `json:"risos"`
		Topic   string `json:"topic"`
		Process string `json:"process"`
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

	var mrmModel registerModels.MRMModel

	mrmModel.Create(registerTypes.MRM{
		Id: mrmModel.GenerateUniqueId(),
		No: mrmModel.GenerateUniqueNo(),
		RISOS: tableComponentTypes.DropDownListItem{
			Id: body.RISOS,
		},
		Topic: tableComponentTypes.DropDownListItem{
			Id: body.Topic,
		},
		Process: tableComponentTypes.DropDownListItem{
			Id: body.Process,
		},
		DbStatus:     "active",
		DbLastStatus: "active",
	})

	c.IndentedJSON(201, gin.H{})
}

func (*MRMHandler) Update(c *gin.Context) {
	Id := c.Param("id")

	var mrmModel registerModels.MRMModel

	currentMRM, _ := mrmModel.GetById(Id)

	if currentMRM.IsEmpty() {
		c.IndentedJSON(404, gin.H{})
		return
	}

	var body struct {
		RISOS   string `json:"risos"`
		Topic   string `json:"topic"`
		Process string `json:"process"`
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

	mrmModel.Update(Id, map[string]interface{}{
		"risos":   body.RISOS,
		"topic":   body.Topic,
		"process": body.Process,
	})

	c.JSON(200, gin.H{})
}

func (*MRMHandler) Archive(c *gin.Context) {
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

	var mrmModel registerModels.MRMModel

	for _, Id := range body.Ids {
		currentMRM, _ := mrmModel.GetById(Id)
		if currentMRM.IsEmpty() {
			continue
		}

		if currentMRM.DbStatus != "active" {
			continue
		}

		mrmModel.Update(Id, map[string]interface{}{
			"dbStatus":     "archived",
			"dbLastStatus": currentMRM.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*MRMHandler) Unarchive(c *gin.Context) {
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

	var mrmModel registerModels.MRMModel

	for _, Id := range body.Ids {
		currentMRM, _ := mrmModel.GetById(Id)
		if currentMRM.IsEmpty() {
			continue
		}

		if currentMRM.DbStatus != "archived" {
			continue
		}

		mrmModel.Update(Id, map[string]interface{}{
			"dbStatus":     "active",
			"dbLastStatus": currentMRM.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*MRMHandler) Delete(c *gin.Context) {
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

	var mrmModel registerModels.MRMModel

	for _, Id := range body.Ids {
		currentMRM, _ := mrmModel.GetById(Id)
		if currentMRM.IsEmpty() {
			continue
		}

		if currentMRM.DbStatus == "deleted" {
			continue
		}

		mrmModel.Update(Id, map[string]interface{}{
			"dbStatus":     "deleted",
			"dbLastStatus": currentMRM.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*MRMHandler) Undelete(c *gin.Context) {
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

	var mrmModel registerModels.MRMModel

	for _, Id := range body.Ids {
		currentMRM, _ := mrmModel.GetById(Id)
		if currentMRM.IsEmpty() {
			continue
		}

		if currentMRM.DbStatus != "deleted" {
			continue
		}

		mrmModel.Update(Id, map[string]interface{}{
			"dbStatus":     currentMRM.DbLastStatus,
			"dbLastStatus": currentMRM.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}
