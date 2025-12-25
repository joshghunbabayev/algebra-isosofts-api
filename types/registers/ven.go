package registerTypes

import (
	registerComponentTypes "algebra-isosofts-api/types/registers/components"
	tableComponentTypes "algebra-isosofts-api/types/tableComponents"
)

type VEN struct {
	Id               string                               `json:"id"`
	No               string                               `json:"no"`
	Name             string                               `json:"name"`
	RegNumber        string                               `json:"regNumber"`
	Scope1           tableComponentTypes.DropDownListItem `json:"scope1"`
	Scope2           tableComponentTypes.DropDownListItem `json:"scope2"`
	Scope3           tableComponentTypes.DropDownListItem `json:"scope3"`
	RegistrationDate string                               `json:"registrationDate"`
	ReviewDate       string                               `json:"reviewDate"`
	Approved         int8                                 `json:"approved"`
	QGS              int8                                 `json:"qgs"`
	Communication    int8                                 `json:"communication"`
	OTD              int8                                 `json:"otd"`
	Documentation    int8                                 `json:"documentation"`
	HS               int8                                 `json:"hs"`
	Environment      int8                                 `json:"environment"`
	DbStatus         string                               `json:"dbStatus"`
	DbLastStatus     string                               `json:"-"`
	Actions          []registerComponentTypes.Action      `json:"actions"`
}

func (ven VEN) IsEmpty() bool {
	return ven.Id == ""
}
