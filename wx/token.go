package wx

import (
	"encoding/json"
	"github.com/songouku/gwx/constant"
	"github.com/songouku/gwx/util"
)

type Token struct {
	AccessToken string `json:"access_token,omitempty"`
	ExpiresIn   int    `json:"expires_in,omitempty"`
	ErrCode     int    `json:"errcode,omitempty"`
	ErrMsg      string `json:"errmsg,omitempty"`
}

func (c *Config) GetToken() (*Token, error) {
	params := map[string]interface{}{
		"grant_type": "client_credential",
		"appid":      c.AppId,
		"secret":     c.Secret,
	}
	res, err := util.Get(constant.Token, params)
	if err != nil {
		return nil, err
	}
	var token Token
	err = json.Unmarshal(res, &token)
	if err != nil {
		return nil, err
	}
	return &token, nil
}
