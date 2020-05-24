package constant

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
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

//<editor-fold desc="消息回复体">
type TextMessageResponse struct {
	ToUserName   string
	FromUserName string
	CreateTime   string `content:"消息创建时间 （整型）"`
	MsgType      string
	Content      string
}

type ImageMessageResponse struct {
	ToUserName   string
	FromUserName string
	CreateTime   string `content:"消息创建时间 （整型）"`
	MsgType      string
	MediaId      string
}

type VoiceMessageResponse struct {
	ToUserName   string
	FromUserName string
	CreateTime   string `content:"消息创建时间 （整型）"`
	MsgType      string
	MediaId      string
}

type VideoMessageResponse struct {
	ToUserName   string
	FromUserName string
	CreateTime   string `content:"消息创建时间 （整型）"`
	MsgType      string
	MediaId      string
	Title        string
	Description  string
}

type MusicMessageResponse struct {
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

//临时素材
type WxImage struct {
	Type      string `json:"type"`
	MediaId   string `json:"media_id"`
	CreatedAt int    `json:"created_at"`
	ErrCode   string `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
}
