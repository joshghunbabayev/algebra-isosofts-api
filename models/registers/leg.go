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

type LEGModel struct {
}

func (*LEGModel) GenerateUniqueId() string {
	Id := modules.GenerateRandomString(30)

	var legModel LEGModel

	leg, _ := legModel.GetById(Id)

	if leg.IsEmpty() {
		return Id
	} else {
		return legModel.GenerateUniqueId()
	}
}

func (*LEGModel) GenerateUniqueNo() string {
	db := database.GetDatabase()

	year := time.Now().Format("06")

	var lastNo string
	db.QueryRow(`
		SELECT "no" 
		FROM legregisters 
		WHERE "no" LIKE ? 
		ORDER BY "no" DESC 
		LIMIT 1
		`,
		"LEGR/"+year+"/%",
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

	newNo := fmt.Sprintf("LEGR/%s/%04d", year, nextNumber)
	return newNo
}

func (*LEGModel) GetById(Id string) (registerTypes.LEG, error) {
	db := database.GetDatabase()
	row := db.QueryRow(`
			SELECT * 
			FROM legregisters
			WHERE id = ?
		`,
		Id,
	)

	var leg registerTypes.LEG
	var dropDownListItemModel tableComponentModels.DropDownListItemModel
	var actionModel registerComponentModels.ActionModel

	err := row.Scan(
		&leg.Id,
		&leg.No,
		&leg.Process.Id,
		&leg.Legislation,
		&leg.Section,
		&leg.Requirement,
		&leg.RiskOfViolation,
		&leg.AffectedPositions.Id,
		&leg.InitialRiskSeverity,
		&leg.InitialRiskLikelyhood,
		&leg.ResidualRiskSeverity,
		&leg.ResidualRiskLikelyhood,
		&leg.DbStatus,
		&leg.DbLastStatus,
	)
	leg.Process, _ = dropDownListItemModel.GetById(leg.Process.Id)
	leg.AffectedPositions, _ = dropDownListItemModel.GetById(leg.AffectedPositions.Id)
	leg.Actions, _ = actionModel.GetAll(map[string]interface{}{
		"registerId": leg.Id,
	})

	return leg, err
}

func (*LEGModel) GetAll(filters map[string]interface{}) ([]registerTypes.LEG, error) {
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
			SELECT * FROM legregisters %s
		`,
		whereClause,
	)
	rows, err := db.Query(query, values...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var legs []registerTypes.LEG

	for rows.Next() {
		var leg registerTypes.LEG
		var dropDownListItemModel tableComponentModels.DropDownListItemModel
		var actionModel registerComponentModels.ActionModel

		rows.Scan(
			&leg.Id,
			&leg.No,
			&leg.Process.Id,
			&leg.Legislation,
			&leg.Section,
			&leg.Requirement,
			&leg.RiskOfViolation,
			&leg.AffectedPositions.Id,
			&leg.InitialRiskSeverity,
			&leg.InitialRiskLikelyhood,
			&leg.ResidualRiskSeverity,
			&leg.ResidualRiskLikelyhood,
			&leg.DbStatus,
			&leg.DbLastStatus,
		)
		leg.Process, _ = dropDownListItemModel.GetById(leg.Process.Id)
		leg.AffectedPositions, _ = dropDownListItemModel.GetById(leg.AffectedPositions.Id)
		leg.Actions, _ = actionModel.GetAll(map[string]interface{}{
			"registerId": leg.Id,
		})

		legs = append(legs, leg)
	}

	return legs, nil
}

func (*LEGModel) Create(leg registerTypes.LEG) error {
	db := database.GetDatabase()
	_, err := db.Exec(`
			INSERT INTO legregisters ( 
				"id",
				"no",
				"process",
				"legislation", 
				"section", 
				"requirement", 
				"riskOfViolation", 
				"affectedPositions", 
				"initialRiskSeverity", 
				"initialRiskLikelyhood", 
				"residualRiskSeverity", 
				"residualRiskLikelyhood",
				"dbStatus",
				"dbLastStatus"
			) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`,
		leg.Id,
		leg.No,
		leg.Process.Id,
		leg.Legislation,
		leg.Section,
		leg.Requirement,
		leg.RiskOfViolation,
		leg.AffectedPositions.Id,
		leg.InitialRiskSeverity,
		leg.InitialRiskLikelyhood,
		leg.ResidualRiskSeverity,
		leg.ResidualRiskLikelyhood,
		leg.DbStatus,
		leg.DbLastStatus,
	)

	if err != nil {
		return err
	}

	return nil
}

func (*LEGModel) Update(Id string, fields map[string]interface{}) error {
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
			UPDATE legregisters 
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
