package wx

import (
	"crypto/sha1"
	"encoding/hex"
	"sort"
)

type Config struct {
	AppId  string `json:"appId"`
	Secret string `json:"secret"`
	Token  string `json:"token"`
}

func NewConfig(appId string, secret string, token string) *Config {
	return &Config{
		AppId:  appId,
		Secret: secret,
		Token:  token,
	}
}

func (c *Config) Sign(param []string, signature string) bool {
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
