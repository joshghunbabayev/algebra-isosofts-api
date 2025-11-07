package registerModels

import (
	"algebra-isosofts-api/database"
	tableComponentModels "algebra-isosofts-api/models/tableComponents"
	"algebra-isosofts-api/modules"
	registerTypes "algebra-isosofts-api/types/registers"
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

func (*BrModel) GetById(Id string) (registerTypes.Br, error) {
	db := database.GetDatabase()
	row := db.QueryRow(`
		SELECT * 
		FROM brregisters
		WHERE id = ?`,
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
	)
	br.Swot, _ = dropDownListItemModel.GetById(br.Swot.Id)
	br.Pestle, _ = dropDownListItemModel.GetById(br.Pestle.Id)
	br.InterestedParty, _ = dropDownListItemModel.GetById(br.InterestedParty.Id)
	br.Process, _ = dropDownListItemModel.GetById(br.Process.Id)

	return br, err
}

func (*BrModel) GetAll() ([]registerTypes.Br, error) {
	db := database.GetDatabase()
	rows, err := db.Query(`
		SELECT * 
		FROM brregisters`,
	)

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
		)
		br.Swot, _ = dropDownListItemModel.GetById(br.Swot.Id)
		br.Pestle, _ = dropDownListItemModel.GetById(br.Pestle.Id)
		br.InterestedParty, _ = dropDownListItemModel.GetById(br.InterestedParty.Id)
		br.Process, _ = dropDownListItemModel.GetById(br.Process.Id)

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
			"residualRiskLikelyhood"
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		br.Id,
		br.No,
		br.Swot.Id,
		br.Pestle.Id,
		br.InterestedParty.Id,
		br.RiskOpportunity,
		br.Objective,
		br.KPI,
		br.Process.Id,
		br.ERMEOA,
		br.InitialRiskSeverity,
		br.InitialRiskLikelyhood,
		br.ResidualRiskSeverity,
		br.ResidualRiskLikelyhood,
	)

	if err != nil {
		return err
	}

	return nil
}

func (*BrModel) Update(Id string, br registerTypes.Br) error {
	db := database.GetDatabase()

	_, err := db.Exec(`
		UPDATE brregisters 
		SET 
			"swot" = ?,
			"pestle" = ?,
			"interestedParty" = ?,
			"riskOpportunity" = ?,
			"objective" = ?,
			"kpi" = ?,
			"process" = ?,
			"ermeoa" = ?,
			"initialRiskSeverity" = ?,
			"initialRiskLikelyhood" = ?,
			"residualRiskSeverity" = ?,
			"residualRiskLikelyhood" = ?
		WHERE "id" = ?`,
		br.Swot.Id,
		br.Pestle.Id,
		br.InterestedParty.Id,
		br.RiskOpportunity,
		br.Objective,
		br.KPI,
		br.Process.Id,
		br.ERMEOA,
		br.InitialRiskSeverity,
		br.InitialRiskLikelyhood,
		br.ResidualRiskSeverity,
		br.ResidualRiskLikelyhood,
		Id,
	)

	if err != nil {
		return err
	}

	return nil
}
