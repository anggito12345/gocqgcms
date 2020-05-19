package models

type RequestModel struct {
	Signal int         `json:"signal" bson:"signal"`
	Data   interface{} `json:"data" bson:"data"`
}
