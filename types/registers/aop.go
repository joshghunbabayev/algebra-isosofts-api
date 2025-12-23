package registerTypes

import (
	registerComponentTypes "algebra-isosofts-api/types/registerComponents"
	tableComponentTypes "algebra-isosofts-api/types/tableComponents"
)

type AOP struct {
	Id                  string                               `json:"id"`
	No                  string                               `json:"no"`
	ActivityDescription tableComponentTypes.DropDownListItem `json:"activityDescription"`
	AuditorInspector    string                               `json:"auditorInspector"`
	AuditeeInspectee    string                               `json:"auditeeInspectee"`
	ReviewedPremises    tableComponentTypes.DropDownListItem `json:"reviewedPremises"`
	ReviewedProcess     tableComponentTypes.DropDownListItem `json:"reviewedProcess"`
	RTIC                string                               `json:"rtic"`
	Frequency           string                               `json:"frequency"`
	AuditDate           string                               `json:"AuditDate"`
	InspectionFrequency tableComponentTypes.DropDownListItem `json:"inspectionFrequency"`
	NextAuditDate       string                               `json:"nextAuditDate"`
	AuditStatus         string                               `json:"auditStatus"`
	DbStatus            string                               `json:"dbStatus"`
	DbLastStatus        string                               `json:"-"`
	Actions             []registerComponentTypes.Action      `json:"actions"`
}

func (aop AOP) IsEmpty() bool {
	return aop.Id == ""
}
