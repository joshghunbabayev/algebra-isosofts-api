package dashboardModels

import (
	"algebra-isosofts-api/database"
	tableComponentModels "algebra-isosofts-api/models/tableComponents"
	"algebra-isosofts-api/modules"
	dashboardTypes "algebra-isosofts-api/types/dashboards"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type OpKPIModel struct{}

func (*OpKPIModel) GenerateUniqueId() string {
	id := modules.GenerateRandomString(30)
	var opKPIModel OpKPIModel
	opKPI, _ := opKPIModel.GetById(id)

	if opKPI.IsEmpty() {
		return id
	}
	return opKPIModel.GenerateUniqueId()
}

func (*OpKPIModel) GenerateUniqueNo() string {
	db := database.GetDatabase()

	year := time.Now().Format("06")

	var lastNo string
	db.QueryRow(`
		SELECT "no" 
		FROM opkpis 
		WHERE "no" LIKE ? 
		ORDER BY "no" DESC 
		LIMIT 1
		`,
		"OKİ/"+year+"/%",
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

	newNo := fmt.Sprintf("OKİ/%s/%04d", year, nextNumber)
	return newNo
}

func (*OpKPIModel) GetById(id string) (dashboardTypes.OpKPI, error) {
	db := database.GetDatabase()
	row := db.QueryRow(`
			SELECT * FROM opkpis
			WHERE id = ?
		`, id)

	var opKPI dashboardTypes.OpKPI
	var dropDownListItemModel tableComponentModels.DropDownListItemModel

	err := row.Scan(
		&opKPI.Id,
		&opKPI.CompanyId,
		&opKPI.No,
		&opKPI.Title,
		&opKPI.Function.Id,
		&opKPI.LYKPI,
		&opKPI.AnnualTarget,
		&opKPI.January,
		&opKPI.February,
		&opKPI.March,
		&opKPI.April,
		&opKPI.May,
		&opKPI.June,
		&opKPI.July,
		&opKPI.August,
		&opKPI.September,
		&opKPI.October,
		&opKPI.November,
		&opKPI.December,
		&opKPI.DbStatus,
		&opKPI.DbLastStatus,
	)
	opKPI.Function, _ = dropDownListItemModel.GetById(opKPI.Function.Id)

	return opKPI, err
}

func (*OpKPIModel) GetAll(filters map[string]interface{}) ([]dashboardTypes.OpKPI, error) {
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

	query := fmt.Sprintf(`SELECT * FROM opkpis %s`, whereClause)
	rows, err := db.Query(query, values...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var opKPIs []dashboardTypes.OpKPI
	for rows.Next() {
		var opKPI dashboardTypes.OpKPI
		var dropDownListItemModel tableComponentModels.DropDownListItemModel

		err := rows.Scan(
			&opKPI.Id,
			&opKPI.CompanyId,
			&opKPI.No,
			&opKPI.Title,
			&opKPI.Function.Id,
			&opKPI.LYKPI,
			&opKPI.AnnualTarget,
			&opKPI.January,
			&opKPI.February,
			&opKPI.March,
			&opKPI.April,
			&opKPI.May,
			&opKPI.June,
			&opKPI.July,
			&opKPI.August,
			&opKPI.September,
			&opKPI.October,
			&opKPI.November,
			&opKPI.December,
			&opKPI.DbStatus,
			&opKPI.DbLastStatus,
		)

		opKPI.Function, _ = dropDownListItemModel.GetById(opKPI.Function.Id)
		if err != nil {
			return nil, err
		}
		opKPIs = append(opKPIs, opKPI)
	}

	return opKPIs, nil
}

func (*OpKPIModel) Create(opKPI dashboardTypes.OpKPI) error {
	db := database.GetDatabase()
	_, err := db.Exec(`
			INSERT INTO opkpis (
				"id", "companyId", "no", "title", "function", 
				"lykpi", "annualTarget",
				"january", "february", "march", "april", "may", "june",
				"july", "august", "september", "october", "november", "december", 
				"dbStatus", "dbLastStatus" 
			) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`,
		opKPI.Id, opKPI.CompanyId, opKPI.No, opKPI.Title, opKPI.Function.Id,
		opKPI.LYKPI, opKPI.AnnualTarget,
		opKPI.January, opKPI.February, opKPI.March, opKPI.April, opKPI.May, opKPI.June,
		opKPI.July, opKPI.August, opKPI.September, opKPI.October, opKPI.November, opKPI.December,
		opKPI.DbStatus, opKPI.DbLastStatus,
	)
	return err
}

func (*OpKPIModel) Update(id string, fields map[string]interface{}) error {
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
	query := fmt.Sprintf(`UPDATE opkpis SET %s WHERE "id" = ?`, setClause)
	values = append(values, id)

	db := database.GetDatabase()
	_, err := db.Exec(query, values...)
	return err
}
