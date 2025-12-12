package registerComponentModels

import (
	"algebra-isosofts-api/database"
	tableComponentModels "algebra-isosofts-api/models/tableComponents"
	"algebra-isosofts-api/modules"
	registerComponentTypes "algebra-isosofts-api/types/registers/components"
	"fmt"
	"strings"
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
		&action.RegisterId,
		&action.Title,
		&action.RaiseDate,
		&action.Resources,
		&action.Currency,
		&action.RelativeFunction.Id,
		&action.Responsible.Id,
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
		&action.DbStatus,
		&action.DbLastStatus,
	)
	action.RelativeFunction, _ = dropDownListItemModel.GetById(action.RelativeFunction.Id)
	action.Responsible, _ = dropDownListItemModel.GetById(action.Responsible.Id)
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
			&action.RegisterId,
			&action.Title,
			&action.RaiseDate,
			&action.Resources,
			&action.Currency,
			&action.RelativeFunction.Id,
			&action.Responsible.Id,
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
			&action.DbStatus,
			&action.DbLastStatus,
		)
		action.RelativeFunction, _ = dropDownListItemModel.GetById(action.RelativeFunction.Id)
		action.Responsible, _ = dropDownListItemModel.GetById(action.Responsible.Id)
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
				"registerId",
				"title",
				"raiseDate",
				"resources",
				"currency",
				"relativeFunction",
				"responsible",
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
				"dbStatus",
				"dbLastStatus"
			) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`,
		action.Id,
		action.RegisterId,
		action.Title,
		action.RaiseDate,
		action.Resources,
		action.Currency,
		action.RelativeFunction.Id,
		action.Responsible.Id,
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
		action.DbStatus,
		action.DbLastStatus,
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
