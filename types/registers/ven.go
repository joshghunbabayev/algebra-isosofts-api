package registerTypes

import (
	registerComponentTypes "algebra-isosofts-api/types/registers/components"
	tableComponentTypes "algebra-isosofts-api/types/tableComponents"
)

type VEN struct {
	Id               string                               `json:"id"`
	CompanyId        string                               `json:"companyId"`
	No               string                               `json:"no"`
	Name             string                               `json:"name"`
	RegNumber        string                               `json:"regNumber"`
	Scope1           tableComponentTypes.DropDownListItem `json:"scope1"`
	Scope2           tableComponentTypes.DropDownListItem `json:"scope2"`
	Scope3           tableComponentTypes.DropDownListItem `json:"scope3"`
	RegistrationDate string                               `json:"registrationDate"`
	NRD              string                               `json:"nrd"`
	EvaluationDone   int8                                 `json:"evaluationDone"`
	QGS              float32                              `json:"qgs"`
	Communication    float32                              `json:"communication"`
	OTD              float32                              `json:"otd"`
	Documentation    float32                              `json:"documentation"`
	HS               float32                              `json:"hs"`
	Environment      float32                              `json:"environment"`
	Comment          string                               `json:"comment"`
	DbStatus         string                               `json:"dbStatus"`
	DbLastStatus     string                               `json:"-"`
	Actions          []registerComponentTypes.Action      `json:"actions"`
}

func (ven VEN) IsEmpty() bool {
	return ven.Id == ""
}
