package registerTypes

import (
	registerComponentTypes "algebra-isosofts-api/types/registers/components"
	tableComponentTypes "algebra-isosofts-api/types/tableComponents"
)

type MRM struct {
	Id           string                               `json:"id"`
	No           string                               `json:"no"`
	RISOS        tableComponentTypes.DropDownListItem `json:"risos"`
	Topic        tableComponentTypes.DropDownListItem `json:"topic"`
	Process      tableComponentTypes.DropDownListItem `json:"process"`
	DbStatus     string                               `json:"dbStatus"`
	DbLastStatus string                               `json:"-"`
	Actions      []registerComponentTypes.Action      `json:"actions"`
}

func (mrm MRM) IsEmpty() bool {
	return mrm.Id == ""
}
