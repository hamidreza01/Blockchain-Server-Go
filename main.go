package main

import (
	"net/http"

	ErR "github.com/Developix-ir/Developix-Blockchain-Server/addon"
	"github.com/Developix-ir/Developix-Blockchain-Server/index"
	"github.com/Developix-ir/Developix-Blockchain-Server/src/blockchain"
	"github.com/Developix-ir/Developix-Blockchain-Server/src/network"
)

func main() {
	var pull *network.NodesPull
	var blockchainApp *blockchain.Blockchain
	rootServer := http.NewServeMux()
	nodeServer := http.NewServeMux()
	rootServer.HandleFunc("/", index.StartNode(pull, blockchainApp, CONFIG.ip, CONFIG.nodePort))
	nodeServer.HandleFunc("/", index.StartBlockchain(pull, blockchainApp))
	go func() {
		err := http.ListenAndServe(CONFIG.ip+CONFIG.nodePort, nodeServer)
		ErR.ErrorCheck(err, 2)
	}()
	err := http.ListenAndServe(CONFIG.ip+CONFIG.port, rootServer)
	ErR.ErrorCheck(err, 2)
}
