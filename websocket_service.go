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
	WebsocketKeepalive = true
)

// getWsEndpoint return the base endpoint of the WS according the UseTestnet flag
func getWsEndpoint(private bool) string {
	if private {
		return baseWsPrivateURL
	}
	return baseWsPublicURL
}

// ACCCOUNT WEBSOCKET (PUBLIC)

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
	cfg := newWsConfig(endpoint, reqData, "", "", "")
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

// ACCCOUNT WEBSOCKET (PRIVATE)

type WsAccountsEvent struct {
	Arg  map[string]string `json:"arg"`
	Data []*WsAccount      `json:"data"`
}

type WsAccount struct {
	UTime       string             `json:"uTime"`
	TotalEq     string             `json:"totalEq"`
	IsoEq       string             `json:"isoEq"`
	AdjEq       string             `json:"adjEq"`
	OrdFroz     string             `json:"ordFroz"`
	Imr         string             `json:"imr"`
	Mmr         string             `json:"mmr"`
	NotionalUsd string             `json:"notionalUsd"`
	MgnRatio    string             `json:"mgnRatio"`
	Details     []*WsAccountDetail `json:"details"`
}

type WsAccountDetail struct {
	AvailBal      string `json:"availBal"`
	AvailEq       string `json:"availEq"`
	Ccy           string `json:"ccy"`
	CashBal       string `json:"cashBal"`
	UTime         string `json:"uTime"`
	DisEq         string `json:"disEq"`
	Eq            string `json:"eq"`
	EqUsd         string `json:"eqUsd"`
	FrozenBal     string `json:"frozenBal"`
	Interest      string `json:"interest"`
	IsoEq         string `json:"isoEq"`
	Liab          string `json:"liab"`
	MaxLoan       string `json:"maxLoan"`
	MgnRatio      string `json:"mgnRatio"`
	NotionalLever string `json:"notionalLever"`
	OrdFrozen     string `json:"ordFrozen"`
	Upl           string `json:"upl"`
	UplLiab       string `json:"uplLiab"`
	CrossLiab     string `json:"crossLiab"`
	IsoLiab       string `json:"isoLiab"`
	CoinUsdPrice  string `json:"coinUsdPrice"`
	StgyEq        string `json:"stgyEq"`
	IsoUpl        string `json:"isoUpl"`
}

// WsAccounts handle websocket instrument message
type WsAccountsHandler func(event *WsAccountsEvent)

func WsAccountsServe(ccy string, apikey string, apisecret string, passphrase string, handler WsAccountsHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := getWsEndpoint(true) // get private endpoint
	return wsAccountsServe(endpoint, ccy, apikey, apisecret, passphrase, ccy, handler, errHandler)
}

// WsAccountsServe serve websocket
func wsAccountsServe(endpoint string, ccy string, apiKey string, secretKey string, passPhrase string, instType string, handler WsAccountsHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	arg := map[string]string{
		"channel": "account",
	}
	if ccy != "" {
		arg["ccy"] = ccy
	}
	var args []map[string]string
	args = append(args, arg)
	reqData := ReqData{Op: "subscribe",
		Args: args,
	}
	//fmt.Println(reqData)
	cfg := newWsConfig(endpoint, reqData, apiKey, secretKey, passPhrase)
	wsHandler := func(message []byte) {
		event := new(WsAccountsEvent)
		err = json.Unmarshal(message, event)
		if err != nil {
			errHandler(err)
			return
		}

		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// POSITIONS WEBSOCKET (PRIVATE)

type WsPositionsEvent struct {
	Arg  map[string]string `json:"arg"`
	Data []*PositionDetail `json:"data"`
}

// WsPositions handle websocket instrument message
type WsPositionsHandler func(event *WsPositionsEvent)

func WsPositionsServe(instType string, uly string, instId string, apikey string, apisecret string, passphrase string, handler WsPositionsHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := getWsEndpoint(true) // get private endpoint
	return wsPositionsServe(endpoint, instType, uly, instId, apikey, apisecret, passphrase, handler, errHandler)
}

// WsAccountsServe serve websocket
func wsPositionsServe(endpoint string, instType string, uly string, InstId string, apiKey string, secretKey string, passPhrase string, handler WsPositionsHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	arg := map[string]string{
		"channel":  "positions",
		"instType": instType,
	}
	if uly != "" {
		arg["uly"] = uly
	}
	if InstId != "" {
		arg["InstId"] = InstId
	}
	var args []map[string]string
	args = append(args, arg)
	reqData := ReqData{Op: "subscribe",
		Args: args,
	}
	//fmt.Println(reqData)
	cfg := newWsConfig(endpoint, reqData, apiKey, secretKey, passPhrase)
	wsHandler := func(message []byte) {
		event := new(WsPositionsEvent)
		err = json.Unmarshal(message, event)
		if err != nil {
			errHandler(err)
			return
		}

		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}
