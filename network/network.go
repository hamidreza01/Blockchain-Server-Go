package network

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func errorCheck(err error, t int) {
	if err != nil {
		if t == 1 {
			log.Println(err)
		} else if t == 2 {
			log.Fatalln(err)
		}
	}
}
func logFunc(r *http.Request, body []byte) {
	fmt.Printf("ip: %s\npath: %s\nmethod: %s\nbody: %s\n", r.RemoteAddr, r.URL.Path, r.Method, string(body))
}

func SendAllData(pull NodesPull, url string, hash string, serverIp string, nodePort string) {
	nodes := pull.getNodes(serverIp, nodePort)
	// chain := blockchain.getChain()
	data, err := json.Marshal(
		map[string]interface{}{
			"nodes": nodes,
			"hash":  hash,
		},
	)
	errorCheck(err, 1)
	pull.brodcastAnode(url, "welcome", data)
}

func Start(pull NodesPull, serverIp string, nodeServerPort string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := r.Body.Close(); err != nil {
				log.Println(err)
			}
			w.Write([]byte("ok"))
		}()
		// add node
		if r.URL.Path == "/addMe" {
			body, err := io.ReadAll(r.Body)
			logFunc(r, body)
			errorCheck(err, 1)
			var bodyJSON map[string]interface{}
			err = json.Unmarshal(body, &bodyJSON)
			errorCheck(err, 1)
			nodeData := bodyJSON["data"].(map[string]interface{})
			ip := strings.Split(r.RemoteAddr, ":")[0]
			nodePort := nodeData["port"].(float64)
			rootPort := nodeData["port"].(float64)
			hash := nodeData["hash"].(string)
			// sendData `nodes` and `chain
			SendAllData(pull, ip+":"+fmt.Sprintf("%.0f", rootPort), hash, serverIp, nodeServerPort)
			pull.nodes = append(pull.nodes, nodes{hash: hash, ip: ip, port: rootPort, nodePort: nodePort})
		}

	}

}
