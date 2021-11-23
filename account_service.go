package okex

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetBalanceService get account balance
type GetBalanceService struct {
	c   *Client
	ccy *string
}

// Set currency ccy
func (s *GetBalanceService) Currencies(ccy string) *GetBalanceService {
	s.ccy = &ccy
	return s
}

// Do send request
func (s *GetBalanceService) Do(ctx context.Context, opts ...RequestOption) (res *GetBalanceServiceResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/account/balance",
		secType:  secTypeSigned,
	}

	if s.ccy != nil {
		r.setParam("ccy", *s.ccy)
		//r.query.Add("ccy", *s.ccy)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(GetBalanceServiceResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Balance define user balance of your account
type GetBalanceServiceResponse struct {
	Code string     `json:"code"`
	Data []*Balance `json:"data"`
	Msg  string     `json:"msg"`
}

// Balance define account info
type BalanceDetail struct {
	AvailBal      string `json:"availBal"`
	AvailEq       string `json:"availEq"`
	CashBal       string `json:"cashBal"`
	Ccy           string `json:"ccy"`
	CrossLiab     string `json:"crossLiab"`
	DisEq         string `json:"disEq"`
	Eq            string `json:"eq"`
	EqUsd         string `json:"eqUsd"`
	FrozenBal     string `json:"frozenBal"`
	Interest      string `json:"interest"`
	IsoEq         string `json:"isoEq"`
	IsoLiab       string `json:"isoLiab"`
	Liab          string `json:"liab"`
	MaxLoan       string `json:"maxLoan"`
	MgnRatio      string `json:"mgnRatio"`
	NotionalLever string `json:"notionalLever"`
	OrdFrozen     string `json:"ordFrozen"`
	Twap          string `json:"twap"`
	UTime         string `json:"uTime"`
	Upl           string `json:"upl"`
	PplLiab       string `json:"uplLiab"`
	StgyEq        string `json:"stgyEq"`
}

type Balance struct {
	AdjEq       string           `json:"adjEq"`
	Details     []*BalanceDetail `json:"details"`
	Imr         string           `json:"imr"`
	IsoEq       string           `json:"isoEq"`
	MgnRatio    string           `json:"mgnRatio"`
	Mnr         string           `json:"mnr"`
	NotionalUsd string           `json:"notionalUsd"`
	OrdFroz     string           `json:"ordFroz"`
	TotalEq     string           `json:"totalEq"`
	UTime       string           `json:"uTime"`
}

// GetPositionsService
type GetPositionsService struct {
	c        *Client
	instType *string
	instId   *string
	posId    *string
}

// Set Instrument Type
func (s *GetPositionsService) InstrumentType(instType string) *GetPositionsService {
	s.instType = &instType
	return s
}

// Set Instrument Id
func (s *GetPositionsService) InstrumentId(instId string) *GetPositionsService {
	s.instId = &instId
	return s
}

// Set Position Id
func (s *GetPositionsService) PositionId(posId string) *GetPositionsService {
	s.posId = &posId
	return s
}

// Do send request
func (s *GetPositionsService) Do(ctx context.Context, opts ...RequestOption) (res *GetPositionsServiceResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/account/positions",
		secType:  secTypeSigned,
	}

	if s.instType != nil {
		r.setParam("instType", *s.instType)
	}
	if s.instId != nil {
		r.setParam("instId", *s.instId)
	}
	if s.posId != nil {
		r.setParam("posId", *s.posId)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(GetPositionsServiceResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetPositionsServiceResponse struct as https://www.okex.com/docs-v5/en/#rest-api-account-get-positions
type GetPositionsServiceResponse struct {
	Code string            `json:"code"`
	Data []*PositionDetail `json:"data"`
	Msg  string            `json:"msg"`
}

// Position detail inside the positions
type PositionDetail struct {
	Adl         string `json:"adl"`
	AvailPos    string `json:"availPos"`
	AvgPx       string `json:"avgPx"`
	CTime       string `json:"cTime"`
	Ccy         string `json:"ccy"`
	DeltaBS     string `json:"deltaBS"`
	DeltaPA     string `json:"deltaPA"`
	GammaBS     string `json:"gammaBS"`
	GammaPA     string `json:"gammaPA"`
	Imr         string `json:"imr"`
	InstId      string `json:"instId"`
	InstType    string `json:"instType"`
	Interest    string `json:"interest"`
	Last        string `json:"last"`
	Lever       string `json:"lever"`
	Liab        string `json:"liab"`
	LiabCcy     string `json:"liabCcy"`
	LiqPx       string `json:"liqPx"`
	Margin      string `json:"margin"`
	MgnMode     string `json:"mgnMode"`
	MgnRatio    string `json:"mgnRatio"`
	Mmr         string `json:"mmr"`
	NotionalCcy string `json:"notionalCcy"` // this is for AccountAndPositionRisk
	NotionalUsd string `json:"notionalUsd"`
	OptVal      string `json:"optVal"`
	PTime       string `json:"pTime"`
	Pos         string `json:"pos"`
	PosCcy      string `json:"posCcy"`
	PosId       string `json:"posId"`
	PosSide     string `json:"posSide"`
	ThetaBS     string `json:"thetaBS"`
	TradeId     string `json:"tradeId"`
	UTime       string `json:"uTime"`
	Upl         string `json:"upl"`
	UplRatio    string `json:"uplRatio"`
	VegaBS      string `json:"vegaBS"`
	VegaPA      string `json:"vegaPA"`
}

// GetAccountAndPositionRiskService
type GetAccountAndPositionRiskService struct {
	c        *Client
	instType *string
}

// Do send request
func (s *GetAccountAndPositionRiskService) Do(ctx context.Context, opts ...RequestOption) (res *GetAccountAndPositionRiskServiceResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/account/account-position-risk",
		secType:  secTypeSigned,
	}

	if s.instType != nil {
		r.setParam("instType", *s.instType)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(GetAccountAndPositionRiskServiceResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Response to GetAccountAndPositionRiskService
type GetAccountAndPositionRiskServiceResponse struct {
	Code string                    `json:"code"`
	Data []*AccountAndPositionRisk `json:"data"`
	Msg  string                    `json:"msg"`
}

type AccountAndPositionRisk struct {
	Ts      string           `json:"ts"`
	AdjEq   string           `json:"adjEq"`
	BalData []BalanceDetail  `json:"balData"`
	PosData []PositionDetail `json:"posData"`
}

// GetAccountConfigurationService
type GetAccountConfigurationService struct {
	c *Client
}

// Do send request
func (s *GetAccountConfigurationService) Do(ctx context.Context, opts ...RequestOption) (res *GetAccountConfigurationServiceResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/account/config",
		secType:  secTypeSigned,
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(GetAccountConfigurationServiceResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Response to GetAccountConfigurationService
type GetAccountConfigurationServiceResponse struct {
	Code string                  `json:"code"`
	Data []*AccountConfiguration `json:"data"`
	Msg  string                  `json:"msg"`
}

type AccountConfiguration struct {
	Uid        string `json:"udi"`
	AcctLv     string `json:"acctLc"`
	PosMode    string `json:"posMode"`
	AutoLoan   bool   `json:"autoLoan"`
	GreeksType string `json:"greeksType"`
	Level      string `json:"level"`
	LevelTmp   string `json:"levelTmp"`
}

// SetAccountPositionModeService
type SetAccountPositionModeService struct {
	c       *Client
	posMode string
}

// Set pos Mode 'long_short_mode' or 'net_mode'
func (s *SetAccountPositionModeService) PosMode(posMode string) *SetAccountPositionModeService {
	s.posMode = posMode
	return s
}

// Do send request
func (s *SetAccountPositionModeService) Do(ctx context.Context, opts ...RequestOption) (res *SetAccountPositionModeServiceResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/api/v5/account/set-position-mode",
		secType:  secTypeSigned,
	}

	r.setBodyParam("posMode", s.posMode)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SetAccountPositionModeServiceResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Response to SetAccountPositionModeServiceService
type SetAccountPositionModeServiceResponse struct {
	Code string         `json:"code"`
	Data []*interface{} `json:"data"`
	Msg  string         `json:"msg"`
}

// // GetAccountSnapshotService all account orders; active, canceled, or filled
// type GetAccountSnapshotService struct {
// 	c           *Client
// 	accountType string
// 	startTime   *int64
// 	endTime     *int64
// 	limit       *int
// }

// // Type set account type ("SPOT", "MARGIN", "FUTURES")
// func (s *GetAccountSnapshotService) Type(accountType string) *GetAccountSnapshotService {
// 	s.accountType = accountType
// 	return s
// }

// // StartTime set starttime
// func (s *GetAccountSnapshotService) StartTime(startTime int64) *GetAccountSnapshotService {
// 	s.startTime = &startTime
// 	return s
// }

// // EndTime set endtime
// func (s *GetAccountSnapshotService) EndTime(endTime int64) *GetAccountSnapshotService {
// 	s.endTime = &endTime
// 	return s
// }

// // Limit set limit
// func (s *GetAccountSnapshotService) Limit(limit int) *GetAccountSnapshotService {
// 	s.limit = &limit
// 	return s
// }

// // Do send request
// func (s *GetAccountSnapshotService) Do(ctx context.Context, opts ...RequestOption) (res *Snapshot, err error) {
// 	r := &request{
// 		method:   http.MethodGet,
// 		endpoint: "/sapi/v1/accountSnapshot",
// 		secType:  secTypeSigned,
// 	}
// 	r.setParam("type", s.accountType)

// 	if s.startTime != nil {
// 		r.setParam("startTime", *s.startTime)
// 	}
// 	if s.endTime != nil {
// 		r.setParam("endTime", *s.endTime)
// 	}
// 	if s.limit != nil {
// 		r.setParam("limit", *s.limit)
// 	}
// 	data, err := s.c.callAPI(ctx, r, opts...)
// 	if err != nil {
// 		return &Snapshot{}, err
// 	}
// 	res = new(Snapshot)
// 	err = json.Unmarshal(data, &res)
// 	if err != nil {
// 		return &Snapshot{}, err
// 	}
// 	return res, nil
// }

// // Snapshot define snapshot
// type Snapshot struct {
// 	Code     int            `json:"code"`
// 	Msg      string         `json:"msg"`
// 	Snapshot []*SnapshotVos `json:"snapshotVos"`
// }

// // SnapshotVos define content of a snapshot
// type SnapshotVos struct {
// 	Data       *SnapshotData `json:"data"`
// 	Type       string        `json:"type"`
// 	UpdateTime int64         `json:"updateTime"`
// }

// // SnapshotData define content of a snapshot
// type SnapshotData struct {
// 	MarginLevel         string `json:"marginLevel"`
// 	TotalAssetOfBtc     string `json:"totalAssetOfBtc"`
// 	TotalLiabilityOfBtc string `json:"totalLiabilityOfBtc"`
// 	TotalNetAssetOfBtc  string `json:"totalNetAssetOfBtc"`

// 	Balances   []*SnapshotBalances   `json:"balances"`
// 	UserAssets []*SnapshotUserAssets `json:"userAssets"`
// 	Assets     []*SnapshotAssets     `json:"assets"`
// 	Positions  []*SnapshotPositions  `json:"position"`
// }

// // SnapshotBalances define snapshot balances
// type SnapshotBalances struct {
// 	Asset  string `json:"asset"`
// 	Free   string `json:"free"`
// 	Locked string `json:"locked"`
// }

// // SnapshotUserAssets define snapshot user assets
// type SnapshotUserAssets struct {
// 	Asset    string `json:"asset"`
// 	Borrowed string `json:"borrowed"`
// 	Free     string `json:"free"`
// 	Interest string `json:"interest"`
// 	Locked   string `json:"locked"`
// 	NetAsset string `json:"netAsset"`
// }

// // SnapshotAssets define snapshot assets
// type SnapshotAssets struct {
// 	Asset         string `json:"asset"`
// 	MarginBalance string `json:"marginBalance"`
// 	WalletBalance string `json:"walletBalance"`
// }

// // SnapshotPositions define snapshot positions
// type SnapshotPositions struct {
// 	EntryPrice       string `json:"entryPrice"`
// 	MarkPrice        string `json:"markPrice"`
// 	PositionAmt      string `json:"positionAmt"`
// 	Symbol           string `json:"symbol"`
// 	UnRealizedProfit string `json:"unRealizedProfit"`
// }
