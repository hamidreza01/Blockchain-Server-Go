package blockchain

const _DEFUALT_DIFFICULTY int = 10

type _CONFIG struct {
	DEFUALT_DIFFICULTY int
	rewardValue        int
	reward             struct{ address string }
	genesis            Block
}

var CONFIG _CONFIG = _CONFIG{
	DEFUALT_DIFFICULTY: 10,
	rewardValue:        10,
	reward:             struct{ address string }{address: "/*s*//*s*/DPX Blockchain/*s*//*s*/"},
	genesis: Block{
		Hash:       "DEFAULT-DPX-GENESIS-HASH",
		LastHash:   "DEFAULT-DPX-LAST-HASH",
		Nonce:      0,
		Difficulty: _DEFUALT_DIFFICULTY,
		Timestamp:  0,
		Data: struct {
			Transaction /*s*/ []Transaction `json:"transaction"`
		}{},
	},
}
