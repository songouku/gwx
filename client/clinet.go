package client

import (
	"crypto/sha1"
	"encoding/hex"
	"sort"
)

type WxConfig struct {
	AppId  string `json:"appId"`
	Secret string `json:"secret"`
	Token  string `json:"token"`
}

func NewConfig(appId string, secret string, token string) *WxConfig {
	return &WxConfig{
		AppId:  appId,
		Secret: secret,
		Token:  token,
	}
}

func (c *WxConfig) Sign(param []string, signature string) bool {
	param = append(param, c.Token)
	sort.Strings(param)
	var str string
	for i := range param {
		str = str + param[i]
	}
	valid := sha1.Sum([]byte(str))
	sign := hex.EncodeToString(valid[:])
	return sign == signature
}
