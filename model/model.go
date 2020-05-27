package model

type Message struct {
	ToUserName   string
	FromUserName string
	CreateTime   string
	MsgType      string
	Content      string
	MsgId        string
	PicUrl       string
	MediaId      string
	Event        string
	Format       string
	Recognition  string
	EventKey     string
	ThumbMediaId string
	LocationX    string `json:"Location_X"`
	LocationY    string `json:"Location_Y"`
	Scale        string
	Label        string
	Title        string
	Description  string
	Url          string
}

//<editor-fold desc="消息">
type TxtMessage struct {
	ToUserName   string
	FromUserName string
	CreateTime   string
	MsgType      string
	Content      string
	MsgId        string
}

type ImageMessage struct {
	ToUserName   string
	FromUserName string
	CreateTime   string
	MsgType      string
	PicUrl       string
	MediaId      string
	MsgId        string
}

type VoiceMessage struct {
	ToUserName   string
	FromUserName string
	CreateTime   string
	MsgType      string
	MediaId      string
	MsgId        string
	Format       string
}

type VideoMessage struct {
	ToUserName   string
	FromUserName string
	CreateTime   string
	MsgType      string
	MediaId      string
	ThumbMediaId string
	MsgId        string
}

type LocationMessage struct {
	ToUserName   string
	FromUserName string
	CreateTime   string
	MsgType      string
	LocationX    string `json:"Location_X"`
	LocationY    string `json:"Location_Y"`
	Scale        string
	Label        string
	MsgId        string
}

type LinkMessage struct {
	ToUserName   string
	FromUserName string
	CreateTime   string
	MsgType      string
	Title        string
	Description  string
	Url          string
	MsgId        string
}

//</editor-fold>

//<editor-fold desc="事件消息">
type SubscribeEventMessage struct {
	ToUserName   string
	FromUserName string
	CreateTime   string
	MsgType      string
	Event        string
	EventKey     string
	Ticket       string
}

type LocationEventMessage struct {
	ToUserName   string
	FromUserName string
	CreateTime   string
	MsgType      string
	Event        string
	Latitude     string
	Longitude    string
	Precision    string
}

type ClickEventMessage struct {
	ToUserName   string
	FromUserName string
	CreateTime   string
	MsgType      string
	Event        string
	EventKey     string
}

type ViewEventMessage struct {
	ToUserName   string
	FromUserName string
	CreateTime   string
	MsgType      string
	Event        string
	EventKey     string
}

//</editor-fold>

type Response struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type Token struct {
	AccessToken string `json:"access_token,omitempty"`
	ExpiresIn   int    `json:"expires_in,omitempty"`
	ErrCode     int    `json:"errcode,omitempty"`
	ErrMsg      string `json:"errmsg,omitempty"`
}

type MenuItem struct {
	Type      string     `json:"type"`
	Name      string     `json:"name"`
	Key       string     `json:"key"`
	Url       string     `json:"url"`
	AppId     string     `json:"appid"`
	PagePath  string     `json:"pagepath"`
	SubButton []MenuItem `json:"sub_button"`
}

type MenuModel struct {
}

//永久素材
type WxMedia struct {
	MediaId      string `json:"media_id"`
	Url          string `json:"url"`
	Title        string `json:"title"`
	Introduction string `json:"introduction"`
	ErrCode      string `json:"errcode"`
	ErrMsg       string `json:"errmsg"`
}

//临时素材
type WxImage struct {
	Type      string `json:"type"`
	MediaId   string `json:"media_id"`
	CreatedAt int    `json:"created_at"`
	ErrCode   string `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
}
