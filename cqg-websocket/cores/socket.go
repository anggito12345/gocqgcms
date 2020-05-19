package cores

import (
	"encoding/hex"
	"log"
	"time"

	pt "kano/cqg/cqg-websocket/proto-gen/proto"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
)

type TrafficConnection interface {
	Receive(pt.ClientMessage, []byte) error
	ReceiveClient([]byte) error
	SetSignal(int)
}

type CoreHubTraffic struct {
	ws     *websocket.Conn
	Signal int
	token  string
}

func NewCoreHubTraffic() *CoreHubTraffic {
	return &CoreHubTraffic{
		Signal: STANDBY_SIGNAL,
		token:  "",
	}
}

func (cht *CoreHubTraffic) Receive(cl pt.ClientMessage, msg []byte) error {
	return nil
}

func (cht *CoreHubTraffic) SetSignal(Signal int) {
	//cht.Signal = Signal
}

func (cht *CoreHubTraffic) ReceiveClient(msg []byte) error {

	return nil
}

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Mfncaximum message size allowed from peer.
	maxMessageSize = 512

	STANDBY_SIGNAL        = 0
	LOGON_SIGNAL          = 1
	LOGOFF_SIGNAL         = 2
	PING_SIGNAL           = 3
	CREATECUSTOMER_SIGNAL = 4
	CREATEACCOUNT_SIGNAL  = 5
	CREATEBALANCE_SIGNAL  = 6
	UPDATEALANCE_SIGNAL   = 7
	BALANCEINFO_SIGNAL    = 8
)

var SIGNAL_READ = make(chan TrafficConnection)

var CQG_CONNECTION *websocket.Conn

func Run() (*websocket.Conn, error) {
	url := "wss://democmsapi.cqg.com"

	httpHeader := map[string][]string{}

	dialer := websocket.Dialer{}

	ws, _, err := dialer.Dial(url, httpHeader)
	if err != nil {
		return nil, err
	}

	ws.SetCloseHandler(func(i int, text string) error {
		log.Println("Websocket close", text)

		return nil
	})

	CQG_CONNECTION = ws

	go ReadMessageStream(ws)

	return ws, nil
}

func ReadMessageStream(ws *websocket.Conn) {
	go func() {
		for {
			select {
			case traffic := <-SIGNAL_READ:

				cl := pt.ClientMessage{}

				_, msg, _ := ws.ReadMessage()

				proto.Unmarshal(msg, &cl)

				if cl.String() == "" {
					log.Println("[INFO]", string(msg))
				} else {
					str := hex.EncodeToString([]byte(cl.String()))
					x, _ := hex.DecodeString(str)
					log.Println("[INFO]", string(x))
				}

				go traffic.Receive(cl, msg)
			}

		}
		log.Println("END?")
	}()
}

func Close() {
	close(SIGNAL_READ)
	//CQG_CONNECTION.Close()
}
