package registerModels

import (
	"algebra-isosofts-api/database"
	"fmt"
)

type CommonModel struct {
}

func (*CommonModel) GetRegNo(Id string, Type string) (string, error) {
	db := database.GetDatabase()

	var no string
	tableName := Type + "registers"

	query := fmt.Sprintf("SELECT no FROM %s WHERE id = ?", tableName)
	err := db.QueryRow(query, Id).Scan(&no)

	if err != nil {
		return "", err
	}

	return no, nil
}
