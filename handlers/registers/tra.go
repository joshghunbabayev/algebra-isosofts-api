package registerHandlers

import (
	"algebra-isosofts-api/middlewares"
	registerModels "algebra-isosofts-api/models/registers"
	registerTypes "algebra-isosofts-api/types/registers"

	"github.com/gin-gonic/gin"
)

type TRAHandler struct {
}

func (*TRAHandler) GetAll(c *gin.Context) {
	account, _ := c.MustGet("account").(middlewares.RemoteAccount)
	status := c.Query("status")

	if status == "" {
		status = "active"
	}

	var traModel registerModels.TRAModel

	tras, err := traModel.GetAll(map[string]interface{}{
		"dbStatus":  status,
		"companyId": account.CompanyId,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, tras)
}

func (*TRAHandler) Create(c *gin.Context) {
	var body struct {
		EmployeeName   string `json:"employeeName"`
		Position       string `json:"position"`
		TCLN           string `json:"tcln"`
		TCLID          string `json:"nvcd"`
		CLNumber       string `json:"clnumber"`
		NCD            string `json:"ncd"`
		ValidityStatus int8   `json:"validityStatus"`
		Effectiveness  int8   `json:"effectiveness"`
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

	account, _ := c.MustGet("account").(middlewares.RemoteAccount)

	traModel.Create(registerTypes.TRA{
		Id:             traModel.GenerateUniqueId(),
		CompanyId:      account.CompanyId,
		No:             traModel.GenerateUniqueNo(),
		EmployeeName:   body.EmployeeName,
		Position:       body.Position,
		TCLN:           body.TCLN,
		TCLID:          body.TCLID,
		CLNumber:       body.CLNumber,
		NCD:            body.NCD,
		ValidityStatus: body.ValidityStatus,
		Effectiveness:  body.Effectiveness,
		DbStatus:       "active",
		DbLastStatus:   "active",
	})

	c.IndentedJSON(201, gin.H{})
}

func (*TRAHandler) Update(c *gin.Context) {
	account, _ := c.MustGet("account").(middlewares.RemoteAccount)
	Id := c.Param("id")

	var traModel registerModels.TRAModel

	currentTra, _ := traModel.GetById(Id)

	if currentTra.IsEmpty() || currentTra.CompanyId != account.CompanyId {
		c.IndentedJSON(404, gin.H{})
		return
	}

	var body struct {
		EmployeeName   string `json:"employeeName"`
		Position       string `json:"position"`
		TCLN           string `json:"tcln"`
		TCLID          string `json:"nvcd"`
		CLNumber       string `json:"clnumber"`
		NCD            string `json:"ncd"`
		ValidityStatus int8   `json:"validityStatus"`
		Effectiveness  int8   `json:"effectiveness"`
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
		"employeeName":   body.EmployeeName,
		"position":       body.Position,
		"tcln":           body.TCLN,
		"nvcd":           body.TCLID,
		"clnumber":       body.CLNumber,
		"ncd":            body.NCD,
		"validityStatus": body.ValidityStatus,
		"effectiveness":  body.Effectiveness,
	})

	c.JSON(200, gin.H{})
}

func (*TRAHandler) Archive(c *gin.Context) {
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

	var traModel registerModels.TRAModel

	for _, Id := range body.Ids {
		currentTra, _ := traModel.GetById(Id)
		if currentTra.IsEmpty() || currentTra.CompanyId != account.CompanyId {
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

	var traModel registerModels.TRAModel

	for _, Id := range body.Ids {
		currentTra, _ := traModel.GetById(Id)
		if currentTra.IsEmpty() || currentTra.CompanyId != account.CompanyId {
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

	var traModel registerModels.TRAModel

	for _, Id := range body.Ids {
		currentTra, _ := traModel.GetById(Id)
		if currentTra.IsEmpty() || currentTra.CompanyId != account.CompanyId {
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

	var traModel registerModels.TRAModel

	for _, Id := range body.Ids {
		currentTra, _ := traModel.GetById(Id)
		if currentTra.IsEmpty() || currentTra.CompanyId != account.CompanyId {
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
