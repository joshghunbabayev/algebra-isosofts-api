package registerModels

import (
	"algebra-isosofts-api/database"
	registerComponentModels "algebra-isosofts-api/models/registerComponents"
	tableComponentModels "algebra-isosofts-api/models/tableComponents"
	"algebra-isosofts-api/modules"
	registerTypes "algebra-isosofts-api/types/registers"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type EAIModel struct {
}

func (*EAIModel) GenerateUniqueId() string {
	Id := modules.GenerateRandomString(30)

	var eaiModel EAIModel

	eai, _ := eaiModel.GetById(Id)

	if eai.IsEmpty() {
		return Id
	} else {
		return eaiModel.GenerateUniqueId()
	}
}

func (*EAIModel) GenerateUniqueNo() string {
	db := database.GetDatabase()

	year := time.Now().Format("06")

	var lastNo string
	db.QueryRow(`
		SELECT "no" 
		FROM eairegisters 
		WHERE "no" LIKE ? 
		ORDER BY "no" DESC 
		LIMIT 1
		`,
		"EAIR/"+year+"/%",
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

	newNo := fmt.Sprintf("EAIR/%s/%04d", year, nextNumber)
	return newNo
}

func (*EAIModel) GetById(Id string) (registerTypes.EAI, error) {
	db := database.GetDatabase()
	row := db.QueryRow(`
			SELECT * 
			FROM eairegisters
			WHERE id = ?
		`,
		Id,
	)

	var eai registerTypes.EAI
	var dropDownListItemModel tableComponentModels.DropDownListItemModel
	var actionModel registerComponentModels.ActionModel

	err := row.Scan(
		&eai.Id,
		&eai.No,
		&eai.Process.Id,
		&eai.Aspect.Id,
		&eai.Impact,
		&eai.AffectedReceptors.Id,
		&eai.ExistingControls,
		&eai.IDOSProbability,
		&eai.IDOSSeverity,
		&eai.IDOSDuration,
		&eai.IDOSScale,
		&eai.RDOSProbability,
		&eai.RDOSSeverity,
		&eai.RDOSDuration,
		&eai.RDOSScale,
		&eai.DbStatus,
		&eai.DbLastStatus,
	)
	eai.Process, _ = dropDownListItemModel.GetById(eai.Process.Id)
	eai.Aspect, _ = dropDownListItemModel.GetById(eai.Aspect.Id)
	eai.AffectedReceptors, _ = dropDownListItemModel.GetById(eai.AffectedReceptors.Id)
	eai.Actions, _ = actionModel.GetAll(map[string]interface{}{
		"registerId": eai.Id,
	})

	return eai, err
}

func (*EAIModel) GetAll(filters map[string]interface{}) ([]registerTypes.EAI, error) {
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
			SELECT * FROM eairegisters %s
		`,
		whereClause,
	)
	rows, err := db.Query(query, values...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var eais []registerTypes.EAI

	for rows.Next() {
		var eai registerTypes.EAI
		var dropDownListItemModel tableComponentModels.DropDownListItemModel
		var actionModel registerComponentModels.ActionModel

		rows.Scan(
			&eai.Id,
			&eai.No,
			&eai.Process.Id,
			&eai.Aspect.Id,
			&eai.Impact,
			&eai.AffectedReceptors.Id,
			&eai.ExistingControls,
			&eai.IDOSProbability,
			&eai.IDOSSeverity,
			&eai.IDOSDuration,
			&eai.IDOSScale,
			&eai.RDOSProbability,
			&eai.RDOSSeverity,
			&eai.RDOSDuration,
			&eai.RDOSScale,
			&eai.DbStatus,
			&eai.DbLastStatus,
		)
		eai.Process, _ = dropDownListItemModel.GetById(eai.Process.Id)
		eai.Aspect, _ = dropDownListItemModel.GetById(eai.Aspect.Id)
		eai.AffectedReceptors, _ = dropDownListItemModel.GetById(eai.AffectedReceptors.Id)
		eai.Actions, _ = actionModel.GetAll(map[string]interface{}{
			"registerId": eai.Id,
		})

		eais = append(eais, eai)
	}

	return eais, nil
}

func (*EAIModel) Create(eai registerTypes.EAI) error {
	db := database.GetDatabase()
	_, err := db.Exec(`
			INSERT INTO eairegisters ( 
				"id",
				"no",
				"process",
				"aspect",
				"impact",
				"affectedReceptors", 
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
			) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`,
		eai.Id,
		eai.No,
		eai.Process.Id,
		eai.Aspect.Id,
		eai.Impact,
		eai.AffectedReceptors.Id,
		eai.ExistingControls,
		eai.IDOSProbability,
		eai.IDOSSeverity,
		eai.IDOSDuration,
		eai.IDOSScale,
		eai.RDOSProbability,
		eai.RDOSSeverity,
		eai.RDOSDuration,
		eai.RDOSScale,
		eai.DbStatus,
		eai.DbLastStatus,
	)

	if err != nil {
		return err
	}

	return nil
}

func (*EAIModel) Update(Id string, fields map[string]interface{}) error {
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
			UPDATE eairegisters 
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
