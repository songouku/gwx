package wx

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"github.com/songouku/gwx/model"
	"github.com/songouku/gwx/util"
	"sort"
)

type Config struct {
	AppId       string `json:"appId"`
	Secret      string `json:"secret"`
	Token       string `json:"token"`
	EncodingKey string

	txtMessage      model.TxtMessage
	imageMessage    model.ImageMessage
	voiceMessage    model.VoiceMessage
	videoMessage    model.VideoMessage
	locationMessage model.LocationMessage
	linkMessage     model.LinkMessage

	subscribeEvent model.SubscribeEventMessage
	locationEvent  model.LocationEventMessage
	clickEvent     model.ClickEventMessage
	viewEvent      model.ViewEventMessage

	currentMsgType string
}

func (c *Config) SetTxtMessage(msg model.TxtMessage) {
	c.txtMessage = msg
}

func (c *Config) GetTxtMessage() model.TxtMessage {
	return c.txtMessage
}

func (c *Config) SetImageMessage(msg model.ImageMessage) {
	c.imageMessage = msg
}

func (c *Config) GetImageMessage() model.ImageMessage {
	return c.imageMessage
}

func (c *Config) SetVoiceMessage(msg model.VoiceMessage) {
	c.voiceMessage = msg
}

func (c *Config) GetVoiceMessage() model.VoiceMessage {
	return c.voiceMessage
}

func (c *Config) SetVideoMessage(msg model.VideoMessage) {
	c.videoMessage = msg
}

func (c *Config) GetVideoMessage() model.VideoMessage {
	return c.videoMessage
}

func (c *Config) SetLocationMessage(msg model.LocationMessage) {
	c.locationMessage = msg
}

func (c *Config) GetLocationMessage() model.LocationMessage {
	return c.locationMessage
}

func (c *Config) SetLinkMessage(msg model.LinkMessage) {
	c.linkMessage = msg
}

func (c *Config) GetLinkMessage() model.LinkMessage {
	return c.linkMessage

}

func (c *Config) SetSubscribeEventMessage(event model.SubscribeEventMessage) {
	c.subscribeEvent = event
}

func (c *Config) GetSubscribeEventMessage() model.SubscribeEventMessage {
	return c.subscribeEvent
}

func (c *Config) SetLocationEventMessage(event model.LocationEventMessage) {
	c.locationEvent = event
}

func (c *Config) GetLocationEventMessage() model.LocationEventMessage {
	return c.locationEvent
}

func (c *Config) SetClickEventMessage(event model.ClickEventMessage) {
	c.clickEvent = event
}

func (c *Config) GetClickEventMessage() model.ClickEventMessage {
	return c.clickEvent
}

func (c *Config) SetViewEventMessage(event model.ViewEventMessage) {
	c.viewEvent = event
}

func (c *Config) GetViewEventMessage() model.ViewEventMessage {
	return c.viewEvent
}

func (c *Config) SetCurrentMsg(msgType string) {
	c.currentMsgType = msgType
}

func (c *Config) GetCurrentMsg() string {
	return c.currentMsgType
}

func NewConfig(appId, secret, token, EncodingKey string) *Config {
	return &Config{
		AppId:       appId,
		Secret:      secret,
		Token:       token,
		EncodingKey: EncodingKey,
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

func (c *Config) Decrypt(content string) (*model.Message, error) {
	aesKey, err := base64.StdEncoding.DecodeString(c.EncodingKey + "=")
	if err != nil {
		return nil, err
	}
	tmp, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		return nil, err
	}
	res, err := util.AesDecrypt(tmp, aesKey)
	if err != nil {
		return nil, err
	}
	s := string(res[20 : len(res)-18])
	var result model.Message
	err = xml.Unmarshal([]byte(s), &tmp)
	if err != nil {
		fmt.Errorf("error is %s\n", err.Error())
		return nil, err
	}
	return &result, err
}
