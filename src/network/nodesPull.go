package network

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	addon "github.com/Developix-ir/Developix-Blockchain-Server/addon"
)

type Nodes struct {
	Hash     string
	Ip       string
	Port     float64
	NodePort float64
}

type data interface{}

type NodesPull struct {
	Nodes []Nodes
}

func (n *NodesPull) Brodcast(channel string, da data) {
	for _, v := range n.Nodes {
		sendingData := make(map[string]interface{})
		sendingData["data"] = da
		sendingData["data"].(map[string]interface{})["hash"] = v.Hash
		d, err := json.Marshal(sendingData)
		addon.ErrorCheck(err, 2)
		_, err = http.Post("http://"+v.Ip+":"+fmt.Sprintf("%.0f", v.Port)+"/"+channel, "application/json", bytes.NewBuffer(d))
		addon.ErrorCheck(err, 1)
	}
}

func (n NodesPull) BrodcastAnode(url string, channel string, da data) {
	sendingData := make(map[string]interface{})
	sendingData["data"] = da
	d, err := json.Marshal(sendingData)
	addon.ErrorCheck(err, 2)
	_, err = http.Post("http://"+url+"/"+channel, "application/json", bytes.NewBuffer(d))
	addon.ErrorCheck(err, 1)
}

func (n *NodesPull) GetNodes(ip string, nodeServerPort string) []string {
	var sendingData []string
	sendingData = append(sendingData, ip+nodeServerPort)
	for _, v := range n.Nodes {
		sendingData = append(sendingData, v.Ip+":"+fmt.Sprintf("%.0f", v.NodePort))
	}
	return sendingData
}
