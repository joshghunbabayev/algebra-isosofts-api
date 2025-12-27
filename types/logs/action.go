package logTypes

import (
	registerComponentTypes "algebra-isosofts-api/types/registers/components"
)

type Action struct {
	Id      string                          `json:"id"`
	No      string                          `json:"no"`
	Actions []registerComponentTypes.Action `json:"actions"`
}

func (action Action) IsEmpty() bool {
	return action.Id == ""
}
