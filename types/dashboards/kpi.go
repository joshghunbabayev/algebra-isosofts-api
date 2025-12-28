package dashboardTypes

import (
	registerComponentTypes "algebra-isosofts-api/types/registers/components"
)

type KPI struct {
	Id      string                          `json:"id"`
	No      string                          `json:"no"`
	Actions []registerComponentTypes.Action `json:"actions"`
}

func (kpi KPI) IsEmpty() bool {
	return kpi.Id == ""
}
