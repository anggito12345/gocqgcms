package services

import (
	"kano/cqg/cqg-websocket/cores"
	pt "kano/cqg/cqg-websocket/proto-gen/proto"
	"log"

	"github.com/gorilla/websocket"
)

type operation_service struct {
	ws *websocket.Conn
}

func NewOperationService(ws *websocket.Conn) *operation_service {
	return &operation_service{
		ws: ws,
	}
}

func (os *operation_service) CreateCustomer(requestId *uint32, customer pt.Customer, traffic cores.TrafficConnection) {

	clcreatecs := pt.ClientMessage{}

	createCustomer := pt.CreateCustomer{}
	createCustomer.Customer = &customer

	operationRequest := pt.OperationRequest{
		Id:             requestId,
		CreateCustomer: &createCustomer,
	}

	log.Println("INFO", *requestId)

	clcreatecs.OperationRequest = append(clcreatecs.OperationRequest, &operationRequest)
	traffic.SetSignal(cores.CREATECUSTOMER_SIGNAL)

	cores.SIGNAL_READ <- traffic

	SendMessage(clcreatecs, os.ws)
}
