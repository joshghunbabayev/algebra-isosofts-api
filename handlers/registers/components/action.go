package registerComponentHandlers

import (
	"algebra-isosofts-api/mailer"
	"algebra-isosofts-api/middlewares"
	registerModels "algebra-isosofts-api/models/registers"
	registerComponentModels "algebra-isosofts-api/models/registers/components"
	tableComponentModels "algebra-isosofts-api/models/tableComponents"
	"algebra-isosofts-api/modules"
	registerComponentTypes "algebra-isosofts-api/types/registers/components"
	tableComponentTypes "algebra-isosofts-api/types/tableComponents"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type ActionHandler struct {
}

func (*ActionHandler) GetAll(c *gin.Context) {
	registerId := c.Query("registerId")
	status := c.Query("status")

	if status == "" {
		status = "active"
	}

	var actionModel registerComponentModels.ActionModel
	var actions []registerComponentTypes.Action
	var err error

	if registerId == "" {
		actions, err = actionModel.GetAll(map[string]interface{}{
			"dbStatus": status,
		})
	} else {
		actions, err = actionModel.GetAll(map[string]interface{}{
			"registerId": registerId,
			"dbStatus":   status,
		})
	}

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, actions)
}

func (*ActionHandler) Create(c *gin.Context) {
	var body struct {
		RegisterId         string `json:"registerId"`
		RegisterType       string `json:"registerType"`
		Title              string `json:"title"`
		RaiseDate          string `json:"raiseDate"`
		Resources          string `json:"resources"`
		Currency           string `json:"currency"`
		RelativeFunction   string `json:"relativeFunction"`
		ResponsibleId      string `json:"responsibleId"`
		ApproverId         string `json:"approverId"`
		Deadline           string `json:"deadline"`
		Confirmation       string `json:"confirmation"`
		Status             string `json:"status"`
		CompletionDate     string `json:"completionDate"`
		VerificationStatus string `json:"verificationStatus"`
		Comment            string `json:"comment"`
		January            string `json:"january"`
		February           string `json:"february"`
		March              string `json:"march"`
		April              string `json:"april"`
		May                string `json:"may"`
		June               string `json:"june"`
		July               string `json:"july"`
		August             string `json:"august"`
		September          string `json:"september"`
		October            string `json:"october"`
		November           string `json:"november"`
		December           string `json:"december"`
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

	var commonModel registerModels.CommonModel
	var actionModel registerComponentModels.ActionModel

	account, _ := c.MustGet("account").(middlewares.RemoteAccount)
	regNo, _ := commonModel.GetRegNo(body.RegisterId, body.RegisterType)
	actionNo := actionModel.GenerateUniqueNo(account.CompanyId, regNo)

	actionModel.Create(registerComponentTypes.Action{
		Id:           actionModel.GenerateUniqueId(),
		CompanyId:    account.CompanyId,
		RegisterId:   body.RegisterId,
		RegisterType: body.RegisterType,
		No:           actionNo,
		Title:        body.Title,
		RaiseDate:    body.RaiseDate,
		Resources:    body.Resources,
		Currency:     body.Currency,
		RelativeFunction: tableComponentTypes.DropDownListItem{
			Id: body.RelativeFunction,
		},
		ResponsibleId: body.ResponsibleId,
		ApproverId:    body.ApproverId,
		Deadline:      body.Deadline,
		Confirmation: tableComponentTypes.DropDownListItem{
			Id: body.Confirmation,
		},
		Status: tableComponentTypes.DropDownListItem{
			Id: body.Status,
		},
		CompletionDate: body.CompletionDate,
		VerificationStatus: tableComponentTypes.DropDownListItem{
			Id: body.VerificationStatus,
		},
		Comment: body.Comment,
		January: tableComponentTypes.DropDownListItem{
			Id: body.January,
		},
		February: tableComponentTypes.DropDownListItem{
			Id: body.February,
		},
		March: tableComponentTypes.DropDownListItem{
			Id: body.March,
		},
		April: tableComponentTypes.DropDownListItem{
			Id: body.April,
		},
		May: tableComponentTypes.DropDownListItem{
			Id: body.May,
		},
		June: tableComponentTypes.DropDownListItem{
			Id: body.June,
		},
		July: tableComponentTypes.DropDownListItem{
			Id: body.July,
		},
		August: tableComponentTypes.DropDownListItem{
			Id: body.August,
		},
		September: tableComponentTypes.DropDownListItem{
			Id: body.September,
		},
		October: tableComponentTypes.DropDownListItem{
			Id: body.October,
		},
		November: tableComponentTypes.DropDownListItem{
			Id: body.November,
		},
		December: tableComponentTypes.DropDownListItem{
			Id: body.December,
		},
		CreatedById:          account.Id,
		DbStatus:             "active",
		DbLastStatus:         "active",
		LastNotificationDate: time.Now().Format("2006-01-02"),
	})

	if strings.TrimSpace(body.ResponsibleId) != "" {
		responsible := modules.GetAccountById(body.ResponsibleId)

		toContacts := []mailer.EmailContact{
			{Email: responsible.Email, Name: responsible.Name + " " + responsible.Surname},
		}

		ccContacts := []mailer.EmailContact{}

		lineManager := modules.GetAccountById(responsible.LineManagerId)
		if !lineManager.IsEmpty() {
			ccContacts = append(ccContacts, mailer.EmailContact{Email: lineManager.Email, Name: lineManager.Name + " " + lineManager.Surname})
		}

		createdBy := modules.GetAccountById(account.Id)
		if !createdBy.IsEmpty() {
			ccContacts = append(ccContacts, mailer.EmailContact{Email: createdBy.Email, Name: createdBy.Name + " " + createdBy.Surname})
		}

		var dropDownListItemModel tableComponentModels.DropDownListItemModel
		relativeFunction, _ := dropDownListItemModel.GetById(body.RelativeFunction)
		status, _ := dropDownListItemModel.GetById(body.Status)

		emailSubject := fmt.Sprintf("New Action Assigned: %s - %s", actionNo, body.Title)

		emailBody := fmt.Sprintf(`
				<p>Dear %s,</p>
				<p>We would like to inform you that the following action has been assigned to you for implementation:</p>
				<br/>
				<table style="border-collapse: collapse; width: 100%%; max-width: 600px;">
					<tr><td style="padding: 5px; font-weight: bold; width: 150px;">Action No:</td><td style="padding: 5px;">%s</td></tr>
					<tr><td style="padding: 5px; font-weight: bold;">Action Description:</td><td style="padding: 5px;">%s</td></tr>
					<tr><td style="padding: 5px; font-weight: bold;">Action Raised Date:</td><td style="padding: 5px;">%s</td></tr>
					<tr><td style="padding: 5px; font-weight: bold;">Resources:</td><td style="padding: 5px;">%s (%s)</td></tr>
					<tr><td style="padding: 5px; font-weight: bold;">Related Function:</td><td style="padding: 5px;">%s</td></tr>
					<tr><td style="padding: 5px; font-weight: bold;">Responsible Person:</td><td style="padding: 5px;">%s %s</td></tr>
					<tr><td style="padding: 5px; font-weight: bold;">Action Status:</td><td style="padding: 5px;">%s</td></tr>
					<tr><td style="padding: 5px; font-weight: bold;">Deadline:</td><td style="padding: 5px;">%s</td></tr>
				</table>
				<br/>
				<p>Kindly accept the action and reply to all recipients of this email.</p>
				<p>If the action is rejected, so please appropriately reply with brief explanation to all recipients of this email.</p>
				<br/>
				<p>Thanks for the prompt response.</p>
			`,
			responsible.Name+" "+responsible.Surname,
			actionNo,
			body.Title,
			body.RaiseDate,
			body.Resources, body.Currency,
			relativeFunction.Value,
			responsible.Name, responsible.Surname,
			status.Value,
			body.Deadline,
		)

		err := mailer.SendEmail(
			toContacts,
			ccContacts,
			emailSubject,
			emailBody,
		)

		fmt.Println("errr", err)
	}

	c.IndentedJSON(201, gin.H{})
}

func (*ActionHandler) Update(c *gin.Context) {
	Id := c.Param("id")

	var actionModel registerComponentModels.ActionModel

	currentAction, _ := actionModel.GetById(Id)

	if currentAction.IsEmpty() {
		c.IndentedJSON(404, gin.H{})
		return
	}

	var body struct {
		Title              string `json:"title"`
		RaiseDate          string `json:"raiseDate"`
		Resources          string `json:"resources"`
		Currency           string `json:"currency"`
		RelativeFunction   string `json:"relativeFunction"`
		ResponsibleId      string `json:"responsibleId"`
		ApproverId         string `json:"approverId"`
		Deadline           string `json:"deadline"`
		Confirmation       string `json:"confirmation"`
		Status             string `json:"status"`
		CompletionDate     string `json:"completionDate"`
		VerificationStatus string `json:"verificationStatus"`
		Comment            string `json:"comment"`
		January            string `json:"january"`
		February           string `json:"february"`
		March              string `json:"march"`
		April              string `json:"april"`
		May                string `json:"may"`
		June               string `json:"june"`
		July               string `json:"july"`
		August             string `json:"august"`
		September          string `json:"september"`
		October            string `json:"october"`
		November           string `json:"november"`
		December           string `json:"december"`
		SendNotification   int    `json:"sendNotification"`
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

	actionModel.Update(Id, map[string]interface{}{
		"title":              body.Title,
		"raiseDate":          body.RaiseDate,
		"resources":          body.Resources,
		"currency":           body.Currency,
		"relativeFunction":   body.RelativeFunction,
		"responsibleId":      body.ResponsibleId,
		"approverId":         body.ApproverId,
		"deadline":           body.Deadline,
		"confirmation":       body.Confirmation,
		"status":             body.Status,
		"completionDate":     body.CompletionDate,
		"verificationStatus": body.VerificationStatus,
		"comment":            body.Comment,
		"january":            body.January,
		"february":           body.February,
		"march":              body.March,
		"april":              body.April,
		"may":                body.May,
		"june":               body.June,
		"july":               body.July,
		"august":             body.August,
		"september":          body.September,
		"october":            body.October,
		"november":           body.November,
		"december":           body.December,
	})

	if strings.TrimSpace(body.ResponsibleId) != "" && body.SendNotification == 1 {
		actionModel.Update(Id, map[string]interface{}{
			"lastNotificationDate": time.Now().Format("2006-01-02"),
		})

		responsible := modules.GetAccountById(body.ResponsibleId)

		toContacts := []mailer.EmailContact{
			{Email: responsible.Email, Name: responsible.Name + " " + responsible.Surname},
		}

		ccContacts := []mailer.EmailContact{}

		lineManager := modules.GetAccountById(responsible.LineManagerId)
		if !lineManager.IsEmpty() {
			ccContacts = append(ccContacts, mailer.EmailContact{Email: lineManager.Email, Name: lineManager.Name + " " + lineManager.Surname})
		}

		createdBy := modules.GetAccountById(currentAction.CreatedById)
		if !createdBy.IsEmpty() {
			ccContacts = append(ccContacts, mailer.EmailContact{Email: createdBy.Email, Name: createdBy.Name + " " + createdBy.Surname})
		}

		var dropDownListItemModel tableComponentModels.DropDownListItemModel
		relativeFunction, _ := dropDownListItemModel.GetById(body.RelativeFunction)
		status, _ := dropDownListItemModel.GetById(body.Status)

		emailSubject := fmt.Sprintf("Action Updated Notification: %s - %s", currentAction.No, body.Title)

		emailBody := fmt.Sprintf(`
				<p>Dear %s,</p>
				<p>We would like to inform that some information related to the action assigned to you for implementation has been updated:</p>
				<br/>
				<table style="border-collapse: collapse; width: 100%%; max-width: 600px;">
						<tr><td style="padding: 5px; font-weight: bold; width: 150px;">Action No:</td><td style="padding: 5px;">%s</td></tr>
						<tr><td style="padding: 5px; font-weight: bold;">Action Description:</td><td style="padding: 5px;">%s</td></tr>
						<tr><td style="padding: 5px; font-weight: bold;">Action Raised Date:</td><td style="padding: 5px;">%s</td></tr>
						<tr><td style="padding: 5px; font-weight: bold;">Resources:</td><td style="padding: 5px;">%s (%s)</td></tr>
						<tr><td style="padding: 5px; font-weight: bold;">Related Function:</td><td style="padding: 5px;">%s</td></tr>
						<tr><td style="padding: 5px; font-weight: bold;">Responsible Person:</td><td style="padding: 5px;">%s %s</td></tr>
						<tr><td style="padding: 5px; font-weight: bold;">Action Status:</td><td style="padding: 5px;">%s</td></tr>
						<tr><td style="padding: 5px; font-weight: bold;">Deadline:</td><td style="padding: 5px;">%s</td></tr>
				</table>
				<br/>
				<p>Thanks for the prompt response.</p>
			`,
			responsible.Name+" "+responsible.Surname,
			currentAction.No,
			body.Title,
			body.RaiseDate,
			body.Resources, body.Currency,
			relativeFunction.Value,
			responsible.Name, responsible.Surname,
			status.Value,
			body.Deadline,
		)

		err := mailer.SendEmail(
			toContacts,
			ccContacts,
			emailSubject,
			emailBody,
		)

		if strings.Contains(status.Value, "100") {
			approver := modules.GetAccountById(body.ApproverId)

			if !approver.IsEmpty() {
				approverContacts := []mailer.EmailContact{
					{Email: approver.Email, Name: approver.Name + " " + approver.Surname},
				}

				approverSubject := fmt.Sprintf("Action Implementation Verification Required: %s", currentAction.No)
				approverBody := fmt.Sprintf(`
						<p>Dear %s,</p>
						<p>I would like to inform you that the following action has been implemented:</p>
						<br/>
						<table style="border-collapse: collapse; width: 100%%; max-width: 600px;">
								<tr><td style="padding: 5px; font-weight: bold; width: 150px;">Action No:</td><td style="padding: 5px;">%s</td></tr>
								<tr><td style="padding: 5px; font-weight: bold;">Action Description:</td><td style="padding: 5px;">%s</td></tr>
								<tr><td style="padding: 5px; font-weight: bold;">Action Raised Date:</td><td style="padding: 5px;">%s</td></tr>
								<tr><td style="padding: 5px; font-weight: bold;">Resources:</td><td style="padding: 5px;">%s (%s)</td></tr>
								<tr><td style="padding: 5px; font-weight: bold;">Related Function:</td><td style="padding: 5px;">%s</td></tr>
								<tr><td style="padding: 5px; font-weight: bold;">Responsible Person:</td><td style="padding: 5px;">%s %s</td></tr>
								<tr><td style="padding: 5px; font-weight: bold;">Action Status:</td><td style="padding: 5px;">%s</td></tr>
								<tr><td style="padding: 5px; font-weight: bold;">Deadline:</td><td style="padding: 5px;">%s</td></tr>
						</table>
						<br/>
						<p>Please verify the completion of the action by highlighting the appropriate completion status:</p>
						<p>Completed, Completed with Delay, Rework Required.</p>
						<br/>
						<p>Thanks for the prompt response</p>
					`,
					approver.Name+" "+approver.Surname,
					currentAction.No,
					body.Title,
					body.RaiseDate,
					body.Resources, body.Currency,
					relativeFunction.Value,
					responsible.Name, responsible.Surname,
					status.Value,
					body.Deadline,
				)

				mailer.SendEmail(
					approverContacts,
					[]mailer.EmailContact{},
					approverSubject,
					approverBody,
				)
			}
		}

		fmt.Println("errr", err)
	}

	c.JSON(200, gin.H{})
}

func (*ActionHandler) Delete(c *gin.Context) {
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

	var actionModel registerComponentModels.ActionModel

	for _, Id := range body.Ids {
		currentAction, _ := actionModel.GetById(Id)
		if currentAction.IsEmpty() {
			continue
		}

		if currentAction.DbStatus == "deleted" {
			continue
		}

		actionModel.Update(Id, map[string]interface{}{
			"dbStatus":     "deleted",
			"dbLastStatus": currentAction.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}

func (*ActionHandler) Undelete(c *gin.Context) {
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

	var actionModel registerComponentModels.ActionModel

	for _, Id := range body.Ids {
		currentAction, _ := actionModel.GetById(Id)
		if currentAction.IsEmpty() {
			continue
		}

		if currentAction.DbStatus != "deleted" {
			continue
		}

		actionModel.Update(Id, map[string]interface{}{
			"dbStatus":     currentAction.DbLastStatus,
			"dbLastStatus": currentAction.DbStatus,
		})
	}

	c.JSON(200, gin.H{})
}
