package registerHandlers

import (
	"algebra-isosofts-api/middlewares"
	registerModels "algebra-isosofts-api/models/registers"
	registerTypes "algebra-isosofts-api/types/registers"
	tableComponentTypes "algebra-isosofts-api/types/tableComponents"

	"github.com/gin-gonic/gin"
)

type FINHandler struct {
}

func (*FINHandler) GetAll(c *gin.Context) {
	account, _ := c.MustGet("account").(middlewares.RemoteAccount)
	status := c.Query("status")

	if status == "" {
		status = "active"
	}

	var finModel registerModels.FINModel

	fins, err := finModel.GetAll(map[string]interface{}{
		"dbStatus":  status,
		"companyId": account.CompanyId,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, fins)
}

func (*FINHandler) Create(c *gin.Context) {
	account, _ := c.MustGet("account").(middlewares.RemoteAccount)
	var body struct {
		Issuer            string `json:"issuer"`
		FindingDate       string `json:"findingDate"`
		JobNumber         string `json:"jobNumber"`
		Process           string `json:"process"`
		CategoryOfFinding string `json:"categoryOfFinding"`
		TypeOfFinding     string `json:"typeOfFinding"`
		SourceOfFinding   string `json:"sourceOfFinding"`
		CustomerId        string `json:"customerId"`
		VendorId          string `json:"vendorId"`
		Description       string `json:"description"`
		ContainmentAction string `json:"containmentAction"`
		RootCauses        string `json:"rootCauses"`
		FindingStatus     string `json:"findingStatus"`
		Comment           string `json:"comment"`
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

	var finModel registerModels.FINModel

	finModel.Create(registerTypes.FIN{
		Id:          finModel.GenerateUniqueId(),
		CompanyId:   account.CompanyId,
		No:          finModel.GenerateUniqueNo(),
		Issuer:      body.Issuer,
		FindingDate: body.FindingDate,
		JobNumber:   body.JobNumber,
		Process: tableComponentTypes.DropDownListItem{
			Id: body.Process,
		},
		CategoryOfFinding: tableComponentTypes.DropDownListItem{
			Id: body.CategoryOfFinding,
		},
		TypeOfFinding: tableComponentTypes.DropDownListItem{
			Id: body.TypeOfFinding,
		},
		SourceOfFinding: tableComponentTypes.DropDownListItem{
			Id: body.SourceOfFinding,
		},
		CustomerId:        body.CustomerId,
		VendorId:          body.VendorId,
		Description:       body.Description,
		ContainmentAction: body.ContainmentAction,
		RootCauses:        body.RootCauses,
		FindingStatus: tableComponentTypes.DropDownListItem{
			Id: body.FindingStatus,
		},
		Comment:      body.Comment,
		DbStatus:     "active",
		DbLastStatus: "active",
	})

	c.IndentedJSON(201, gin.H{})
}

func (*FINHandler) Update(c *gin.Context) {
	account, _ := c.MustGet("account").(middlewares.RemoteAccount)
	Id := c.Param("id")

	var finModel registerModels.FINModel

	currentFin, _ := finModel.GetById(Id)

	if currentFin.IsEmpty() || currentFin.CompanyId != account.CompanyId {
		c.IndentedJSON(404, gin.H{})
		return
	}

	var body struct {
		Issuer            string `json:"issuer"`
		FindingDate       string `json:"findingDate"`
		JobNumber         string `json:"jobNumber"`
		Process           string `json:"process"`
		CategoryOfFinding string `json:"categoryOfFinding"`
		TypeOfFinding     string `json:"typeOfFinding"`
		SourceOfFinding   string `json:"sourceOfFinding"`
		CustomerId        string `json:"customerId"`
		VendorId          string `json:"vendorId"`
		Description       string `json:"description"`
		ContainmentAction string `json:"containmentAction"`
		RootCauses        string `json:"rootCauses"`
		FindingStatus     string `json:"findingStatus"`
		Comment           string `json:"comment"`
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

	finModel.Update(Id, map[string]interface{}{
		"issuer":            body.Issuer,
		"findingDate":       body.FindingDate,
		"jobNumber":         body.JobNumber,
		"process":           body.Process,
		"categoryOfFinding": body.CategoryOfFinding,
		"typeOfFinding":     body.TypeOfFinding,
		"sourceOfFinding":   body.SourceOfFinding,
		"customerId":        body.CustomerId,
		"vendorId":          body.VendorId,
		"description":       body.Description,
		"containmentAction": body.ContainmentAction,
		"rootCauses":        body.RootCauses,
		"findingStatus":     body.FindingStatus,
		"comment":           body.Comment,
	})

	c.JSON(200, gin.H{})
}

func (*FINHandler) Archive(c *gin.Context) {
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

	var finModel registerModels.FINModel

	for _, Id := range body.Ids {
		currentFin, _ := finModel.GetById(Id)
		if currentFin.IsEmpty() || currentFin.CompanyId != account.CompanyId {
			continue
		}

		if currentFin.DbStatus != "active" {
			continue
		}

		finModel.Update(Id, map[string]interface{}{
			"dbStatus":     "archived",
			"dbLastStatus": currentFin.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*FINHandler) Unarchive(c *gin.Context) {
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

	var finModel registerModels.FINModel

	for _, Id := range body.Ids {
		currentFin, _ := finModel.GetById(Id)
		if currentFin.IsEmpty() || currentFin.CompanyId != account.CompanyId {
			continue
		}

		if currentFin.DbStatus != "archived" {
			continue
		}

		finModel.Update(Id, map[string]interface{}{
			"dbStatus":     "active",
			"dbLastStatus": currentFin.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*FINHandler) Delete(c *gin.Context) {
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

	var finModel registerModels.FINModel

	for _, Id := range body.Ids {
		currentFin, _ := finModel.GetById(Id)
		if currentFin.IsEmpty() || currentFin.CompanyId != account.CompanyId {
			continue
		}

		if currentFin.DbStatus == "deleted" {
			continue
		}

		finModel.Update(Id, map[string]interface{}{
			"dbStatus":     "deleted",
			"dbLastStatus": currentFin.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*FINHandler) Undelete(c *gin.Context) {
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

	var finModel registerModels.FINModel

	for _, Id := range body.Ids {
		currentFin, _ := finModel.GetById(Id)
		if currentFin.IsEmpty() || currentFin.CompanyId != account.CompanyId {
			continue
		}

		if currentFin.DbStatus != "deleted" {
			continue
		}

		finModel.Update(Id, map[string]interface{}{
			"dbStatus":     currentFin.DbLastStatus,
			"dbLastStatus": currentFin.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}
