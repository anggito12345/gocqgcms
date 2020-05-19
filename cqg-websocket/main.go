package main

import (
	"flag"
	"kano/cqg/cqg-websocket/cores"
	"kano/cqg/cqg-websocket/hubs"
	"log"
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/gorilla/websocket"
)

var ws *websocket.Conn
var store = sessions.NewCookieStore(securecookie.GenerateRandomKey(32))

//var token string = ""

func main() {

	var err error

	ws, err = cores.Run()
	defer func() {
		cores.Close()
	}()

	if err != nil {
		log.Println("connection fail:", err)
		return
	}

	var addr = flag.String("addr", "localhost:8080", "http service address")

	http.HandleFunc("/cqg-hub", hubs.NewCqgHub().SessionReadMessage)

	log.Fatal(http.ListenAndServe(*addr, nil))

}
