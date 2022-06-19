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
	reward:             struct{ address string }{address: "**DPX Blockchain**"},
	genesis: Block{
		hash:       "DEFAULT-DPX-GENESIS-HASH",
		lastHash:   "DEFAULT-DPX-LAST-HASH",
		nonce:      0,
		difficulty: _DEFUALT_DIFFICULTY,
		timestamp:  0,
		data:       struct{ Transaction []Transaction }{},
	},
}
