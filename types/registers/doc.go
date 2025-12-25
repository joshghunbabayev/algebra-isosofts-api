package registerTypes

import (
	registerComponentTypes "algebra-isosofts-api/types/registers/components"
	tableComponentTypes "algebra-isosofts-api/types/tableComponents"
)

type DOC struct {
	Id                string                               `json:"id"`
	No                string                               `json:"no"`
	Name              string                               `json:"name"`
	Origin            tableComponentTypes.DropDownListItem `json:"origin"`
	Number            string                               `json:"number"`
	DepntFunctionName tableComponentTypes.DropDownListItem `json:"depntFunctionName"`
	Type              tableComponentTypes.DropDownListItem `json:"type"`
	SerialNumber      string                               `json:"serialNumber"`
	RevNumber         string                               `json:"revNumber"`
	Issuer            string                               `json:"issuer"`
	Approver          string                               `json:"approver"`
	IssueDate         string                               `json:"issueDate"`
	NextReviewDate    string                               `json:"nextReviewDate"`
	Actual            int8                                 `json:"actual"`
	DbStatus          string                               `json:"dbStatus"`
	DbLastStatus      string                               `json:"-"`
	Actions           []registerComponentTypes.Action      `json:"actions"`
}

func (doc DOC) IsEmpty() bool {
	return doc.Id == ""
}
