package blockchain

type Transaction struct {
	Id        string         `json:"id"`
	OutputMap map[string]int `json:"outputMap"`
	InputMap  struct {
		Address string `json:"address"`
	} `json:"inputMap"`
}
