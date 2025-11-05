package tableComponentTypes

type DropDownListItem struct {
	Id         string `json:"id"`
	Type       string `json:"type"`
	Value      string `json:"value"`
	ShortValue string `json:"shortValue"`
}

func (dropDownListItem DropDownListItem) IsEmpty() bool {
	return dropDownListItem.Id == ""
}

func GroupDropDownListItems(dropDownListItems []DropDownListItem) map[string][]DropDownListItem {
	groupedDropDownListItem := make(map[string][]DropDownListItem)

	for _, dropDownListItem := range dropDownListItems {
		groupedDropDownListItem[dropDownListItem.Type] = append(groupedDropDownListItem[dropDownListItem.Type], dropDownListItem)
	}

	return groupedDropDownListItem
}
