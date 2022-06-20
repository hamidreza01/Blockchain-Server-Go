package blockchain

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

type Blockchain struct {
	Chain           []Block
	TransactionPull []Transaction
}

func (b *Blockchain) isValid(chain []Block) bool {
	if chain[0].Difficulty != CONFIG.genesis.Difficulty && chain[0].Hash != CONFIG.genesis.Hash && chain[0].LastHash != CONFIG.genesis.LastHash && chain[0].Nonce != CONFIG.genesis.Nonce && chain[0].Timestamp != CONFIG.genesis.Timestamp && len(chain[0].Data.Transaction) != 0 {
		fmt.Printf("%s\n", "invalid genesis block")
		return false
	}
	if len(chain) <= len(b.Chain) {
		fmt.Printf("%s\n", "invalid chain length")
		return false
	}
	for i := 1; i < len(chain); i++ {
		if chain[i].Hash != Hash(chain[i].LastHash,
			`"`+toJson(chain[i].Data)+`"`,
			strconv.Itoa(chain[i].Difficulty),
			strconv.Itoa(chain[i].Nonce),
			strconv.Itoa(chain[i].Timestamp)) {
			fmt.Printf("\njs hash: %s | go hash: %s\n", chain[i].Hash, Hash(chain[i].LastHash,
				toJson(chain[i].Data),
				strconv.Itoa(chain[i].Difficulty),
				strconv.Itoa(chain[i].Nonce),
				strconv.Itoa(chain[i].Timestamp)))
			fmt.Printf("js number1: %d | go number1: %s\n", chain[i].Difficulty, strconv.Itoa(chain[i].Difficulty))
			fmt.Printf("js number2: %d | go number2: %s\n", chain[i].Nonce, strconv.Itoa(chain[i].Nonce))
			fmt.Printf("js number3: %d | go number3: %s\n", chain[i].Timestamp, strconv.Itoa(chain[i].Timestamp))

			fmt.Printf("go json: %s\n", toJson(chain[i].Data))
			fmt.Printf("%s\n", "invalid hash")
			return false
		}
		if chain[i-1].Hash != chain[i].LastHash {
			fmt.Printf("%s\n", "invalid last hash")
			return false
		}
		if math.Abs(float64(chain[i-1].Difficulty)-float64(chain[i].Difficulty)) > 1 {
			fmt.Printf("%s\n", "invalid difficulty")
			return false
		}
	}
	return true
}

func (_ Blockchain) validTransactionData(chain []Block) bool {
	for i := 1; i < len(chain); i++ {
		if len(chain[i].Data.Transaction) < 1 {
			fmt.Printf("%s\n", "invalid transaction data")
			return false
		}
		for _, transaction := range chain[i].Data.Transaction {
			rewardNumber := 0
			if transaction.InputMap.Address == CONFIG.reward.address {
				fmt.Printf("%s\n", "invalid reward transaction")
				rewardNumber++
				if rewardNumber > 1 {
					fmt.Printf("%s\n", "invalid reward transaction 2")
					return false
				}
				if transaction.OutputMap[CONFIG.reward.address] > CONFIG.rewardValue {
					fmt.Printf("%s\n", "invalid reward transaction 3")
					return false
				}
			} else {
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
			}
		}
	}
	return true
}

func errorCheck(err error, t int) {
	if err != nil {
		if t == 1 {
			log.Println(t)
		} else if t == 2 {
			log.Fatalln(t)
		}
	}
}

func logFunc(r *http.Request, body []byte) {
	fmt.Printf("ip: %s\npath: %s\nmethod: %s\nbody: %s\n", r.RemoteAddr, r.URL.Path, r.Method, string(body))
}
