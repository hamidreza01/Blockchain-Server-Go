package blockchain

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Developix-ir/Developix-Blockchain-Server/network"
)

func Start(pull *network.NodesPull) func(w http.ResponseWriter, r *http.Request) {
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
			// request from node
			var validNode bool
			nodeIp := strings.Split(r.RemoteAddr, ":")[0]
			// create a hello world function
			for _, v := range *pull.Nodes {
				if *v.Ip == nodeIp {
					validNode = true
					break
				}
			}
			if validNode {
				if r.URL.Path == "/chain" {
					var data []Block
					err := json.Unmarshal(body, &data)
					errorCheck(err, 1)
					if blockchain.isValid(data) && blockchain.validTransactionData(data) {
						blockchain.clearTransaction(data)
						blockchain.Chain = data
					}
				} else if r.URL.Path == "/transaction" {
					// if Transaction.isValid()
				}
			}
		}
	}
}
