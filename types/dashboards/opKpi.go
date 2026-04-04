package dashboardTypes

import tableComponentTypes "algebra-isosofts-api/types/tableComponents"

type OpKPI struct {
	Id           string                               `json:"id"`
	CompanyId    string                               `json:"companyId"`
	No           string                               `json:"no"`
	Title        string                               `json:"title"`
	Function     tableComponentTypes.DropDownListItem `json:"function"`
	LYKPI        int64                                `json:"lykpi"`
	ActualKPI    int64                                `json:"actualKPI"`
	AnnualTarget int64                                `json:"annualTarget"`
	January      int64                                `json:"january"`
	February     int64                                `json:"february"`
	March        int64                                `json:"march"`
	April        int64                                `json:"april"`
	May          int64                                `json:"may"`
	June         int64                                `json:"june"`
	July         int64                                `json:"july"`
	August       int64                                `json:"august"`
	September    int64                                `json:"september"`
	October      int64                                `json:"october"`
	November     int64                                `json:"november"`
	December     int64                                `json:"december"`
	DbStatus     string                               `json:"dbStatus"`
	DbLastStatus string                               `json:"-"`
}

func (opKPI OpKPI) IsEmpty() bool {
	return opKPI.Id == ""
}
