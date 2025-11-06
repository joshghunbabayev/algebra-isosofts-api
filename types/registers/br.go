package registerTypes

import tableComponentTypes "algebra-isosofts-api/types/tableComponents"

type Br struct {
	Id                     string `json:"id"`
	No                     string `json:"no"`
	Swot                   tableComponentTypes.DropDownListItem
	Pestle                 tableComponentTypes.DropDownListItem
	InterestedParty        tableComponentTypes.DropDownListItem
	RiskOpportunity        string `json:"riskOpportunity"`
	Objective              string `json:"objective"`
	KPI                    string `json:"kpi"`
	Process                tableComponentTypes.DropDownListItem
	ERMEOA                 string `json:"ermeoa"`
	InitialRiskSeverity    int8   `json:"initialRiskSeverity"`
	InitialRiskLikelyhood  int8   `json:"initialRiskLikelyhood"`
	ResidualRiskSeverity   int8   `json:"residualRiskSeverity"`
	ResidualRiskLikelyhood int8   `json:"residualRiskLikelyhood"`
}

func (br Br) IsEmpty() bool {
	return br.Id == ""
}
