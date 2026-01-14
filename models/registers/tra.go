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

type TRAModel struct {
}

func (*TRAModel) GenerateUniqueId() string {
	Id := modules.GenerateRandomString(30)

	var traModel TRAModel

	tra, _ := traModel.GetById(Id)

	if tra.IsEmpty() {
		return Id
	} else {
		return traModel.GenerateUniqueId()
	}
}

func (*TRAModel) GenerateUniqueNo() string {
	db := database.GetDatabase()

	year := time.Now().Format("06")

	var lastNo string
	db.QueryRow(`
		SELECT "no" 
		FROM traregisters 
		WHERE "no" LIKE ? 
		ORDER BY "no" DESC 
		LIMIT 1
		`,
		"TRAR/"+year+"/%",
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

	newNo := fmt.Sprintf("TRAR/%s/%04d", year, nextNumber)
	return newNo
}

func (*TRAModel) GetById(Id string) (registerTypes.TRA, error) {
	db := database.GetDatabase()
	row := db.QueryRow(`
			SELECT * 
			FROM traregisters
			WHERE id = ?
		`,
		Id,
	)

	var tra registerTypes.TRA
	var actionModel registerComponentModels.ActionModel

	err := row.Scan(
		&tra.Id,
		&tra.No,
		&tra.EmployeeName,
		&tra.Position,
		&tra.CLName,
		&tra.TCLID,
		&tra.CLNumber,
		&tra.NCD,
		&tra.CompetencyStatus,
		&tra.Effectiveness,
		&tra.DbStatus,
		&tra.DbLastStatus,
	)
	tra.Actions, _ = actionModel.GetAll(map[string]interface{}{
		"registerId": tra.Id,
		"dbStatus":   "active",
	})

	return tra, err
}

func (*TRAModel) GetAll(filters map[string]interface{}) ([]registerTypes.TRA, error) {
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
			SELECT * FROM traregisters %s
		`,
		whereClause,
	)
	rows, err := db.Query(query, values...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tras []registerTypes.TRA

	for rows.Next() {
		var tra registerTypes.TRA
		var actionModel registerComponentModels.ActionModel

		rows.Scan(
			&tra.Id,
			&tra.No,
			&tra.EmployeeName,
			&tra.Position,
			&tra.CLName,
			&tra.TCLID,
			&tra.CLNumber,
			&tra.NCD,
			&tra.CompetencyStatus,
			&tra.Effectiveness,
			&tra.DbStatus,
			&tra.DbLastStatus,
		)
		tra.Actions, _ = actionModel.GetAll(map[string]interface{}{
			"registerId": tra.Id,
			"dbStatus":   "active",
		})

		tras = append(tras, tra)
	}

	return tras, nil
}

func (*TRAModel) Create(tra registerTypes.TRA) error {
	db := database.GetDatabase()
	_, err := db.Exec(`
			INSERT INTO traregisters ( 
				"id",
				"no",
				"employeeName",
				"position",
				"clname",
				"nvcd",
				"clnumber",
				"ncd",
				"competencyStatus",
				"effectiveness",
				"dbStatus",
				"dbLastStatus"
			) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`,
		tra.Id,
		tra.No,
		tra.EmployeeName,
		tra.Position,
		tra.CLName,
		tra.TCLID,
		tra.CLNumber,
		tra.NCD,
		tra.CompetencyStatus,
		tra.Effectiveness,
		tra.DbStatus,
		tra.DbLastStatus,
	)

	if err != nil {
		return err
	}

	return nil
}

func (*TRAModel) Update(Id string, fields map[string]interface{}) error {
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
			UPDATE traregisters 
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
