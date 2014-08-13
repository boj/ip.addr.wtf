package main

import (
	"io"
	"log"
	"log/syslog"
	"net/http"
	"os"
	"strings"
)

var Logger *log.Logger

func init() {
	sl, err := syslog.Dial("", "", syslog.LOG_ERR, "")
	if err != nil {
		log.Print(err.Error())
		os.Exit(1)
	}
	Logger = log.New(sl, "[ip.addr.wtf] ", 0)
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	addr := strings.Split(r.RemoteAddr, ":")[0]
	Logger.Print(addr)
	io.WriteString(w, addr)
}

func main() {
	http.HandleFunc("/", RootHandler)
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal("ip.addr.wtf: ", err)
	}
}
