package services

import (
	"log"

	"github.com/golang/protobuf/proto"

	"github.com/gorilla/websocket"

	pt "kano/cqg/cqg-websocket/proto-gen/proto"
)

func SendMessage(clientMessage pt.ClientMessage, ws *websocket.Conn) {

	var err error

	payload, _ := proto.Marshal(&clientMessage)

	err = ws.WriteMessage(websocket.BinaryMessage, payload)

	if err != nil {
		log.Fatal("ERROR REQUEST: " + err.Error())
	}

}
