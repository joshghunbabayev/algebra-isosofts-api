package registerTypes

import (
	registerComponentTypes "algebra-isosofts-api/types/registers/components"
	tableComponentTypes "algebra-isosofts-api/types/tableComponents"
)

type HSR struct {
	Id                     string                               `json:"id"`
	No                     string                               `json:"no"`
	Process                tableComponentTypes.DropDownListItem `json:"process"`
	Hazard                 tableComponentTypes.DropDownListItem `json:"hazard"`
	Risk                   tableComponentTypes.DropDownListItem `json:"risk"`
	AffectedPositions      tableComponentTypes.DropDownListItem `json:"affectedPositions"`
	ERMA                   string                               `json:"erma"`
	InitialRiskSeverity    int8                                 `json:"initialRiskSeverity"`
	InitialRiskLikelyhood  int8                                 `json:"initialRiskLikelyhood"`
	ResidualRiskSeverity   int8                                 `json:"residualRiskSeverity"`
	ResidualRiskLikelyhood int8                                 `json:"residualRiskLikelyhood"`
	DbStatus               string                               `json:"dbStatus"`
	DbLastStatus           string                               `json:"-"`
	Actions                []registerComponentTypes.Action      `json:"actions"`
}

func (hsr HSR) IsEmpty() bool {
	return hsr.Id == ""
}
