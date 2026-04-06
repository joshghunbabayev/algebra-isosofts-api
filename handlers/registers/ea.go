package registerHandlers

import (
	"algebra-isosofts-api/middlewares"
	registerModels "algebra-isosofts-api/models/registers"
	registerTypes "algebra-isosofts-api/types/registers"
	tableComponentTypes "algebra-isosofts-api/types/tableComponents"

	"github.com/gin-gonic/gin"
)

type EAHandler struct {
}

func (*EAHandler) GetAll(c *gin.Context) {
	account, _ := c.MustGet("account").(middlewares.RemoteAccount)
	status := c.Query("status")

	if status == "" {
		status = "active"
	}

	var eaModel registerModels.EAModel

	eas, err := eaModel.GetAll(map[string]interface{}{
		"dbStatus":  status,
		"companyId": account.CompanyId,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, eas)
}

func (*EAHandler) Create(c *gin.Context) {
	account, _ := c.MustGet("account").(middlewares.RemoteAccount)
	var body struct {
		EmployeeName             string `json:"employeeName"`
		Position                 string `json:"position"`
		LineManager              string `json:"lineManager"`
		ESD                      string `json:"esd"`
		AppraisalDate            string `json:"appraisalDate"`
		NextAppraisalDate        string `json:"nextAppraisalDate"`
		AppraisalType            string `json:"appraisalType"`
		TCA                      string `json:"tca"`
		SkillsAppraisal          string `json:"skillsAppraisal"`
		JobQuality               int8   `json:"jobQuality"`
		LeadershipSkills         int8   `json:"leadershipSkills"`
		ManagementSkills         int8   `json:"managementSkills"`
		BehavioralSkills         int8   `json:"behavioralSkills"`
		EffectivenessOfTrainings int8   `json:"effectivenessOfTrainings"`
		EVS                      int8   `json:"evs"`
		Comment                  string `json:"comment"`
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

	var eaModel registerModels.EAModel

	eaModel.Create(registerTypes.EA{
		Id:                eaModel.GenerateUniqueId(),
		CompanyId:         account.CompanyId,
		No:                eaModel.GenerateUniqueNo(account.CompanyId),
		EmployeeName:      body.EmployeeName,
		Position:          body.Position,
		LineManager:       body.LineManager,
		ESD:               body.ESD,
		AppraisalDate:     body.AppraisalDate,
		NextAppraisalDate: body.NextAppraisalDate,
		AppraisalType: tableComponentTypes.DropDownListItem{
			Id: body.AppraisalType,
		},
		TCA:                      body.TCA,
		SkillsAppraisal:          body.SkillsAppraisal,
		JobQuality:               body.JobQuality,
		LeadershipSkills:         body.LeadershipSkills,
		ManagementSkills:         body.ManagementSkills,
		BehavioralSkills:         body.BehavioralSkills,
		EffectivenessOfTrainings: body.EffectivenessOfTrainings,
		EVS:                      body.EVS,
		Comment:                  body.Comment,
		DbStatus:                 "active",
		DbLastStatus:             "active",
	})

	c.IndentedJSON(201, gin.H{})
}

func (*EAHandler) Update(c *gin.Context) {
	account, _ := c.MustGet("account").(middlewares.RemoteAccount)
	Id := c.Param("id")

	var eaModel registerModels.EAModel

	currentEa, _ := eaModel.GetById(Id)

	if currentEa.IsEmpty() || currentEa.CompanyId != account.CompanyId {
		c.IndentedJSON(404, gin.H{})
		return
	}

	var body struct {
		EmployeeName             string `json:"employeeName"`
		Position                 string `json:"position"`
		LineManager              string `json:"lineManager"`
		ESD                      string `json:"esd"`
		AppraisalDate            string `json:"appraisalDate"`
		NextAppraisalDate        string `json:"nextAppraisalDate"`
		AppraisalType            string `json:"appraisalType"`
		TCA                      string `json:"tca"`
		SkillsAppraisal          string `json:"skillsAppraisal"`
		JobQuality               int8   `json:"jobQuality"`
		LeadershipSkills         int8   `json:"leadershipSkills"`
		ManagementSkills         int8   `json:"managementSkills"`
		BehavioralSkills         int8   `json:"behavioralSkills"`
		EffectivenessOfTrainings int8   `json:"effectivenessOfTrainings"`
		EVS                      int8   `json:"evs"`
		Comment                  string `json:"comment"`
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

	eaModel.Update(Id, map[string]interface{}{
		"employeeName":             body.EmployeeName,
		"position":                 body.Position,
		"lineManager":              body.LineManager,
		"esd":                      body.ESD,
		"appraisalDate":            body.AppraisalDate,
		"nextAppraisalDate":        body.NextAppraisalDate,
		"appraisalType":            body.AppraisalType,
		"tca":                      body.TCA,
		"skillsAppraisal":          body.SkillsAppraisal,
		"jobQuality":               body.JobQuality,
		"leadershipSkills":         body.LeadershipSkills,
		"managementSkills":         body.ManagementSkills,
		"behavioralSkills":         body.BehavioralSkills,
		"effectivenessOfTrainings": body.EffectivenessOfTrainings,
		"evs":                      body.EVS,
		"comment":                  body.Comment,
	})

	c.JSON(200, gin.H{})
}

func (*EAHandler) Archive(c *gin.Context) {
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

	var eaModel registerModels.EAModel

	for _, Id := range body.Ids {
		currentEa, _ := eaModel.GetById(Id)
		if currentEa.IsEmpty() || currentEa.CompanyId != account.CompanyId {
			continue
		}

		if currentEa.DbStatus != "active" {
			continue
		}

		eaModel.Update(Id, map[string]interface{}{
			"dbStatus":     "archived",
			"dbLastStatus": currentEa.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*EAHandler) Unarchive(c *gin.Context) {
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

	var eaModel registerModels.EAModel

	for _, Id := range body.Ids {
		currentEa, _ := eaModel.GetById(Id)
		if currentEa.IsEmpty() || currentEa.CompanyId != account.CompanyId {
			continue
		}

		if currentEa.DbStatus != "archived" {
			continue
		}

		eaModel.Update(Id, map[string]interface{}{
			"dbStatus":     "active",
			"dbLastStatus": currentEa.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*EAHandler) Delete(c *gin.Context) {
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

	var eaModel registerModels.EAModel

	for _, Id := range body.Ids {
		currentEa, _ := eaModel.GetById(Id)
		if currentEa.IsEmpty() || currentEa.CompanyId != account.CompanyId {
			continue
		}

		if currentEa.DbStatus == "deleted" {
			continue
		}

		eaModel.Update(Id, map[string]interface{}{
			"dbStatus":     "deleted",
			"dbLastStatus": currentEa.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*EAHandler) Undelete(c *gin.Context) {
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

	var eaModel registerModels.EAModel

	for _, Id := range body.Ids {
		currentEa, _ := eaModel.GetById(Id)
		if currentEa.IsEmpty() || currentEa.CompanyId != account.CompanyId {
			continue
		}

		if currentEa.DbStatus != "deleted" {
			continue
		}

		eaModel.Update(Id, map[string]interface{}{
			"dbStatus":     currentEa.DbLastStatus,
			"dbLastStatus": currentEa.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}
