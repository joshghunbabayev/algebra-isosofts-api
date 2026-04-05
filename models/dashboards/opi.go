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

type OPIModel struct{}

func (*OPIModel) GenerateUniqueId() string {
	id := modules.GenerateRandomString(30)
	var opiModel OPIModel
	opi, _ := opiModel.GetById(id)

	if opi.IsEmpty() {
		return id
	}
	return opiModel.GenerateUniqueId()
}

func (*OPIModel) GenerateUniqueNo() string {
	db := database.GetDatabase()

	year := time.Now().Format("06")

	var lastNo string
	db.QueryRow(`
		SELECT "no" 
		FROM opis 
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

func (*OPIModel) GetById(id string) (dashboardTypes.OPI, error) {
	db := database.GetDatabase()
	row := db.QueryRow(`
			SELECT * FROM opis
			WHERE id = ?
		`, id)

	var opi dashboardTypes.OPI
	var dropDownListItemModel tableComponentModels.DropDownListItemModel

	err := row.Scan(
		&opi.Id,
		&opi.CompanyId,
		&opi.No,
		&opi.Title,
		&opi.Function.Id,
		&opi.LYKPI,
		&opi.AnnualTarget,
		&opi.January,
		&opi.February,
		&opi.March,
		&opi.April,
		&opi.May,
		&opi.June,
		&opi.July,
		&opi.August,
		&opi.September,
		&opi.October,
		&opi.November,
		&opi.December,
		&opi.DbStatus,
		&opi.DbLastStatus,
	)
	opi.Function, _ = dropDownListItemModel.GetById(opi.Function.Id)

	return opi, err
}

func (*OPIModel) GetAll(filters map[string]interface{}) ([]dashboardTypes.OPI, error) {
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

	query := fmt.Sprintf(`SELECT * FROM opis %s`, whereClause)
	rows, err := db.Query(query, values...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var opis []dashboardTypes.OPI
	for rows.Next() {
		var opi dashboardTypes.OPI
		var dropDownListItemModel tableComponentModels.DropDownListItemModel

		err := rows.Scan(
			&opi.Id,
			&opi.CompanyId,
			&opi.No,
			&opi.Title,
			&opi.Function.Id,
			&opi.LYKPI,
			&opi.AnnualTarget,
			&opi.January,
			&opi.February,
			&opi.March,
			&opi.April,
			&opi.May,
			&opi.June,
			&opi.July,
			&opi.August,
			&opi.September,
			&opi.October,
			&opi.November,
			&opi.December,
			&opi.DbStatus,
			&opi.DbLastStatus,
		)

		opi.Function, _ = dropDownListItemModel.GetById(opi.Function.Id)
		if err != nil {
			return nil, err
		}
		opis = append(opis, opi)
	}

	return opis, nil
}

func (*OPIModel) Create(opi dashboardTypes.OPI) error {
	db := database.GetDatabase()
	_, err := db.Exec(`
			INSERT INTO opis (
				"id", "companyId", "no", "title", "function", 
				"lykpi", "annualTarget",
				"january", "february", "march", "april", "may", "june",
				"july", "august", "september", "october", "november", "december", 
				"dbStatus", "dbLastStatus" 
			) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`,
		opi.Id, opi.CompanyId, opi.No, opi.Title, opi.Function.Id,
		opi.LYKPI, opi.AnnualTarget,
		opi.January, opi.February, opi.March, opi.April, opi.May, opi.June,
		opi.July, opi.August, opi.September, opi.October, opi.November, opi.December,
		opi.DbStatus, opi.DbLastStatus,
	)
	return err
}

func (*OPIModel) Update(id string, fields map[string]interface{}) error {
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
	query := fmt.Sprintf(`UPDATE opis SET %s WHERE "id" = ?`, setClause)
	values = append(values, id)

	db := database.GetDatabase()
	_, err := db.Exec(query, values...)
	return err
}
