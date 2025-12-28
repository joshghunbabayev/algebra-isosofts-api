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

type EIModel struct {
}

func (*EIModel) GenerateUniqueId() string {
	Id := modules.GenerateRandomString(30)

	var eiModel EIModel

	ei, _ := eiModel.GetById(Id)

	if ei.IsEmpty() {
		return Id
	} else {
		return eiModel.GenerateUniqueId()
	}
}

func (*EIModel) GenerateUniqueNo() string {
	db := database.GetDatabase()

	year := time.Now().Format("06")

	var lastNo string
	db.QueryRow(`
		SELECT "no" 
		FROM eiregisters 
		WHERE "no" LIKE ? 
		ORDER BY "no" DESC 
		LIMIT 1
		`,
		"EIR/"+year+"/%",
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

	newNo := fmt.Sprintf("EIR/%s/%04d", year, nextNumber)
	return newNo
}

func (*EIModel) GetById(Id string) (registerTypes.EI, error) {
	db := database.GetDatabase()
	row := db.QueryRow(`
			SELECT * 
			FROM eiregisters
			WHERE id = ?
		`,
		Id,
	)

	var ei registerTypes.EI
	var dropDownListItemModel tableComponentModels.DropDownListItemModel
	var actionModel registerComponentModels.ActionModel

	err := row.Scan(
		&ei.Id,
		&ei.No,
		&ei.Name,
		&ei.SerialNumber,
		&ei.CertificateNo,
		&ei.InspectionFrequency.Id,
		&ei.ICD,
		&ei.NVCD,
		&ei.SafeToUse,
		&ei.DbStatus,
		&ei.DbLastStatus,
	)
	ei.InspectionFrequency, _ = dropDownListItemModel.GetById(ei.InspectionFrequency.Id)
	ei.Actions, _ = actionModel.GetAll(map[string]interface{}{
		"registerId": ei.Id,
		"dbStatus":   "active",
	})

	return ei, err
}

func (*EIModel) GetAll(filters map[string]interface{}) ([]registerTypes.EI, error) {
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
			SELECT * FROM eiregisters %s
		`,
		whereClause,
	)
	rows, err := db.Query(query, values...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var eis []registerTypes.EI

	for rows.Next() {
		var ei registerTypes.EI
		var dropDownListItemModel tableComponentModels.DropDownListItemModel
		var actionModel registerComponentModels.ActionModel

		rows.Scan(
			&ei.Id,
			&ei.No,
			&ei.Name,
			&ei.SerialNumber,
			&ei.CertificateNo,
			&ei.InspectionFrequency.Id,
			&ei.ICD,
			&ei.NVCD,
			&ei.SafeToUse,
			&ei.DbStatus,
			&ei.DbLastStatus,
		)
		ei.InspectionFrequency, _ = dropDownListItemModel.GetById(ei.InspectionFrequency.Id)
		ei.Actions, _ = actionModel.GetAll(map[string]interface{}{
			"registerId": ei.Id,
			"dbStatus":   "active",
		})

		eis = append(eis, ei)
	}

	return eis, nil
}

func (*EIModel) Create(ei registerTypes.EI) error {
	db := database.GetDatabase()
	_, err := db.Exec(`
			INSERT INTO eiregisters ( 
				"id",
				"no",
				"name",
				"serialNumber",
				"certificateNo",
				"inspectionFrequency",
				"icd",
				"nvcd",
				"safeToUse",
				"dbStatus",
				"dbLastStatus"
			) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`,
		ei.Id,
		ei.No,
		ei.Name,
		ei.SerialNumber,
		ei.CertificateNo,
		ei.InspectionFrequency.Id,
		ei.ICD,
		ei.NVCD,
		ei.SafeToUse,
		ei.DbStatus,
		ei.DbLastStatus,
	)

	if err != nil {
		return err
	}

	return nil
}

func (*EIModel) Update(Id string, fields map[string]interface{}) error {
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
			UPDATE eiregisters 
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
