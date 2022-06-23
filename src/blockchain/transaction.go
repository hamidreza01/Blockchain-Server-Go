package blockchain

import "fmt"

type Transaction struct {
	Id        string         `json:"id"`
	OutputMap map[string]int `json:"outputMap"`
	InputMap  struct {
		Address string `json:"address"`
	} `json:"inputMap"`
}

func (t Transaction) IsValid(transaction Transaction) bool {
	if transaction.Id == "" {
		fmt.Printf("%s\n", "invalid transaction id")
		return false
	}
	if transaction.InputMap.Address == "" {
		fmt.Printf("%s\n", "invalid transaction input address")
		return false
	}
	if len(transaction.OutputMap) < 1 {
		fmt.Printf("%s\n", "invalid transaction output map")
		return false
	}
	for address, value := range transaction.OutputMap {
		if address == "" {
			fmt.Printf("%s\n", "invalid transaction output address")
			return false
		}
		if value < 0 {
			fmt.Printf("%s\n", "invalid transaction output value")
			return false
		}
	}
	return true
}
