package main

import (
	"fmt"
	"net/http"

	ErR "github.com/Developix-ir/Developix-Blockchain-Server/addon"
	"github.com/Developix-ir/Developix-Blockchain-Server/index"
	"github.com/Developix-ir/Developix-Blockchain-Server/src/blockchain"
	"github.com/Developix-ir/Developix-Blockchain-Server/src/network"
)

func main() {
	var pull /*s*/ network.NodesPull
	var blockchainApp /*s*/ blockchain.Blockchain
	rootServer := http.NewServeMux()
	nodeServer := http.NewServeMux()
	rootServer.HandleFunc("/", index.StartNode(pull, blockchainApp, CONFIG.ip, CONFIG.nodePort))
	nodeServer.HandleFunc("/", index.StartBlockchain(pull, blockchainApp))
	go func() {
		fmt.Println("node is running")
		err := http.ListenAndServe(CONFIG.ip+CONFIG.nodePort, nodeServer)
		ErR.ErrorCheck(err, 2)
	}()
	fmt.Println("root server is running")
	err := http.ListenAndServe(CONFIG.ip+CONFIG.port, rootServer)
	ErR.ErrorCheck(err, 2)
}
