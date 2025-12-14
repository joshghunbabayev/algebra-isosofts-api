package registerHandlers

import (
	registerModels "algebra-isosofts-api/models/registers"
	registerTypes "algebra-isosofts-api/types/registers"
	tableComponentTypes "algebra-isosofts-api/types/tableComponents"

	"github.com/gin-gonic/gin"
)

type EIHandler struct {
}

func (*EIHandler) GetAll(c *gin.Context) {
	status := c.Query("status")

	if status == "" {
		status = "active"
	}

	var eiModel registerModels.EIModel

	eis, err := eiModel.GetAll(map[string]interface{}{
		"dbStatus": status,
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
		SerialNumber        string `json:"serialNumber"`
		CertificateNo       string `json:"certificateNo"`
		InspectionFrequency string `json:"inspectionFrequency"`
		ICD                 string `json:"icd"`
		NVCD                string `json:"nvcd"`
		SafeToUse           int8   `json:"safeToUse"`
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

	eiModel.Create(registerTypes.EI{
		Id:            eiModel.GenerateUniqueId(),
		No:            eiModel.GenerateUniqueNo(),
		Name:          body.Name,
		SerialNumber:  body.SerialNumber,
		CertificateNo: body.CertificateNo,
		InspectionFrequency: tableComponentTypes.DropDownListItem{
			Id: body.InspectionFrequency,
		},
		ICD:          body.ICD,
		NVCD:         body.NVCD,
		SafeToUse:    body.SafeToUse,
		DbStatus:     "active",
		DbLastStatus: "active",
	})

	c.IndentedJSON(201, gin.H{})
}

func (*EIHandler) Update(c *gin.Context) {
	Id := c.Param("id")

	var eiModel registerModels.EIModel

	currentEi, _ := eiModel.GetById(Id)

	if currentEi.IsEmpty() {
		c.IndentedJSON(404, gin.H{})
		return
	}

	var body struct {
		Name                string `json:"name"`
		SerialNumber        string `json:"serialNumber"`
		CertificateNo       string `json:"certificateNo"`
		InspectionFrequency string `json:"inspectionFrequency"`
		ICD                 string `json:"icd"`
		NVCD                string `json:"nvcd"`
		SafeToUse           int8   `json:"safeToUse"`
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
		"serialNumber":        body.SerialNumber,
		"certificateNo":       body.CertificateNo,
		"inspectionFrequency": body.InspectionFrequency,
		"icd":                 body.ICD,
		"nvcd":                body.NVCD,
		"safeToUse":           body.SafeToUse,
	})

	c.JSON(200, gin.H{})
}

func (*EIHandler) Archive(c *gin.Context) {
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
		if currentEi.IsEmpty() {
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
		if currentEi.IsEmpty() {
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
		if currentEi.IsEmpty() {
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
		if currentEi.IsEmpty() {
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
