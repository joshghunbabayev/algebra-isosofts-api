package registerModels

import (
	"algebra-isosofts-api/database"
	"algebra-isosofts-api/modules"
	registerTypes "algebra-isosofts-api/types/registers"
)

type BrModel struct {
}

func (*BrModel) GenerateIdForBr() string {
	Id := modules.GenerateRandomString(30)

	var brModel BrModel

	br, _ := brModel.GetById(Id)

	if br.IsEmpty() {
		return Id
	} else {
		return brModel.GenerateIdForBr()
	}
}

func (*BrModel) GetById(Id string) (registerTypes.Br, error) {
	db := database.GetDatabase()
	row := db.QueryRow("SELECT * FROM brregisters WHERE id = ?", Id)

	var br registerTypes.Br
	err := row.Scan(&br.Id, &br.Swot.Id, &br.Pestle.Id)

	return br, err
}

func (*BrModel) GetAll() ([]registerTypes.Br, error) {
	db := database.GetDatabase()
	rows, err := db.Query("SELECT * FROM brregisters")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var brs []registerTypes.Br

	for rows.Next() {
		var br registerTypes.Br
		rows.Scan(&br.Id, &br.No, &br.Swot.Id, &br.Pestle.Id, &br.InterestedParty.Id, &br.RiskOpportunity, &br.Objective, &br.KPI, &br.Process.Id, &br.ERMEOA, &br.InitialRiskSeverity, &br.InitialRiskLikelyhood, &br.ResidualRiskSeverity, &br.ResidualRiskLikelyhood)
		brs = append(brs, br)
	}

	return brs, nil
}

func (*BrModel) Create(br registerTypes.Br) error {
	db := database.GetDatabase()

	_, err := db.Exec(`INSERT INTO brregisters ("id", "swot", "pestle") VALUES (?, ?, ?)`,
		br.Id, br.Swot, br.Pestle)

	if err != nil {
		return err
	}

	return nil
}
