package blockchain

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Start() func(w http.ResponseWriter, r *http.Request) {
	fmt.Println(Hash(`{"transaction":[{"id":"d0et570l4n4f9pc","outputMap":{"040ee51e4b467451a4aeb0ced2808d335cf5682ee38d14fc6f599f1096dd21d2722a65bc3cffce7bf98d7a6f4152ba93c04a93c78da762eec31f942d5a8c3f0959":10},"inputMap":{"address":"**DPX Blockchain**"}}]}`))
	return func(w http.ResponseWriter, r *http.Request) {
		var blockchain Blockchain
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		defer func() {
			err := r.Body.Close()
			errorCheck(err, 1)
			_, err = w.Write([]byte("ok"))
			errorCheck(err, 1)
		}()
		body, err := io.ReadAll(r.Body)
		errorCheck(err, 1)
		logFunc(r, body)
		if r.Method == "POST" && r.Body != nil {
			if r.URL.Path == "/chain" {
				var data []Block
				err := json.Unmarshal(body, &data)
				errorCheck(err, 1)
				fmt.Printf("\n%t\n", blockchain.isValid(data) && blockchain.validTransactionData(data))
				// var data []map[string]interface{}
				// _ = json.NewDecoder(r.Body).Decode(&data)
				// var chain []Block
				// for _, v := range data {
				// 	var transactions []Transaction
				// 	for i := 1; i < len(v["data"].(map[string]interface{})["transaction"].([]map[string]interface{})); i++ {
				// 		v2 := v["data"].(map[string]interface{})["transaction"].([]map[string]interface{})[i]
				// 		transactions = append(transactions, Transaction{
				// 			id: v2["id"].(string),
				// 			inputMap: struct{ address string }{
				// 				address: v2["inputMap"].(map[string]interface{})["address"].(string),
				// 			},
				// 			outputMap: v["outputMap"].(map[string]int),
				// 		})
				// 	}
				// 	chain = append(chain, Block{
				// 		hash:     v["hash"].(string),
				// 		lastHash: v["lastHash"].(string),
				// 		data: struct{ Transaction []Transaction }{
				// 			Transaction: transactions,
				// 		},
				// 		difficulty: v["difficulty"].(int),
				// 		timestamp:  v["timestamp"].(int),
				// 		nonce:      v["nonce"].(int),
				// 	})
				// }
			} else if r.URL.Path == "/transaction" {

			}
		}
	}
}
