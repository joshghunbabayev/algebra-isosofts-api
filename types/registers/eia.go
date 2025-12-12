package registerTypes

import (
	registerComponentTypes "algebra-isosofts-api/types/registers/components"
	tableComponentTypes "algebra-isosofts-api/types/tableComponents"
)

type EIA struct {
	Id                string                               `json:"id"`
	No                string                               `json:"no"`
	Process           tableComponentTypes.DropDownListItem `json:"process"`
	Aspect            tableComponentTypes.DropDownListItem `json:"aspect"`
	Impact            string                               `json:"impact"`
	AffectedReceptors tableComponentTypes.DropDownListItem `json:"affectedReceptors"`
	ExistingControls  string                               `json:"existingControls"`
	IDOSProbability   int8                                 `json:"idosProbability"`
	IDOSSeverity      int8                                 `json:"idosSeverity"`
	IDOSDuration      int8                                 `json:"idosDuration"`
	IDOSScale         int8                                 `json:"idosScale"`
	RDOSProbability   int8                                 `json:"rdosProbability"`
	RDOSSeverity      int8                                 `json:"rdosSeverity"`
	RDOSDuration      int8                                 `json:"rdosDuration"`
	RDOSScale         int8                                 `json:"rdosScale"`
	DbStatus          string                               `json:"dbStatus"`
	DbLastStatus      string                               `json:"-"`
	Actions           []registerComponentTypes.Action      `json:"actions"`
}

func (eia EIA) IsEmpty() bool {
	return eia.Id == ""
}
