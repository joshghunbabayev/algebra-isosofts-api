package registerTypes

import (
	registerComponentTypes "algebra-isosofts-api/types/registers/components"
	tableComponentTypes "algebra-isosofts-api/types/tableComponents"
)

type CUS struct {
	Id               string                               `json:"id"`
	CompanyId        string                               `json:"companyId"`
	No               string                               `json:"no"`
	Name             string                               `json:"name"`
	RegNumber        string                               `json:"regNumber"`
	Scope1           tableComponentTypes.DropDownListItem `json:"scope1"`
	Scope2           tableComponentTypes.DropDownListItem `json:"scope2"`
	Scope3           tableComponentTypes.DropDownListItem `json:"scope3"`
	RegistrationDate string                               `json:"registrationDate"`
	ReviewDate       string                               `json:"reviewDate"`
	EvaluationDone   int8                                 `json:"evaluationDone"`
	QGS              int8                                 `json:"qgs"`
	Communication    int8                                 `json:"communication"`
	OTD              int8                                 `json:"otd"`
	Documentation    int8                                 `json:"documentation"`
	HS               int8                                 `json:"hs"`
	Environment      int8                                 `json:"environment"`
	Comment          string                               `json:"comment"`
	DbStatus         string                               `json:"dbStatus"`
	DbLastStatus     string                               `json:"-"`
	Actions          []registerComponentTypes.Action      `json:"actions"`
}

func (cus CUS) IsEmpty() bool {
	return cus.Id == ""
}
