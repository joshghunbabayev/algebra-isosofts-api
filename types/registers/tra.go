package registerTypes

import (
	registerComponentTypes "algebra-isosofts-api/types/registers/components"
	tableComponentTypes "algebra-isosofts-api/types/tableComponents"
)

type TRA struct {
	Id                string                               `json:"id"`
	CompanyId         string                               `json:"companyId"`
	No                string                               `json:"no"`
	EmployeeName      string                               `json:"employeeName"`
	Position          string                               `json:"position"`
	TCLN              string                               `json:"tcln"`
	TCLID             string                               `json:"nvcd"`
	CLNumber          string                               `json:"clnumber"`
	TrainingFrequency tableComponentTypes.DropDownListItem `json:"trainingFrequency"`
	NCD               string                               `json:"ncd"`
	ValidityStatus    int8                                 `json:"validityStatus"`
	Effectiveness     int8                                 `json:"effectiveness"`
	DbStatus          string                               `json:"dbStatus"`
	DbLastStatus      string                               `json:"-"`
	Actions           []registerComponentTypes.Action      `json:"actions"`
}

func (tra TRA) IsEmpty() bool {
	return tra.Id == ""
}
