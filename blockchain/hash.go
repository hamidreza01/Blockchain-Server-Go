package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"sort"
)

func Hash(data ...string) string {
	sort.Strings(data)
	hash := sha256.New()
	for _, v := range data {
		hash.Write([]byte(v))
	}
	return hex.EncodeToString(hash.Sum(nil))
}
