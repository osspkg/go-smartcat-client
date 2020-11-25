package smartcatclient

import (
	"net/http"
)

//go:generate easyjson

const (
	uriAccount          = "/api/integration/v1/account"
	uriAccountMTEngines = "/api/integration/v1/account/mtengines"
)

//easyjson:json
type (
	Account struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		IsPersonal bool   `json:"isPersonal"`
		Type       string `json:"type"`
	}
	AccountMTEngine struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
	AccountMTEngines []AccountMTEngine
)

//Account Receiving the account details
func (v *Client) Account() (o Account, err error) {
	err, _ = v.call(http.MethodGet, uriAccount, nil, &o)
	return
}

//AccountMtengines Receiving MT engines available for the account
func (v *Client) AccountMTEngines() (o AccountMTEngines, err error) {
	err, _ = v.call(http.MethodGet, uriAccountMTEngines, nil, &o)
	return
}