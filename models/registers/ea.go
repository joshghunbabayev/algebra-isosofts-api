package registerModels

import (
	"algebra-isosofts-api/database"
	registerComponentModels "algebra-isosofts-api/models/registers/components"
	"algebra-isosofts-api/modules"
	registerTypes "algebra-isosofts-api/types/registers"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type EAModel struct {
}

func (*EAModel) GenerateUniqueId() string {
	Id := modules.GenerateRandomString(30)

	var eaModel EAModel

	ea, _ := eaModel.GetById(Id)

	if ea.IsEmpty() {
		return Id
	} else {
		return eaModel.GenerateUniqueId()
	}
}

func (*EAModel) GenerateUniqueNo() string {
	db := database.GetDatabase()

	year := time.Now().Format("06")

	var lastNo string
	db.QueryRow(`
		SELECT "no" 
		FROM earegisters 
		WHERE "no" LIKE ? 
		ORDER BY "no" DESC 
		LIMIT 1
		`,
		"EAR/"+year+"/%",
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

	newNo := fmt.Sprintf("EAR/%s/%04d", year, nextNumber)
	return newNo
}

func (*EAModel) GetById(Id string) (registerTypes.EA, error) {
	db := database.GetDatabase()
	row := db.QueryRow(`
			SELECT * 
			FROM earegisters
			WHERE id = ?
		`,
		Id,
	)

	var ea registerTypes.EA
	var actionModel registerComponentModels.ActionModel

	err := row.Scan(
		&ea.Id,
		&ea.No,
		&ea.EmployeeName,
		&ea.Position,
		&ea.LineManager,
		&ea.ESD,
		&ea.AppraisalDate,
		&ea.AppraisalType,
		&ea.TCA,
		&ea.SkillsAppraisal,
		&ea.DbStatus,
		&ea.DbLastStatus,
	)
	ea.Actions, _ = actionModel.GetAll(map[string]interface{}{
		"registerId": ea.Id,
		"dbStatus":   "active",
	})

	return ea, err
}

func (*EAModel) GetAll(filters map[string]interface{}) ([]registerTypes.EA, error) {
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
			SELECT * FROM earegisters %s
		`,
		whereClause,
	)
	rows, err := db.Query(query, values...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var eas []registerTypes.EA

	for rows.Next() {
		var ea registerTypes.EA
		var actionModel registerComponentModels.ActionModel

		rows.Scan(
			&ea.Id,
			&ea.No,
			&ea.EmployeeName,
			&ea.Position,
			&ea.LineManager,
			&ea.ESD,
			&ea.AppraisalDate,
			&ea.AppraisalType,
			&ea.TCA,
			&ea.SkillsAppraisal,
			&ea.DbStatus,
			&ea.DbLastStatus,
		)
		ea.Actions, _ = actionModel.GetAll(map[string]interface{}{
			"registerId": ea.Id,
			"dbStatus":   "active",
		})

		eas = append(eas, ea)
	}

	return eas, nil
}

func (*EAModel) Create(ea registerTypes.EA) error {
	db := database.GetDatabase()
	_, err := db.Exec(`
			INSERT INTO earegisters ( 
				"id",
				"no",
				"employeeName", 
				"position", 
				"lineManager", 
				"esd", 
				"appraisalDate", 
				"appraisalType", 
				"tca", 
				"skillsAppraisal", 
				"dbStatus",
				"dbLastStatus"
			) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`,
		ea.Id,
		ea.No,
		ea.EmployeeName,
		ea.Position,
		ea.LineManager,
		ea.ESD,
		ea.AppraisalDate,
		ea.AppraisalType,
		ea.TCA,
		ea.SkillsAppraisal,
		ea.DbStatus,
		ea.DbLastStatus,
	)

	if err != nil {
		return err
	}

	return nil
}

func (*EAModel) Update(Id string, fields map[string]interface{}) error {
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
			UPDATE earegisters 
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
