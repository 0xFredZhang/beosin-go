package beosin

const (
	urlAccountBalance = "/api/v1/package/info"
)

func (c *Beosin) AccountBalanceQuery() (resp AccountBalanceResp, err error) {
	_, err = c.client.R().
		SetHeader(HeaderContentType, ContentTypeJson).
		SetHeader(HeaderAppID, c.appId).
		SetHeader(HeaderAppSecret, c.appSecret).
		SetResult(&resp).
		Get(c.host + urlAccountBalance)
	return
}

type AccountBalanceResp struct {
	Code float64 `json:"code"`
	Msg  string  `json:"msg"`
	Data struct {
		SurplusIntegral int64 `json:"surplusIntegral"`
		EquityStartDate int64 `json:"equityStartDate"`
		EquityEndDate   int64 `json:"equityEndDate"`
	} `json:"data"`
}
