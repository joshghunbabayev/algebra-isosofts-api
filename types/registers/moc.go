package registerTypes

import (
	registerComponentTypes "algebra-isosofts-api/types/registers/components"
	tableComponentTypes "algebra-isosofts-api/types/tableComponents"
)

type MOC struct {
	Id                     string                               `json:"id"`
	No                     string                               `json:"no"`
	Issuer                 string                               `json:"issuer"`
	ReasonOfChange         string                               `json:"reasonOfChange"`
	Process                tableComponentTypes.DropDownListItem `json:"process"`
	Action                 string                               `json:"action"`
	Risks                  string                               `json:"risks"`
	InitialRiskSeverity    int8                                 `json:"initialRiskSeverity"`
	InitialRiskLikelyhood  int8                                 `json:"initialRiskLikelyhood"`
	ResidualRiskSeverity   int8                                 `json:"residualRiskSeverity"`
	ResidualRiskLikelyhood int8                                 `json:"residualRiskLikelyhood"`
	DbStatus               string                               `json:"dbStatus"`
	DbLastStatus           string                               `json:"-"`
	Actions                []registerComponentTypes.Action      `json:"actions"`
}

func (moc MOC) IsEmpty() bool {
	return moc.Id == ""
}
