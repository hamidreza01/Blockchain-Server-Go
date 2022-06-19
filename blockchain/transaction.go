package blockchain

type Transaction struct {
	id        string
	outputMap map[string]int
	inputMap  struct {
		address string
		// signature elliptic.CurveParams
	}
}

func (_ Transaction) isValid(transaction Transaction) bool {
	if transaction.id == "" {
		return false
	}
	if len(transaction.outputMap) < 1 {
		return false
	}
	if transaction.inputMap.address == "" {
		return false
	}
	return true
}
