package blockchain

import (
	"encoding/json"
	"log"
)

func toJson(data any) string {
	d, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	return string(d)
}
