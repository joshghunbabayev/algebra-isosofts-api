package registerTypes

import (
	registerComponentTypes "algebra-isosofts-api/types/registers/components"
	tableComponentTypes "algebra-isosofts-api/types/tableComponents"
)

type FIN struct {
	Id                string                               `json:"id"`
	CompanyId         string                               `json:"companyId"`
	No                string                               `json:"no"`
	Issuer            string                               `json:"issuer"`
	FindingDate       string                               `json:"findingDate"`
	JobNumber         string                               `json:"jobNumber"`
	Process           tableComponentTypes.DropDownListItem `json:"process"`
	CategoryOfFinding tableComponentTypes.DropDownListItem `json:"categoryOfFinding"`
	TypeOfFinding     tableComponentTypes.DropDownListItem `json:"typeOfFinding"`
	SourceOfFinding   tableComponentTypes.DropDownListItem `json:"sourceOfFinding"`
	CustomerId        string                               `json:"customerId"`
	VendorId          string                               `json:"vendorId"`
	Description       string                               `json:"description"`
	ContainmentAction string                               `json:"containmentAction"`
	RootCauses        string                               `json:"rootCauses"`
	FindingStatus     tableComponentTypes.DropDownListItem `json:"findingStatus"`
	Comment           string                               `json:"comment"`
	DbStatus          string                               `json:"dbStatus"`
	DbLastStatus      string                               `json:"-"`
	Actions           []registerComponentTypes.Action      `json:"actions"`
}

func (fin FIN) IsEmpty() bool {
	return fin.Id == ""
}
