package registerTypes

type Br struct {
	Id     string `json:"id"`
	Swot   string `json:"swot"`
	Pestle string `json:"pestle"`
}

func (br Br) IsEmpty() bool {
	return br.Id == ""
}
