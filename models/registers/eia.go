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

type EIAModel struct {
}

func (*EIAModel) GenerateUniqueId() string {
	Id := modules.GenerateRandomString(30)

	var eiaModel EIAModel

	eia, _ := eiaModel.GetById(Id)

	if eia.IsEmpty() {
		return Id
	} else {
		return eiaModel.GenerateUniqueId()
	}
}

func (*EIAModel) GenerateUniqueNo() string {
	db := database.GetDatabase()

	year := time.Now().Format("06")

	var lastNo string
	db.QueryRow(`
		SELECT "no" 
		FROM eiaregisters 
		WHERE "no" LIKE ? 
		ORDER BY "no" DESC 
		LIMIT 1
		`,
		"EIAR/"+year+"/%",
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

	newNo := fmt.Sprintf("EIAR/%s/%04d", year, nextNumber)
	return newNo
}

func (*EIAModel) GetById(Id string) (registerTypes.EIA, error) {
	db := database.GetDatabase()
	row := db.QueryRow(`
			SELECT * 
			FROM eiaregisters
			WHERE id = ?
		`,
		Id,
	)

	var eia registerTypes.EIA
	var dropDownListItemModel tableComponentModels.DropDownListItemModel
	var actionModel registerComponentModels.ActionModel

	err := row.Scan(
		&eia.Id,
		&eia.No,
		&eia.Process.Id,
		&eia.Aspect.Id,
		&eia.Impact,
		&eia.ExistingControls,
		&eia.IDOSProbability,
		&eia.IDOSSeverity,
		&eia.IDOSDuration,
		&eia.IDOSScale,
		&eia.RDOSProbability,
		&eia.RDOSSeverity,
		&eia.RDOSDuration,
		&eia.RDOSScale,
		&eia.DbStatus,
		&eia.DbLastStatus,
	)
	eia.Process, _ = dropDownListItemModel.GetById(eia.Process.Id)
	eia.Aspect, _ = dropDownListItemModel.GetById(eia.Aspect.Id)
	eia.Actions, _ = actionModel.GetAll(map[string]interface{}{
		"registerId": eia.Id,
	})

	return eia, err
}

func (*EIAModel) GetAll(filters map[string]interface{}) ([]registerTypes.EIA, error) {
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
			SELECT * FROM eiaregisters %s
		`,
		whereClause,
	)
	rows, err := db.Query(query, values...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var eias []registerTypes.EIA

	for rows.Next() {
		var eia registerTypes.EIA
		var dropDownListItemModel tableComponentModels.DropDownListItemModel
		var actionModel registerComponentModels.ActionModel

		rows.Scan(
			&eia.Id,
			&eia.No,
			&eia.Process.Id,
			&eia.Aspect.Id,
			&eia.Impact,
			&eia.ExistingControls,
			&eia.IDOSProbability,
			&eia.IDOSSeverity,
			&eia.IDOSDuration,
			&eia.IDOSScale,
			&eia.RDOSProbability,
			&eia.RDOSSeverity,
			&eia.RDOSDuration,
			&eia.RDOSScale,
			&eia.DbStatus,
			&eia.DbLastStatus,
		)
		eia.Process, _ = dropDownListItemModel.GetById(eia.Process.Id)
		eia.Aspect, _ = dropDownListItemModel.GetById(eia.Aspect.Id)
		eia.Actions, _ = actionModel.GetAll(map[string]interface{}{
			"registerId": eia.Id,
		})

		eias = append(eias, eia)
	}

	return eias, nil
}

func (*EIAModel) Create(eia registerTypes.EIA) error {
	db := database.GetDatabase()
	_, err := db.Exec(`
			INSERT INTO eiaregisters ( 
				"id",
				"no",
				"process",
				"aspect",
				"impact",
				"existingControls",
				"idosProbability",
				"idosSeverity",
				"idosDuration",
				"idosScale",
				"rdosProbability",
				"rdosSeverity",
				"rdosDuration",
				"rdosScale",
				"dbStatus",
				"dbLastStatus"
			) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`,
		eia.Id,
		eia.No,
		eia.Process.Id,
		eia.Aspect.Id,
		eia.Impact,
		eia.ExistingControls,
		eia.IDOSProbability,
		eia.IDOSSeverity,
		eia.IDOSDuration,
		eia.IDOSScale,
		eia.RDOSProbability,
		eia.RDOSSeverity,
		eia.RDOSDuration,
		eia.RDOSScale,
		eia.DbStatus,
		eia.DbLastStatus,
	)

	if err != nil {
		return err
	}

	return nil
}

func (*EIAModel) Update(Id string, fields map[string]interface{}) error {
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
			UPDATE eiaregisters 
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
