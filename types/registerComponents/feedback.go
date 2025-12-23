package registerComponentTypes

import (
	tableComponentTypes "algebra-isosofts-api/types/tableComponents"
)

type Feedback struct {
	Id                 string                               `json:"id"`
	RegisterId         string                               `json:"registerId"`
	Scope              string                               `json:"title"`
	Vendor             string                               `json:"vendor"`
	TypeOfFinding      tableComponentTypes.DropDownListItem `json:"typeOfFinding"`
	Resources          int64                                `json:"resources"`
	Currency           string                               `json:"currency"`
	RelativeFunction   tableComponentTypes.DropDownListItem `json:"relativeFunction"`
	Responsible        tableComponentTypes.DropDownListItem `json:"responsible"`
	Deadline           string                               `json:"deadline"`
	Confirmation       tableComponentTypes.DropDownListItem `json:"confirmation"`
	Status             tableComponentTypes.DropDownListItem `json:"status"`
	CompletionDate     string                               `json:"completionDate"`
	VerificationStatus tableComponentTypes.DropDownListItem `json:"verificationStatus"`
	Comment            string                               `json:"comment"`
	January            tableComponentTypes.DropDownListItem `json:"january"`
	February           tableComponentTypes.DropDownListItem `json:"february"`
	March              tableComponentTypes.DropDownListItem `json:"march"`
	April              tableComponentTypes.DropDownListItem `json:"april"`
	May                tableComponentTypes.DropDownListItem `json:"may"`
	June               tableComponentTypes.DropDownListItem `json:"june"`
	July               tableComponentTypes.DropDownListItem `json:"july"`
	August             tableComponentTypes.DropDownListItem `json:"august"`
	September          tableComponentTypes.DropDownListItem `json:"september"`
	October            tableComponentTypes.DropDownListItem `json:"october"`
	November           tableComponentTypes.DropDownListItem `json:"november"`
	December           tableComponentTypes.DropDownListItem `json:"december"`
	DbStatus           string                               `json:"dbStatus"`
	DbLastStatus       string                               `json:"-"`
}

func (feedback Feedback) IsEmpty() bool {
	return feedback.Id == ""
}
