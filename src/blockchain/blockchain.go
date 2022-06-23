package blockchain

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

type Blockchain struct {
	Chain           []Block
	TransactionPull []Transaction
}

func (b *Blockchain) ClearTransaction(chain []Block) {
	var cleanTx []Transaction
	for _, v := range b.TransactionPull {
		has := false
		for _, v2 := range chain {
			if has {
				break
			}
			for _, v3 := range *v2.Data.Transaction {
				if v3.Id == v.Id {
					has = true
					break
				}
			}
		}
		if !has {
			cleanTx = append(cleanTx, v)
		}
	}
	b.TransactionPull = cleanTx
}

func (b *Blockchain) IsValid(chain []Block) bool {
	if chain[0].Difficulty != CONFIG.genesis.Difficulty && chain[0].Hash != CONFIG.genesis.Hash && chain[0].LastHash != CONFIG.genesis.LastHash && chain[0].Nonce != CONFIG.genesis.Nonce && chain[0].Timestamp != CONFIG.genesis.Timestamp && len(*chain[0].Data.Transaction) != 0 {
		fmt.Printf("%s\n", "invalid genesis block")
		return false
	}
	if len(chain) <= len(b.Chain) {
		fmt.Printf("%s\n", "invalid chain length")
		return false
	}
	for i := 1; i < len(chain); i++ {
		if chain[i].Hash != Hash(
			chain[i].LastHash,
			strconv.Itoa(chain[i].Nonce),
			strconv.Itoa(chain[i].Timestamp),
			strconv.Itoa(chain[i].Difficulty),
			ToJson(chain[i].Data)) {

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

func (_ Blockchain) ValidTransactionData(chain []Block) bool {
	for i := 1; i < len(chain); i++ {
		if len(*chain[i].Data.Transaction) < 1 {
			fmt.Printf("%s\n", "invalid transaction data")
			return false
		}
		for _, transaction := range *chain[i].Data.Transaction {
			rewardNumber := 0
			if transaction.InputMap.Address == CONFIG.reward.address {
				rewardNumber++
				if rewardNumber > 1 {
					fmt.Printf("%s\n", "invalid reward transaction 1")
					return false
				}
				if transaction.OutputMap[CONFIG.reward.address] > CONFIG.rewardValue {
					fmt.Printf("%s\n", "invalid reward transaction 2")
					return false
				}
			} else {
				if !transaction.IsValid(transaction) {
					return false
				}
			}
		}
	}
	return true
}

func ErrorCheck(err error, t int) {
	if err != nil {
		if t == 1 {
			log.Println(t)
		} else if t == 2 {
			log.Fatalln(t)
		}
	}
}
