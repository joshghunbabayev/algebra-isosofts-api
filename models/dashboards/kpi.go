package dashboardModels

import (
	"algebra-isosofts-api/database"
	"algebra-isosofts-api/modules"
	dashboardTypes "algebra-isosofts-api/types/dashboards"
	"fmt"
	"strings"
)

type KPIModel struct{}

func (*KPIModel) GenerateUniqueId() string {
	id := modules.GenerateRandomString(30)
	var kpiModel KPIModel
	kpi, _ := kpiModel.GetById(id)

	if kpi.IsEmpty() {
		return id
	}
	return kpiModel.GenerateUniqueId()
}

func (*KPIModel) GetById(id string) (dashboardTypes.KPI, error) {
	db := database.GetDatabase()
	row := db.QueryRow(`
			SELECT * FROM kpis
			WHERE id = ?
		`, id)

	var kpi dashboardTypes.KPI
	err := row.Scan(
		&kpi.Id,
		&kpi.CompanyId,
		&kpi.SNo,
		&kpi.No,
		&kpi.Title,
		&kpi.Function,
		&kpi.LYKPI,
		&kpi.AnnualTarget,
		&kpi.January,
		&kpi.February,
		&kpi.March,
		&kpi.April,
		&kpi.May,
		&kpi.June,
		&kpi.July,
		&kpi.August,
		&kpi.September,
		&kpi.October,
		&kpi.November,
		&kpi.December,
	)

	return kpi, err
}

func (*KPIModel) GetAll(filters map[string]interface{}) ([]dashboardTypes.KPI, error) {
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

	query := fmt.Sprintf(`SELECT * FROM kpis %s`, whereClause)
	rows, err := db.Query(query, values...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var kpis []dashboardTypes.KPI
	for rows.Next() {
		var kpi dashboardTypes.KPI
		err := rows.Scan(
			&kpi.Id,
			&kpi.CompanyId,
			&kpi.SNo,
			&kpi.No,
			&kpi.Title,
			&kpi.Function,
			&kpi.LYKPI,
			&kpi.AnnualTarget,
			&kpi.January,
			&kpi.February,
			&kpi.March,
			&kpi.April,
			&kpi.May,
			&kpi.June,
			&kpi.July,
			&kpi.August,
			&kpi.September,
			&kpi.October,
			&kpi.November,
			&kpi.December,
		)
		if err != nil {
			return nil, err
		}
		kpis = append(kpis, kpi)
	}

	return kpis, nil
}

func (*KPIModel) Create(kpi dashboardTypes.KPI) error {
	db := database.GetDatabase()
	_, err := db.Exec(`
			INSERT INTO kpis (
				"id", "companyId", "sno", "no", "title", "function", 
				"lykpi", "annualTarget",
				"january", "february", "march", "april", "may", "june",
				"july", "august", "september", "october", "november", "december"
			) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`,
		kpi.Id, kpi.CompanyId, kpi.SNo, kpi.No, kpi.Title, kpi.Function,
		kpi.LYKPI, kpi.AnnualTarget,
		kpi.January, kpi.February, kpi.March, kpi.April, kpi.May, kpi.June,
		kpi.July, kpi.August, kpi.September, kpi.October, kpi.November, kpi.December,
	)
	return err
}

func (*KPIModel) Update(id string, fields map[string]interface{}) error {
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
	query := fmt.Sprintf(`UPDATE kpis SET %s WHERE "id" = ?`, setClause)
	values = append(values, id)

	db := database.GetDatabase()
	_, err := db.Exec(query, values...)
	return err
}

func (*KPIModel) DuplicateDefaults(companyId string) error {
	db := database.GetDatabase()

	// 1. Şablondakı (default) KPI-ları gətiririk
	rows, err := db.Query(`
			SELECT sno, no, title 
			FROM defaultkpis
		`)
	if err != nil {
		return err
	}
	defer rows.Close()

	var defaultKPIs []dashboardTypes.KPI
	var kpiModel KPIModel

	for rows.Next() {
		var defaultKPI dashboardTypes.KPI
		// Şəkildəki struktura uyğun olaraq scan edirik
		if err := rows.Scan(&defaultKPI.SNo, &defaultKPI.No, &defaultKPI.Title); err != nil {
			return err
		}
		defaultKPIs = append(defaultKPIs, defaultKPI)
	}

	// 2. Hər birini yeni şirkət ID-si ilə bazaya yazırıq
	for _, kpi := range defaultKPIs {
		// Create metodu vasitəsilə yeni rəqəmlər 0 olaraq (və ya NULL) yaranacaq
		if err := kpiModel.Create(dashboardTypes.KPI{
			Id:           kpiModel.GenerateUniqueId(),
			CompanyId:    companyId,
			SNo:          kpi.SNo,
			No:           kpi.No,
			Title:        kpi.Title,
			Function:     "",
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
