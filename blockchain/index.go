package blockchain

import (
	"fmt"
	"io"
	"net/http"
)

func Start() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := r.Body.Close()
			errorCheck(err, 1)
			_, err = w.Write([]byte("ok"))
			errorCheck(err, 1)
		}()
		body, err := io.ReadAll(r.Body)
		errorCheck(err, 1)
		logFunc(r, body)
		if r.Method == "POST" && r.Body != nil {
			if r.URL.Path == "/chain" {
				fmt.Printf("new chain %s")
			} else if r.URL.Path == "/transaction" {

			}
		}
	}
}
