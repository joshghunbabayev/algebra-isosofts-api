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

type MOCModel struct {
}

func (*MOCModel) GenerateUniqueId() string {
	Id := modules.GenerateRandomString(30)

	var mocModel MOCModel

	moc, _ := mocModel.GetById(Id)

	if moc.IsEmpty() {
		return Id
	} else {
		return mocModel.GenerateUniqueId()
	}
}

func (*MOCModel) GenerateUniqueNo() string {
	db := database.GetDatabase()

	year := time.Now().Format("06")

	var lastNo string
	db.QueryRow(`
		SELECT "no" 
		FROM mocregisters 
		WHERE "no" LIKE ? 
		ORDER BY "no" DESC 
		LIMIT 1
		`,
		"MOCR/"+year+"/%",
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

	newNo := fmt.Sprintf("MOCR/%s/%04d", year, nextNumber)
	return newNo
}

func (*MOCModel) GetById(Id string) (registerTypes.MOC, error) {
	db := database.GetDatabase()
	row := db.QueryRow(`
			SELECT * 
			FROM mocregisters
			WHERE id = ?
		`,
		Id,
	)

	var moc registerTypes.MOC
	var dropDownListItemModel tableComponentModels.DropDownListItemModel
	var actionModel registerComponentModels.ActionModel

	err := row.Scan(
		&moc.Id,
		&moc.No,
		&moc.Issuer,
		&moc.ReasonOfChange,
		&moc.Process.Id,
		&moc.Action,
		&moc.Risks,
		&moc.InitialRiskSeverity,
		&moc.InitialRiskLikelyhood,
		&moc.ResidualRiskSeverity,
		&moc.ResidualRiskLikelyhood,
		&moc.DbStatus,
		&moc.DbLastStatus,
	)
	moc.Process, _ = dropDownListItemModel.GetById(moc.Process.Id)
	moc.Actions, _ = actionModel.GetAll(map[string]interface{}{
		"registerId": moc.Id,
		"dbStatus":   "active",
	})

	return moc, err
}

func (*MOCModel) GetAll(filters map[string]interface{}) ([]registerTypes.MOC, error) {
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
			SELECT * FROM mocregisters %s
		`,
		whereClause,
	)
	rows, err := db.Query(query, values...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var mocs []registerTypes.MOC

	for rows.Next() {
		var moc registerTypes.MOC
		var dropDownListItemModel tableComponentModels.DropDownListItemModel
		var actionModel registerComponentModels.ActionModel

		rows.Scan(
			&moc.Id,
			&moc.No,
			&moc.Issuer,
			&moc.ReasonOfChange,
			&moc.Process.Id,
			&moc.Action,
			&moc.Risks,
			&moc.InitialRiskSeverity,
			&moc.InitialRiskLikelyhood,
			&moc.ResidualRiskSeverity,
			&moc.ResidualRiskLikelyhood,
			&moc.DbStatus,
			&moc.DbLastStatus,
		)
		moc.Process, _ = dropDownListItemModel.GetById(moc.Process.Id)
		moc.Actions, _ = actionModel.GetAll(map[string]interface{}{
			"registerId": moc.Id,
			"dbStatus":   "active",
		})

		mocs = append(mocs, moc)
	}

	return mocs, nil
}

func (*MOCModel) Create(moc registerTypes.MOC) error {
	db := database.GetDatabase()
	_, err := db.Exec(`
			INSERT INTO mocregisters ( 
				"id",
				"no",
				"issuer", 
				"reasonOfChange", 
				"process", 
				"action", 
				"risks", 
				"initialRiskSeverity", 
				"initialRiskLikelyhood", 
				"residualRiskSeverity", 
				"residualRiskLikelyhood",
				"dbStatus",
				"dbLastStatus"
			) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`,
		moc.Id,
		moc.No,
		moc.Issuer,
		moc.ReasonOfChange,
		moc.Process.Id,
		moc.Action,
		moc.Risks,
		moc.InitialRiskSeverity,
		moc.InitialRiskLikelyhood,
		moc.ResidualRiskSeverity,
		moc.ResidualRiskLikelyhood,
		moc.DbStatus,
		moc.DbLastStatus,
	)

	if err != nil {
		return err
	}

	return nil
}

func (*MOCModel) Update(Id string, fields map[string]interface{}) error {
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
			UPDATE mocregisters 
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
