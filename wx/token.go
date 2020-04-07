package wx

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/songouku/gwx/constant"
	"github.com/songouku/gwx/model"
	"github.com/songouku/gwx/util"
)

func (c *WxConfig) GetToken() (*model.Token, error) {
	params := map[string]interface{}{
		"grant_type": "client_credential",
		"appid":      c.AppId,
		"secret":     c.Secret,
	}
	res, err := util.Get(params, constant.Token)
	if err != nil {
		return nil, errors.New("")
	}
	fmt.Printf("result is %s", res)
	var token model.Token
	err = json.Unmarshal([]byte(res), &token)
	if err != nil {
		fmt.Errorf("decode failed , error is %v", err)
		return nil, errors.New(err.Error())
	}
	return &token, nil
}