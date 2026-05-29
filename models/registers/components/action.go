package registerComponentModels

import (
	"algebra-isosofts-api/database"
	"algebra-isosofts-api/mailer"
	tableComponentModels "algebra-isosofts-api/models/tableComponents"
	"algebra-isosofts-api/modules"
	registerComponentTypes "algebra-isosofts-api/types/registers/components"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type ActionModel struct {
}

func (*ActionModel) GenerateUniqueId() string {
	Id := modules.GenerateRandomString(30)

	var actionModel ActionModel

	action, _ := actionModel.GetById(Id)

	if action.IsEmpty() {
		return Id
	} else {
		return actionModel.GenerateUniqueId()
	}
}

func (*ActionModel) GenerateUniqueNo(companyId string, regNo string) string {
	db := database.GetDatabase()

	var lastNo string
	err := db.QueryRow(`
        SELECT "no" 
        FROM actions 
        WHERE companyId = ? AND "no" LIKE ? 
        ORDER BY "no" DESC 
        LIMIT 1
        `,
		companyId,
		regNo+"/%",
	).Scan(&lastNo)

	var nextNumber int
	if err != nil || lastNo == "" {
		nextNumber = 1
	} else {
		parts := strings.Split(lastNo, "/")
		if len(parts) == 4 {
			numPart := parts[3]
			num, _ := strconv.Atoi(numPart)
			nextNumber = num + 1
		} else {
			nextNumber = 1
		}
	}

	newNo := fmt.Sprintf("%s/%02d", regNo, nextNumber)
	return newNo
}

func (*ActionModel) GetById(Id string) (registerComponentTypes.Action, error) {
	db := database.GetDatabase()
	row := db.QueryRow(`
			SELECT * 
			FROM actions
			WHERE id = ?
		`,
		Id,
	)

	var action registerComponentTypes.Action
	var dropDownListItemModel tableComponentModels.DropDownListItemModel

	err := row.Scan(
		&action.Id,
		&action.CompanyId,
		&action.RegisterId,
		&action.RegisterType,
		&action.No,
		&action.Title,
		&action.RaiseDate,
		&action.Resources,
		&action.Currency,
		&action.RelativeFunction.Id,
		&action.ResponsibleId,
		&action.ApproverId,
		&action.Deadline,
		&action.Confirmation.Id,
		&action.Status.Id,
		&action.CompletionDate,
		&action.VerificationStatus.Id,
		&action.Comment,
		&action.January.Id,
		&action.February.Id,
		&action.March.Id,
		&action.April.Id,
		&action.May.Id,
		&action.June.Id,
		&action.July.Id,
		&action.August.Id,
		&action.September.Id,
		&action.October.Id,
		&action.November.Id,
		&action.December.Id,
		&action.CreatedById,
		&action.DbStatus,
		&action.DbLastStatus,
		&action.LastNotificationDate,
	)
	action.RelativeFunction, _ = dropDownListItemModel.GetById(action.RelativeFunction.Id)
	action.Confirmation, _ = dropDownListItemModel.GetById(action.Confirmation.Id)
	action.Status, _ = dropDownListItemModel.GetById(action.Status.Id)
	action.VerificationStatus, _ = dropDownListItemModel.GetById(action.VerificationStatus.Id)
	action.January, _ = dropDownListItemModel.GetById(action.January.Id)
	action.February, _ = dropDownListItemModel.GetById(action.February.Id)
	action.March, _ = dropDownListItemModel.GetById(action.March.Id)
	action.April, _ = dropDownListItemModel.GetById(action.April.Id)
	action.May, _ = dropDownListItemModel.GetById(action.May.Id)
	action.June, _ = dropDownListItemModel.GetById(action.June.Id)
	action.July, _ = dropDownListItemModel.GetById(action.July.Id)
	action.August, _ = dropDownListItemModel.GetById(action.August.Id)
	action.September, _ = dropDownListItemModel.GetById(action.September.Id)
	action.October, _ = dropDownListItemModel.GetById(action.October.Id)
	action.November, _ = dropDownListItemModel.GetById(action.November.Id)
	action.December, _ = dropDownListItemModel.GetById(action.December.Id)

	return action, err
}

func (*ActionModel) GetAll(filters map[string]interface{}) ([]registerComponentTypes.Action, error) {
	db := database.GetDatabase()
	whereClause := ""
	values := []interface{}{}

	if len(filters) > 0 {
		whereParts := []string{}
		for key, val := range filters {
			whereParts = append(whereParts, fmt.Sprintf(`"%s" = ?`, key))
			values = append(values, val)
		}
		whereClause = "WHERE " + strings.Join(whereParts, " AND ")
	}

	query := fmt.Sprintf(`
			SELECT * FROM actions %s
		`,
		whereClause,
	)
	rows, err := db.Query(query, values...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var actions []registerComponentTypes.Action

	for rows.Next() {
		var action registerComponentTypes.Action
		var dropDownListItemModel tableComponentModels.DropDownListItemModel

		rows.Scan(
			&action.Id,
			&action.CompanyId,
			&action.RegisterId,
			&action.RegisterType,
			&action.No,
			&action.Title,
			&action.RaiseDate,
			&action.Resources,
			&action.Currency,
			&action.RelativeFunction.Id,
			&action.ResponsibleId,
			&action.ApproverId,
			&action.Deadline,
			&action.Confirmation.Id,
			&action.Status.Id,
			&action.CompletionDate,
			&action.VerificationStatus.Id,
			&action.Comment,
			&action.January.Id,
			&action.February.Id,
			&action.March.Id,
			&action.April.Id,
			&action.May.Id,
			&action.June.Id,
			&action.July.Id,
			&action.August.Id,
			&action.September.Id,
			&action.October.Id,
			&action.November.Id,
			&action.December.Id,
			&action.CreatedById,
			&action.DbStatus,
			&action.DbLastStatus,
			&action.LastNotificationDate,
		)
		action.RelativeFunction, _ = dropDownListItemModel.GetById(action.RelativeFunction.Id)
		action.Confirmation, _ = dropDownListItemModel.GetById(action.Confirmation.Id)
		action.Status, _ = dropDownListItemModel.GetById(action.Status.Id)
		action.VerificationStatus, _ = dropDownListItemModel.GetById(action.VerificationStatus.Id)
		action.January, _ = dropDownListItemModel.GetById(action.January.Id)
		action.February, _ = dropDownListItemModel.GetById(action.February.Id)
		action.March, _ = dropDownListItemModel.GetById(action.March.Id)
		action.April, _ = dropDownListItemModel.GetById(action.April.Id)
		action.May, _ = dropDownListItemModel.GetById(action.May.Id)
		action.June, _ = dropDownListItemModel.GetById(action.June.Id)
		action.July, _ = dropDownListItemModel.GetById(action.July.Id)
		action.August, _ = dropDownListItemModel.GetById(action.August.Id)
		action.September, _ = dropDownListItemModel.GetById(action.September.Id)
		action.October, _ = dropDownListItemModel.GetById(action.October.Id)
		action.November, _ = dropDownListItemModel.GetById(action.November.Id)
		action.December, _ = dropDownListItemModel.GetById(action.December.Id)

		actions = append(actions, action)
	}

	return actions, nil
}

func (*ActionModel) Create(action registerComponentTypes.Action) error {
	db := database.GetDatabase()
	_, err := db.Exec(`
			INSERT INTO actions ( 
				"id",
				"companyId",
				"registerId",
				"registerType",
				"no",
				"title",
				"raiseDate",
				"resources",
				"currency",
				"relativeFunction",
				"responsibleId",
				"approverId",
				"deadline",
				"confirmation",
				"status",
				"completionDate",
				"verificationStatus",
				"comment",
				"january",
				"february",
				"march",
				"april",
				"may",
				"june",
				"july",
				"august",
				"september",
				"october",
				"november",
				"december",
				"createdById",
				"dbStatus",
				"dbLastStatus",
				"lastNotificationDate"
			) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`,
		action.Id,
		action.CompanyId,
		action.RegisterId,
		action.RegisterType,
		action.No,
		action.Title,
		action.RaiseDate,
		action.Resources,
		action.Currency,
		action.RelativeFunction.Id,
		action.ResponsibleId,
		action.ApproverId,
		action.Deadline,
		action.Confirmation.Id,
		action.Status.Id,
		action.CompletionDate,
		action.VerificationStatus.Id,
		action.Comment,
		action.January.Id,
		action.February.Id,
		action.March.Id,
		action.April.Id,
		action.May.Id,
		action.June.Id,
		action.July.Id,
		action.August.Id,
		action.September.Id,
		action.October.Id,
		action.November.Id,
		action.December.Id,
		action.CreatedById,
		action.DbStatus,
		action.DbLastStatus,
		action.LastNotificationDate,
	)

	fmt.Println(err)

	return err
}

func (*ActionModel) Update(Id string, fields map[string]interface{}) error {
	if len(fields) == 0 {
		return nil
	}

	setClause := ""
	values := []interface{}{}

	for key, val := range fields {
		setClause += fmt.Sprintf(` "%s" = ?,`, key)
		values = append(values, val)
	}

	setClause = strings.TrimSuffix(setClause, ",")
	query := fmt.Sprintf(`
			UPDATE actions 
			SET %s 
			WHERE "id" = ?
		`,
		setClause,
	)
	values = append(values, Id)

	db := database.GetDatabase()
	_, err := db.Exec(query, values...)
	return err
}

func (*ActionModel) SendDailyNotifications() {
	var actionModel ActionModel
	var dropDownModel tableComponentModels.DropDownListItemModel

	actions, err := actionModel.GetAll(map[string]interface{}{
		"dbStatus": "active",
	})

	if err != nil {
		return
	}

	today := time.Now().Truncate(24 * time.Hour)
	tomorrow := today.AddDate(0, 0, 1)
	isEndOfMonth := today.Month() != tomorrow.Month()

	for _, action := range actions {
		confirmation, _ := dropDownModel.GetById(action.Confirmation.Id)
		status, _ := dropDownModel.GetById(action.Status.Id)
		relativeFunc, _ := dropDownModel.GetById(action.RelativeFunction.Id)

		responsible := modules.GetAccountById(action.ResponsibleId)
		toContacts := []mailer.EmailContact{{Email: responsible.Email, Name: responsible.Name + " " + responsible.Surname}}

		ccContacts := []mailer.EmailContact{}
		lineManager := modules.GetAccountById(responsible.LineManagerId)
		if !lineManager.IsEmpty() {
			ccContacts = append(ccContacts, mailer.EmailContact{Email: lineManager.Email, Name: lineManager.Name + " " + lineManager.Surname})
		}

		createdBy := modules.GetAccountById(action.CreatedById)
		if !createdBy.IsEmpty() {
			ccContacts = append(ccContacts, mailer.EmailContact{Email: createdBy.Email, Name: createdBy.Name + " " + createdBy.Surname})
		}

		// DÜZƏLİŞ 1: "confirmed" əvəzinə "agreed" yazıldı
		isConfirmed := strings.ToLower(confirmation.Value) == "agreed" || strings.ToLower(confirmation.Value) == "rejected"

		lastNotifDate := action.LastNotificationDate
		if lastNotifDate == "" {
			lastNotifDate = action.RaiseDate
		}

		lastNotif, err := time.Parse("2006-01-02", lastNotifDate)
		if err == nil {
			daysSinceNotif := int(today.Sub(lastNotif).Hours() / 24)

			if !isConfirmed && daysSinceNotif >= 3 {
				subject := fmt.Sprintf("Action Assignment Notification: %s - %s", action.No, action.Title)
				body := fmt.Sprintf(`
                    <p>Dear Recipient,</p>
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
                    <p>Thanks for the prompt response</p>
                `, action.No, action.Title, action.RaiseDate, action.Resources, action.Currency, relativeFunc.Value, responsible.Name, responsible.Surname, status.Value, action.Deadline)

				mailer.SendEmail(toContacts, ccContacts, subject, body)

				actionModel.Update(action.Id, map[string]interface{}{
					"lastNotificationDate": today.Format("2006-01-02"),
				})

				// DÜZƏLİŞ 2: 'continue' silindi ki, kod aşağıdakı Deadline yoxlamalarına keçə bilsin
			}
		}

		deadline, err := time.Parse("2006-01-02", action.Deadline)
		if err == nil {
			daysToDeadline := int(deadline.Sub(today).Hours() / 24)
			isCompleted := strings.Contains(status.Value, "100") || strings.ToLower(status.Value) == "completed"

			if !isCompleted {
				isOneMonthPrior := deadline.AddDate(0, -1, 0).Equal(today)

				if isEndOfMonth || isOneMonthPrior || daysToDeadline == 7 || daysToDeadline == 3 || daysToDeadline == 0 {
					subject := fmt.Sprintf("Action Reminder: %s - Upcoming Deadline", action.No)
					body := fmt.Sprintf(`
                        <p>Dear Recipient,</p>
                        <p>I would like to remind you that the following action has been assigned to you for implementation:</p>
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
                        <p>Please provide an update on the status of the action with evidences of the action progress and highlight the percentage that appropriately reflects the progress status from the list.</p>
                        <p>10%% 20%% 30%% 40%% 50%% 60%% 70%% 80%% 90%% 100%%</p>
                        <br/>
                        <p>Thanks for the prompt response</p>
                    `, action.No, action.Title, action.RaiseDate, action.Resources, action.Currency, relativeFunc.Value, responsible.Name, responsible.Surname, status.Value, action.Deadline)

					mailer.SendEmail(toContacts, ccContacts, subject, body)
				}

				if daysToDeadline < 0 {
					overdueDays := -daysToDeadline
					if overdueDays%7 == 0 {
						subject := fmt.Sprintf("DELAYED ACTION WARNING: %s - Overdue by %d Days", action.No, overdueDays)
						body := fmt.Sprintf(`
                            <p>Dear Recipient,</p>
                            <p>I would like to inform you that the following action has been delayed for %d days:</p>
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
                            <p>Please provide an update on the status of the action with evidences of the action progress and highlight the percentage that appropriately reflects the progress status from the list.</p>
                            <p>10%% 20%% 30%% 40%% 50%% 60%% 70%% 80%% 90%% 100%%</p>
                            <br/>
                            <p>Thanks for the prompt response</p>
                        `, overdueDays, action.No, action.Title, action.RaiseDate, action.Resources, action.Currency, relativeFunc.Value, responsible.Name, responsible.Surname, status.Value, action.Deadline)

						mailer.SendEmail(toContacts, ccContacts, subject, body)
					}
				}
			}
		}
	}
}
