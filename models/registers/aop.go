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

type AOPModel struct {
}

func (*AOPModel) GenerateUniqueId() string {
	Id := modules.GenerateRandomString(30)

	var aopModel AOPModel

	aop, _ := aopModel.GetById(Id)

	if aop.IsEmpty() {
		return Id
	} else {
		return aopModel.GenerateUniqueId()
	}
}

func (*AOPModel) GenerateUniqueNo() string {
	db := database.GetDatabase()

	year := time.Now().Format("06")

	var lastNo string
	db.QueryRow(`
		SELECT "no" 
		FROM aopregisters 
		WHERE "no" LIKE ? 
		ORDER BY "no" DESC 
		LIMIT 1
		`,
		"AOPR/"+year+"/%",
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

	newNo := fmt.Sprintf("AOPR/%s/%04d", year, nextNumber)
	return newNo
}

func (*AOPModel) GetById(Id string) (registerTypes.AOP, error) {
	db := database.GetDatabase()
	row := db.QueryRow(`
			SELECT * 
			FROM aopregisters
			WHERE id = ?
		`,
		Id,
	)

	var aop registerTypes.AOP
	var dropDownListItemModel tableComponentModels.DropDownListItemModel
	var actionModel registerComponentModels.ActionModel

	err := row.Scan(
		&aop.Id,
		&aop.No,
		&aop.ActivityDescription.Id,
		&aop.AuditorInspector,
		&aop.AuditeeInspectee,
		&aop.ReviewedPremises.Id,
		&aop.ReviewedProcess.Id,
		&aop.RTIC,
		&aop.Frequency,
		&aop.AuditDate,
		&aop.InspectionFrequency.Id,
		&aop.NextAuditDate,
		&aop.AuditStatus,
		&aop.DbStatus,
		&aop.DbLastStatus,
	)
	aop.ActivityDescription, _ = dropDownListItemModel.GetById(aop.ActivityDescription.Id)
	aop.ReviewedPremises, _ = dropDownListItemModel.GetById(aop.ReviewedPremises.Id)
	aop.ReviewedProcess, _ = dropDownListItemModel.GetById(aop.ReviewedProcess.Id)
	aop.InspectionFrequency, _ = dropDownListItemModel.GetById(aop.InspectionFrequency.Id)
	aop.Actions, _ = actionModel.GetAll(map[string]interface{}{
		"registerId": aop.Id,
	})

	return aop, err
}

func (*AOPModel) GetAll(filters map[string]interface{}) ([]registerTypes.AOP, error) {
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
			SELECT * FROM aopregisters %s
		`,
		whereClause,
	)
	rows, err := db.Query(query, values...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var aops []registerTypes.AOP

	for rows.Next() {
		var aop registerTypes.AOP
		var dropDownListItemModel tableComponentModels.DropDownListItemModel
		var actionModel registerComponentModels.ActionModel

		rows.Scan(
			&aop.Id,
			&aop.No,
			&aop.ActivityDescription.Id,
			&aop.AuditorInspector,
			&aop.AuditeeInspectee,
			&aop.ReviewedPremises.Id,
			&aop.ReviewedProcess.Id,
			&aop.RTIC,
			&aop.Frequency,
			&aop.AuditDate,
			&aop.InspectionFrequency.Id,
			&aop.NextAuditDate,
			&aop.AuditStatus,
			&aop.DbStatus,
			&aop.DbLastStatus,
		)
		aop.ActivityDescription, _ = dropDownListItemModel.GetById(aop.ActivityDescription.Id)
		aop.ReviewedPremises, _ = dropDownListItemModel.GetById(aop.ReviewedPremises.Id)
		aop.ReviewedProcess, _ = dropDownListItemModel.GetById(aop.ReviewedProcess.Id)
		aop.InspectionFrequency, _ = dropDownListItemModel.GetById(aop.InspectionFrequency.Id)
		aop.Actions, _ = actionModel.GetAll(map[string]interface{}{
			"registerId": aop.Id,
		})

		aops = append(aops, aop)
	}

	return aops, nil
}

func (*AOPModel) Create(aop registerTypes.AOP) error {
	db := database.GetDatabase()
	_, err := db.Exec(`
			INSERT INTO aopregisters ( 
				"id",
				"no",
				"activityDescription", 
				"auditorInspector", 
				"auditeeInspectee", 
				"reviewedPremises", 
				"reviewedProcess", 
				"rtic", 
				"frequency", 
				"auditDate", 
				"inspectionFrequency", 
				"nextAuditDate", 
				"auditStatus", 
				"dbStatus",
				"dbLastStatus"
			) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`,
		aop.Id,
		aop.No,
		aop.ActivityDescription.Id,
		aop.AuditorInspector,
		aop.AuditeeInspectee,
		aop.ReviewedPremises.Id,
		aop.ReviewedProcess.Id,
		aop.RTIC,
		aop.Frequency,
		aop.AuditDate,
		aop.InspectionFrequency.Id,
		aop.NextAuditDate,
		aop.AuditStatus,
		aop.DbStatus,
		aop.DbLastStatus,
	)

	if err != nil {
		return err
	}

	return nil
}

func (*AOPModel) Update(Id string, fields map[string]interface{}) error {
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
			UPDATE aopregisters 
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
