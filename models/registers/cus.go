package registerModels

import (
	"algebra-isosofts-api/database"
	registerComponentModels "algebra-isosofts-api/models/registers/components"
	tableComponentModels "algebra-isosofts-api/models/tableComponents"
	"algebra-isosofts-api/modules"
	registerTypes "algebra-isosofts-api/types/registers"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type CUSModel struct {
}

func (*CUSModel) GenerateUniqueId() string {
	Id := modules.GenerateRandomString(30)

	var cusModel CUSModel

	cus, _ := cusModel.GetById(Id)

	if cus.IsEmpty() {
		return Id
	} else {
		return cusModel.GenerateUniqueId()
	}
}

func (*CUSModel) GenerateUniqueNo() string {
	db := database.GetDatabase()

	year := time.Now().Format("06")

	var lastNo string
	db.QueryRow(`
		SELECT "no" 
		FROM cusregisters 
		WHERE "no" LIKE ? 
		ORDER BY "no" DESC 
		LIMIT 1
		`,
		"CUSR/"+year+"/%",
	).Scan(&lastNo)

	var nextNumber int
	if lastNo == "" {
		nextNumber = 1
	} else {
		parts := strings.Split(lastNo, "/")
		numPart := parts[2]
		num, _ := strconv.Atoi(numPart)
		nextNumber = num + 1
	}

	newNo := fmt.Sprintf("CUSR/%s/%04d", year, nextNumber)
	return newNo
}

func (*CUSModel) GetById(Id string) (registerTypes.CUS, error) {
	db := database.GetDatabase()
	row := db.QueryRow(`
			SELECT * 
			FROM cusregisters
			WHERE id = ?
		`,
		Id,
	)

	var cus registerTypes.CUS
	var dropDownListItemModel tableComponentModels.DropDownListItemModel
	var actionModel registerComponentModels.ActionModel

	err := row.Scan(
		&cus.Id,
		&cus.No,
		&cus.Name,
		&cus.RegNumber,
		&cus.Scope1.Id,
		&cus.Scope2.Id,
		&cus.Scope3.Id,
		&cus.RegistrationDate,
		&cus.ReviewDate,
		&cus.Actual,
		&cus.QGS,
		&cus.Communication,
		&cus.OTD,
		&cus.Documentation,
		&cus.HS,
		&cus.Environment,
		&cus.DbStatus,
		&cus.DbLastStatus,
	)
	cus.Scope1, _ = dropDownListItemModel.GetById(cus.Scope1.Id)
	cus.Scope2, _ = dropDownListItemModel.GetById(cus.Scope2.Id)
	cus.Scope3, _ = dropDownListItemModel.GetById(cus.Scope3.Id)
	cus.Actions, _ = actionModel.GetAll(map[string]interface{}{
		"registerId": cus.Id,
	})

	return cus, err
}

func (*CUSModel) GetAll(filters map[string]interface{}) ([]registerTypes.CUS, error) {
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
			SELECT * FROM cusregisters %s
		`,
		whereClause,
	)
	rows, err := db.Query(query, values...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cuss []registerTypes.CUS

	for rows.Next() {
		var cus registerTypes.CUS
		var dropDownListItemModel tableComponentModels.DropDownListItemModel
		var actionModel registerComponentModels.ActionModel

		rows.Scan(
			&cus.Id,
			&cus.No,
			&cus.Name,
			&cus.RegNumber,
			&cus.Scope1.Id,
			&cus.Scope2.Id,
			&cus.Scope3.Id,
			&cus.RegistrationDate,
			&cus.ReviewDate,
			&cus.Actual,
			&cus.QGS,
			&cus.Communication,
			&cus.OTD,
			&cus.Documentation,
			&cus.HS,
			&cus.Environment,
			&cus.DbStatus,
			&cus.DbLastStatus,
		)
		cus.Scope1, _ = dropDownListItemModel.GetById(cus.Scope1.Id)
		cus.Scope2, _ = dropDownListItemModel.GetById(cus.Scope2.Id)
		cus.Scope3, _ = dropDownListItemModel.GetById(cus.Scope3.Id)
		cus.Actions, _ = actionModel.GetAll(map[string]interface{}{
			"registerId": cus.Id,
		})

		cuss = append(cuss, cus)
	}

	return cuss, nil
}

func (*CUSModel) Create(cus registerTypes.CUS) error {
	db := database.GetDatabase()
	_, err := db.Exec(`
			INSERT INTO cusregisters ( 
				"id",
				"no",
				"name", 
				"regNumber", 
				"scope1", 
				"scope2", 
				"scope3", 
				"registrationDate", 
				"reviewDate", 
				"actual", 
				"qgs", 
				"communication", 
				"otd", 
				"documentation", 
				"hs", 
				"environment", 
				"dbStatus",
				"dbLastStatus"
			) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`,
		cus.Id,
		cus.No,
		cus.Name,
		cus.RegNumber,
		cus.Scope1.Id,
		cus.Scope2.Id,
		cus.Scope3.Id,
		cus.RegistrationDate,
		cus.ReviewDate,
		cus.Actual,
		0,
		0,
		0,
		0,
		0,
		0,
		cus.DbStatus,
		cus.DbLastStatus,
	)

	if err != nil {
		return err
	}

	return nil
}

func (*CUSModel) Update(Id string, fields map[string]interface{}) error {
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
			UPDATE cusregisters 
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
