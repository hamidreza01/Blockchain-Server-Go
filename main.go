package main

import (
	"log"
	"net/http"

	"github.com/Developix-ir/Developix-Blockchain-Server/blockchain"
	"github.com/Developix-ir/Developix-Blockchain-Server/network"
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

func main() {
	var pull network.NodesPull
	rootServer := http.NewServeMux()
	nodeServer := http.NewServeMux()
	rootServer.HandleFunc("/", network.Start(&pull, CONFIG.ip, CONFIG.nodePort))
	nodeServer.HandleFunc("/", blockchain.Start(&pull))
	go func() {
		err := http.ListenAndServe(CONFIG.ip+CONFIG.nodePort, nodeServer)
		errorCheck(err, 2)
	}()
	err := http.ListenAndServe(CONFIG.ip+CONFIG.port, rootServer)
	errorCheck(err, 2)
}
