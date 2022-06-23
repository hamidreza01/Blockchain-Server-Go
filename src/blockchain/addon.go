package blockchain

import (
	"encoding/json"
	"log"
)

func ToJson(data any) string {
	d, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	return string(d)
}
