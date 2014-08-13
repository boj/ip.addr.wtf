package main

import (
	"io"
	"log"
	"net/http"
	"strings"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	addr := strings.Split(r.RemoteAddr, ":")[0]
	io.WriteString(w, addr)
}

func main() {
	http.HandleFunc("/", RootHandler)
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal("ip.addr.wtf: ", err)
	}
}
