package services

import (
	"kano/cqg/cqg-websocket/cores"
	pt "kano/cqg/cqg-websocket/proto-gen/proto"

	"github.com/gorilla/websocket"
)

type traderouting_service struct {
	ws *websocket.Conn
}

func NewTradeRouting(ws *websocket.Conn) *traderouting_service {
	return &traderouting_service{
		ws: ws,
	}
}

func (ts *traderouting_service) CreateAccount(requestId *uint32, account pt.Account, traffic cores.TrafficConnection) {

	traffic.SetSignal(cores.CREATEACCOUNT_SIGNAL)

	cores.SIGNAL_READ <- traffic

	clcreatecs := pt.ClientMessage{}

	classValue := uint32(1)
	account.Class = &classValue

	createAccount := pt.CreateAccount{}
	createAccount.Account = &account

	accountScopeRequest := pt.AccountScopeRequest{
		CreateAccount: &createAccount,
	}

	tradeRoutingRequest := pt.TradeRoutingRequest{
		Id:                  requestId,
		AccountScopeRequest: &accountScopeRequest,
	}

	clcreatecs.TradeRoutingRequest = append(clcreatecs.TradeRoutingRequest, &tradeRoutingRequest)

	SendMessage(clcreatecs, ts.ws)
}

func (ts *traderouting_service) CreateBalance(requestId *uint32, createBalanceRecord pt.CreateBalanceRecord, traffic cores.TrafficConnection) {

	traffic.SetSignal(cores.CREATEACCOUNT_SIGNAL)

	cores.SIGNAL_READ <- traffic

	clcreatecs := pt.ClientMessage{}
	accountScopeRequest := pt.AccountScopeRequest{
		CreateBalanceRecord: &createBalanceRecord,
	}

	tradeRoutingRequest := pt.TradeRoutingRequest{
		Id:                  requestId,
		AccountScopeRequest: &accountScopeRequest,
	}

	clcreatecs.TradeRoutingRequest = append(clcreatecs.TradeRoutingRequest, &tradeRoutingRequest)

	SendMessage(clcreatecs, ts.ws)
}

func (ts *traderouting_service) BalanceInfo(requestId *uint32, balanceRecord pt.BalanceRecordsRequest, traffic cores.TrafficConnection) {

	traffic.SetSignal(cores.BALANCEINFO_SIGNAL)

	cores.SIGNAL_READ <- traffic

	clcreatecs := pt.ClientMessage{}
	accountScopeRequest := pt.AccountScopeRequest{
		BalanceRecordsRequest: &balanceRecord,
	}

	tradeRoutingRequest := pt.TradeRoutingRequest{
		Id:                  requestId,
		AccountScopeRequest: &accountScopeRequest,
	}

	clcreatecs.TradeRoutingRequest = append(clcreatecs.TradeRoutingRequest, &tradeRoutingRequest)

	SendMessage(clcreatecs, ts.ws)
}

func (ts *traderouting_service) UpdateBalance(requestId *uint32, balanceRecord pt.UpdateBalanceRecord, traffic cores.TrafficConnection) {

	traffic.SetSignal(cores.BALANCEINFO_SIGNAL)

	cores.SIGNAL_READ <- traffic

	clcreatecs := pt.ClientMessage{}
	accountScopeRequest := pt.AccountScopeRequest{
		UpdateBalanceRecord: &balanceRecord,
	}

	tradeRoutingRequest := pt.TradeRoutingRequest{
		Id:                  requestId,
		AccountScopeRequest: &accountScopeRequest,
	}

	clcreatecs.TradeRoutingRequest = append(clcreatecs.TradeRoutingRequest, &tradeRoutingRequest)

	SendMessage(clcreatecs, ts.ws)
}
