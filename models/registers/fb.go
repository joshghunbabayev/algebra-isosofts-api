package registerModels

import (
	"algebra-isosofts-api/database"
	tableComponentModels "algebra-isosofts-api/models/tableComponents"
	"algebra-isosofts-api/modules"
	registerTypes "algebra-isosofts-api/types/registers"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type FBModel struct {
}

func (*FBModel) GenerateUniqueId() string {
	Id := modules.GenerateRandomString(30)

	var fbModel FBModel

	fb, _ := fbModel.GetById(Id)

	if fb.IsEmpty() {
		return Id
	} else {
		return fbModel.GenerateUniqueId()
	}
}

func (*FBModel) GenerateUniqueNo() string {
	db := database.GetDatabase()

	year := time.Now().Format("06")

	var lastNo string
	db.QueryRow(`
		SELECT "no" 
		FROM fbregisters 
		WHERE "no" LIKE ? 
		ORDER BY "no" DESC 
		LIMIT 1
		`,
		"FBR/"+year+"/%",
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

	newNo := fmt.Sprintf("FBR/%s/%04d", year, nextNumber)
	return newNo
}

func (*FBModel) GetById(Id string) (registerTypes.FB, error) {
	db := database.GetDatabase()
	row := db.QueryRow(`
			SELECT * 
			FROM fbregisters
			WHERE id = ?
		`,
		Id,
	)

	var fb registerTypes.FB
	var dropDownListItemModel tableComponentModels.DropDownListItemModel

	err := row.Scan(
		&fb.Id,
		&fb.No,
		&fb.JobNumber,
		&fb.JobStartDate,
		&fb.JobCompletionDate,
		&fb.Scope,
		&fb.CustomerId,
		&fb.TypeOfFinding,
		&fb.QGS,
		&fb.Communication,
		&fb.OTD,
		&fb.Documentation,
		&fb.HS,
		&fb.Environment,
		&fb.DbStatus,
		&fb.DbLastStatus,
	)
	fb.Scope, _ = dropDownListItemModel.GetById(fb.Scope.Id)
	fb.TypeOfFinding, _ = dropDownListItemModel.GetById(fb.TypeOfFinding.Id)

	return fb, err
}

func (*FBModel) GetAll(filters map[string]interface{}) ([]registerTypes.FB, error) {
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
			SELECT * FROM fbregisters %s
		`,
		whereClause,
	)
	rows, err := db.Query(query, values...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var fbs []registerTypes.FB

	for rows.Next() {
		var fb registerTypes.FB
		var dropDownListItemModel tableComponentModels.DropDownListItemModel

		rows.Scan(
			&fb.Id,
			&fb.No,
			&fb.JobNumber,
			&fb.JobStartDate,
			&fb.JobCompletionDate,
			&fb.Scope,
			&fb.CustomerId,
			&fb.TypeOfFinding,
			&fb.QGS,
			&fb.Communication,
			&fb.OTD,
			&fb.Documentation,
			&fb.HS,
			&fb.Environment,
			&fb.DbStatus,
			&fb.DbLastStatus,
		)
		fb.Scope, _ = dropDownListItemModel.GetById(fb.Scope.Id)
		fb.TypeOfFinding, _ = dropDownListItemModel.GetById(fb.TypeOfFinding.Id)

		fbs = append(fbs, fb)
	}

	return fbs, nil
}

func (*FBModel) Create(fb registerTypes.FB) error {
	db := database.GetDatabase()
	_, err := db.Exec(`
			INSERT INTO fbregisters ( 
				"id",
				"no",
				"jobNumber",
				"jobStartDate",
				"jobCompletionDate",
				"scope",
				"customerId",
				"typeOfFinding",
				"qgs",
				"communication",
				"otd",
				"documentation",
				"hs",
				"environment",
				"dbStatus",
				"dbLastStatus"
			) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`,
		fb.Id,
		fb.No,
		fb.JobNumber,
		fb.JobStartDate,
		fb.JobCompletionDate,
		fb.Scope.Id,
		fb.CustomerId,
		fb.TypeOfFinding.Id,
		fb.QGS,
		fb.Communication,
		fb.OTD,
		fb.Documentation,
		fb.HS,
		fb.Environment,
		fb.DbStatus,
		fb.DbLastStatus,
	)

	if err != nil {
		return err
	}

	return nil
}

func (*FBModel) Update(Id string, fields map[string]interface{}) error {
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
			UPDATE fbregisters 
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
