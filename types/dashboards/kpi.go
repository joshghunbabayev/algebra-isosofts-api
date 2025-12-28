package dashboardTypes

type KPI struct {
	TotalActions int64 `json:"totalActions"`

	Status struct {
		Average float64          `json:"average"`
		Counts  map[string]int64 `json:"counts"`
	} `json:"status"`

	VerificationStatus map[string]int64 `json:"verificationStatus"`
	Confirmation       map[string]int64 `json:"confirmation"`

	MonthlyProgress struct {
		January   float64 `json:"january"`
		February  float64 `json:"february"`
		March     float64 `json:"march"`
		April     float64 `json:"april"`
		May       float64 `json:"may"`
		June      float64 `json:"june"`
		July      float64 `json:"july"`
		August    float64 `json:"august"`
		September float64 `json:"september"`
		October   float64 `json:"october"`
		November  float64 `json:"november"`
		December  float64 `json:"december"`
	} `json:"monthlyProgress"`
}
