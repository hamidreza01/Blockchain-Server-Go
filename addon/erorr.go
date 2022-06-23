package error

import (
	"log"
	"net/http"
)

func LogFunc(r /*s*/ http.Request, body []byte) {
	log.Printf("ip: %s\npath: %s\nmethod: %s\nbody: %s\n", r.RemoteAddr, r.URL.Path, r.Method, string(body))
}
func ErrorCheck(err error, mod int) {
	if err != nil {
		if mod == 1 {
			log.Println(err)
		} else {
			log.Fatalln(err)
		}
	}
}
