package registerTypes

import (
	registerComponentTypes "algebra-isosofts-api/types/registers/components"
	tableComponentTypes "algebra-isosofts-api/types/tableComponents"
)

type EA struct {
	Id                       string                               `json:"id"`
	CompanyId                string                               `json:"companyId"`
	No                       string                               `json:"no"`
	EmployeeName             string                               `json:"employeeName"`
	Position                 string                               `json:"position"`
	LineManager              string                               `json:"lineManager"`
	ESD                      string                               `json:"esd"`
	AppraisalDate            string                               `json:"appraisalDate"`
	NextAppraisalDate        string                               `json:"nextAppraisalDate"`
	AppraisalType            tableComponentTypes.DropDownListItem `json:"appraisalType"`
	TCA                      string                               `json:"tca"`
	SkillsAppraisal          string                               `json:"skillsAppraisal"`
	JobQuality               int8                                 `json:"jobQuality"`
	LeadershipSkills         int8                                 `json:"leadershipSkills"`
	ManagementSkills         int8                                 `json:"managementSkills"`
	BehavioralSkills         int8                                 `json:"behavioralSkills"`
	EffectivenessOfTrainings int8                                 `json:"effectivenessOfTrainings"`
	EVS                      int8                                 `json:"evs"`
	DbStatus                 string                               `json:"dbStatus"`
	DbLastStatus             string                               `json:"-"`
	Actions                  []registerComponentTypes.Action      `json:"actions"`
}

func (ea EA) IsEmpty() bool {
	return ea.Id == ""
}
