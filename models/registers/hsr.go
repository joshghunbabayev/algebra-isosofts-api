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

type HSRModel struct {
}

func (*HSRModel) GenerateUniqueId() string {
	Id := modules.GenerateRandomString(30)

	var hsrModel HSRModel

	hsr, _ := hsrModel.GetById(Id)

	if hsr.IsEmpty() {
		return Id
	} else {
		return hsrModel.GenerateUniqueId()
	}
}

func (*HSRModel) GenerateUniqueNo() string {
	db := database.GetDatabase()

	year := time.Now().Format("06")

	var lastNo string
	db.QueryRow(`
		SELECT "no" 
		FROM hsrregisters 
		WHERE "no" LIKE ? 
		ORDER BY "no" DESC 
		LIMIT 1
		`,
		"HSRR/"+year+"/%",
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

	newNo := fmt.Sprintf("HSRR/%s/%04d", year, nextNumber)
	return newNo
}

func (*HSRModel) GetById(Id string) (registerTypes.HSR, error) {
	db := database.GetDatabase()
	row := db.QueryRow(`
			SELECT * 
			FROM hsrregisters
			WHERE id = ?
		`,
		Id,
	)

	var hsr registerTypes.HSR
	var dropDownListItemModel tableComponentModels.DropDownListItemModel
	var actionModel registerComponentModels.ActionModel

	err := row.Scan(
		&hsr.Id,
		&hsr.No,
		&hsr.Process.Id,
		&hsr.Hazard.Id,
		&hsr.Risk.Id,
		&hsr.AffectedPositions.Id,
		&hsr.ERMA,
		&hsr.InitialRiskSeverity,
		&hsr.InitialRiskLikelyhood,
		&hsr.ResidualRiskSeverity,
		&hsr.ResidualRiskLikelyhood,
		&hsr.DbStatus,
		&hsr.DbLastStatus,
	)
	hsr.Process, _ = dropDownListItemModel.GetById(hsr.Process.Id)
	hsr.Hazard, _ = dropDownListItemModel.GetById(hsr.Hazard.Id)
	hsr.Risk, _ = dropDownListItemModel.GetById(hsr.Risk.Id)
	hsr.AffectedPositions, _ = dropDownListItemModel.GetById(hsr.AffectedPositions.Id)
	hsr.Actions, _ = actionModel.GetAll(map[string]interface{}{
		"registerId": hsr.Id,
	})

	return hsr, err
}

func (*HSRModel) GetAll(filters map[string]interface{}) ([]registerTypes.HSR, error) {
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
			SELECT * FROM hsrregisters %s
		`,
		whereClause,
	)
	rows, err := db.Query(query, values...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var hsrs []registerTypes.HSR

	for rows.Next() {
		var hsr registerTypes.HSR
		var dropDownListItemModel tableComponentModels.DropDownListItemModel
		var actionModel registerComponentModels.ActionModel

		rows.Scan(
			&hsr.Id,
			&hsr.No,
			&hsr.Process.Id,
			&hsr.Hazard.Id,
			&hsr.Risk.Id,
			&hsr.AffectedPositions.Id,
			&hsr.ERMA,
			&hsr.InitialRiskSeverity,
			&hsr.InitialRiskLikelyhood,
			&hsr.ResidualRiskSeverity,
			&hsr.ResidualRiskLikelyhood,
			&hsr.DbStatus,
			&hsr.DbLastStatus,
		)
		hsr.Process, _ = dropDownListItemModel.GetById(hsr.Process.Id)
		hsr.Hazard, _ = dropDownListItemModel.GetById(hsr.Hazard.Id)
		hsr.Risk, _ = dropDownListItemModel.GetById(hsr.Risk.Id)
		hsr.AffectedPositions, _ = dropDownListItemModel.GetById(hsr.AffectedPositions.Id)
		hsr.Actions, _ = actionModel.GetAll(map[string]interface{}{
			"registerId": hsr.Id,
		})

		hsrs = append(hsrs, hsr)
	}

	return hsrs, nil
}

func (*HSRModel) Create(hsr registerTypes.HSR) error {
	db := database.GetDatabase()
	_, err := db.Exec(`
			INSERT INTO hsrregisters ( 
				"id",
				"no",
				"process", 
				"hazard", 
				"risk", 
				"affectedPositions",
				"erma",
				"initialRiskSeverity", 
				"initialRiskLikelyhood", 
				"residualRiskSeverity", 
				"residualRiskLikelyhood",
				"dbStatus",
				"dbLastStatus"
			) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`,
		hsr.Id,
		hsr.No,
		hsr.Process.Id,
		hsr.Hazard.Id,
		hsr.Risk.Id,
		hsr.AffectedPositions.Id,
		hsr.ERMA,
		hsr.InitialRiskSeverity,
		hsr.InitialRiskLikelyhood,
		hsr.ResidualRiskSeverity,
		hsr.ResidualRiskLikelyhood,
		hsr.DbStatus,
		hsr.DbLastStatus,
	)

	if err != nil {
		return err
	}

	return nil
}

func (*HSRModel) Update(Id string, fields map[string]interface{}) error {
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
			UPDATE hsrregisters 
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
