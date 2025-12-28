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

type MRMModel struct {
}

func (*MRMModel) GenerateUniqueId() string {
	Id := modules.GenerateRandomString(30)

	var mrmModel MRMModel

	mrm, _ := mrmModel.GetById(Id)

	if mrm.IsEmpty() {
		return Id
	} else {
		return mrmModel.GenerateUniqueId()
	}
}

func (*MRMModel) GenerateUniqueNo() string {
	db := database.GetDatabase()

	year := time.Now().Format("06")

	var lastNo string
	db.QueryRow(`
		SELECT "no" 
		FROM mrmregisters 
		WHERE "no" LIKE ? 
		ORDER BY "no" DESC 
		LIMIT 1
		`,
		"MRMR/"+year+"/%",
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

	newNo := fmt.Sprintf("MRMR/%s/%04d", year, nextNumber)
	return newNo
}

func (*MRMModel) GetById(Id string) (registerTypes.MRM, error) {
	db := database.GetDatabase()
	row := db.QueryRow(`
			SELECT * 
			FROM mrmregisters
			WHERE id = ?
		`,
		Id,
	)

	var mrm registerTypes.MRM
	var dropDownListItemModel tableComponentModels.DropDownListItemModel
	var actionModel registerComponentModels.ActionModel

	err := row.Scan(
		&mrm.Id,
		&mrm.No,
		&mrm.RISOS.Id,
		&mrm.Topic.Id,
		&mrm.Process.Id,
		&mrm.DbStatus,
		&mrm.DbLastStatus,
	)
	mrm.RISOS, _ = dropDownListItemModel.GetById(mrm.RISOS.Id)
	mrm.Topic, _ = dropDownListItemModel.GetById(mrm.Topic.Id)
	mrm.Process, _ = dropDownListItemModel.GetById(mrm.Process.Id)
	mrm.Process, _ = dropDownListItemModel.GetById(mrm.Process.Id)
	mrm.Actions, _ = actionModel.GetAll(map[string]interface{}{
		"registerId": mrm.Id,
		"dbStatus":   "active",
	})

	return mrm, err
}

func (*MRMModel) GetAll(filters map[string]interface{}) ([]registerTypes.MRM, error) {
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
			SELECT * FROM mrmregisters %s
		`,
		whereClause,
	)
	rows, err := db.Query(query, values...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var mrms []registerTypes.MRM

	for rows.Next() {
		var mrm registerTypes.MRM
		var dropDownListItemModel tableComponentModels.DropDownListItemModel
		var actionModel registerComponentModels.ActionModel

		rows.Scan(
			&mrm.Id,
			&mrm.No,
			&mrm.RISOS.Id,
			&mrm.Topic.Id,
			&mrm.Process.Id,
			&mrm.DbStatus,
			&mrm.DbLastStatus,
		)
		mrm.RISOS, _ = dropDownListItemModel.GetById(mrm.RISOS.Id)
		mrm.Topic, _ = dropDownListItemModel.GetById(mrm.Topic.Id)
		mrm.Process, _ = dropDownListItemModel.GetById(mrm.Process.Id)
		mrm.Process, _ = dropDownListItemModel.GetById(mrm.Process.Id)
		mrm.Actions, _ = actionModel.GetAll(map[string]interface{}{
			"registerId": mrm.Id,
			"dbStatus":   "active",
		})

		mrms = append(mrms, mrm)
	}

	return mrms, nil
}

func (*MRMModel) Create(mrm registerTypes.MRM) error {
	db := database.GetDatabase()
	_, err := db.Exec(`
			INSERT INTO mrmregisters ( 
				"id",
				"no",
				"risos", 
				"topic", 
				"process", 
				"dbStatus",
				"dbLastStatus"
			) VALUES (?, ?, ?, ?, ?, ?, ?)
		`,
		mrm.Id,
		mrm.No,
		mrm.RISOS.Id,
		mrm.Topic.Id,
		mrm.Process.Id,
		mrm.DbStatus,
		mrm.DbLastStatus,
	)

	if err != nil {
		return err
	}

	return nil
}

func (*MRMModel) Update(Id string, fields map[string]interface{}) error {
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
			UPDATE mrmregisters 
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
