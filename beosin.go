package beosin

import (
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

// https://beosinofficials-organization.gitbook.io/beosin-api-documentation

type Beosin struct {
	host      string
	appId     string
	appSecret string

	client *resty.Client
}

func NewBeosin(appId, appSecret string) *Beosin {
	client := resty.New()
	client.SetTimeout(30 * time.Second)

	return &Beosin{
		host:      "https://api.beosin.com",
		appId:     appId,
		appSecret: appSecret,
		client:    client,
	}
}

// SetTransport
// Used to set Transport for rate limit
func (c *Beosin) SetTransport(transport http.RoundTripper) {
	c.client.SetTransport(transport)
}

func (c *Beosin) SetTimeout(timeout time.Duration) {
	c.client.SetTimeout(timeout)
}
