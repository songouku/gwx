package model

import "encoding/xml"

//<editor-fold desc="消息回复体">
type TextMessageResponse struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   string `content:"消息创建时间 （整型）"`
	MsgType      string
	Content      string
}

type ImageMessageResponse struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   string `content:"消息创建时间 （整型）"`
	MsgType      string
	MediaId      string
}

type VoiceMessageResponse struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   string `content:"消息创建时间 （整型）"`
	MsgType      string
	MediaId      string
}

type VideoMessageResponse struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   string `content:"消息创建时间 （整型）"`
	MsgType      string
	MediaId      string
	Title        string
	Description  string
}

type MusicMessageResponse struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   string `content:"消息创建时间 （整型）"`
	MsgType      string
	Title        string
	Description  string
	MusicURL     string
	HQMusicUrl   string `content:"高质量音乐链接，WIFI环境优先使用该链接播放音乐"`
	ThumbMediaId string `content:"缩略图的媒体id，通过素材管理中的接口上传多媒体文件，得到的id"`
}

type NewsMessageResponse struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   string `content:"消息创建时间 （整型）"`
	MsgType      string
	ArticleCount string `content:"图文消息个数；当用户发送文本、图片、视频、图文、地理位置这五种消息时，开发者只能回复1条图文消息；其余场景最多可回复8条图文消息"`
	Articles     string `content:"图文消息信息，注意，如果图文数超过限制，则将只发限制内的条数"`
	Title        string
	Description  string
	PicUrl       string
	Url          string
}

//</editor-fold>

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
}

type UserInfoResponse struct {
	ErrCode        int         `json:"errcode"`
	ErrMsg         string      `json:"errmsg"`
	Subscribe      int         `json:"subscribe"`
	OpenId         string      `json:"openid"`
	NickName       string      `json:"nickname"`
	Sex            int         `json:"sex"`
	Language       string      `json:"language"`
	City           string      `json:"city"`
	Province       string      `json:"province"`
	Country        string      `json:"country"`
	HeadImgUrl     string      `json:"headimgurl"`
	SubscribeTime  int64       `json:"subscribe_time"`
	Remark         string      `json:"remark"`
	GroupId        int         `json:"groupid"`
	TagIdList      interface{} `json:"tagid_list"`
	SubscribeScene string      `json:"subscribe_scene"`
	QrScene        int         `json:"qr_scene"`
	QrSceneStr     string
}
