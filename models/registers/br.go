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

type BrModel struct {
}

func (*BrModel) GenerateUniqueId() string {
	Id := modules.GenerateRandomString(30)

	var brModel BrModel

	br, _ := brModel.GetById(Id)

	if br.IsEmpty() {
		return Id
	} else {
		return brModel.GenerateUniqueId()
	}
}

func (*BrModel) GenerateUniqueNo() string {
	db := database.GetDatabase()

	year := time.Now().Format("06")

	var lastNo string
	db.QueryRow(`
		SELECT "no" 
		FROM brregisters 
		WHERE "no" LIKE ? 
		ORDER BY "no" DESC 
		LIMIT 1
		`,
		"BRR/"+year+"/%",
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

	newNo := fmt.Sprintf("BRR/%s/%04d", year, nextNumber)
	return newNo
}

func (*BrModel) GetById(Id string) (registerTypes.Br, error) {
	db := database.GetDatabase()
	row := db.QueryRow(`
			SELECT * 
			FROM brregisters
			WHERE id = ?
		`,
		Id,
	)

	var br registerTypes.Br
	var dropDownListItemModel tableComponentModels.DropDownListItemModel
	var actionModel registerComponentModels.ActionModel

	err := row.Scan(
		&br.Id,
		&br.No,
		&br.Swot.Id,
		&br.Pestle.Id,
		&br.InterestedParty.Id,
		&br.RiskOpportunity,
		&br.Objective,
		&br.KPI,
		&br.Process.Id,
		&br.ERMEOA,
		&br.InitialRiskSeverity,
		&br.InitialRiskLikelyhood,
		&br.ResidualRiskSeverity,
		&br.ResidualRiskLikelyhood,
		&br.DbStatus,
		&br.DbLastStatus,
	)
	br.Swot, _ = dropDownListItemModel.GetById(br.Swot.Id)
	br.Pestle, _ = dropDownListItemModel.GetById(br.Pestle.Id)
	br.InterestedParty, _ = dropDownListItemModel.GetById(br.InterestedParty.Id)
	br.Process, _ = dropDownListItemModel.GetById(br.Process.Id)
	br.Actions, _ = actionModel.GetAll(map[string]interface{}{
		"registerId": br.Id,
	})

	return br, err
}

func (*BrModel) GetAll(filters map[string]interface{}) ([]registerTypes.Br, error) {
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
			SELECT * FROM brregisters %s
		`,
		whereClause,
	)
	rows, err := db.Query(query, values...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var brs []registerTypes.Br

	for rows.Next() {
		var br registerTypes.Br
		var dropDownListItemModel tableComponentModels.DropDownListItemModel
		var actionModel registerComponentModels.ActionModel

		rows.Scan(
			&br.Id,
			&br.No,
			&br.Swot.Id,
			&br.Pestle.Id,
			&br.InterestedParty.Id,
			&br.RiskOpportunity,
			&br.Objective,
			&br.KPI,
			&br.Process.Id,
			&br.ERMEOA,
			&br.InitialRiskSeverity,
			&br.InitialRiskLikelyhood,
			&br.ResidualRiskSeverity,
			&br.ResidualRiskLikelyhood,
			&br.DbStatus,
			&br.DbLastStatus,
		)
		br.Swot, _ = dropDownListItemModel.GetById(br.Swot.Id)
		br.Pestle, _ = dropDownListItemModel.GetById(br.Pestle.Id)
		br.InterestedParty, _ = dropDownListItemModel.GetById(br.InterestedParty.Id)
		br.Process, _ = dropDownListItemModel.GetById(br.Process.Id)
		br.Actions, _ = actionModel.GetAll(map[string]interface{}{
			"registerId": br.Id,
		})

		brs = append(brs, br)
	}

	return brs, nil
}

func (*BrModel) Create(br registerTypes.Br) error {
	db := database.GetDatabase()
	_, err := db.Exec(`
			INSERT INTO brregisters ( 
				"id",
				"no",
				"swot", 
				"pestle", 
				"interestedParty", 
				"riskOpportunity", 
				"objective", 
				"kpi", 
				"process", 
				"ermeoa", 
				"initialRiskSeverity", 
				"initialRiskLikelyhood", 
				"residualRiskSeverity", 
				"residualRiskLikelyhood",
				"dbStatus",
				"dbLastStatus"
			) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`,
		br.Id,
		br.No,
		br.Swot.Id,
		br.Pestle.Id,
		br.InterestedParty.Id,
		br.RiskOpportunity,
		br.Objective, br.KPI,
		br.Process.Id,
		br.ERMEOA,
		br.InitialRiskSeverity,
		br.InitialRiskLikelyhood,
		br.ResidualRiskSeverity,
		br.ResidualRiskLikelyhood,
		br.DbStatus,
		br.DbLastStatus,
	)

	if err != nil {
		return err
	}

	return nil
}

func (*BrModel) Update(Id string, fields map[string]interface{}) error {
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
			UPDATE brregisters 
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
