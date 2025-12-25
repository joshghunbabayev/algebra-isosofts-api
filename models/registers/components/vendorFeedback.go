package registerComponentModels

import (
	"algebra-isosofts-api/database"
	tableComponentModels "algebra-isosofts-api/models/tableComponents"
	"algebra-isosofts-api/modules"
	registerComponentTypes "algebra-isosofts-api/types/registers/components"
	"fmt"
	"strings"
)

type VendorFeedbackModel struct {
}

func (*VendorFeedbackModel) GenerateUniqueId() string {
	Id := modules.GenerateRandomString(30)

	var vendorFeedbackModel VendorFeedbackModel

	vendorFeedback, _ := vendorFeedbackModel.GetById(Id)

	if vendorFeedback.IsEmpty() {
		return Id
	} else {
		return vendorFeedbackModel.GenerateUniqueId()
	}
}

func (*VendorFeedbackModel) GetById(Id string) (registerComponentTypes.VendorFeedback, error) {
	db := database.GetDatabase()
	row := db.QueryRow(`
			SELECT * 
			FROM vendorFeedbacks
			WHERE id = ?
		`,
		Id,
	)

	var vendorFeedback registerComponentTypes.VendorFeedback
	var dropDownListItemModel tableComponentModels.DropDownListItemModel

	err := row.Scan(
		&vendorFeedback.Id,
		&vendorFeedback.RegisterId,
		&vendorFeedback.Scope.Id,
		&vendorFeedback.VendorId,
		&vendorFeedback.TypeOfFinding.Id,
		&vendorFeedback.QGS,
		&vendorFeedback.Communication,
		&vendorFeedback.OTD,
		&vendorFeedback.Documentation,
		&vendorFeedback.HS,
		&vendorFeedback.Environment,
		&vendorFeedback.DbStatus,
		&vendorFeedback.DbLastStatus,
	)
	vendorFeedback.Scope, _ = dropDownListItemModel.GetById(vendorFeedback.Scope.Id)
	vendorFeedback.TypeOfFinding, _ = dropDownListItemModel.GetById(vendorFeedback.TypeOfFinding.Id)

	return vendorFeedback, err
}

func (*VendorFeedbackModel) GetAll(filters map[string]interface{}) ([]registerComponentTypes.VendorFeedback, error) {
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
			SELECT * FROM vendorFeedbacks %s
		`,
		whereClause,
	)
	rows, err := db.Query(query, values...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vendorFeedbacks []registerComponentTypes.VendorFeedback

	for rows.Next() {
		var vendorFeedback registerComponentTypes.VendorFeedback
		var dropDownListItemModel tableComponentModels.DropDownListItemModel

		rows.Scan(
			&vendorFeedback.Id,
			&vendorFeedback.RegisterId,
			&vendorFeedback.Scope.Id,
			&vendorFeedback.VendorId,
			&vendorFeedback.TypeOfFinding.Id,
			&vendorFeedback.QGS,
			&vendorFeedback.Communication,
			&vendorFeedback.OTD,
			&vendorFeedback.Documentation,
			&vendorFeedback.HS,
			&vendorFeedback.Environment,
			&vendorFeedback.DbStatus,
			&vendorFeedback.DbLastStatus,
		)
		vendorFeedback.Scope, _ = dropDownListItemModel.GetById(vendorFeedback.Scope.Id)
		vendorFeedback.TypeOfFinding, _ = dropDownListItemModel.GetById(vendorFeedback.TypeOfFinding.Id)

		vendorFeedbacks = append(vendorFeedbacks, vendorFeedback)
	}

	return vendorFeedbacks, nil
}

func (*VendorFeedbackModel) Create(vendorFeedback registerComponentTypes.VendorFeedback) error {
	db := database.GetDatabase()
	_, err := db.Exec(`
			INSERT INTO vendorFeedbacks ( 
				"id",
				"registerId",
				"scope",
				"vendorId",
				"typeOfFinding",
				"qgs",
				"communication",
				"otd",
				"documentation",
				"hs",
				"environment",
				"dbStatus",
				"dbLastStatus"
			) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`,
		vendorFeedback.Id,
		vendorFeedback.RegisterId,
		vendorFeedback.Scope.Id,
		vendorFeedback.VendorId,
		vendorFeedback.TypeOfFinding.Id,
		vendorFeedback.QGS,
		vendorFeedback.Communication,
		vendorFeedback.OTD,
		vendorFeedback.Documentation,
		vendorFeedback.HS,
		vendorFeedback.Environment,
		vendorFeedback.DbStatus,
		vendorFeedback.DbLastStatus,
	)

	fmt.Println(err)

	return err
}

func (*VendorFeedbackModel) Update(Id string, fields map[string]interface{}) error {
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
			UPDATE vendorFeedbacks 
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
