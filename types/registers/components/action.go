package registerComponentsTypes

type Action struct {
	RegisterId string `json:"registerId"`
}

func (action Action) IsEmpty() bool {
	return action.RegisterId == ""
}
