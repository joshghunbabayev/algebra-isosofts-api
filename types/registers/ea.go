package registerTypes

import registerComponentTypes "algebra-isosofts-api/types/registers/components"

type EA struct {
	Id              string                          `json:"id"`
	No              string                          `json:"no"`
	EmployeeName    string                          `json:"employeeName"`
	Position        string                          `json:"position"`
	LineManager     string                          `json:"lineManager"`
	ESD             string                          `json:"esd"`
	AppraisalDate   string                          `json:"appraisalDate"`
	AppraisalType   string                          `json:"appraisalType"`
	TCA             string                          `json:"tca"`
	SkillsAppraisal string                          `json:"skillsAppraisal"`
	DbStatus        string                          `json:"dbStatus"`
	DbLastStatus    string                          `json:"-"`
	Actions         []registerComponentTypes.Action `json:"actions"`
}

func (ea EA) IsEmpty() bool {
	return ea.Id == ""
}
