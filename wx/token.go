package wx

import (
	"encoding/json"
	"github.com/songouku/gwx/constant"
	"github.com/songouku/gwx/model"
	"github.com/songouku/gwx/util"
)

func (c *Config) GetToken() (*model.Token, error) {
	params := map[string]interface{}{
		"grant_type": "client_credential",
		"appid":      c.AppId,
		"secret":     c.Secret,
	}
	res, err := util.Get(constant.Token, params)
	if err != nil {
		return nil, err
	}
	var token model.Token
	err = json.Unmarshal(res, &token)
	if err != nil {
		return nil, err
	}
	return &token, nil
}
