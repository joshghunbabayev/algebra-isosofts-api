package registerComponentHandlers

import (
	registerComponentModels "algebra-isosofts-api/models/registers/components"
	registerComponentTypes "algebra-isosofts-api/types/registers/components"
	tableComponentTypes "algebra-isosofts-api/types/tableComponents"

	"github.com/gin-gonic/gin"
)

type VendorFeedbackHandler struct {
}

func (*VendorFeedbackHandler) GetAll(c *gin.Context) {
	registerId := c.Query("registerId")
	status := c.Query("status")

	if status == "" {
		status = "active"
	}

	var vendorFeedbackModel registerComponentModels.VendorFeedbackModel
	var vendorFeedbacks []registerComponentTypes.VendorFeedback
	var err error

	if registerId == "" {
		vendorFeedbacks, err = vendorFeedbackModel.GetAll(map[string]interface{}{
			"dbStatus": status,
		})
	} else {
		vendorFeedbacks, err = vendorFeedbackModel.GetAll(map[string]interface{}{
			"registerId": registerId,
			"dbStatus":   status,
		})
	}

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, vendorFeedbacks)
}

func (*VendorFeedbackHandler) Create(c *gin.Context) {
	var body struct {
		RegisterId    string `json:"registerId"`
		Scope         string `json:"scope"`
		VendorId      string `json:"vendorId"`
		TypeOfFinding string `json:"typeOfFinding"`
		QGS           int8   `json:"qgs"`
		Communication int8   `json:"communication"`
		OTD           int8   `json:"otd"`
		Documentation int8   `json:"documentation"`
		HS            int8   `json:"hs"`
		Environment   int8   `json:"environment"`
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

	// var brModel registerModels.BRModel

	// br, _ := brModel.GetById(body.RegisterId)

	// if br.IsEmpty() {
	// 	c.IndentedJSON(404, gin.H{})
	// 	return
	// }

	// if br.DbStatus != "active" {
	// 	c.IndentedJSON(400, gin.H{})
	// 	return
	// }

	var vendorFeedbackModel registerComponentModels.VendorFeedbackModel

	vendorFeedbackModel.Create(registerComponentTypes.VendorFeedback{
		Id:         vendorFeedbackModel.GenerateUniqueId(),
		RegisterId: body.RegisterId,
		Scope: tableComponentTypes.DropDownListItem{
			Id: body.Scope,
		},
		VendorId: body.VendorId,
		TypeOfFinding: tableComponentTypes.DropDownListItem{
			Id: body.TypeOfFinding,
		},
		QGS:           body.QGS,
		Communication: body.Communication,
		OTD:           body.OTD,
		Documentation: body.Documentation,
		HS:            body.HS,
		Environment:   body.Environment,
		DbStatus:      "active",
		DbLastStatus:  "active",
	})

	c.IndentedJSON(201, gin.H{})
}

func (*VendorFeedbackHandler) Update(c *gin.Context) {
	Id := c.Param("id")

	var vendorFeedbackModel registerComponentModels.VendorFeedbackModel

	currentVendorFeedback, _ := vendorFeedbackModel.GetById(Id)

	if currentVendorFeedback.IsEmpty() {
		c.IndentedJSON(404, gin.H{})
		return
	}

	var body struct {
		RegisterId    string `json:"registerId"`
		Scope         string `json:"scope"`
		VendorId      string `json:"vendorId"`
		TypeOfFinding string `json:"typeOfFinding"`
		QGS           int8   `json:"qgs"`
		Communication int8   `json:"communication"`
		OTD           int8   `json:"otd"`
		Documentation int8   `json:"documentation"`
		HS            int8   `json:"hs"`
		Environment   int8   `json:"environment"`
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

	vendorFeedbackModel.Update(Id, map[string]interface{}{
		"registerId":    body.RegisterId,
		"scope":         body.Scope,
		"vendorId":      body.VendorId,
		"typeOfFinding": body.TypeOfFinding,
		"qgs":           body.QGS,
		"communication": body.Communication,
		"otd":           body.OTD,
		"documentation": body.Documentation,
		"hs":            body.HS,
		"environment":   body.Environment,
	})

	c.JSON(200, gin.H{})
}

func (*VendorFeedbackHandler) Delete(c *gin.Context) {
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

	var vendorFeedbackModel registerComponentModels.VendorFeedbackModel

	for _, Id := range body.Ids {
		currentVendorFeedback, _ := vendorFeedbackModel.GetById(Id)
		if currentVendorFeedback.IsEmpty() {
			continue
		}

		if currentVendorFeedback.DbStatus == "deleted" {
			continue
		}

		vendorFeedbackModel.Update(Id, map[string]interface{}{
			"dbStatus":     "deleted",
			"dbLastStatus": currentVendorFeedback.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*VendorFeedbackHandler) Undelete(c *gin.Context) {
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

	var vendorFeedbackModel registerComponentModels.VendorFeedbackModel

	for _, Id := range body.Ids {
		currentVendorFeedback, _ := vendorFeedbackModel.GetById(Id)
		if currentVendorFeedback.IsEmpty() {
			continue
		}

		if currentVendorFeedback.DbStatus != "deleted" {
			continue
		}

		vendorFeedbackModel.Update(Id, map[string]interface{}{
			"dbStatus":     currentVendorFeedback.DbLastStatus,
			"dbLastStatus": currentVendorFeedback.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}
