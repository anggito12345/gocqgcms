package hubs

import (
	"encoding/hex"
	"encoding/json"
	"kano/cqg/cqg-websocket/cores"
	"kano/cqg/cqg-websocket/services"
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	pt "kano/cqg/cqg-websocket/proto-gen/proto"

	"kano/cqg/cqg-websocket/models"
)

type HubTraffic struct {
	*cores.CoreHubTraffic
	ws        *websocket.Conn
	token     string
	Signal    int
	RequestId uint32
}

func (cht *HubTraffic) Receive(cl pt.ClientMessage, msg []byte) error {

	if cl.String() == "" {
		cht.ws.WriteMessage(websocket.BinaryMessage, msg)
	} else {
		str := hex.EncodeToString([]byte(cl.String()))
		x, _ := hex.DecodeString(str)
		cht.ws.WriteMessage(websocket.BinaryMessage, []byte(string(x)))
	}

	if cht.Signal == cores.LOGON_SIGNAL {
		//when logged in do ping
		if cl.Logon != nil {
			cht.token = *cl.Logon.ClientAppId
		}

	} else if cht.Signal == cores.CREATECUSTOMER_SIGNAL {
	} else if cht.Signal == cores.BALANCEINFO_SIGNAL {
		log.Println(string(msg))
	}

	return nil
}

func (cht *HubTraffic) ReceiveClient(msg []byte) error {

	requestData := models.RequestModel{}

	err := json.Unmarshal(msg, &requestData)

	if err != nil {
		return err
	}

	if requestData.Signal == cores.CREATECUSTOMER_SIGNAL {
		customer := pt.Customer{}
		byteData, _ := json.Marshal(requestData.Data)

		json.Unmarshal(byteData, &customer)

		services.NewOperationService(cores.CQG_CONNECTION).CreateCustomer(&cht.RequestId, customer, cht)

		cht.RequestId++
	} else if requestData.Signal == cores.CREATEACCOUNT_SIGNAL {

		account := pt.Account{}
		byteData, _ := json.Marshal(requestData.Data)

		json.Unmarshal(byteData, &account)

		services.NewTradeRouting(cores.CQG_CONNECTION).CreateAccount(&cht.RequestId, account, cht)

		cht.RequestId++
	} else if requestData.Signal == cores.CREATEBALANCE_SIGNAL {
		createBalance := pt.CreateBalanceRecord{}
		byteData, _ := json.Marshal(requestData.Data)

		json.Unmarshal(byteData, &createBalance)

		services.NewTradeRouting(cores.CQG_CONNECTION).CreateBalance(&cht.RequestId, createBalance, cht)

		cht.RequestId++
	} else if requestData.Signal == cores.BALANCEINFO_SIGNAL {
		balanceRecordsRequest := pt.BalanceRecordsRequest{}
		byteData, _ := json.Marshal(requestData.Data)

		json.Unmarshal(byteData, &balanceRecordsRequest)

		services.NewTradeRouting(cores.CQG_CONNECTION).BalanceInfo(&cht.RequestId, balanceRecordsRequest, cht)

		cht.RequestId++
	} else if requestData.Signal == cores.UPDATEALANCE_SIGNAL {
		balanceRecordsRequest := pt.UpdateBalanceRecord{}
		byteData, _ := json.Marshal(requestData.Data)

		json.Unmarshal(byteData, &balanceRecordsRequest)

		services.NewTradeRouting(cores.CQG_CONNECTION).UpdateBalance(&cht.RequestId, balanceRecordsRequest, cht)

		cht.RequestId++
	}

	return nil
}

func (cht *HubTraffic) SetSignal(signal int) {
	cht.Signal = signal
}

type cqg_hub struct {
	conn    *websocket.Conn
	send    chan []byte
	hub     *Hub
	traffic HubTraffic
}

func NewCqgHub() *cqg_hub {
	initRequestId := uint32(0)

	return &cqg_hub{
		hub: &Hub{},
		traffic: HubTraffic{
			Signal:    cores.STANDBY_SIGNAL,
			RequestId: initRequestId,
		},
	}
}

func (sh *cqg_hub) SessionReadMessage(w http.ResponseWriter, r *http.Request) {

	go h.Run()

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	sh.traffic.ws = ws

	ws.SetCloseHandler(func(i int, text string) error {
		log.Println("Websocket client close")

		return nil
	})

	room := r.URL.Query()["room"][0]

	c := &connection{send: make(chan []byte, 256), ws: sh.traffic.ws}
	s := subscription{c, room}
	h.register <- s
	go s.writePump()

	if sh.traffic.token == "" {
		services.NewSessionService(cores.CQG_CONNECTION).Logon(&sh.traffic)

	} else {
		sh.traffic.ws.WriteMessage(websocket.BinaryMessage, []byte("Logged In"))
	}

	s.readPump(&sh.traffic)
}
