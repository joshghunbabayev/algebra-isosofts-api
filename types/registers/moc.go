package registerTypes

import (
	registerComponentTypes "algebra-isosofts-api/types/registers/components"
	tableComponentTypes "algebra-isosofts-api/types/tableComponents"
)

type MOC struct {
	Id                     string                               `json:"id"`
	CompanyId              string                               `json:"companyId"`
	No                     string                               `json:"no"`
	Issuer                 string                               `json:"issuer"`
	IssuerDate             string                               `json:"issuerDate"`
	ReasonOfChange         string                               `json:"reasonOfChange"`
	Process                tableComponentTypes.DropDownListItem `json:"process"`
	ChangeDescription      string                               `json:"changeDescription"`
	Risks                  string                               `json:"risks"`
	ECM                    string                               `json:"ecm"`
	Approval               int8                                 `json:"approval"`
	InitialRiskSeverity    int8                                 `json:"initialRiskSeverity"`
	InitialRiskLikelihood  int8                                 `json:"initialRiskLikelihood"`
	ACM                    string                               `json:"acm"`
	ResidualRiskSeverity   int8                                 `json:"residualRiskSeverity"`
	ResidualRiskLikelihood int8                                 `json:"residualRiskLikelihood"`
	Comment                string                               `json:"comment"`
	DbStatus               string                               `json:"dbStatus"`
	DbLastStatus           string                               `json:"-"`
	Actions                []registerComponentTypes.Action      `json:"actions"`
}

func (moc MOC) IsEmpty() bool {
	return moc.Id == ""
}
