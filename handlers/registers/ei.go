package registerHandlers

import (
	"algebra-isosofts-api/middlewares"
	registerModels "algebra-isosofts-api/models/registers"
	registerTypes "algebra-isosofts-api/types/registers"
	tableComponentTypes "algebra-isosofts-api/types/tableComponents"

	"github.com/gin-gonic/gin"
)

type EIHandler struct {
}

func (*EIHandler) GetAll(c *gin.Context) {
	account, _ := c.MustGet("account").(middlewares.RemoteAccount)
	status := c.Query("status")

	if status == "" {
		status = "active"
	}

	var eiModel registerModels.EIModel

	eis, err := eiModel.GetAll(map[string]interface{}{
		"dbStatus":  status,
		"companyId": account.CompanyId,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, eis)
}

func (*EIHandler) Create(c *gin.Context) {
	var body struct {
		Name                string `json:"name"`
		Type                string `json:"type"`
		SerialNumber        string `json:"serialNumber"`
		CertificateNo       string `json:"certificateNo"`
		CalibrationRequired int8   `json:"calibrationRequired"`
		InspectionFrequency string `json:"inspectionFrequency"`
		ICD                 string `json:"icd"`
		NVCD                string `json:"nvcd"`
		EIS                 int8   `json:"eis"`
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

	var eiModel registerModels.EIModel

	account, _ := c.MustGet("account").(middlewares.RemoteAccount)

	eiModel.Create(registerTypes.EI{
		Id:        eiModel.GenerateUniqueId(),
		CompanyId: account.CompanyId,
		No:        eiModel.GenerateUniqueNo(account.CompanyId),
		Name:      body.Name,
		Type: tableComponentTypes.DropDownListItem{
			Id: body.Type,
		},
		SerialNumber:        body.SerialNumber,
		CertificateNo:       body.CertificateNo,
		CalibrationRequired: body.CalibrationRequired,
		InspectionFrequency: tableComponentTypes.DropDownListItem{
			Id: body.InspectionFrequency,
		},
		ICD:          body.ICD,
		NVCD:         body.NVCD,
		EIS:          body.EIS,
		Comment:      body.Comment,
		DbStatus:     "active",
		DbLastStatus: "active",
	})

	c.IndentedJSON(201, gin.H{})
}

func (*EIHandler) Update(c *gin.Context) {
	account, _ := c.MustGet("account").(middlewares.RemoteAccount)
	Id := c.Param("id")

	var eiModel registerModels.EIModel

	currentEi, _ := eiModel.GetById(Id)

	if currentEi.IsEmpty() || currentEi.CompanyId != account.CompanyId {
		c.IndentedJSON(404, gin.H{})
		return
	}

	var body struct {
		Name                string `json:"name"`
		Type                string `json:"type"`
		SerialNumber        string `json:"serialNumber"`
		CertificateNo       string `json:"certificateNo"`
		CalibrationRequired int8   `json:"calibrationRequired"`
		InspectionFrequency string `json:"inspectionFrequency"`
		ICD                 string `json:"icd"`
		NVCD                string `json:"nvcd"`
		EIS                 int8   `json:"eis"`
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

	eiModel.Update(Id, map[string]interface{}{
		"name":                body.Name,
		"type":                body.Type,
		"serialNumber":        body.SerialNumber,
		"certificateNo":       body.CertificateNo,
		"calibrationRequired": body.CalibrationRequired,
		"inspectionFrequency": body.InspectionFrequency,
		"icd":                 body.ICD,
		"nvcd":                body.NVCD,
		"eis":                 body.EIS,
		"comment":             body.Comment,
	})

	c.JSON(200, gin.H{})
}

func (*EIHandler) Archive(c *gin.Context) {
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

	var eiModel registerModels.EIModel

	for _, Id := range body.Ids {
		currentEi, _ := eiModel.GetById(Id)
		if currentEi.IsEmpty() || currentEi.CompanyId != account.CompanyId {
			continue
		}

		if currentEi.DbStatus != "active" {
			continue
		}

		eiModel.Update(Id, map[string]interface{}{
			"dbStatus":     "archived",
			"dbLastStatus": currentEi.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*EIHandler) Unarchive(c *gin.Context) {
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

	var eiModel registerModels.EIModel

	for _, Id := range body.Ids {
		currentEi, _ := eiModel.GetById(Id)
		if currentEi.IsEmpty() || currentEi.CompanyId != account.CompanyId {
			continue
		}

		if currentEi.DbStatus != "archived" {
			continue
		}

		eiModel.Update(Id, map[string]interface{}{
			"dbStatus":     "active",
			"dbLastStatus": currentEi.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*EIHandler) Delete(c *gin.Context) {
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

	var eiModel registerModels.EIModel

	for _, Id := range body.Ids {
		currentEi, _ := eiModel.GetById(Id)
		if currentEi.IsEmpty() || currentEi.CompanyId != account.CompanyId {
			continue
		}

		if currentEi.DbStatus == "deleted" {
			continue
		}

		eiModel.Update(Id, map[string]interface{}{
			"dbStatus":     "deleted",
			"dbLastStatus": currentEi.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*EIHandler) Undelete(c *gin.Context) {
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

	var eiModel registerModels.EIModel

	for _, Id := range body.Ids {
		currentEi, _ := eiModel.GetById(Id)
		if currentEi.IsEmpty() || currentEi.CompanyId != account.CompanyId {
			continue
		}

		if currentEi.DbStatus != "deleted" {
			continue
		}

		eiModel.Update(Id, map[string]interface{}{
			"dbStatus":     currentEi.DbLastStatus,
			"dbLastStatus": currentEi.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}
