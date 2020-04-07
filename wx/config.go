package wx

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/songouku/gwx/constant"
	"github.com/songouku/gwx/util"
	"sort"
	"time"
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

func (c *Config) Message(data string) (*constant.Message, error) {
	var message constant.Message
	err := util.ConvertXmlToStruct([]byte(data), &message)
	if err != nil {
		return nil, err
	}
	return &message, nil
}

func (c *Config) Response(message *constant.Message) *constant.MsgResponse {
	if message.MsgType == constant.EventMsg.Type {
		return event(message)
	}
	return nil
}

func event(message *constant.Message) *constant.MsgResponse {
	if message.Event == constant.Subscribe.Type {
		content := "感谢关注测试公众号，此公众号正在调试中"
		return subscribe(message, content)
	}
	return nil
}

//扫描关注
func subscribe(message *constant.Message, content string) *constant.MsgResponse {
	return &constant.MsgResponse{
		FromUserName: message.ToUserName,
		ToUserName:   message.FromUserName,
		CreateTime:   time.Now().Unix(),
		MsgType:      constant.TextMsg.Type,
		Content:      content,
	}
}
