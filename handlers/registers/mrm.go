package registerHandlers

import (
	"algebra-isosofts-api/middlewares"
	registerModels "algebra-isosofts-api/models/registers"
	registerTypes "algebra-isosofts-api/types/registers"
	tableComponentTypes "algebra-isosofts-api/types/tableComponents"

	"github.com/gin-gonic/gin"
)

type MRMHandler struct {
}

func (*MRMHandler) GetAll(c *gin.Context) {
	account, _ := c.MustGet("account").(middlewares.RemoteAccount)
	status := c.Query("status")

	if status == "" {
		status = "active"
	}

	var mrmModel registerModels.MRMModel

	mrms, err := mrmModel.GetAll(map[string]interface{}{
		"dbStatus":  status,
		"companyId": account.CompanyId,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, mrms)
}

func (*MRMHandler) Create(c *gin.Context) {
	var body struct {
		Topic   string `json:"topic"`
		RISOS   string `json:"risos"`
		Process string `json:"process"`
		Comment string `json:"comment"`
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

	account, _ := c.MustGet("account").(middlewares.RemoteAccount)

	mrmModel.Create(registerTypes.MRM{
		Id:        mrmModel.GenerateUniqueId(),
		CompanyId: account.CompanyId,
		No:        mrmModel.GenerateUniqueNo(account.CompanyId),
		Topic: tableComponentTypes.DropDownListItem{
			Id: body.Topic,
		},
		RISOS: body.RISOS,
		Process: tableComponentTypes.DropDownListItem{
			Id: body.Process,
		},
		Comment:      body.Comment,
		DbStatus:     "active",
		DbLastStatus: "active",
	})

	c.IndentedJSON(201, gin.H{})
}

func (*MRMHandler) Update(c *gin.Context) {
	account, _ := c.MustGet("account").(middlewares.RemoteAccount)
	Id := c.Param("id")

	var mrmModel registerModels.MRMModel

	currentMrm, _ := mrmModel.GetById(Id)

	if currentMrm.IsEmpty() || currentMrm.CompanyId != account.CompanyId {
		c.IndentedJSON(404, gin.H{})
		return
	}

	var body struct {
		Topic   string `json:"topic"`
		RISOS   string `json:"risos"`
		Process string `json:"process"`
		Comment string `json:"comment"`
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
		"topic":   body.Topic,
		"risos":   body.RISOS,
		"process": body.Process,
		"comment": body.Comment,
	})

	c.JSON(200, gin.H{})
}

func (*MRMHandler) Archive(c *gin.Context) {
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

	var mrmModel registerModels.MRMModel

	for _, Id := range body.Ids {
		currentMrm, _ := mrmModel.GetById(Id)
		if currentMrm.IsEmpty() || currentMrm.CompanyId != account.CompanyId {
			continue
		}

		if currentMrm.DbStatus != "active" {
			continue
		}

		mrmModel.Update(Id, map[string]interface{}{
			"dbStatus":     "archived",
			"dbLastStatus": currentMrm.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*MRMHandler) Unarchive(c *gin.Context) {
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

	var mrmModel registerModels.MRMModel

	for _, Id := range body.Ids {
		currentMrm, _ := mrmModel.GetById(Id)
		if currentMrm.IsEmpty() || currentMrm.CompanyId != account.CompanyId {
			continue
		}

		if currentMrm.DbStatus != "archived" {
			continue
		}

		mrmModel.Update(Id, map[string]interface{}{
			"dbStatus":     "active",
			"dbLastStatus": currentMrm.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*MRMHandler) Delete(c *gin.Context) {
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

	var mrmModel registerModels.MRMModel

	for _, Id := range body.Ids {
		currentMrm, _ := mrmModel.GetById(Id)
		if currentMrm.IsEmpty() || currentMrm.CompanyId != account.CompanyId {
			continue
		}

		if currentMrm.DbStatus == "deleted" {
			continue
		}

		mrmModel.Update(Id, map[string]interface{}{
			"dbStatus":     "deleted",
			"dbLastStatus": currentMrm.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*MRMHandler) Undelete(c *gin.Context) {
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

	var mrmModel registerModels.MRMModel

	for _, Id := range body.Ids {
		currentMrm, _ := mrmModel.GetById(Id)
		if currentMrm.IsEmpty() || currentMrm.CompanyId != account.CompanyId {
			continue
		}

		if currentMrm.DbStatus != "deleted" {
			continue
		}

		mrmModel.Update(Id, map[string]interface{}{
			"dbStatus":     currentMrm.DbLastStatus,
			"dbLastStatus": currentMrm.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}
