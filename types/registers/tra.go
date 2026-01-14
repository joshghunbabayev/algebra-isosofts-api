package registerTypes

import (
	registerComponentTypes "algebra-isosofts-api/types/registers/components"
)

type TRA struct {
	Id               string                          `json:"id"`
	No               string                          `json:"no"`
	EmployeeName     string                          `json:"employeeName"`
	Position         string                          `json:"position"`
	CLName           string                          `json:"clname"`
	TCLID            string                          `json:"nvcd"`
	CLNumber         string                          `json:"clnumber"`
	NCD              string                          `json:"ncd"`
	CompetencyStatus int8                            `json:"competencyStatus"`
	Effectiveness    string                          `json:"effectiveness"`
	DbStatus         string                          `json:"dbStatus"`
	DbLastStatus     string                          `json:"-"`
	Actions          []registerComponentTypes.Action `json:"actions"`
}

func (tra TRA) IsEmpty() bool {
	return tra.Id == ""
}
