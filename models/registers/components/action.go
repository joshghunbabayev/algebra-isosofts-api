package registerModels

import (
	"algebra-isosofts-api/database"
	tableComponentModels "algebra-isosofts-api/models/tableComponents"
	"algebra-isosofts-api/modules"
	registerTypes "algebra-isosofts-api/types/registers"
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

func (*ActionModel) GetById(Id string) (registerTypes.Br, error) {
	db := database.GetDatabase()
	row := db.QueryRow(`
			SELECT * 
			FROM actions
			WHERE id = ?
		`,
		Id,
	)

	var br registerTypes.Br
	var dropDownListItemModel tableComponentModels.DropDownListItemModel

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

	return br, err
}

func (*ActionModel) GetAll(filters map[string]interface{}) ([]registerTypes.Br, error) {
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

		brs = append(brs, br)
	}

	return brs, nil
}

func (*ActionModel) Create(br registerTypes.Br) error {
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
