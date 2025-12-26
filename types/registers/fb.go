package registerTypes

import (
	registerComponentTypes "algebra-isosofts-api/types/registers/components"
	tableComponentTypes "algebra-isosofts-api/types/tableComponents"
)

type FB struct {
	Id                string                                  `json:"id"`
	No                string                                  `json:"no"`
	JobNumber         string                                  `json:"jobNumber"`
	JobStartDate      string                                  `json:"jobStartDate"`
	JobCompletionDate string                                  `json:"jobCompletionDate"`
	Scope             tableComponentTypes.DropDownListItem    `json:"scope"`
	CustomerId        string                                  `json:"customerId"`
	TypeOfFinding     tableComponentTypes.DropDownListItem    `json:"typeOfFinding"`
	QGS               int8                                    `json:"qgs"`
	Communication     int8                                    `json:"communication"`
	OTD               int8                                    `json:"otd"`
	Documentation     int8                                    `json:"documentation"`
	HS                int8                                    `json:"hs"`
	Environment       int8                                    `json:"environment"`
	DbStatus          string                                  `json:"dbStatus"`
	DbLastStatus      string                                  `json:"-"`
	VendorFeedbacks   []registerComponentTypes.VendorFeedback `json:"vendorFeedbacks"`
}

func (fb FB) IsEmpty() bool {
	return fb.Id == ""
}
