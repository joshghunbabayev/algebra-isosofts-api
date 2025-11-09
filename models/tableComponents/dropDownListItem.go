package tableComponentModels

import (
	"algebra-isosofts-api/database"
	"algebra-isosofts-api/modules"
	tableComponentTypes "algebra-isosofts-api/types/tableComponents"
)

type DropDownListItemModel struct {
}

func (*DropDownListItemModel) GenerateUniqueId() string {
	Id := modules.GenerateRandomString(30)

	var dropDownListItemModel DropDownListItemModel

	dropDownListItem, _ := dropDownListItemModel.GetById(Id)

	if dropDownListItem.IsEmpty() {
		return Id
	} else {
		return dropDownListItemModel.GenerateUniqueId()
	}
}

func (*DropDownListItemModel) GetById(Id string) (tableComponentTypes.DropDownListItem, error) {
	db := database.GetDatabase()
	row := db.QueryRow(`
			SELECT * 
			FROM dropdownlistitems 
			WHERE id = ?
		`,
		Id,
	)

	var dropDownListItem tableComponentTypes.DropDownListItem
	err := row.Scan(
		&dropDownListItem.Id,
		&dropDownListItem.Type,
		&dropDownListItem.Value,
		&dropDownListItem.ShortValue,
	)

	return dropDownListItem, err
}

func (*DropDownListItemModel) GetAll() ([]tableComponentTypes.DropDownListItem, error) {
	db := database.GetDatabase()
	rows, err := db.Query(`
			SELECT * 
			FROM dropdownlistitems
		`,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dropDownListItems []tableComponentTypes.DropDownListItem

	for rows.Next() {
		var dropDownListItem tableComponentTypes.DropDownListItem
		rows.Scan(
			&dropDownListItem.Id,
			&dropDownListItem.Type,
			&dropDownListItem.Value,
			&dropDownListItem.ShortValue,
		)
		dropDownListItems = append(dropDownListItems, dropDownListItem)
	}

	return dropDownListItems, nil
}

func (*DropDownListItemModel) Create(dropDownListItem tableComponentTypes.DropDownListItem) error {

	db := database.GetDatabase()

	_, err := db.Exec(`
			INSERT INTO dropdownlistitems (
				"id", 
				"type", 
				"value", 
				"shortValue"
			) 
			VALUES (?, ?, ?, ?)
		`,
		dropDownListItem.Id,
		dropDownListItem.Type,
		dropDownListItem.Value,
		dropDownListItem.ShortValue,
	)

	if err != nil {
		return err
	}

	return nil
}

func (*DropDownListItemModel) DuplicateDefaults() error {
	db := database.GetDatabase()
	rows, err := db.Query(`
			SELECT * 
			FROM defaultdropdownlistitems
		`,
	)

	if err != nil {
		return err
	}
	defer rows.Close()

	var defaultDropDownListItems []tableComponentTypes.DropDownListItem
	var dropDownListItemModel DropDownListItemModel

	for rows.Next() {
		var defaultDropDownListItem tableComponentTypes.DropDownListItem
		rows.Scan(
			&defaultDropDownListItem.Type,
			&defaultDropDownListItem.Value,
			&defaultDropDownListItem.ShortValue,
		)
		defaultDropDownListItems = append(defaultDropDownListItems, defaultDropDownListItem)
	}

	for _, defaultDropDownListItem := range defaultDropDownListItems {
		defaultDropDownListItem.Id = dropDownListItemModel.GenerateUniqueId()
		dropDownListItemModel.Create(defaultDropDownListItem)
	}

	return nil
}
