package main

import (
	"net/http"
)

import _ "net/http/pprof"

func HiHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi"))
}

func main() {
	//r := http.NewServeMux()
	http.HandleFunc("/", HiHandler)

	//r.HandleFunc("/debug/pprof/", pprof.Index)
	//r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	//r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	//r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	//r.HandleFunc("/debug/pprof/trace", pprof.Trace)

	http.ListenAndServe(":8888", nil)
}
