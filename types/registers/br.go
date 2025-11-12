package registerTypes

import (
	registerComponentTypes "algebra-isosofts-api/types/registers/components"
	tableComponentTypes "algebra-isosofts-api/types/tableComponents"
)

type Br struct {
	Id                     string                               `json:"id"`
	No                     string                               `json:"no"`
	Swot                   tableComponentTypes.DropDownListItem `json:"swot"`
	Pestle                 tableComponentTypes.DropDownListItem `json:"pestle"`
	InterestedParty        tableComponentTypes.DropDownListItem `json:"interestedParty"`
	RiskOpportunity        string                               `json:"riskOpportunity"`
	Objective              string                               `json:"objective"`
	KPI                    string                               `json:"kpi"`
	Process                tableComponentTypes.DropDownListItem `json:"process"`
	ERMEOA                 string                               `json:"ermeoa"`
	InitialRiskSeverity    int8                                 `json:"initialRiskSeverity"`
	InitialRiskLikelyhood  int8                                 `json:"initialRiskLikelyhood"`
	ResidualRiskSeverity   int8                                 `json:"residualRiskSeverity"`
	ResidualRiskLikelyhood int8                                 `json:"residualRiskLikelyhood"`
	DbStatus               string                               `json:"dbStatus"`
	DbLastStatus           string                               `json:"-"`
	Actions                []registerComponentTypes.Action      `json:"actions"`
}

func (br Br) IsEmpty() bool {
	return br.Id == ""
}
