package index

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	addon "github.com/Developix-ir/Developix-Blockchain-Server/addon"
	"github.com/Developix-ir/Developix-Blockchain-Server/src/blockchain"
	"github.com/Developix-ir/Developix-Blockchain-Server/src/network"
)

func SendAllData(pull /*s*/ network.NodesPull, bl /*s*/ blockchain.Blockchain, url string, hash string, serverIp string, nodePort string) {
	fmt.Println("_______________________________\n send all data")
	nodes := pull.GetNodes(serverIp, nodePort)
	chain := bl.Chain
	transaction := bl.TransactionPull
	data, err := json.Marshal(
		map[string]interface{}{
			"nodes":       nodes,
			"chain":       chain,
			"transaction": transaction,
			"hash":        hash,
		},
	)
	addon.ErrorCheck(err, 1)
	pull.BrodcastAnode(url, "welcome", data)
}

func StartNode(pull /*s*/ network.NodesPull, bl /*s*/ blockchain.Blockchain, serverIp string, nodeServerPort string) func(w http.ResponseWriter, r /*s*/ *http.Request) {
	return func(w http.ResponseWriter, r /*s*/ *http.Request) {
		defer func() {
			if err := r.Body.Close(); err != nil {
				log.Println(err)
			}
			w.Write([]byte("ok"))
		}()
		if r.Method == "POST" && r.Body != nil {
			if r.URL.Path == "/addMe" {
				body, err := io.ReadAll(r.Body)
				addon.LogFunc(*r, body)
				addon.ErrorCheck(err, 1)
				var bodyJSON map[string]interface{}
				err = json.Unmarshal(body, &bodyJSON)
				addon.ErrorCheck(err, 1)
				nodeData := bodyJSON["data"].(map[string]interface{})
				ip := strings.Split(r.RemoteAddr, ":")[0]
				nodePort := nodeData["port"].(float64)
				rootPort := nodeData["port"].(float64)
				hash := nodeData["hash"].(string)

				// sendData `nodes` and `chain
				// if ip == "172.20.165.179" {
				// 	serverIp = "172.20.160.1"
				// }

				SendAllData(pull, bl, ip+":"+fmt.Sprintf("%.0f", rootPort), hash, serverIp, nodeServerPort)
				pull.Nodes = append(pull.Nodes, network.Nodes{Hash: hash, Ip: ip, Port: rootPort, NodePort: nodePort})
			} else {
				var validNode bool
				nodeIp := strings.Split(r.RemoteAddr, ":")[0]
				for _, v := range pull.Nodes {
					if v.Ip == nodeIp {
						validNode = true
						break
					}
				}
				if validNode {

				}
			}
		}

	}

}
