package registerTypes

import (
	registerComponentTypes "algebra-isosofts-api/types/registers/components"
	tableComponentTypes "algebra-isosofts-api/types/tableComponents"
)

type LEG struct {
	Id                     string                               `json:"id"`
	No                     string                               `json:"no"`
	Process                tableComponentTypes.DropDownListItem `json:"process"`
	Legislation            string                               `json:"legislation"`
	Section                string                               `json:"section"`
	Requirement            string                               `json:"requirement"`
	RiskOfViolation        string                               `json:"riskOfViolation"`
	InitialRiskSeverity    int8                                 `json:"initialRiskSeverity"`
	InitialRiskLikelyhood  int8                                 `json:"initialRiskLikelyhood"`
	ResidualRiskSeverity   int8                                 `json:"residualRiskSeverity"`
	ResidualRiskLikelyhood int8                                 `json:"residualRiskLikelyhood"`
	DbStatus               string                               `json:"dbStatus"`
	DbLastStatus           string                               `json:"-"`
	Actions                []registerComponentTypes.Action      `json:"actions"`
}

func (leg LEG) IsEmpty() bool {
	return leg.Id == ""
}
