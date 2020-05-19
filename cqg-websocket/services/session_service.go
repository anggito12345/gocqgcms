package services

import (
	"kano/cqg/cqg-websocket/cores"
	pt "kano/cqg/cqg-websocket/proto-gen/proto"

	"github.com/gorilla/websocket"
)

type session_service struct {
	ws *websocket.Conn
}

func NewSessionService(ws *websocket.Conn) *session_service {
	return &session_service{
		ws: ws,
	}
}

func (os *session_service) Logon(traffic cores.TrafficConnection) *pt.LogonResult {

	traffic.SetSignal(cores.LOGON_SIGNAL)

	cores.SIGNAL_READ <- traffic

	username := "ClearRiskSIM"
	password := "1Password!"
	clientAppId := "CmsApiTest"
	clientApiTest := "CmsApiTest"
	protocolVersionMajor := uint32(pt.ProtocolVersion_value["PROTOCOL_VERSION_MAJOR"])
	protocolVersionMinor := uint32(pt.ProtocolVersion_value["PROTOCOL_VERSION_MINOR"])

	logon := pt.Logon{}
	logon.UserName = &username
	logon.Password = &password
	logon.ClientAppId = &clientAppId
	logon.ClientVersion = &clientApiTest
	logon.ProtocolVersionMajor = &protocolVersionMajor
	logon.ProtocolVersionMinor = &protocolVersionMinor

	cl := pt.ClientMessage{}
	cl.Logon = &logon

	SendMessage(cl, os.ws)

	return nil
}

func (os *session_service) Logoff(traffic cores.TrafficConnection) *pt.LogonResult {

	traffic.SetSignal(cores.LOGOFF_SIGNAL)

	cores.SIGNAL_READ <- traffic

	logoff := pt.Logoff{}

	cl := pt.ClientMessage{}
	cl.Logoff = &logoff

	SendMessage(cl, os.ws)

	return nil
}

func (os *session_service) Ping(traffic cores.TrafficConnection) *pt.LogonResult {

	traffic.SetSignal(cores.PING_SIGNAL)

	cores.SIGNAL_READ <- traffic

	ping := pt.Ping{}

	cl := pt.ClientMessage{}
	cl.Ping = &ping

	SendMessage(cl, os.ws)

	return nil
}
