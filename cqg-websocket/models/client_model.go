package models

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	Ws   *websocket.Conn
	Data struct {
		ClientID string         `json:"clientid" bson:"clientid"`
		Customer *CustomerModel `json:"customer" bson:"customer"`
	} `json:"data" bson:"bson"`
}
