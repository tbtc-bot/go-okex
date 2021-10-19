package okex

import (
	"context"
	"log"
	"time"

	//"github.com/gorilla/websocket"
	. "github.com/tbtc-bot/go-okex/interface"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

// WsHandler handle raw websocket message
type WsHandler func(message []byte)

// ErrHandler handles errors
type ErrHandler func(err error)

// WsConfig webservice configuration
type WsConfig struct {
	Endpoint string
	WsOp     WSReqData
}

func newWsConfig(endpoint string, wsop WSReqData) *WsConfig {
	return &WsConfig{
		Endpoint: endpoint,
		WsOp:     wsop,
	}
}

// websocket manager for public endpoint
var wsServe = func(cfg *WsConfig, handler WsHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	ctx, cancel := context.WithCancel(context.Background())
	c, _, err := websocket.Dial(ctx, cfg.Endpoint, nil)
	if err != nil {
		log.Fatal("Error to dial websocket: ", err)
		cancel()
		return nil, nil, err
	}
	c.SetReadLimit(655350)
	err = wsjson.Write(ctx, c, []byte(cfg.WsOp.ToString()))
	if err != nil {
		log.Fatal("Error sending Op to websocket: ", err)
	}
	//c.WriteMessage(websocket.TextMessage, []byte(cfg.WsOp.ToString()))
	doneC = make(chan struct{})
	stopC = make(chan struct{})

	go func() {
		// This function will exit either on error from
		// websocket.Conn.ReadMessage or when the stopC channel is
		// closed by the client.
		defer close(doneC)
		defer cancel()
		if WebsocketKeepalive {
			go keepAlive(ctx, c, WebsocketTimeout)
		}
		// Wait for the stopC channel to be closed.  We do that in a
		// separate goroutine because ReadMessage is a blocking
		// operation.
		silent := false
		go func() {
			select {
			case <-stopC:
				silent = true
			case <-doneC:
			}
			_ = c.Close(websocket.StatusNormalClosure, "normal closure")
		}()
		for {
			_, message, readErr := c.Read(ctx)
			if readErr != nil {
				if !silent {
					errHandler(readErr)
				}
				return
			}
			handler(message)
		}
	}()
	return
}

func keepAlive(ctx context.Context, c *websocket.Conn, d time.Duration) {
	t := time.NewTimer(d)
	defer t.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-t.C:
		}

		err := c.Ping(ctx)
		if err != nil {
			return
		}

		t.Reset(d)
	}
}
