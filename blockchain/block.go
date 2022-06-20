package blockchain

type Block struct {
	Hash     string `json:"hash"`
	LastHash string `json:"lastHash"`
	Data     struct {
		Transaction []Transaction `json:"transaction"`
	} `json:"data"`
	Difficulty int `json:"difficulty"`
	Timestamp  int `json:"timestamp"`
	Nonce      int `json:"nonce"`
}
