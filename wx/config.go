package wx

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/songouku/gwx/constant"
	"sort"
)

type Config struct {
	AppId  string `json:"appId"`
	Secret string `json:"secret"`
	Token  string `json:"token"`

	txtMessage      constant.TxtMessage
	imageMessage    constant.ImageMessage
	voiceMessage    constant.VoiceMessage
	videoMessage    constant.VideoMessage
	locationMessage constant.LocationMessage
	linkMessage     constant.LinkMessage

	subscribeEvent constant.SubscribeEventMessage
	locationEvent  constant.LocationEventMessage
	clickEvent     constant.ClickEventMessage
	viewEvent      constant.ViewEventMessage

	currentMsgType string
}

func (c *Config) SetTxtMessage(msg constant.TxtMessage) {
	c.txtMessage = msg
}

func (c *Config) GetTxtMessage() constant.TxtMessage {
	return c.txtMessage
}

func (c *Config) SetImageMessage(msg constant.ImageMessage) {
	c.imageMessage = msg
}

func (c *Config) GetImageMessage() constant.ImageMessage {
	return c.imageMessage
}

func (c *Config) SetVoiceMessage(msg constant.VoiceMessage) {
	c.voiceMessage = msg
}

func (c *Config) GetVoiceMessage() constant.VoiceMessage {
	return c.voiceMessage
}

func (c *Config) SetVideoMessage(msg constant.VideoMessage) {
	c.videoMessage = msg
}

func (c *Config) GetVideoMessage() constant.VideoMessage {
	return c.videoMessage
}

func (c *Config) SetLocationMessage(msg constant.LocationMessage) {
	c.locationMessage = msg
}

func (c *Config) GetLocationMessage() constant.LocationMessage {
	return c.locationMessage
}

func (c *Config) SetLinkMessage(msg constant.LinkMessage) {
	c.linkMessage = msg
}

func (c *Config) GetLinkMessage() constant.LinkMessage {
	return c.linkMessage

}

func (c *Config) SetSubscribeEventMessage(event constant.SubscribeEventMessage) {
	c.subscribeEvent = event
}

func (c *Config) GetSubscribeEventMessage() constant.SubscribeEventMessage {
	return c.subscribeEvent
}

func (c *Config) SetLocationEventMessage(event constant.LocationEventMessage) {
	c.locationEvent = event
}

func (c *Config) GetLocationEventMessage() constant.LocationEventMessage {
	return c.locationEvent
}

func (c *Config) SetClickEventMessage(event constant.ClickEventMessage) {
	c.clickEvent = event
}

func (c *Config) GetClickEventMessage() constant.ClickEventMessage {
	return c.clickEvent
}

func (c *Config) SetViewEventMessage(event constant.ViewEventMessage) {
	c.viewEvent = event
}

func (c *Config) GetViewEventMessage() constant.ViewEventMessage {
	return c.viewEvent
}

func (c *Config) SetCurrentMsg(msgType string) {
	c.currentMsgType = msgType
}

func (c *Config) GetCurrentMsg() string {
	return c.currentMsgType
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
