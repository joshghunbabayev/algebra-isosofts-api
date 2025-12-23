package registerModels

import (
	"algebra-isosofts-api/database"
	registerComponentModels "algebra-isosofts-api/models/registerComponents"
	tableComponentModels "algebra-isosofts-api/models/tableComponents"
	"algebra-isosofts-api/modules"
	registerTypes "algebra-isosofts-api/types/registers"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type FINModel struct {
}

func (*FINModel) GenerateUniqueId() string {
	Id := modules.GenerateRandomString(30)

	var finModel FINModel

	fin, _ := finModel.GetById(Id)

	if fin.IsEmpty() {
		return Id
	} else {
		return finModel.GenerateUniqueId()
	}
}

func (*FINModel) GenerateUniqueNo() string {
	db := database.GetDatabase()

	year := time.Now().Format("06")

	var lastNo string
	db.QueryRow(`
		SELECT "no" 
		FROM finregisters 
		WHERE "no" LIKE ? 
		ORDER BY "no" DESC 
		LIMIT 1
		`,
		"FINR/"+year+"/%",
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

	newNo := fmt.Sprintf("FINR/%s/%04d", year, nextNumber)
	return newNo
}

func (*FINModel) GetById(Id string) (registerTypes.FIN, error) {
	db := database.GetDatabase()
	row := db.QueryRow(`
			SELECT * 
			FROM finregisters
			WHERE id = ?
		`,
		Id,
	)

	var fin registerTypes.FIN
	var dropDownListItemModel tableComponentModels.DropDownListItemModel
	var actionModel registerComponentModels.ActionModel

	err := row.Scan(
		&fin.Id,
		&fin.No,
		&fin.Issuer,
		&fin.Process.Id,
		&fin.CategoryOfFinding.Id,
		&fin.TypeOfFinding.Id,
		&fin.SourceOfFinding.Id,
		&fin.Customer,
		&fin.Vendor,
		&fin.Description,
		&fin.ContainmentAction,
		&fin.RootCauses,
		&fin.DbStatus,
		&fin.DbLastStatus,
	)
	fin.Process, _ = dropDownListItemModel.GetById(fin.Process.Id)
	fin.CategoryOfFinding, _ = dropDownListItemModel.GetById(fin.CategoryOfFinding.Id)
	fin.TypeOfFinding, _ = dropDownListItemModel.GetById(fin.TypeOfFinding.Id)
	fin.SourceOfFinding, _ = dropDownListItemModel.GetById(fin.SourceOfFinding.Id)
	fin.Actions, _ = actionModel.GetAll(map[string]interface{}{
		"registerId": fin.Id,
	})

	return fin, err
}

func (*FINModel) GetAll(filters map[string]interface{}) ([]registerTypes.FIN, error) {
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
			SELECT * FROM finregisters %s
		`,
		whereClause,
	)
	rows, err := db.Query(query, values...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var fins []registerTypes.FIN

	for rows.Next() {
		var fin registerTypes.FIN
		var dropDownListItemModel tableComponentModels.DropDownListItemModel
		var actionModel registerComponentModels.ActionModel

		rows.Scan(
			&fin.Id,
			&fin.No,
			&fin.Issuer,
			&fin.Process.Id,
			&fin.CategoryOfFinding.Id,
			&fin.TypeOfFinding.Id,
			&fin.SourceOfFinding.Id,
			&fin.Customer,
			&fin.Vendor,
			&fin.Description,
			&fin.ContainmentAction,
			&fin.RootCauses,
			&fin.DbStatus,
			&fin.DbLastStatus,
		)
		fin.Process, _ = dropDownListItemModel.GetById(fin.Process.Id)
		fin.CategoryOfFinding, _ = dropDownListItemModel.GetById(fin.CategoryOfFinding.Id)
		fin.TypeOfFinding, _ = dropDownListItemModel.GetById(fin.TypeOfFinding.Id)
		fin.SourceOfFinding, _ = dropDownListItemModel.GetById(fin.SourceOfFinding.Id)
		fin.Actions, _ = actionModel.GetAll(map[string]interface{}{
			"registerId": fin.Id,
		})

		fins = append(fins, fin)
	}

	return fins, nil
}

func (*FINModel) Create(fin registerTypes.FIN) error {
	db := database.GetDatabase()
	_, err := db.Exec(`
			INSERT INTO finregisters ( 
				"id",
				"no",
				"issuer", 
				"process", 
				"categoryOfFinding", 
				"typeOfFinding", 
				"sourceOfFinding", 
				"customer", 
				"vendor", 
				"description", 
				"containmentAction", 
				"rootCauses", 
				"dbStatus",
				"dbLastStatus"
			) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`,
		fin.Id,
		fin.No,
		fin.Issuer,
		fin.Process.Id,
		fin.CategoryOfFinding.Id,
		fin.TypeOfFinding.Id,
		fin.SourceOfFinding.Id,
		fin.Customer,
		fin.Vendor,
		fin.Description,
		fin.ContainmentAction,
		fin.RootCauses,
		fin.DbStatus,
		fin.DbLastStatus,
	)

	if err != nil {
		return err
	}

	return nil
}

func (*FINModel) Update(Id string, fields map[string]interface{}) error {
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
			UPDATE finregisters 
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
