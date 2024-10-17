package beosin

const (
	urlDepositTransactionAssessment    = "/api/v2/kyt/tx/deposit"
	urlWithdrawalTransactionAssessment = "/api/v2/kyt/tx/withdraw"
	urlEOAAddressRiskAssessment        = "/api/v2/kyt/address/risk"
	urlTokenRiskAssessment             = "/api/v2/kyt/token/risk"
	urlMaliciousAddressQuery           = "/api/v2/kyt/tag/malicious"
	urlVASPQuery                       = "/api/v2/kyt/tag/vasp"
	urlEntityQuery                     = "/api/v2/kyt/tag/entity"
)

func (c *Beosin) DepositTransactionAssessment(chainId, txHash string) (resp TransactionAssessmentResp, err error) {
	_, err = c.client.R().
		SetHeader(HeaderContentType, ContentTypeJson).
		SetHeader(HeaderAppID, c.appId).
		SetHeader(HeaderAppSecret, c.appSecret).
		SetQueryParam("chainId", chainId).
		SetQueryParam("hash", txHash).
		SetResult(&resp).
		Get(c.host + urlDepositTransactionAssessment)
	return
}

func (c *Beosin) WithdrawalTransactionAssessment(chainId, txHash string) (resp TransactionAssessmentResp, err error) {
	_, err = c.client.R().
		SetHeader(HeaderContentType, ContentTypeJson).
		SetHeader(HeaderAppID, c.appId).
		SetHeader(HeaderAppSecret, c.appSecret).
		SetQueryParam("chainId", chainId).
		SetQueryParam("hash", txHash).
		SetResult(&resp).
		Get(c.host + urlWithdrawalTransactionAssessment)
	return
}

func (c *Beosin) EOAAddressRiskAssessment(chainId, address string) (resp EOAAddressRiskResp, err error) {
	_, err = c.client.R().
		SetHeader(HeaderContentType, ContentTypeJson).
		SetHeader(HeaderAppID, c.appId).
		SetHeader(HeaderAppSecret, c.appSecret).
		SetQueryParam("chainId", chainId).
		SetQueryParam("address", address).
		SetResult(&resp).
		Get(c.host + urlEOAAddressRiskAssessment)
	return
}

func (c *Beosin) TokenRiskAssessment(chainId, address string) (resp TokenRiskResp, err error) {
	_, err = c.client.R().
		SetHeader(HeaderContentType, ContentTypeJson).
		SetHeader(HeaderAppID, c.appId).
		SetHeader(HeaderAppSecret, c.appSecret).
		SetQueryParam("chainId", chainId).
		SetQueryParam("address", address).
		SetResult(&resp).
		Get(c.host + urlTokenRiskAssessment)
	return
}

func (c *Beosin) MaliciousAddressQuery(chainId, address string) (resp MaliciousAddressResp, err error) {
	_, err = c.client.R().
		SetHeader(HeaderContentType, ContentTypeJson).
		SetHeader(HeaderAppID, c.appId).
		SetHeader(HeaderAppSecret, c.appSecret).
		SetQueryParam("chainId", chainId).
		SetQueryParam("address", address).
		SetResult(&resp).
		Get(c.host + urlMaliciousAddressQuery)
	return
}

func (c *Beosin) VASPQuery(chainId, address string) (resp VASPResp, err error) {
	_, err = c.client.R().
		SetHeader(HeaderContentType, ContentTypeJson).
		SetHeader(HeaderAppID, c.appId).
		SetHeader(HeaderAppSecret, c.appSecret).
		SetQueryParam("chainId", chainId).
		SetQueryParam("address", address).
		SetResult(&resp).
		Get(c.host + urlVASPQuery)
	return
}

func (c *Beosin) EntityQuery(chainId, address string) (resp EntityResp, err error) {
	_, err = c.client.R().
		SetHeader(HeaderContentType, ContentTypeJson).
		SetHeader(HeaderAppID, c.appId).
		SetHeader(HeaderAppSecret, c.appSecret).
		SetQueryParam("chainId", chainId).
		SetQueryParam("address", address).
		SetResult(&resp).
		Get(c.host + urlEntityQuery)
	return
}

type TransactionAssessmentResp struct {
	Code float64 `json:"code"`
	Msg  string  `json:"msg"`
	Data struct {
		Score     float64 `json:"score"`
		RiskLevel string  `json:"riskLevel"`
		Risks     []struct {
			RiskStrategy string `json:"riskStrategy"`
			RiskDetails  []struct {
				RiskName string  `json:"riskName"`
				Rate     float64 `json:"rate"`
				Amount   float64 `json:"amount"`
			} `json:"riskDetails"`
		} `json:"risks"`
	} `json:"data"`
}

type EOAAddressRiskResp struct {
	Code float64 `json:"code"`
	Msg  string  `json:"msg"`
	Data struct {
		Score          float64 `json:"score"`
		RiskLevel      string  `json:"riskLevel"`
		IncomingScore  float64 `json:"incomingScore"`
		IncomingLevel  string  `json:"incomingLevel"`
		IncomingDetail []struct {
			StrategyName string `json:"strategyName"`
			RiskDetails  []struct {
				RiskName string  `json:"riskName"`
				Rate     float64 `json:"rate"`
				Amount   float64 `json:"amount"`
			} `json:"riskDetails"`
		} `json:"incomingDetail"`
		OutgoingScore  float64 `json:"outgoingScore"`
		OutgoingLevel  string  `json:"outgoingLevel"`
		OutgoingDetail []struct {
			StrategyName string `json:"strategyName"`
			RiskDetails  []struct {
				RiskName string  `json:"riskName"`
				Rate     float64 `json:"rate"`
				Amount   float64 `json:"amount"`
			} `json:"riskDetails"`
		} `json:"outgoingDetail"`
	} `json:"data"`
}

type TokenRiskResp struct {
	Code float64 `json:"code"`
	Msg  string  `json:"msg"`
	Data struct {
		ContractName         string `json:"contractName"`
		RiskLevel            string `json:"riskLevel"`
		IsFrailty            bool   `json:"isFrailty"`
		DeployerIsBlack      bool   `json:"deployerIsBlack"`
		OwnerIsBlack         bool   `json:"ownerIsBlack"`
		IsMintable           bool   `json:"isMintable"`
		CanTakeBackOwnership bool   `json:"canTakeBackOwnership"`
		SelfDestruct         bool   `json:"selfDestruct"`
		ExternalCall         bool   `json:"externalCall"`
		GasAbuse             bool   `json:"gasAbuse"`
		IsAntiWhale          bool   `json:"isAntiWhale"`
		SlippageModifiable   bool   `json:"slippageModifiable"`
		OwnerChangeBalance   bool   `json:"ownerChangeBalance"`
		HiddenOwner          bool   `json:"hiddenOwner"`
	} `json:"data"`
}

type MaliciousAddressResp struct {
	Code float64 `json:"code"`
	Msg  string  `json:"msg"`
	Data struct {
		Address      string `json:"address"`
		IsMalicious  bool   `json:"isMalicious"`
		MaliceDetail struct {
			Source     string `json:"source"`
			MaliceTags []struct {
				TagType string `json:"tagType"`
				Tag     string `json:"tag"`
			} `json:"maliceTags"`
		} `json:"maliceDetail"`
		IsSanction     bool `json:"isSanction"`
		SanctionDetail struct {
			Standard string `json:"standard"`
			Tag      string `json:"tag"`
			Entity   string `json:"entity"`
			Country  string `json:"country"`
			Source   string `json:"source"`
		} `json:"sanctionDetail"`
		IsInCustomerBlackList bool `json:"isInCustomerBlackList"`
	} `json:"data"`
}

type VASPResp struct {
	Code float64 `json:"code"`
	Msg  string  `json:"msg"`
	Data struct {
		Address  string   `json:"address"`
		IsVasp   bool     `json:"isVasp"`
		VaspTags []string `json:"vaspTags"`
	} `json:"data"`
}

type EntityResp struct {
	Code float64 `json:"code"`
	Msg  string  `json:"msg"`
	Data struct {
		Address  string   `json:"address"`
		Entities []string `json:"entities"`
	} `json:"data"`
}
