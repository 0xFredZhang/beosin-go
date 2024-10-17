package beosin

import (
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func getBeosin() Beosin {
	appId, appSecret := os.Getenv("BeosinAppId"), os.Getenv("BeosinAppSecret")
	return NewBeosin(appId, appSecret)
}

func TestDepositTransactionAssessment(t *testing.T) {
	data, err := getBeosin().DepositTransactionAssessment("1", "0xb9ac975f438de40b09cf683acc7e3e319c2a1822b0f7d2f1b0de2bbd377658f9")
	assert.Nil(t, err)
	spew.Dump(data)
}

func TestWithdrawalTransactionAssessment(t *testing.T) {
	data, err := getBeosin().WithdrawalTransactionAssessment("1", "0x27b4cd34f43f70a6ef5737376b1621defa1a3bb82845864c94992d157d2fa4f5")
	assert.Nil(t, err)
	spew.Dump(data)
}

func TestEOAAddressRiskAssessment(t *testing.T) {
	data, err := getBeosin().EOAAddressRiskAssessment("1", "0x4fe666ecc5263f5dbb34adb8bc1c8cbb9bbcd1cc")
	assert.Nil(t, err)
	spew.Dump(data)
}

func TestTokenRiskAssessment(t *testing.T) {
	data, err := getBeosin().TokenRiskAssessment("56", "0xea51801b8f5b88543ddad3d1727400c15b209d8f")
	assert.Nil(t, err)
	spew.Dump(data)
}

func TestMaliciousAddressQuery(t *testing.T) {
	data, err := getBeosin().MaliciousAddressQuery("1", "0x901bb9583b24d97e995513c6778dc6888ab6870e")
	assert.Nil(t, err)
	spew.Dump(data)
}

func TestVASPQuery(t *testing.T) {
	data, err := getBeosin().VASPQuery("1", "0xec6ad3cb0e62cd7c8e75d2919f12c3195d998002")
	assert.Nil(t, err)
	spew.Dump(data)
}

func TestEntityQuery(t *testing.T) {
	data, err := getBeosin().EntityQuery("1", "0x58820fbd8e7ac77145360318432a1ce12502f553")
	assert.Nil(t, err)
	spew.Dump(data)
}

func TestAccountBalanceQuery(t *testing.T) {
	data, err := getBeosin().AccountBalanceQuery()
	assert.Nil(t, err)
	spew.Dump(data)
}

func TestBlackAddressScreen(t *testing.T) {
	data, err := getBeosin().BlackAddressScreening("bsc", "0x3cffd56b47b7b41c56258d9c7731abadc360e073")
	assert.Nil(t, err)
	spew.Dump(data)
}
