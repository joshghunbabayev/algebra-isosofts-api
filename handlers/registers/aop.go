package registerHandlers

import (
	"algebra-isosofts-api/middlewares"
	registerModels "algebra-isosofts-api/models/registers"
	registerTypes "algebra-isosofts-api/types/registers"
	tableComponentTypes "algebra-isosofts-api/types/tableComponents"

	"github.com/gin-gonic/gin"
)

type AOPHandler struct {
}

func (*AOPHandler) GetAll(c *gin.Context) {
	account, _ := c.MustGet("account").(middlewares.RemoteAccount)
	status := c.Query("status")

	if status == "" {
		status = "active"
	}

	var aopModel registerModels.AOPModel

	aops, err := aopModel.GetAll(map[string]interface{}{
		"dbStatus":  status,
		"companyId": account.CompanyId,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, aops)
}

func (*AOPHandler) Create(c *gin.Context) {
	account, _ := c.MustGet("account").(middlewares.RemoteAccount)
	var body struct {
		ActivityDescription string `json:"activityDescription"`
		AuditorInspector    string `json:"auditorInspector"`
		AuditeeInspectee    string `json:"auditeeInspectee"`
		ReviewedPremises    string `json:"reviewedPremises"`
		ReviewedProcess     string `json:"reviewedProcess"`
		RTIC                string `json:"rtic"`
		Frequency           string `json:"frequency"`
		AOADate             string `json:"aoaDate"`
		InspectionFrequency string `json:"inspectionFrequency"`
		NextAoaDate         string `json:"nextAoaDate"`
		AOAStatus           string `json:"aoaStatus"`
		Comment             string `json:"comment"`
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

	var aopModel registerModels.AOPModel

	aopModel.Create(registerTypes.AOP{
		Id:        aopModel.GenerateUniqueId(),
		CompanyId: account.CompanyId,
		No:        aopModel.GenerateUniqueNo(),
		ActivityDescription: tableComponentTypes.DropDownListItem{
			Id: body.ActivityDescription,
		},
		AuditorInspector: body.AuditorInspector,
		AuditeeInspectee: body.AuditeeInspectee,
		ReviewedPremises: tableComponentTypes.DropDownListItem{
			Id: body.ReviewedPremises,
		},
		ReviewedProcess: tableComponentTypes.DropDownListItem{
			Id: body.ReviewedProcess,
		},
		RTIC:      body.RTIC,
		Frequency: body.Frequency,
		AOADate:   body.AOADate,
		InspectionFrequency: tableComponentTypes.DropDownListItem{
			Id: body.InspectionFrequency,
		},
		NextAoaDate:  body.NextAoaDate,
		AOAStatus:    body.AOAStatus,
		Comment:      body.Comment,
		DbStatus:     "active",
		DbLastStatus: "active",
	})

	c.IndentedJSON(201, gin.H{})
}

func (*AOPHandler) Update(c *gin.Context) {
	account, _ := c.MustGet("account").(middlewares.RemoteAccount)
	Id := c.Param("id")

	var aopModel registerModels.AOPModel

	currentAop, _ := aopModel.GetById(Id)

	if currentAop.IsEmpty() || currentAop.CompanyId != account.CompanyId {
		c.IndentedJSON(404, gin.H{})
		return
	}

	var body struct {
		ActivityDescription string `json:"activityDescription"`
		AuditorInspector    string `json:"auditorInspector"`
		AuditeeInspectee    string `json:"auditeeInspectee"`
		ReviewedPremises    string `json:"reviewedPremises"`
		ReviewedProcess     string `json:"reviewedProcess"`
		RTIC                string `json:"rtic"`
		Frequency           string `json:"frequency"`
		AOADate             string `json:"aoaDate"`
		InspectionFrequency string `json:"inspectionFrequency"`
		NextAoaDate         string `json:"nextAoaDate"`
		AOAStatus           string `json:"aoaStatus"`
		Comment             string `json:"comment"`
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

	aopModel.Update(Id, map[string]interface{}{
		"activityDescription": body.ActivityDescription,
		"auditorInspector":    body.AuditorInspector,
		"auditeeInspectee":    body.AuditeeInspectee,
		"reviewedPremises":    body.ReviewedPremises,
		"reviewedProcess":     body.ReviewedProcess,
		"rtic":                body.RTIC,
		"frequency":           body.Frequency,
		"aoaDate":             body.AOADate,
		"inspectionFrequency": body.InspectionFrequency,
		"nextAoaDate":         body.NextAoaDate,
		"aoaStatus":           body.AOAStatus,
		"comment":             body.Comment,
	})

	c.JSON(200, gin.H{})
}

func (*AOPHandler) Archive(c *gin.Context) {
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

	var aopModel registerModels.AOPModel

	for _, Id := range body.Ids {
		currentAop, _ := aopModel.GetById(Id)
		if currentAop.IsEmpty() || currentAop.CompanyId != account.CompanyId {
			continue
		}

		if currentAop.DbStatus != "active" {
			continue
		}

		aopModel.Update(Id, map[string]interface{}{
			"dbStatus":     "archived",
			"dbLastStatus": currentAop.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*AOPHandler) Unarchive(c *gin.Context) {
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

	var aopModel registerModels.AOPModel

	for _, Id := range body.Ids {
		currentAop, _ := aopModel.GetById(Id)
		if currentAop.IsEmpty() || currentAop.CompanyId != account.CompanyId {
			continue
		}

		if currentAop.DbStatus != "archived" {
			continue
		}

		aopModel.Update(Id, map[string]interface{}{
			"dbStatus":     "active",
			"dbLastStatus": currentAop.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*AOPHandler) Delete(c *gin.Context) {
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

	var aopModel registerModels.AOPModel

	for _, Id := range body.Ids {
		currentAop, _ := aopModel.GetById(Id)
		if currentAop.IsEmpty() || currentAop.CompanyId != account.CompanyId {
			continue
		}

		if currentAop.DbStatus == "deleted" {
			continue
		}

		aopModel.Update(Id, map[string]interface{}{
			"dbStatus":     "deleted",
			"dbLastStatus": currentAop.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*AOPHandler) Undelete(c *gin.Context) {
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

	var aopModel registerModels.AOPModel

	for _, Id := range body.Ids {
		currentAop, _ := aopModel.GetById(Id)
		if currentAop.IsEmpty() || currentAop.CompanyId != account.CompanyId {
			continue
		}

		if currentAop.DbStatus != "deleted" {
			continue
		}

		aopModel.Update(Id, map[string]interface{}{
			"dbStatus":     currentAop.DbLastStatus,
			"dbLastStatus": currentAop.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}
