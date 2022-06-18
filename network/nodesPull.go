package network

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type nodes struct {
	hash     string
	ip       string
	port     float64
	nodePort float64
}

type data interface{}

type NodesPull struct {
	nodes []nodes
}

func (n *NodesPull) brodcast(channel string, da data) {
	for _, v := range n.nodes {
		sendingData := make(map[string]interface{})
		sendingData["data"] = da
		sendingData["data"].(map[string]interface{})["hash"] = v.hash
		d, err := json.Marshal(sendingData)
		errorCheck(err, 2)
		_, err = http.Post("http://"+v.ip+":"+fmt.Sprintf("%.0f", v.port)+"/"+channel, "application/json", bytes.NewBuffer(d))
		errorCheck(err, 1)
	}
}

func (n NodesPull) brodcastAnode(url string, channel string, da data) {
	sendingData := make(map[string]interface{})
	sendingData["data"] = da
	d, err := json.Marshal(sendingData)
	errorCheck(err, 2)
	_, err = http.Post("http://"+url+"/"+channel, "application/json", bytes.NewBuffer(d))
	errorCheck(err, 1)
}

func (n *NodesPull) getNodes(ip string, nodeServerPort string) []string {
	var sendingData []string
	sendingData = append(sendingData, ip+nodeServerPort)
	for _, v := range n.nodes {
		sendingData = append(sendingData, v.ip+":"+fmt.Sprintf("%.0f", v.nodePort))
	}
	return sendingData
}
