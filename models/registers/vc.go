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

type VCModel struct {
}

func (*VCModel) GenerateUniqueId() string {
	Id := modules.GenerateRandomString(30)

	var vcModel VCModel

	vc, _ := vcModel.GetById(Id)

	if vc.IsEmpty() {
		return Id
	} else {
		return vcModel.GenerateUniqueId()
	}
}

func (*VCModel) GenerateUniqueNo() string {
	db := database.GetDatabase()

	year := time.Now().Format("06")

	var lastNo string
	db.QueryRow(`
		SELECT "no" 
		FROM vcregisters 
		WHERE "no" LIKE ? 
		ORDER BY "no" DESC 
		LIMIT 1
		`,
		"VCR/"+year+"/%",
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

	newNo := fmt.Sprintf("VCR/%s/%04d", year, nextNumber)
	return newNo
}

func (*VCModel) GetById(Id string) (registerTypes.VC, error) {
	db := database.GetDatabase()
	row := db.QueryRow(`
			SELECT * 
			FROM vcregisters
			WHERE id = ?
		`,
		Id,
	)

	var vc registerTypes.VC
	var dropDownListItemModel tableComponentModels.DropDownListItemModel
	var actionModel registerComponentModels.ActionModel

	err := row.Scan(
		&vc.Id,
		&vc.No,
		&vc.Name,
		&vc.RegNumber,
		&vc.Scope1.Id,
		&vc.Scope2.Id,
		&vc.Scope3.Id,
		&vc.RegistrationDate,
		&vc.ReviewDate,
		&vc.ApprovedActual,
		&vc.QGS,
		&vc.Communication,
		&vc.OTD,
		&vc.Documentation,
		&vc.HS,
		&vc.Environment,
		&vc.DbStatus,
		&vc.DbLastStatus,
	)
	vc.Scope1, _ = dropDownListItemModel.GetById(vc.Scope1.Id)
	vc.Scope2, _ = dropDownListItemModel.GetById(vc.Scope2.Id)
	vc.Scope3, _ = dropDownListItemModel.GetById(vc.Scope3.Id)
	vc.Actions, _ = actionModel.GetAll(map[string]interface{}{
		"registerId": vc.Id,
	})

	return vc, err
}

func (*VCModel) GetAll(filters map[string]interface{}) ([]registerTypes.VC, error) {
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
			SELECT * FROM vcregisters %s
		`,
		whereClause,
	)
	rows, err := db.Query(query, values...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vcs []registerTypes.VC

	for rows.Next() {
		var vc registerTypes.VC
		var dropDownListItemModel tableComponentModels.DropDownListItemModel
		var actionModel registerComponentModels.ActionModel

		rows.Scan(
			&vc.Id,
			&vc.No,
			&vc.Name,
			&vc.RegNumber,
			&vc.Scope1.Id,
			&vc.Scope2.Id,
			&vc.Scope3.Id,
			&vc.RegistrationDate,
			&vc.ReviewDate,
			&vc.ApprovedActual,
			&vc.QGS,
			&vc.Communication,
			&vc.OTD,
			&vc.Documentation,
			&vc.HS,
			&vc.Environment,
			&vc.DbStatus,
			&vc.DbLastStatus,
		)
		vc.Scope1, _ = dropDownListItemModel.GetById(vc.Scope1.Id)
		vc.Scope2, _ = dropDownListItemModel.GetById(vc.Scope2.Id)
		vc.Scope3, _ = dropDownListItemModel.GetById(vc.Scope3.Id)
		vc.Actions, _ = actionModel.GetAll(map[string]interface{}{
			"registerId": vc.Id,
		})

		vcs = append(vcs, vc)
	}

	return vcs, nil
}

func (*VCModel) Create(vc registerTypes.VC) error {
	db := database.GetDatabase()
	_, err := db.Exec(`
			INSERT INTO vcregisters ( 
				"id",
				"no",
				"name", 
				"regNumber", 
				"scope1", 
				"scope2", 
				"scope3", 
				"registrationDate", 
				"reviewDate", 
				"approvedActual", 
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
		vc.Id,
		vc.No,
		vc.Name,
		vc.RegNumber,
		vc.Scope1.Id,
		vc.Scope2.Id,
		vc.Scope3.Id,
		vc.RegistrationDate,
		vc.ReviewDate,
		vc.ApprovedActual,
		vc.QGS,
		vc.Communication,
		vc.OTD,
		vc.Documentation,
		vc.HS,
		vc.Environment,
		vc.DbStatus,
		vc.DbLastStatus,
	)

	if err != nil {
		return err
	}

	return nil
}

func (*VCModel) Update(Id string, fields map[string]interface{}) error {
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
			UPDATE vcregisters 
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
