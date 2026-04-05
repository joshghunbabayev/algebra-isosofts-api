package registerTypes

import (
	registerComponentTypes "algebra-isosofts-api/types/registers/components"
	tableComponentTypes "algebra-isosofts-api/types/tableComponents"
)

type HSR struct {
	Id                     string                               `json:"id"`
	CompanyId              string                               `json:"companyId"`
	No                     string                               `json:"no"`
	Process                tableComponentTypes.DropDownListItem `json:"process"`
	Hazard                 tableComponentTypes.DropDownListItem `json:"hazard"`
	Risk                   tableComponentTypes.DropDownListItem `json:"risk"`
	AffectedPositions      tableComponentTypes.DropDownListItem `json:"affectedPositions"`
	ECM                    string                               `json:"ecm"`
	InitialRiskSeverity    int8                                 `json:"initialRiskSeverity"`
	InitialRiskLikelihood  int8                                 `json:"initialRiskLikelihood"`
	ACM                    string                               `json:"acm"`
	ResidualRiskSeverity   int8                                 `json:"residualRiskSeverity"`
	ResidualRiskLikelihood int8                                 `json:"residualRiskLikelihood"`
	Comment                string                               `json:"comment"`
	DbStatus               string                               `json:"dbStatus"`
	DbLastStatus           string                               `json:"-"`
	Actions                []registerComponentTypes.Action      `json:"actions"`
}

func (hsr HSR) IsEmpty() bool {
	return hsr.Id == ""
}
