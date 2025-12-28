package dashboardTypes

import (
	registerComponentTypes "algebra-isosofts-api/types/registers/components"
)

type ActionLog struct {
	Id      string                          `json:"id"`
	No      string                          `json:"no"`
	Actions []registerComponentTypes.Action `json:"actions"`
}

func (actionLog ActionLog) IsEmpty() bool {
	return actionLog.Id == ""
}
