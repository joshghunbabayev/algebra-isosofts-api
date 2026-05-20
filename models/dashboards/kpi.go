package dashboardModels

import (
	"algebra-isosofts-api/database"
	registerModels "algebra-isosofts-api/models/registers"
	registerComponentModels "algebra-isosofts-api/models/registers/components"
	tableComponentModels "algebra-isosofts-api/models/tableComponents"
	"algebra-isosofts-api/modules"
	dashboardTypes "algebra-isosofts-api/types/dashboards"
	tableComponentTypes "algebra-isosofts-api/types/tableComponents"
	"fmt"
	"strings"
	"time"
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
	var dropDownListItemModel tableComponentModels.DropDownListItemModel

	err := row.Scan(
		&kpi.Id,
		&kpi.CompanyId,
		&kpi.SNo,
		&kpi.No,
		&kpi.Title,
		&kpi.Function.Id,
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
	kpi.Function, _ = dropDownListItemModel.GetById(kpi.Function.Id)

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
		var dropDownListItemModel tableComponentModels.DropDownListItemModel

		err := rows.Scan(
			&kpi.Id,
			&kpi.CompanyId,
			&kpi.SNo,
			&kpi.No,
			&kpi.Title,
			&kpi.Function.Id,
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

		kpi.Function, _ = dropDownListItemModel.GetById(kpi.Function.Id)
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
		kpi.Id, kpi.CompanyId, kpi.SNo, kpi.No, kpi.Title, kpi.Function.Id,
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
		if err := rows.Scan(&defaultKPI.SNo, &defaultKPI.No, &defaultKPI.Title); err != nil {
			return err
		}
		defaultKPIs = append(defaultKPIs, defaultKPI)
	}

	// 2. Hər birini yeni şirkət ID-si ilə bazaya yazırıq
	for _, kpi := range defaultKPIs {
		// Create metodu vasitəsilə yeni rəqəmlər 0 olaraq (və ya NULL) yaranacaq
		if err := kpiModel.Create(dashboardTypes.KPI{
			Id:        kpiModel.GenerateUniqueId(),
			CompanyId: companyId,
			SNo:       kpi.SNo,
			No:        kpi.No,
			Title:     kpi.Title,
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

func (*KPIModel) UpdateMonthsByCompanyId(companyId string) error {
	var kpiModel KPIModel
	// var test string = kpis[i].CompanyId // sjfsd2fnsf string
	kpis, err := kpiModel.GetAll(map[string]interface{}{
		"companyId": companyId, // GetAll a Company Id ni Query Kimi Verir.
	})

	if err != nil {
		return err
	}

	for i := range kpis {
		var calculatedValue int64 = 0

		switch kpis[i].SNo {

		case 1: // Number of Not inspected Inventory/Equipment
			var eiModel registerModels.EIModel

			eis, _ := eiModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, ei := range eis {
				if !modules.IsDateBigger(ei.NVCD) {
					calculatedValue++
				}
			}

		case 2: // Number of Overdue Trainings
			var traModel registerModels.TRAModel

			tras, _ := traModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, tra := range tras {
				if !modules.IsDateBigger(tra.NCD) {
					calculatedValue++
				}
			}

		case 3: // Number of Not reviewed Documents
			var docModel registerModels.DOCModel

			docs, _ := docModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, doc := range docs {
				if !modules.IsDateBigger(doc.NextReviewDate) {
					calculatedValue++
				}
			}

		case 4: // Number of Not evaluated Vendors
			var venModel registerModels.VENModel

			vens, _ := venModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, ven := range vens {
				if !modules.IsDateBigger(ven.NRD) {
					calculatedValue++
				}
			}

		case 5: // Number of Not evaluated Customers
			var cusModel registerModels.CUSModel

			cuss, _ := cusModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, cus := range cuss {
				if !modules.IsDateBigger(cus.ReviewDate) {
					calculatedValue++
				}
			}

		case 6: // Number of Not evaluated Employees
			var eaModel registerModels.EAModel

			eas, _ := eaModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, ea := range eas {
				if !modules.IsDateBigger(ea.NextAppraisalDate) {
					calculatedValue++
				}
			}

		case 7: // Number of Open Findings
			var finModel registerModels.FINModel

			fins, _ := finModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, fin := range fins {
				if fin.FindingStatus.Value == "Open" {
					calculatedValue++
				}
			}

		case 8: // Number of Not completed A&O Activities
			var aopModel registerModels.AOPModel

			aops, _ := aopModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, aop := range aops {
				if !modules.IsDateBigger(aop.NextAoaDate) {
					calculatedValue++
				}
			}

		case 9: // Number of Overdue Actions
			var actionModel registerComponentModels.ActionModel

			actions, _ := actionModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, action := range actions {
				if action.VerificationStatus.Value == "Delayed" {
					calculatedValue++
				}
			}

		case 10: // Number of Residual High Business Risks / Opportunity Level
			var brModel registerModels.BRModel

			brs, _ := brModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, br := range brs {
				if br.ResidualRiskSeverity*br.ResidualRiskLikelihood >= 12 {
					calculatedValue++
				}
			}

		case 11: // Number of Residual High H&S Risks
			var hsrModel registerModels.HSRModel

			hsrs, _ := hsrModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, hsr := range hsrs {
				if hsr.ResidualRiskSeverity*hsr.ResidualRiskLikelihood >= 12 {
					calculatedValue++
				}
			}

		case 12: // Number of Legal Requirment with Residual High Risk
			var legModel registerModels.LEGModel

			legs, _ := legModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, leg := range legs {
				if leg.ResidualRiskSeverity*leg.ResidualRiskLikelihood >= 12 {
					calculatedValue++
				}
			}

		case 13: // Number of E&A Aspects with High Residual Significance Level
			var eaiModel registerModels.EAIModel

			eais, _ := eaiModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, eai := range eais {
				if eai.RDOSProbability*eai.RDOSSeverity*eai.RDOSDuration*eai.RDOSScale >= 80 {
					calculatedValue++
				}
			}

		case 14: // Equipment Safety Rate %
			var eiModel registerModels.EIModel

			var num int64 = 0

			eis, _ := eiModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, ei := range eis {
				if modules.IsDateBigger(ei.NVCD) {
					num++
				}
			}

			if len(eis) > 0 {
				calculatedValue = int64(float64(num) / float64(len(eis)) * 100)
			} else {
				calculatedValue = 0
			}

		case 15: // Training Validity Rate %
			var traModel registerModels.TRAModel

			var num int64 = 0

			tras, _ := traModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, tra := range tras {
				if modules.IsDateBigger(tra.NCD) {
					num++
				}
			}

			if len(tras) > 0 {
				calculatedValue = int64(float64(num) / float64(len(tras)) * 100)
			} else {
				calculatedValue = 0
			}

		case 16: // Documents Review Rate %
			var docModel registerModels.DOCModel

			var num int64 = 0

			docs, _ := docModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, doc := range docs {
				if modules.IsDateBigger(doc.NextReviewDate) {
					num++
				}
			}

			if len(docs) > 0 {
				calculatedValue = int64(float64(num) / float64(len(docs)) * 100)
			} else {
				calculatedValue = 0
			}

		case 17: // Vendors Evaluation Rate %
			var venModel registerModels.VENModel

			var num int64 = 0

			vens, _ := venModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, ven := range vens {
				if modules.IsDateBigger(ven.NRD) {
					num++
				}
			}

			if len(vens) > 0 {
				calculatedValue = int64(float64(num) / float64(len(vens)) * 100)
			} else {
				calculatedValue = 0
			}

		case 18: // Average Vendors Satisfuction Score
			var venModel registerModels.VENModel

			var num int64 = 0

			vens, _ := venModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			if len(vens) > 0 {
				for _, ven := range vens {
					num += int64(ven.QGS + ven.Communication + ven.OTD + ven.Documentation + ven.HS + ven.Environment)
				}

				calculatedValue = int64(float64(num) / float64(len(vens)))
			} else {
				calculatedValue = 0
			}

		case 19: // Customer Feedback Evaluation Rate %
			var cusModel registerModels.CUSModel

			var num int64 = 0

			cuss, _ := cusModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, cus := range cuss {
				if modules.IsDateBigger(cus.ReviewDate) {
					num++
				}
			}

			if len(cuss) > 0 {
				calculatedValue = int64(float64(num) / float64(len(cuss)) * 100)
			} else {
				calculatedValue = 0
			}

		case 20: // Average Customer Satisfuction Score
			var cusModel registerModels.CUSModel

			var num int64 = 0

			cuss, _ := cusModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			if len(cuss) > 0 {
				for _, cus := range cuss {
					num += int64(cus.QGS + cus.Communication + cus.OTD + cus.Documentation + cus.HS + cus.Environment)
				}

				calculatedValue = int64(float64(num) / float64(len(cuss)))
			} else {
				calculatedValue = 0
			}

		case 21: // Number of Jobs
			var fbModel registerModels.FBModel

			fbs, _ := fbModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			calculatedValue = int64(len(fbs))

		case 22: // Employee Performance Appraisal Status Rate %
			var eaModel registerModels.EAModel

			var num int64 = 0

			eas, _ := eaModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, ea := range eas {
				if modules.IsDateBigger(ea.NextAppraisalDate) {
					num++
				}
			}

			if len(eas) > 0 {
				calculatedValue = int64(float64(num) / float64(len(eas)) * 100)
			} else {
				calculatedValue = 0
			}

		case 23: // Average Employee Skills Appraisal Score
			var eaModel registerModels.EAModel

			var num int64 = 0

			eas, _ := eaModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			if len(eas) > 0 {
				for _, ea := range eas {
					num += int64(ea.JobQuality + ea.LeadershipSkills + ea.ManagementSkills + ea.BehavioralSkills + ea.EffectivenessOfTrainings)
				}

				calculatedValue = int64(float64(num) / float64(len(eas)))
			} else {
				calculatedValue = 0
			}

		case 24: // Number of Residual High MoC Risks
			var mocModel registerModels.MOCModel

			mocs, _ := mocModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, moc := range mocs {
				if moc.ResidualRiskSeverity*moc.ResidualRiskLikelihood >= 12 {
					calculatedValue++
				}
			}

		case 25: // Findings Closure Rate
			var finModel registerModels.FINModel

			var num int64 = 0

			fins, _ := finModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, fin := range fins {
				if fin.FindingStatus.Value == "Closed" {
					num++
				}
			}

			if len(fins) > 0 {
				calculatedValue = int64(float64(num) / float64(len(fins)) * 100)
			} else {
				calculatedValue = 0
			}

		case 26: // Number of Non-Conformancies
			var finModel registerModels.FINModel

			fins, _ := finModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, fin := range fins {
				if fin.CategoryOfFinding.Value == "Non conformance" {
					calculatedValue++
				}
			}

		case 27: // Number of Opportunities for Improvement
			var finModel registerModels.FINModel

			fins, _ := finModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, fin := range fins {
				if fin.CategoryOfFinding.Value == "Opportunity for Improvement" {
					calculatedValue++
				}
			}

		case 28: // Number of Internal Complaints
			var finModel registerModels.FINModel

			fins, _ := finModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, fin := range fins {
				if fin.CategoryOfFinding.Value == "Internal complaint" {
					calculatedValue++
				}
			}

		case 29: // Number of External Complaints
			var finModel registerModels.FINModel

			fins, _ := finModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, fin := range fins {
				if fin.CategoryOfFinding.Value == "External Complaint" {
					calculatedValue++
				}
			}

		case 30: // Number of Good Practices
			var finModel registerModels.FINModel

			fins, _ := finModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, fin := range fins {
				if fin.CategoryOfFinding.Value == "Good practice" {
					calculatedValue++
				}
			}

		case 31: // Number of Near-Misses
			var finModel registerModels.FINModel

			fins, _ := finModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, fin := range fins {
				if fin.CategoryOfFinding.Value == "Near miss" {
					calculatedValue++
				}
			}

		case 32: // Number of Incidents
			var finModel registerModels.FINModel

			fins, _ := finModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, fin := range fins {
				if fin.CategoryOfFinding.Value == "Incident" {
					calculatedValue++
				}
			}

		case 33: // Rate of Incidents %
			var finModel registerModels.FINModel

			var num int64 = 0

			fins, _ := finModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, fin := range fins {
				if fin.CategoryOfFinding.Value == "Incident" {
					num++
				}
			}

			if len(fins) > 0 {
				calculatedValue = int64(float64(num) / float64(len(fins)) * 100)
			} else {
				calculatedValue = 0
			}

		case 34: // Number of Accident
			var finModel registerModels.FINModel

			fins, _ := finModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, fin := range fins {
				if fin.CategoryOfFinding.Value == "Accident" {
					calculatedValue++
				}
			}

		case 35: // Number of Unsafe Actions
			var finModel registerModels.FINModel

			fins, _ := finModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, fin := range fins {
				if fin.CategoryOfFinding.Value == "Unsafe actions" {
					calculatedValue++
				}
			}

		case 36: // Number of Unsafe Conditions
			var finModel registerModels.FINModel

			fins, _ := finModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, fin := range fins {
				if fin.CategoryOfFinding.Value == "Unsafe condition" {
					calculatedValue++
				}
			}

		case 37: // Number of Environmental Incidents
			var finModel registerModels.FINModel

			fins, _ := finModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, fin := range fins {
				if fin.CategoryOfFinding.Value == "Environmental Incident" {
					calculatedValue++
				}
			}

		case 38: // Assurance & Oversight Plan Implementation Rate %
			var aopModel registerModels.AOPModel

			var num int64 = 0

			aops, _ := aopModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, aop := range aops {
				if modules.IsDateBigger(aop.NextAoaDate) {
					num++
				}
			}

			if len(aops) > 0 {
				calculatedValue = int64(float64(num) / float64(len(aops)) * 100)
			} else {
				calculatedValue = 0
			}

		case 39: // Actions Closure Rate
			var actionModel registerComponentModels.ActionModel

			var num int64 = 0

			actions, _ := actionModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, action := range actions {
				if action.Status.Value == "100" {
					num++
				}
			}

			if len(actions) > 0 {
				calculatedValue = int64(float64(num) / float64(len(actions)) * 100)
			} else {
				calculatedValue = 0
			}

		default:
			calculatedValue = 0
		}

		kpiModel.Update(kpis[i].Id, map[string]interface{}{
			strings.ToLower(time.Now().Month().String()): calculatedValue,
		})
	}
	return nil
}

func (*KPIModel) UpdateMonthsAll() error {
	var kpiModel KPIModel
	// var test string = kpis[i].CompanyId // sjfsd2fnsf string
	kpis, err := kpiModel.GetAll(map[string]interface{}{})

	if err != nil {
		return err
	}

	for i := range kpis {
		var calculatedValue int64 = 0

		switch kpis[i].SNo {

		case 1: // Number of Not inspected Inventory/Equipment
			var eiModel registerModels.EIModel

			eis, _ := eiModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, ei := range eis {
				if !modules.IsDateBigger(ei.NVCD) {
					calculatedValue++
				}
			}

		case 2: // Number of Overdue Trainings
			var traModel registerModels.TRAModel

			tras, _ := traModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, tra := range tras {
				if !modules.IsDateBigger(tra.NCD) {
					calculatedValue++
				}
			}

		case 3: // Number of Not reviewed Documents
			var docModel registerModels.DOCModel

			docs, _ := docModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, doc := range docs {
				if !modules.IsDateBigger(doc.NextReviewDate) {
					calculatedValue++
				}
			}

		case 4: // Number of Not evaluated Vendors
			var venModel registerModels.VENModel

			vens, _ := venModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, ven := range vens {
				if !modules.IsDateBigger(ven.NRD) {
					calculatedValue++
				}
			}

		case 5: // Number of Not evaluated Customers
			var cusModel registerModels.CUSModel

			cuss, _ := cusModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, cus := range cuss {
				if !modules.IsDateBigger(cus.ReviewDate) {
					calculatedValue++
				}
			}

		case 6: // Number of Not evaluated Employees
			var eaModel registerModels.EAModel

			eas, _ := eaModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, ea := range eas {
				if !modules.IsDateBigger(ea.NextAppraisalDate) {
					calculatedValue++
				}
			}

		case 7: // Number of Open Findings
			var finModel registerModels.FINModel

			fins, _ := finModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, fin := range fins {
				if fin.FindingStatus.Value == "Open" {
					calculatedValue++
				}
			}

		case 8: // Number of Not completed A&O Activities
			var aopModel registerModels.AOPModel

			aops, _ := aopModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, aop := range aops {
				if !modules.IsDateBigger(aop.NextAoaDate) {
					calculatedValue++
				}
			}

		case 9: // Number of Overdue Actions
			var actionModel registerComponentModels.ActionModel

			actions, _ := actionModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, action := range actions {
				if action.VerificationStatus.Value == "Delayed" {
					calculatedValue++
				}
			}

		case 10: // Number of Residual High Business Risks / Opportunity Level
			var brModel registerModels.BRModel

			brs, _ := brModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, br := range brs {
				if br.ResidualRiskSeverity*br.ResidualRiskLikelihood >= 12 {
					calculatedValue++
				}
			}

		case 11: // Number of Residual High H&S Risks
			var hsrModel registerModels.HSRModel

			hsrs, _ := hsrModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, hsr := range hsrs {
				if hsr.ResidualRiskSeverity*hsr.ResidualRiskLikelihood >= 12 {
					calculatedValue++
				}
			}

		case 12: // Number of Legal Requirment with Residual High Risk
			var legModel registerModels.LEGModel

			legs, _ := legModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, leg := range legs {
				if leg.ResidualRiskSeverity*leg.ResidualRiskLikelihood >= 12 {
					calculatedValue++
				}
			}

		case 13: // Number of E&A Aspects with High Residual Significance Level
			var eaiModel registerModels.EAIModel

			eais, _ := eaiModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, eai := range eais {
				if eai.RDOSProbability*eai.RDOSSeverity*eai.RDOSDuration*eai.RDOSScale >= 80 {
					calculatedValue++
				}
			}

		case 14: // Equipment Safety Rate %
			var eiModel registerModels.EIModel

			var num int64 = 0

			eis, _ := eiModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, ei := range eis {
				if modules.IsDateBigger(ei.NVCD) {
					num++
				}
			}

			if len(eis) > 0 {
				calculatedValue = int64(float64(num) / float64(len(eis)) * 100)
			} else {
				calculatedValue = 0
			}

		case 15: // Training Validity Rate %
			var traModel registerModels.TRAModel

			var num int64 = 0

			tras, _ := traModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, tra := range tras {
				if modules.IsDateBigger(tra.NCD) {
					num++
				}
			}

			if len(tras) > 0 {
				calculatedValue = int64(float64(num) / float64(len(tras)) * 100)
			} else {
				calculatedValue = 0
			}

		case 16: // Documents Review Rate %
			var docModel registerModels.DOCModel

			var num int64 = 0

			docs, _ := docModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, doc := range docs {
				if modules.IsDateBigger(doc.NextReviewDate) {
					num++
				}
			}

			if len(docs) > 0 {
				calculatedValue = int64(float64(num) / float64(len(docs)) * 100)
			} else {
				calculatedValue = 0
			}

		case 17: // Vendors Evaluation Rate %
			var venModel registerModels.VENModel

			var num int64 = 0

			vens, _ := venModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, ven := range vens {
				if modules.IsDateBigger(ven.NRD) {
					num++
				}
			}

			if len(vens) > 0 {
				calculatedValue = int64(float64(num) / float64(len(vens)) * 100)
			} else {
				calculatedValue = 0
			}

		case 18: // Average Vendors Satisfuction Score
			var venModel registerModels.VENModel

			var num int64 = 0

			vens, _ := venModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			if len(vens) > 0 {
				for _, ven := range vens {
					num += int64(ven.QGS + ven.Communication + ven.OTD + ven.Documentation + ven.HS + ven.Environment)
				}

				calculatedValue = int64(float64(num) / float64(len(vens)))
			} else {
				calculatedValue = 0
			}

		case 19: // Customer Feedback Evaluation Rate %
			var cusModel registerModels.CUSModel

			var num int64 = 0

			cuss, _ := cusModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, cus := range cuss {
				if modules.IsDateBigger(cus.ReviewDate) {
					num++
				}
			}

			if len(cuss) > 0 {
				calculatedValue = int64(float64(num) / float64(len(cuss)) * 100)
			} else {
				calculatedValue = 0
			}

		case 20: // Average Customer Satisfuction Score
			var cusModel registerModels.CUSModel

			var num int64 = 0

			cuss, _ := cusModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			if len(cuss) > 0 {
				for _, cus := range cuss {
					num += int64(cus.QGS + cus.Communication + cus.OTD + cus.Documentation + cus.HS + cus.Environment)
				}

				calculatedValue = int64(float64(num) / float64(len(cuss)))
			} else {
				calculatedValue = 0
			}

		case 21: // Number of Jobs
			var fbModel registerModels.FBModel

			fbs, _ := fbModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			calculatedValue = int64(len(fbs))

		case 22: // Employee Performance Appraisal Status Rate %
			var eaModel registerModels.EAModel

			var num int64 = 0

			eas, _ := eaModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, ea := range eas {
				if modules.IsDateBigger(ea.NextAppraisalDate) {
					num++
				}
			}

			if len(eas) > 0 {
				calculatedValue = int64(float64(num) / float64(len(eas)) * 100)
			} else {
				calculatedValue = 0
			}

		case 23: // Average Employee Skills Appraisal Score
			var eaModel registerModels.EAModel

			var num int64 = 0

			eas, _ := eaModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			if len(eas) > 0 {
				for _, ea := range eas {
					num += int64(ea.JobQuality + ea.LeadershipSkills + ea.ManagementSkills + ea.BehavioralSkills + ea.EffectivenessOfTrainings)
				}

				calculatedValue = int64(float64(num) / float64(len(eas)))
			} else {
				calculatedValue = 0
			}

		case 24: // Number of Residual High MoC Risks
			var mocModel registerModels.MOCModel

			mocs, _ := mocModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, moc := range mocs {
				if moc.ResidualRiskSeverity*moc.ResidualRiskLikelihood >= 12 {
					calculatedValue++
				}
			}

		case 25: // Findings Closure Rate
			var finModel registerModels.FINModel

			var num int64 = 0

			fins, _ := finModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, fin := range fins {
				if fin.FindingStatus.Value == "Closed" {
					num++
				}
			}

			if len(fins) > 0 {
				calculatedValue = int64(float64(num) / float64(len(fins)) * 100)
			} else {
				calculatedValue = 0
			}

		case 26: // Number of Non-Conformancies
			var finModel registerModels.FINModel

			fins, _ := finModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, fin := range fins {
				if fin.CategoryOfFinding.Value == "Non conformance" {
					calculatedValue++
				}
			}

		case 27: // Number of Opportunities for Improvement
			var finModel registerModels.FINModel

			fins, _ := finModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, fin := range fins {
				if fin.CategoryOfFinding.Value == "Opportunity for Improvement" {
					calculatedValue++
				}
			}

		case 28: // Number of Internal Complaints
			var finModel registerModels.FINModel

			fins, _ := finModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, fin := range fins {
				if fin.CategoryOfFinding.Value == "Internal complaint" {
					calculatedValue++
				}
			}

		case 29: // Number of External Complaints
			var finModel registerModels.FINModel

			fins, _ := finModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, fin := range fins {
				if fin.CategoryOfFinding.Value == "External Complaint" {
					calculatedValue++
				}
			}

		case 30: // Number of Good Practices
			var finModel registerModels.FINModel

			fins, _ := finModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, fin := range fins {
				if fin.CategoryOfFinding.Value == "Good practice" {
					calculatedValue++
				}
			}

		case 31: // Number of Near-Misses
			var finModel registerModels.FINModel

			fins, _ := finModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, fin := range fins {
				if fin.CategoryOfFinding.Value == "Near miss" {
					calculatedValue++
				}
			}

		case 32: // Number of Incidents
			var finModel registerModels.FINModel

			fins, _ := finModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, fin := range fins {
				if fin.CategoryOfFinding.Value == "Incident" {
					calculatedValue++
				}
			}

		case 33: // Rate of Incidents %
			var finModel registerModels.FINModel

			var num int64 = 0

			fins, _ := finModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, fin := range fins {
				if fin.CategoryOfFinding.Value == "Incident" {
					num++
				}
			}

			if len(fins) > 0 {
				calculatedValue = int64(float64(num) / float64(len(fins)) * 100)
			} else {
				calculatedValue = 0
			}

		case 34: // Number of Accident
			var finModel registerModels.FINModel

			fins, _ := finModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, fin := range fins {
				if fin.CategoryOfFinding.Value == "Accident" {
					calculatedValue++
				}
			}

		case 35: // Number of Unsafe Actions
			var finModel registerModels.FINModel

			fins, _ := finModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, fin := range fins {
				if fin.CategoryOfFinding.Value == "Unsafe actions" {
					calculatedValue++
				}
			}

		case 36: // Number of Unsafe Conditions
			var finModel registerModels.FINModel

			fins, _ := finModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, fin := range fins {
				if fin.CategoryOfFinding.Value == "Unsafe condition" {
					calculatedValue++
				}
			}

		case 37: // Number of Environmental Incidents
			var finModel registerModels.FINModel

			fins, _ := finModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, fin := range fins {
				if fin.CategoryOfFinding.Value == "Environmental Incident" {
					calculatedValue++
				}
			}

		case 38: // Assurance & Oversight Plan Implementation Rate %
			var aopModel registerModels.AOPModel

			var num int64 = 0

			aops, _ := aopModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, aop := range aops {
				if modules.IsDateBigger(aop.NextAoaDate) {
					num++
				}
			}

			if len(aops) > 0 {
				calculatedValue = int64(float64(num) / float64(len(aops)) * 100)
			} else {
				calculatedValue = 0
			}

		case 39: // Actions Closure Rate
			var actionModel registerComponentModels.ActionModel

			var num int64 = 0

			actions, _ := actionModel.GetAll(map[string]interface{}{
				"dbStatus":  "active",
				"companyId": kpis[i].CompanyId,
			})

			for _, action := range actions {
				if action.Status.Value == "100" {
					num++
				}
			}

			if len(actions) > 0 {
				calculatedValue = int64(float64(num) / float64(len(actions)) * 100)
			} else {
				calculatedValue = 0
			}

		default:
			calculatedValue = 0
		}

		kpiModel.Update(kpis[i].Id, map[string]interface{}{
			strings.ToLower(time.Now().Month().String()): calculatedValue,
		})
	}
	return nil
}
