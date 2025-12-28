package registerModels

import (
	"algebra-isosofts-api/database"
	registerComponentModels "algebra-isosofts-api/models/registers/components"
	tableComponentModels "algebra-isosofts-api/models/tableComponents"
	"algebra-isosofts-api/modules"
	registerTypes "algebra-isosofts-api/types/registers"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

type VENModel struct {
}

func (*VENModel) GenerateUniqueId() string {
	Id := modules.GenerateRandomString(30)

	var venModel VENModel

	ven, _ := venModel.GetById(Id)

	if ven.IsEmpty() {
		return Id
	} else {
		return venModel.GenerateUniqueId()
	}
}

func (*VENModel) GenerateUniqueNo() string {
	db := database.GetDatabase()

	year := time.Now().Format("06")

	var lastNo string
	db.QueryRow(`
		SELECT "no" 
		FROM venregisters 
		WHERE "no" LIKE ? 
		ORDER BY "no" DESC 
		LIMIT 1
		`,
		"VENR/"+year+"/%",
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

	newNo := fmt.Sprintf("VENR/%s/%04d", year, nextNumber)
	return newNo
}

func (*VENModel) GetById(Id string) (registerTypes.VEN, error) {
	db := database.GetDatabase()
	row := db.QueryRow(`
			SELECT * 
			FROM venregisters
			WHERE id = ?
		`,
		Id,
	)

	var ven registerTypes.VEN
	var dropDownListItemModel tableComponentModels.DropDownListItemModel
	var venModel VENModel
	var actionModel registerComponentModels.ActionModel

	err := row.Scan(
		&ven.Id,
		&ven.No,
		&ven.Name,
		&ven.RegNumber,
		&ven.Scope1.Id,
		&ven.Scope2.Id,
		&ven.Scope3.Id,
		&ven.RegistrationDate,
		&ven.ReviewDate,
		&ven.Approved,
		&ven.DbStatus,
		&ven.DbLastStatus,
	)
	ven.Scope1, _ = dropDownListItemModel.GetById(ven.Scope1.Id)
	ven.Scope2, _ = dropDownListItemModel.GetById(ven.Scope2.Id)
	ven.Scope3, _ = dropDownListItemModel.GetById(ven.Scope3.Id)
	ven.Actions, _ = actionModel.GetAll(map[string]interface{}{
		"registerId": ven.Id,
	})

	venModel.SetScores(ven.Id, &ven)

	return ven, err
}

func (*VENModel) GetAll(filters map[string]interface{}) ([]registerTypes.VEN, error) {
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
			SELECT * FROM venregisters %s
		`,
		whereClause,
	)
	rows, err := db.Query(query, values...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vens []registerTypes.VEN

	for rows.Next() {
		var ven registerTypes.VEN
		var dropDownListItemModel tableComponentModels.DropDownListItemModel
		var venModel VENModel
		var actionModel registerComponentModels.ActionModel

		rows.Scan(
			&ven.Id,
			&ven.No,
			&ven.Name,
			&ven.RegNumber,
			&ven.Scope1.Id,
			&ven.Scope2.Id,
			&ven.Scope3.Id,
			&ven.RegistrationDate,
			&ven.ReviewDate,
			&ven.DbStatus,
			&ven.DbLastStatus,
		)
		ven.Scope1, _ = dropDownListItemModel.GetById(ven.Scope1.Id)
		ven.Scope2, _ = dropDownListItemModel.GetById(ven.Scope2.Id)
		ven.Scope3, _ = dropDownListItemModel.GetById(ven.Scope3.Id)
		ven.Actions, _ = actionModel.GetAll(map[string]interface{}{
			"registerId": ven.Id,
		})

		venModel.SetScores(ven.Id, &ven)

		vens = append(vens, ven)
	}

	return vens, nil
}

func (*VENModel) Create(ven registerTypes.VEN) error {
	db := database.GetDatabase()
	_, err := db.Exec(`
			INSERT INTO venregisters ( 
				"id",
				"no",
				"name", 
				"regNumber", 
				"scope1", 
				"scope2", 
				"scope3", 
				"registrationDate", 
				"reviewDate", 
				"approved", 
				"dbStatus",
				"dbLastStatus"
			) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`,
		ven.Id,
		ven.No,
		ven.Name,
		ven.RegNumber,
		ven.Scope1.Id,
		ven.Scope2.Id,
		ven.Scope3.Id,
		ven.RegistrationDate,
		ven.ReviewDate,
		ven.Approved,
		ven.DbStatus,
		ven.DbLastStatus,
	)

	if err != nil {
		return err
	}

	return nil
}

func (*VENModel) Update(Id string, fields map[string]interface{}) error {
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
			UPDATE venregisters 
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

func (*VENModel) SetScores(Id string, ven *registerTypes.VEN) error {
	var vendorFeedbackModel registerComponentModels.VendorFeedbackModel

	vendorFeedbacks, err := vendorFeedbackModel.GetAll(map[string]interface{}{
		"vendorId": Id,
		"dbStatus": "active",
	})

	if len(vendorFeedbacks) == 0 {
		ven.QGS = 0
		ven.Communication = 0
		ven.OTD = 0
		ven.Documentation = 0
		ven.HS = 0
		ven.Environment = 0
	} else {
		var sumQGS, sumCommunication, sumOTD, sumDocumentation, sumHS, sumEnvironment int

		for _, vendorFeedback := range vendorFeedbacks {
			sumQGS += int(vendorFeedback.QGS)
			sumCommunication += int(vendorFeedback.Communication)
			sumOTD += int(vendorFeedback.OTD)
			sumDocumentation += int(vendorFeedback.Documentation)
			sumHS += int(vendorFeedback.HS)
			sumEnvironment += int(vendorFeedback.Environment)
		}

		count := float64(len(vendorFeedbacks))

		ven.QGS = int8(math.Round(float64(sumQGS) / count))
		ven.Communication = int8(math.Round(float64(sumCommunication) / count))
		ven.OTD = int8(math.Round(float64(sumOTD) / count))
		ven.Documentation = int8(math.Round(float64(sumDocumentation) / count))
		ven.HS = int8(math.Round(float64(sumHS) / count))
		ven.Environment = int8(math.Round(float64(sumEnvironment) / count))
	}

	return err
}
