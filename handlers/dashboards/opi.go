package dashboardHandlers

import (
	"algebra-isosofts-api/middlewares"
	dashboardModels "algebra-isosofts-api/models/dashboards"
	dashboardTypes "algebra-isosofts-api/types/dashboards"
	tableComponentTypes "algebra-isosofts-api/types/tableComponents"

	"github.com/gin-gonic/gin"
)

type OPIHandler struct {
}

func (*OPIHandler) GetAll(c *gin.Context) {
	account, _ := c.MustGet("account").(middlewares.RemoteAccount)
	status := c.Query("status")

	if status == "" {
		status = "active"
	}

	var opiModel dashboardModels.OPIModel

	opis, err := opiModel.GetAll(map[string]interface{}{
		"dbStatus":  status,
		"companyId": account.CompanyId,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, opis)
}

func (*OPIHandler) Create(c *gin.Context) {
	var body struct {
		Title        string `json:"title"`
		Function     string `json:"function"`
		LYKPI        int64  `json:"lykpi"`
		ActualKPI    int64  `json:"actualKPI"`
		AnnualTarget int64  `json:"annualTarget"`
		January      int64  `json:"january"`
		February     int64  `json:"february"`
		March        int64  `json:"march"`
		April        int64  `json:"april"`
		May          int64  `json:"may"`
		June         int64  `json:"june"`
		July         int64  `json:"july"`
		August       int64  `json:"august"`
		September    int64  `json:"september"`
		October      int64  `json:"october"`
		November     int64  `json:"november"`
		December     int64  `json:"december"`
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

	var opiModel dashboardModels.OPIModel

	account, _ := c.MustGet("account").(middlewares.RemoteAccount)

	opiModel.Create(dashboardTypes.OPI{
		Id:        opiModel.GenerateUniqueId(),
		CompanyId: account.CompanyId,
		No:        opiModel.GenerateUniqueNo(account.CompanyId),
		Title:     body.Title,
		Function: tableComponentTypes.DropDownListItem{
			Id: body.Function,
		},
		LYKPI:        body.LYKPI,
		ActualKPI:    body.ActualKPI,
		AnnualTarget: body.AnnualTarget,
		January:      body.January,
		February:     body.February,
		March:        body.March,
		April:        body.April,
		May:          body.May,
		June:         body.June,
		July:         body.July,
		August:       body.August,
		September:    body.September,
		October:      body.October,
		November:     body.November,
		December:     body.December,
		DbStatus:     "active",
		DbLastStatus: "active",
	})

	c.IndentedJSON(201, gin.H{})
}

func (*OPIHandler) Update(c *gin.Context) {
	account, _ := c.MustGet("account").(middlewares.RemoteAccount)
	Id := c.Param("id")

	var opiModel dashboardModels.OPIModel

	currentOPI, _ := opiModel.GetById(Id)

	if currentOPI.IsEmpty() || currentOPI.CompanyId != account.CompanyId {
		c.IndentedJSON(404, gin.H{})
		return
	}

	var body struct {
		Title        string `json:"title"`
		Function     string `json:"function"`
		LYKPI        int64  `json:"lykpi"`
		ActualKPI    int64  `json:"actualKPI"`
		AnnualTarget int64  `json:"annualTarget"`
		January      int64  `json:"january"`
		February     int64  `json:"february"`
		March        int64  `json:"march"`
		April        int64  `json:"april"`
		May          int64  `json:"may"`
		June         int64  `json:"june"`
		July         int64  `json:"july"`
		August       int64  `json:"august"`
		September    int64  `json:"september"`
		October      int64  `json:"october"`
		November     int64  `json:"november"`
		December     int64  `json:"december"`
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

	opiModel.Update(Id, map[string]interface{}{
		"title":        body.Title,
		"function":     body.Function,
		"lykpi":        body.LYKPI,
		"actualKPI":    body.ActualKPI,
		"annualTarget": body.AnnualTarget,
		"january":      body.January,
		"february":     body.February,
		"march":        body.March,
		"april":        body.April,
		"may":          body.May,
		"june":         body.June,
		"july":         body.July,
		"august":       body.August,
		"september":    body.September,
		"october":      body.October,
		"november":     body.November,
		"december":     body.December,
	})

	c.JSON(200, gin.H{})
}

func (*OPIHandler) Archive(c *gin.Context) {
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

	var opiModel dashboardModels.OPIModel

	for _, Id := range body.Ids {
		currentOPI, _ := opiModel.GetById(Id)
		if currentOPI.IsEmpty() || currentOPI.CompanyId != account.CompanyId {
			continue
		}

		if currentOPI.DbStatus != "active" {
			continue
		}

		opiModel.Update(Id, map[string]interface{}{
			"dbStatus":     "archived",
			"dbLastStatus": currentOPI.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*OPIHandler) Unarchive(c *gin.Context) {
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

	var opiModel dashboardModels.OPIModel

	for _, Id := range body.Ids {
		currentOPI, _ := opiModel.GetById(Id)
		if currentOPI.IsEmpty() || currentOPI.CompanyId != account.CompanyId {
			continue
		}

		if currentOPI.DbStatus != "archived" {
			continue
		}

		opiModel.Update(Id, map[string]interface{}{
			"dbStatus":     "active",
			"dbLastStatus": currentOPI.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*OPIHandler) Delete(c *gin.Context) {
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

	var opiModel dashboardModels.OPIModel

	for _, Id := range body.Ids {
		currentOPI, _ := opiModel.GetById(Id)
		if currentOPI.IsEmpty() || currentOPI.CompanyId != account.CompanyId {
			continue
		}

		if currentOPI.DbStatus == "deleted" {
			continue
		}

		opiModel.Update(Id, map[string]interface{}{
			"dbStatus":     "deleted",
			"dbLastStatus": currentOPI.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*OPIHandler) Undelete(c *gin.Context) {
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

	var opiModel dashboardModels.OPIModel

	for _, Id := range body.Ids {
		currentOPI, _ := opiModel.GetById(Id)
		if currentOPI.IsEmpty() || currentOPI.CompanyId != account.CompanyId {
			continue
		}

		if currentOPI.DbStatus != "deleted" {
			continue
		}

		opiModel.Update(Id, map[string]interface{}{
			"dbStatus":     currentOPI.DbLastStatus,
			"dbLastStatus": currentOPI.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}
