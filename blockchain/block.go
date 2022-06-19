package blockchain

type Block struct {
	hash       string
	lastHash   string
	data       struct{ Transaction []Transaction }
	difficulty int
	timestamp  int
	nonce      int
}
