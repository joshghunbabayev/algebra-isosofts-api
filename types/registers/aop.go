package registerTypes

import (
	registerComponentTypes "algebra-isosofts-api/types/registers/components"
	tableComponentTypes "algebra-isosofts-api/types/tableComponents"
)

type AOP struct {
	Id                  string                               `json:"id"`
	CompanyId           string                               `json:"companyId"`
	No                  string                               `json:"no"`
	ActivityDescription tableComponentTypes.DropDownListItem `json:"activityDescription"`
	AuditorInspector    string                               `json:"auditorInspector"`
	AuditeeInspectee    string                               `json:"auditeeInspectee"`
	ReviewedPremises    tableComponentTypes.DropDownListItem `json:"reviewedPremises"`
	ReviewedProcess     tableComponentTypes.DropDownListItem `json:"reviewedProcess"`
	RTIC                string                               `json:"rtic"`
	Frequency           string                               `json:"frequency"`
	AOADate             string                               `json:"aoaDate"`
	InspectionFrequency tableComponentTypes.DropDownListItem `json:"inspectionFrequency"`
	NextAoaDate         string                               `json:"nextAoaDate"`
	AOAStatus           string                               `json:"aoaStatus"`
	Comment             string                               `json:"comment"`
	DbStatus            string                               `json:"dbStatus"`
	DbLastStatus        string                               `json:"-"`
	Actions             []registerComponentTypes.Action      `json:"actions"`
}

func (aop AOP) IsEmpty() bool {
	return aop.Id == ""
}
