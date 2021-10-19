package okex

import (
	"encoding/json"
	"time"

	. "github.com/tbtc-bot/go-okex/impl"
)

// Endpoints
const (
	baseWsPublicURL  = "wss://ws.okex.com:8443/ws/v5/public"
	baseWsPrivateURL = "wss://ws.okex.com:8443/ws/v5/private"
)

var (
	// WebsocketTimeout is an interval for sending ping/pong messages if WebsocketKeepalive is enabled
	WebsocketTimeout = time.Second * 20
	// WebsocketKeepalive enables sending ping/pong messages to check the connection stability
	WebsocketKeepalive = false
)

// getWsEndpoint return the base endpoint of the WS according the UseTestnet flag
func getWsEndpoint(private bool) string {
	if private {
		return baseWsPrivateURL
	}
	return baseWsPublicURL
}

// WsInstruments define websocket struct event
type WsInstrumentsEvent struct {
	Arg  map[string]string `json:"arg"`
	Data []*WsInstrument   `json:"data"`
}

type WsInstrument struct {
	InstType  string `json:"instType"`
	InstId    string `json:"instId"`
	Uly       string `json:"uly"`
	Category  string `json:"category"`
	BaseCcy   string `json:"baseCcy"`
	QuoteCcy  string `json:"quoteCcy"`
	SettleCcy string `json:"settleCcy"`
	CtVal     string `json:"ctVal"`
	CtMult    string `json:"ctMult"`
	CtValCcy  string `json:"ctValCcy"`
	OptType   string `json:"optType"`
	Stk       string `json:"stk"`
	ListTime  string `json:"listTime"`
	ExpTime   string `json:"expTime"`
	TickSz    string `json:"tickSz"`
	LotSz     string `json:"lotSz"`
	MinSz     string `json:"minSz"`
	CtType    string `json:"ctType"`
	Alias     string `json:"alias"`
	State     string `json:"state"`
}

// WsInstruments handle websocket instrument message
type WsInstrumentsHandler func(event *WsInstrumentsEvent)

// WsInstruments as per https://www.okex.com/docs-v5/en/#websocket-api-public-channels-instruments-channel
func WsInstrumentsServe(instType string, handler WsInstrumentsHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := getWsEndpoint(false)
	return wsInstrumentsServe(endpoint, instType, handler, errHandler)
}

// WsInstrumentsServe serve websocket
func wsInstrumentsServe(endpoint string, instType string, handler WsInstrumentsHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	arg := map[string]string{
		"channel":  "instruments",
		"instType": instType,
	}
	var args []map[string]string
	args = append(args, arg)
	reqData := ReqData{Op: "subscribe",
		Args: args,
	}
	//fmt.Println(reqData)
	cfg := newWsConfig(endpoint, reqData)
	wsHandler := func(message []byte) {
		event := new(WsInstrumentsEvent)
		err = json.Unmarshal(message, event)
		if err != nil {
			errHandler(err)
			return
		}

		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}
