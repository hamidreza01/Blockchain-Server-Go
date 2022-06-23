package index

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	addon "github.com/Developix-ir/Developix-Blockchain-Server/addon"
	"github.com/Developix-ir/Developix-Blockchain-Server/src/blockchain"
	"github.com/Developix-ir/Developix-Blockchain-Server/src/network"
)

func StartBlockchain(pull /*s*/ network.NodesPull, bl /*s*/ blockchain.Blockchain) func(w http.ResponseWriter, r /*s*/ *http.Request) {
	return func(w http.ResponseWriter, r /*s*/ *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		defer func() {
			err := r.Body.Close()
			addon.ErrorCheck(err, 1)
			_, err = w.Write([]byte("ok"))
			addon.ErrorCheck(err, 1)
		}()
		body, err := io.ReadAll(r.Body)
		addon.ErrorCheck(err, 1)
		addon.LogFunc(*r, body)
		if r.Method == "POST" && r.Body != nil {
			var validNode bool
			nodeIp := strings.Split(r.RemoteAddr, ":")[0]
			for _, v := range /*s*/ pull.Nodes {
				if v.Ip == nodeIp {
					validNode = true
					break
				}
			}
			if validNode {
				if r.URL.Path == "/chain" {
					var data []blockchain.Block
					err := json.Unmarshal(body, &data)
					addon.ErrorCheck(err, 1)
					if bl.IsValid(data) && bl.ValidTransactionData(data) {
						bl.ClearTransaction(data)
						bl.Chain = data
					}
					fmt.Printf("%+v\n", bl.Chain)
				} else if r.URL.Path == "/transaction" {
					// if Transaction.isValid()
				}
			}
		}
	}
}
