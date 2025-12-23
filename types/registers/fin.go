package registerTypes

import (
	registerComponentTypes "algebra-isosofts-api/types/registerComponents"
	tableComponentTypes "algebra-isosofts-api/types/tableComponents"
)

type FIN struct {
	Id                string                               `json:"id"`
	No                string                               `json:"no"`
	Issuer            string                               `json:"issuer"`
	Process           tableComponentTypes.DropDownListItem `json:"process"`
	CategoryOfFinding tableComponentTypes.DropDownListItem `json:"categoryOfFinding"`
	TypeOfFinding     tableComponentTypes.DropDownListItem `json:"typeOfFinding"`
	SourceOfFinding   tableComponentTypes.DropDownListItem `json:"sourceOfFinding"`
	Customer          string                               `json:"customer"`
	Vendor            string                               `json:"vendor"`
	Description       string                               `json:"description"`
	ContainmentAction string                               `json:"containmentAction"`
	RootCauses        string                               `json:"rootCauses"`
	DbStatus          string                               `json:"dbStatus"`
	DbLastStatus      string                               `json:"-"`
	Actions           []registerComponentTypes.Action      `json:"actions"`
}

func (fin FIN) IsEmpty() bool {
	return fin.Id == ""
}
