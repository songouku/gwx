package constant

type MenuType struct {
	Key  string
	Desc string
}

var (
	Click           = MenuType{Key: "click", Desc: "点击事件"}
	View            = MenuType{Key: "view", Desc: "视图事件"}
	ScanCodePush    = MenuType{Key: "scancode_push", Desc: "扫码推送事件"}
	ScanCodeWaitMsg = MenuType{Key: "scancode_waitmsg", Desc: "扫码等待信息"}
	PicSysPhoto     = MenuType{Key: "pic_sysphoto", Desc: ""}
	PicPhotoOrAlbum = MenuType{Key: "pic_photo_or_album", Desc: ""}
	PicWx           = MenuType{Key: "pic_weixin", Desc: ""}
	LocationSelect  = MenuType{Key: "location_select", Desc: ""}
	MediaId         = MenuType{Key: "media_id", Desc: ""}
	ViewLimited     = MenuType{Key: "view_limited", Desc: ""}
)

type MsgType struct {
	Type string
	Desc string
}

var (
	TextMsg       = MsgType{Type: "text", Desc: "文本消息"}
	EventMsg      = MsgType{Type: "event", Desc: "事件消息"}
	ImageMsg      = MsgType{Type: "image", Desc: "图片消息"}
	VoiceMsg      = MsgType{Type: "voice", Desc: "语音消息"}
	VideoMsg      = MsgType{Type: "video", Desc: "视频消息"}
	ShortVideoMsg = MsgType{Type: "shortvideo", Desc: "短视频消息"}
	LocationMsg   = MsgType{Type: "location", Desc: "地理位置"}
	LinkMsg       = MsgType{Type: "link", Desc: "链接消息"}
)

type Event struct {
	Type string
	Desc string
}

var (
	ViewEvent     = Event{Type: "VIEW", Desc: "点击菜单跳转链接时的事件推送"}
	ClickEvent    = Event{Type: "CLICK", Desc: "自定义菜单事件"}
	LocationEvent = Event{Type: "LOCATION", Desc: "上报地理位置事件"}
	ScanEvent     = Event{Type: "SCAN", Desc: "用户已关注时的事件推送"}
	Subscribe     = Event{Type: "subscribe", Desc: "扫描带参数二维码事件"}
)

type MediaType struct {
	Type string
	Desc string
}

var (
	Image = MediaType{Type: "image", Desc: "图片素材"}
	Voice = MediaType{Type: "voice", Desc: "音频素材"}
	Video = MediaType{Type: "video", Desc: "视频素材"}
	Thumb = MediaType{Type: "thumb", Desc: "缩略图"}
	News  = MediaType{Type: "news", Desc: "图文素材"}
)
