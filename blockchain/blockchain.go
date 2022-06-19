package blockchain

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

type Blockchain struct {
	chain []Block
}

func (_ Blockchain) isValid(chain []Block) bool {
	if chain[0].difficulty != CONFIG.genesis.difficulty && chain[0].hash != CONFIG.genesis.hash && chain[0].lastHash != CONFIG.genesis.lastHash && chain[0].nonce != CONFIG.genesis.nonce && chain[0].timestamp != CONFIG.genesis.timestamp && len(chain[0].data.Transaction) != 0 {
		return false
	}
	for i := 1; i < len(chain); i++ {
		if chain[i].hash != Hash(chain[i].lastHash,
			toJson(chain[i].data),
			strconv.Itoa(chain[i].difficulty),
			strconv.Itoa(chain[i].nonce),
			strconv.Itoa(chain[i].timestamp)) {
			return false
		}
		if chain[i-1].hash != chain[i].lastHash {
			return false
		}
		if math.Abs(float64(chain[i-1].difficulty)-float64(chain[i].difficulty)) > 1 {
			return false
		}
	}
	return true
}

func (_ Blockchain) validTransactionData(chain []Block) bool {
	for i := 1; i < len(chain); i++ {
		if len(chain[i].data.Transaction) < 1 {
			return false
		}
		for _, transaction := range chain[i].data.Transaction {
			rewardNumber := 0
			if transaction.inputMap.address == CONFIG.reward.address {
				rewardNumber++
				if rewardNumber > 1 {
					return false
				}
				if transaction.outputMap[CONFIG.reward.address] > CONFIG.rewardValue {
					return false
				}
			} else {
				// transactionResualt := Transaction.isValid(transaction)
				// if transactionResualt != true {
				// 	return false
				// } else {
				// 	trueValue := Wallet.CalculateBalance(chain, transaction.InputMap.Address)
				// 	if trueValue != transaction.InputMap.Amount {
				// 		return false
				// 	}
				// }
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
