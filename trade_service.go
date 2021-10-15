package okex

import (
	"context"
	"encoding/json"
	"net/http"
)

// START GET BALANCE CODE
type CreateOrderService struct {
	c          *Client
	instId     string
	tdMode     string
	ccy        *string
	clOrdId    *string
	tag        *string
	side       string
	posSide    string
	ordType    string
	sz         string
	px         *string
	reduceOnly *bool
	tgtCcy     *string
}

// Do send request
func (s *CreateOrderService) Do(ctx context.Context, opts ...RequestOption) (res *Orders, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/api/v5/trade/order",
		secType:  secTypeSigned,
	}

	r.setParam("instId", s.instId)
	r.setParam("tdMode", s.tdMode)
	r.setParam("side", s.side)
	r.setParam("posSide", s.posSide)
	r.setParam("ordType", s.ordType)
	r.setParam("sz", s.sz)
	if s.ccy != nil {
		r.setParam("ccy", *s.ccy)
	}
	if s.clOrdId != nil {
		r.setParam("clOrdId", *s.clOrdId)
	}
	if s.tag != nil {
		r.setParam("tag", *s.tag)
	}
	if s.px != nil {
		r.setParam("px", *s.px)
	}
	if s.reduceOnly != nil {
		r.setParam("reduceOnly", *s.reduceOnly)
	}
	if s.tgtCcy != nil {
		r.setParam("tgtCcy", *s.tgtCcy)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(Orders)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// // Set currency ccy
// func (s *GetBalanceService) Currencies(ccy string) *GetBalanceService {
// 	s.ccy = ccy
// 	return s
// }

// Ordets structure

type Orders struct {
	Code string         `json:"code"`
	Msg  string         `json:"msg"`
	Data []*OrderDetail `json:"data"`
}
type OrderDetail struct {
	OrdId   string `json:"ordId"`
	ClOrdId string `json:"clOrdId"`
	Tag     string `json:"tag"`
	SCode   string `json:"sCode"`
	SMsg    string `json:"sMsg"`
}
