package registerTypes

import (
	registerComponentTypes "algebra-isosofts-api/types/registers/components"
	tableComponentTypes "algebra-isosofts-api/types/tableComponents"
)

type EI struct {
	Id                  string                               `json:"id"`
	No                  string                               `json:"no"`
	Name                string                               `json:"name"`
	SerialNumber        string                               `json:"serialNumber"`
	CertificateNo       string                               `json:"certificateNo"`
	InspectionFrequency tableComponentTypes.DropDownListItem `json:"inspectionFrequency"`
	ICD                 string                               `json:"icd"`
	NVCD                string                               `json:"nvcd"`
	SafeToUse           int8                                 `json:"safeToUse"`
	DbStatus            string                               `json:"dbStatus"`
	DbLastStatus        string                               `json:"-"`
	Actions             []registerComponentTypes.Action      `json:"actions"`
}

func (ei EI) IsEmpty() bool {
	return ei.Id == ""
}
