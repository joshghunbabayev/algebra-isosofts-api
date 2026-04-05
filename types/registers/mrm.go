package registerTypes

import (
	registerComponentTypes "algebra-isosofts-api/types/registers/components"
	tableComponentTypes "algebra-isosofts-api/types/tableComponents"
)

type MRM struct {
	Id           string                               `json:"id"`
	CompanyId    string                               `json:"companyId"`
	No           string                               `json:"no"`
	Topic        tableComponentTypes.DropDownListItem `json:"topic"`
	RISOS        string                               `json:"risos"`
	Process      tableComponentTypes.DropDownListItem `json:"process"`
	Comment      string                               `json:"comment"`
	DbStatus     string                               `json:"dbStatus"`
	DbLastStatus string                               `json:"-"`
	Actions      []registerComponentTypes.Action      `json:"actions"`
}

func (mrm MRM) IsEmpty() bool {
	return mrm.Id == ""
}
