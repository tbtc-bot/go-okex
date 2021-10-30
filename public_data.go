package okex

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetInstrumentsService
type GetInstrumentsService struct {
	c        *Client
	instType string
	uly      *string
	instId   *string
}

// Set instrument type
func (s *GetInstrumentsService) InstrumentType(instType string) *GetInstrumentsService {
	s.instType = instType
	return s
}

// Set underlying
func (s *GetInstrumentsService) Underlying(uly string) *GetInstrumentsService {
	s.uly = &uly
	return s
}

// Set instrument id
func (s *GetInstrumentsService) InstrumentId(instId string) *GetInstrumentsService {
	s.instId = &instId
	return s
}

// Do send request
func (s *GetInstrumentsService) Do(ctx context.Context, opts ...RequestOption) (res *GetTickerServiceResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/trade/ticker",
	}

	r.setBodyParam("instId", s.instType)

	if s.uly != nil {
		r.setBodyParam("uly", *s.uly)
	}
	if s.instId != nil {
		r.setBodyParam("instId", *s.instId)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(GetTickerServiceResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Response to GetInstrumentsService
type GetInstrumentsServiceRespone struct {
	Code string              `json:"code"`
	Msg  string              `json:"msg"`
	Data []*InstrumentDetail `json:"data"`
}

type InstrumentDetail struct {
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
	Lever     string `json:"lever"`
	TickSz    string `json:"tickSz"`
	LotSz     string `json:"lotSz"`
	MinSz     string `json:"minSz"`
	CtType    string `json:"ctType"`
	Alias     string `json:"alias"`
	State     string `json:"state"`
}
