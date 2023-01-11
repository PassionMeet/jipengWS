package main

import (
	"log"
	"net/http"

	"github.com/cmfunc/jipengWS/hub"
)

func main() {
	hubO := hub.NewHub()
	go hubO.Run()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		hub.ServeWs(hubO, w, r)
	})
	addr := ":7770"
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
