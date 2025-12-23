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

type DOCModel struct {
}

func (*DOCModel) GenerateUniqueId() string {
	Id := modules.GenerateRandomString(30)

	var docModel DOCModel

	doc, _ := docModel.GetById(Id)

	if doc.IsEmpty() {
		return Id
	} else {
		return docModel.GenerateUniqueId()
	}
}

func (*DOCModel) GenerateUniqueNo() string {
	db := database.GetDatabase()

	year := time.Now().Format("06")

	var lastNo string
	db.QueryRow(`
		SELECT "no" 
		FROM docregisters 
		WHERE "no" LIKE ? 
		ORDER BY "no" DESC 
		LIMIT 1
		`,
		"DOCR/"+year+"/%",
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

	newNo := fmt.Sprintf("DOCR/%s/%04d", year, nextNumber)
	return newNo
}

func (*DOCModel) GenerateUniqueNumber(prefix string) string {
	db := database.GetDatabase()

	var lastNumber string

	db.QueryRow(`
		SELECT "no"
		FROM docregisters
		ORDER BY CAST(SUBSTR("no", LENGTH("no") - 3, 4) AS INTEGER) DESC
		LIMIT 1
	`).Scan(&lastNumber)

	var nextNumber int

	if lastNumber == "" {
		nextNumber = 1
	} else {
		parts := strings.Split(lastNumber, "/")
		numPart := parts[len(parts)-1]
		num, _ := strconv.Atoi(numPart)
		nextNumber = num + 1
	}

	newNumber := fmt.Sprintf("%s/%04d", prefix, nextNumber)
	return newNumber
}

func (*DOCModel) GetById(Id string) (registerTypes.DOC, error) {
	db := database.GetDatabase()
	row := db.QueryRow(`
			SELECT * 
			FROM docregisters
			WHERE id = ?
		`,
		Id,
	)

	var doc registerTypes.DOC
	var dropDownListItemModel tableComponentModels.DropDownListItemModel
	var actionModel registerComponentModels.ActionModel

	err := row.Scan(
		&doc.Id,
		&doc.No,
		&doc.Name,
		&doc.Origin.Id,
		&doc.Number,
		&doc.DepntFunctionName.Id,
		&doc.Type.Id,
		&doc.SerialNumber,
		&doc.RevNumber,
		&doc.Issuer,
		&doc.Approver,
		&doc.IssueDate,
		&doc.NextReviewDate,
		&doc.Actual,
		&doc.DbStatus,
		&doc.DbLastStatus,
	)
	doc.Origin, _ = dropDownListItemModel.GetById(doc.Origin.Id)
	doc.DepntFunctionName, _ = dropDownListItemModel.GetById(doc.DepntFunctionName.Id)
	doc.Type, _ = dropDownListItemModel.GetById(doc.Type.Id)
	doc.Actions, _ = actionModel.GetAll(map[string]interface{}{
		"registerId": doc.Id,
	})

	return doc, err
}

func (*DOCModel) GetAll(filters map[string]interface{}) ([]registerTypes.DOC, error) {
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
			SELECT * FROM docregisters %s
		`,
		whereClause,
	)
	rows, err := db.Query(query, values...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var docs []registerTypes.DOC

	for rows.Next() {
		var doc registerTypes.DOC
		var dropDownListItemModel tableComponentModels.DropDownListItemModel
		var actionModel registerComponentModels.ActionModel

		rows.Scan(
			&doc.Id,
			&doc.No,
			&doc.Name,
			&doc.Origin.Id,
			&doc.Number,
			&doc.DepntFunctionName.Id,
			&doc.Type.Id,
			&doc.SerialNumber,
			&doc.RevNumber,
			&doc.Issuer,
			&doc.Approver,
			&doc.IssueDate,
			&doc.NextReviewDate,
			&doc.Actual,
			&doc.DbStatus,
			&doc.DbLastStatus,
		)
		doc.Origin, _ = dropDownListItemModel.GetById(doc.Origin.Id)
		doc.DepntFunctionName, _ = dropDownListItemModel.GetById(doc.DepntFunctionName.Id)
		doc.Type, _ = dropDownListItemModel.GetById(doc.Type.Id)
		doc.Actions, _ = actionModel.GetAll(map[string]interface{}{
			"registerId": doc.Id,
		})

		docs = append(docs, doc)
	}

	return docs, nil
}

func (*DOCModel) Create(doc registerTypes.DOC) error {
	db := database.GetDatabase()
	_, err := db.Exec(`
			INSERT INTO docregisters ( 
				"id",
				"no",
				"name", 
				"origin", 
				"number", 
				"depntFunctionName", 
				"type", 
				"serialNumber", 
				"revNumber", 
				"issuer", 
				"approver", 
				"issueDate", 
				"nextReviewDate", 
				"actual",
				"dbStatus",
				"dbLastStatus"
			) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`,
		doc.Id,
		doc.No,
		doc.Name,
		doc.Origin.Id,
		doc.Number,
		doc.DepntFunctionName.Id,
		doc.Type.Id,
		doc.SerialNumber,
		doc.RevNumber,
		doc.Issuer,
		doc.Approver,
		doc.IssueDate,
		doc.NextReviewDate,
		doc.Actual,
		doc.DbStatus,
		doc.DbLastStatus,
	)

	if err != nil {
		return err
	}

	return nil
}

func (*DOCModel) Update(Id string, fields map[string]interface{}) error {
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
			UPDATE docregisters 
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
