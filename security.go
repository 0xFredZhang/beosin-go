package beosin

const (
	urlBlackAddressScreening = "/api/v1/tag/black/screening"
)

func (c *Beosin) BlackAddressScreening(platform, address string) (resp BlackAddressResp, err error) {
	_, err = c.client.R().
		SetHeader(HeaderContentType, ContentTypeJson).
		SetHeader(HeaderAppID, c.appId).
		SetHeader(HeaderAppSecret, c.appSecret).
		SetQueryParam("platform", platform).
		SetQueryParam("address", address).
		SetResult(&resp).
		Get(c.host + urlBlackAddressScreening)
	return
}

type BlackAddressResp struct {
	Code float64 `json:"code"`
	Msg  string  `json:"msg"`
	Data struct {
		Lawsuit            bool `json:"lawsuit"`
		Piracy             bool `json:"piracy"`
		Mixing             bool `json:"mixing"`
		Gambling           bool `json:"gambling"`
		Scam               bool `json:"scam"`
		Darknet            bool `json:"darknet"`
		Trojan             bool `json:"trojan"`
		Theft              bool `json:"theft"`
		Terrorist          bool `json:"terrorist"`
		Drug               bool `json:"drug"`
		Hacker             bool `json:"hacker"`
		BusinessBlackList  bool `json:"businessBlackList"`
		Sanction           bool `json:"sanction"`
		ChildAbuseMaterial bool `json:"childAbuseMaterial"`
		Ransomware         bool `json:"ransomware"`
	} `json:"data"`
}
