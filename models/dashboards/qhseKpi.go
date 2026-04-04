package dashboardModels

import (
	"algebra-isosofts-api/database"
	tableComponentModels "algebra-isosofts-api/models/tableComponents"
	"algebra-isosofts-api/modules"
	dashboardTypes "algebra-isosofts-api/types/dashboards"
	tableComponentTypes "algebra-isosofts-api/types/tableComponents"
	"fmt"
	"strings"
)

type QhseKPIModel struct{}

func (*QhseKPIModel) GenerateUniqueId() string {
	id := modules.GenerateRandomString(30)
	var qhseKPIModel QhseKPIModel
	qhseKPI, _ := qhseKPIModel.GetById(id)

	if qhseKPI.IsEmpty() {
		return id
	}
	return qhseKPIModel.GenerateUniqueId()
}

func (*QhseKPIModel) GetById(id string) (dashboardTypes.QhseKPI, error) {
	db := database.GetDatabase()
	row := db.QueryRow(`
			SELECT * FROM qhsekpis
			WHERE id = ?
		`, id)

	var qhseKPI dashboardTypes.QhseKPI
	var dropDownListItemModel tableComponentModels.DropDownListItemModel

	err := row.Scan(
		&qhseKPI.Id,
		&qhseKPI.CompanyId,
		&qhseKPI.SNo,
		&qhseKPI.No,
		&qhseKPI.Title,
		&qhseKPI.Function.Id,
		&qhseKPI.LYKPI,
		&qhseKPI.AnnualTarget,
		&qhseKPI.January,
		&qhseKPI.February,
		&qhseKPI.March,
		&qhseKPI.April,
		&qhseKPI.May,
		&qhseKPI.June,
		&qhseKPI.July,
		&qhseKPI.August,
		&qhseKPI.September,
		&qhseKPI.October,
		&qhseKPI.November,
		&qhseKPI.December,
	)
	qhseKPI.Function, _ = dropDownListItemModel.GetById(qhseKPI.Function.Id)

	return qhseKPI, err
}

func (*QhseKPIModel) GetAll(filters map[string]interface{}) ([]dashboardTypes.QhseKPI, error) {
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

	query := fmt.Sprintf(`SELECT * FROM qhsekpis %s`, whereClause)
	rows, err := db.Query(query, values...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var qhseKPIs []dashboardTypes.QhseKPI
	for rows.Next() {
		var qhseKPI dashboardTypes.QhseKPI
		var dropDownListItemModel tableComponentModels.DropDownListItemModel

		err := rows.Scan(
			&qhseKPI.Id,
			&qhseKPI.CompanyId,
			&qhseKPI.SNo,
			&qhseKPI.No,
			&qhseKPI.Title,
			&qhseKPI.Function.Id,
			&qhseKPI.LYKPI,
			&qhseKPI.AnnualTarget,
			&qhseKPI.January,
			&qhseKPI.February,
			&qhseKPI.March,
			&qhseKPI.April,
			&qhseKPI.May,
			&qhseKPI.June,
			&qhseKPI.July,
			&qhseKPI.August,
			&qhseKPI.September,
			&qhseKPI.October,
			&qhseKPI.November,
			&qhseKPI.December,
		)

		qhseKPI.Function, _ = dropDownListItemModel.GetById(qhseKPI.Function.Id)
		if err != nil {
			return nil, err
		}
		qhseKPIs = append(qhseKPIs, qhseKPI)
	}

	return qhseKPIs, nil
}

func (*QhseKPIModel) Create(qhseKPI dashboardTypes.QhseKPI) error {
	db := database.GetDatabase()
	_, err := db.Exec(`
			INSERT INTO qhsekpis (
				"id", "companyId", "sno", "no", "title", "function", 
				"lykpi", "annualTarget",
				"january", "february", "march", "april", "may", "june",
				"july", "august", "september", "october", "november", "december"
			) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`,
		qhseKPI.Id, qhseKPI.CompanyId, qhseKPI.SNo, qhseKPI.No, qhseKPI.Title, qhseKPI.Function.Id,
		qhseKPI.LYKPI, qhseKPI.AnnualTarget,
		qhseKPI.January, qhseKPI.February, qhseKPI.March, qhseKPI.April, qhseKPI.May, qhseKPI.June,
		qhseKPI.July, qhseKPI.August, qhseKPI.September, qhseKPI.October, qhseKPI.November, qhseKPI.December,
	)
	return err
}

func (*QhseKPIModel) Update(id string, fields map[string]interface{}) error {
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
	query := fmt.Sprintf(`UPDATE qhsekpis SET %s WHERE "id" = ?`, setClause)
	values = append(values, id)

	db := database.GetDatabase()
	_, err := db.Exec(query, values...)
	return err
}

func (*QhseKPIModel) DuplicateDefaults(companyId string) error {
	db := database.GetDatabase()

	rows, err := db.Query(`
			SELECT sno, no, title 
			FROM defaultqhsekpis
		`)
	if err != nil {
		return err
	}
	defer rows.Close()

	var defaultQhseKPIs []dashboardTypes.QhseKPI
	var qhseKPIModel QhseKPIModel

	for rows.Next() {
		var defaultKPI dashboardTypes.QhseKPI
		if err := rows.Scan(&defaultKPI.SNo, &defaultKPI.No, &defaultKPI.Title); err != nil {
			return err
		}
		defaultQhseKPIs = append(defaultQhseKPIs, defaultKPI)
	}

	// 2. Hər birini yeni şirkət ID-si ilə bazaya yazırıq
	for _, qhseKPI := range defaultQhseKPIs {
		// Create metodu vasitəsilə yeni rəqəmlər 0 olaraq (və ya NULL) yaranacaq
		if err := qhseKPIModel.Create(dashboardTypes.QhseKPI{
			Id:        qhseKPIModel.GenerateUniqueId(),
			CompanyId: companyId,
			SNo:       qhseKPI.SNo,
			No:        qhseKPI.No,
			Title:     qhseKPI.Title,
			Function: tableComponentTypes.DropDownListItem{
				Id: "",
			},
			LYKPI:        0,
			AnnualTarget: 0,
			January:      0,
			February:     0,
			March:        0,
			April:        0,
			May:          0,
			June:         0,
			July:         0,
			August:       0,
			September:    0,
			October:      0,
			November:     0,
			December:     0,
		}); err != nil {
			return err
		}
	}

	return nil
}
