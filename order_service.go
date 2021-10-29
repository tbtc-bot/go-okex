package okex

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

// PlaceOrderService places a single order
type PlaceOrderService struct {
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

// Set instrument Id
func (s *PlaceOrderService) InstrumentId(instId string) *PlaceOrderService {
	s.instId = instId
	return s
}

// Set trade mode
func (s *PlaceOrderService) TradeMode(tdMode string) *PlaceOrderService {
	s.tdMode = tdMode
	return s
}

// Set Currency
func (s *PlaceOrderService) Currency(ccy string) *PlaceOrderService {
	s.ccy = &ccy
	return s
}

// Set Client Order Id
func (s *PlaceOrderService) ClientOrderId(clOrdId string) *PlaceOrderService {
	s.clOrdId = &clOrdId
	return s
}

// Set Tag field
func (s *PlaceOrderService) Tag(tag string) *PlaceOrderService {
	s.tag = &tag
	return s
}

// Set side
func (s *PlaceOrderService) Side(side string) *PlaceOrderService {
	s.side = side
	return s
}

// Set position side
func (s *PlaceOrderService) PositionSide(posSide string) *PlaceOrderService {
	s.posSide = posSide
	return s
}

// Set order type
func (s *PlaceOrderService) OrderType(ordType string) *PlaceOrderService {
	s.ordType = ordType
	return s
}

// Set size
func (s *PlaceOrderService) Size(sz string) *PlaceOrderService {
	s.sz = sz
	return s
}

// Set Order Price
func (s *PlaceOrderService) OrderPrice(px string) *PlaceOrderService {
	s.px = &px
	return s
}

// Set ReduceOnly
func (s *PlaceOrderService) ReduceOnly(reduceOnly bool) *PlaceOrderService {
	s.reduceOnly = &reduceOnly
	return s
}

// Set Quantity Type
func (s *PlaceOrderService) QuantityType(tgtCcy string) *PlaceOrderService {
	s.tgtCcy = &tgtCcy
	return s
}

// Do send request
func (s *PlaceOrderService) Do(ctx context.Context, opts ...RequestOption) (res *PlaceOrderResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/api/v5/trade/order",
		secType:  secTypeSigned,
	}

	r.setBodyParam("instId", s.instId)
	r.setBodyParam("tdMode", s.tdMode)
	r.setBodyParam("side", s.side)
	r.setBodyParam("posSide", s.posSide)
	r.setBodyParam("ordType", s.ordType)
	r.setBodyParam("sz", s.sz)
	if s.ccy != nil {
		r.setBodyParam("ccy", *s.ccy)
	}
	if s.clOrdId != nil {
		r.setBodyParam("clOrdId", *s.clOrdId)
	}
	if s.tag != nil {
		r.setBodyParam("tag", *s.tag)
	}
	if s.px != nil {
		r.setBodyParam("px", *s.px)
	}
	if s.reduceOnly != nil {
		r.setBodyParam("reduceOnly", strconv.FormatBool(*s.reduceOnly))
	}
	if s.tgtCcy != nil {
		r.setBodyParam("tgtCcy", *s.tgtCcy)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(PlaceOrderResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Response to CreateOrderService
type PlaceOrderResponse struct {
	Code string         `json:"code"`
	Msg  string         `json:"msg"`
	Data []*OrderDetail `json:"data"`
}

// CancelOrderService cancel an order
type CancelOrderService struct {
	c       *Client
	instId  string
	ordId   *string
	clOrdId *string
}

// Set instrument Id
func (s *CancelOrderService) InstrumentId(instId string) *CancelOrderService {
	s.instId = instId
	return s
}

// Set order Id
func (s *CancelOrderService) OrderId(ordId string) *CancelOrderService {
	s.ordId = &ordId
	return s
}

// Set order Id
func (s *CancelOrderService) ClientOrderId(clOrdId string) *CancelOrderService {
	s.clOrdId = &clOrdId
	return s
}

// Do send request
func (s *CancelOrderService) Do(ctx context.Context, opts ...RequestOption) (res *CancelOrderResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/api/v5/trade/cancel-order",
		secType:  secTypeSigned,
	}

	r.setBodyParam("instId", s.instId)
	if s.ordId != nil {
		r.setBodyParam("ordId", *s.ordId)
	}
	if s.clOrdId != nil {
		r.setBodyParam("clOrdId", *s.clOrdId)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(CancelOrderResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Response to CancelOrderService
type CancelOrderResponse struct {
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

// OrderListService list of all open orders
type OrderListService struct {
	c        *Client
	instType *string
	uly      *string
	instId   *string
	ordType  *string
	state    *string
	after    *string
	before   *string
	limit    *string
}

// Set intrument type
func (s *OrderListService) InstrumentType(instType string) *OrderListService {
	s.instType = &instType
	return s
}

// Set underlying
func (s *OrderListService) Underlying(uly string) *OrderListService {
	s.uly = &uly
	return s
}

// Set instrument id
func (s *OrderListService) InstrumentId(instId string) *OrderListService {
	s.instId = &instId
	return s
}

// Set order type
func (s *OrderListService) OrderType(ordType string) *OrderListService {
	s.ordType = &ordType
	return s
}

// Set state
func (s *OrderListService) State(state string) *OrderListService {
	s.state = &state
	return s
}

// Set after
func (s *OrderListService) After(after string) *OrderListService {
	s.after = &after
	return s
}

// Set before
func (s *OrderListService) Before(before string) *OrderListService {
	s.before = &before
	return s
}

// Set limit
func (s *OrderListService) Limit(limit string) *OrderListService {
	s.limit = &limit
	return s
}

// Do send request
func (s *OrderListService) Do(ctx context.Context, opts ...RequestOption) (res *OrderListServiceResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/trade/orders-pending",
		secType:  secTypeSigned,
	}

	// TODO these filters do not work!

	if s.instType != nil {
		r.setBodyParam("instType", *s.instType)
	}
	if s.uly != nil {
		r.setBodyParam("uly", *s.uly)
	}
	if s.instId != nil {
		r.setBodyParam("instId", *s.instId)
	}
	if s.ordType != nil {
		r.setBodyParam("ordType", *s.ordType)
	}
	if s.state != nil {
		r.setBodyParam("state", *s.state)
	}
	if s.after != nil {
		r.setBodyParam("after", *s.after)
	}
	if s.before != nil {
		r.setBodyParam("before", *s.before)
	}
	if s.limit != nil {
		r.setBodyParam("limit", *s.limit)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(OrderListServiceResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Response to OrderListService
type OrderListServiceResponse struct {
	Code string             `json:"code"`
	Msg  string             `json:"msg"`
	Data []*OrderListDetail `json:"data"`
}

type OrderListDetail struct {
	AccFillSz   string `json:"accFillSz"`
	AvgPx       string `json:"avgPx"`
	CTime       string `json:"cTime"`
	Category    string `json:"category"`
	Ccy         string `json:"ccy"`
	ClOrdId     string `json:"clOrdId"`
	Fee         string `json:"fee"`
	FeeCcy      string `json:"feeCcy"`
	FillPx      string `json:"fillPx"`
	FillSz      string `json:"fillSz"`
	FillTime    string `json:"fillTime"`
	InstId      string `json:"instId"`
	InstType    string `json:"instType"`
	Lever       string `json:"lever"`
	OrdId       string `json:"ordId"`
	OrdType     string `json:"ordType"`
	Pnl         string `json:"pnl"`
	PosSide     string `json:"posSide"`
	Px          string `json:"px"`
	Rebate      string `json:"rebate"`
	RebateCcy   string `json:"rebateCcy"`
	Side        string `json:"side"`
	SlOrdPx     string `json:"slOrdPx"`
	SlTriggerPx string `json:"slTriggerPx"`
	State       string `json:"state"`
	Sz          string `json:"sz"`
	Tag         string `json:"tag"`
	TgtCcy      string `json:"tgtCcy"`
	TdMode      string `json:"tdMode"`
	TpOrdPx     string `json:"tpOrdPx"`
	TpTriggerPx string `json:"tpTriggerPx"`
	TradeId     string `json:"tradeId"`
	UTime       string `json:"uTime"`
}
