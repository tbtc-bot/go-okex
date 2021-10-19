package okex

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
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

// Set instrument Id
func (s *CreateOrderService) InstrumentId(instId string) *CreateOrderService {
	s.instId = instId
	return s
}

// Set trade mode
func (s *CreateOrderService) TradeMode(tdMode string) *CreateOrderService {
	s.tdMode = tdMode
	return s
}

// Set Currency
func (s *CreateOrderService) Currency(ccy string) *CreateOrderService {
	s.ccy = &ccy
	return s
}

// Set Client Order Id
func (s *CreateOrderService) ClientOrderId(clOrdId string) *CreateOrderService {
	s.clOrdId = &clOrdId
	return s
}

// Set Tag field
func (s *CreateOrderService) Tag(tag string) *CreateOrderService {
	s.tag = &tag
	return s
}

// Set side
func (s *CreateOrderService) Side(side string) *CreateOrderService {
	s.side = side
	return s
}

// Set position side
func (s *CreateOrderService) PositionSide(posSide string) *CreateOrderService {
	s.posSide = posSide
	return s
}

// Set order type
func (s *CreateOrderService) OrderType(ordType string) *CreateOrderService {
	s.ordType = ordType
	return s
}

// Set size
func (s *CreateOrderService) Size(sz string) *CreateOrderService {
	s.sz = sz
	return s
}

// Set Order Price
func (s *CreateOrderService) OrderPrice(px string) *CreateOrderService {
	s.px = &px
	return s
}

// Set ReduceOnly
func (s *CreateOrderService) ReduceOnly(reduceOnly bool) *CreateOrderService {
	s.reduceOnly = &reduceOnly
	return s
}

// Set Quantity Type
func (s *CreateOrderService) QuantityType(tgtCcy string) *CreateOrderService {
	s.tgtCcy = &tgtCcy
	return s
}

// Do send request
func (s *CreateOrderService) Do(ctx context.Context, opts ...RequestOption) (res *Orders, err error) {
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
	res = new(Orders)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Orders structure
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
