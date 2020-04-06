package constant

type MsgResponse struct {
	ToUserName   string
	FromUserName string
	Content      string
	CreateTime   int64
	MsgType      string
	Image        Media
	Voice        Media
	Video        Video
	Articles     []struct {
		Item ArticlesItem
	}
}

type Media struct {
	MediaId string
}

type Video struct {
	MediaId     string
	Title       string
	Description string
}

type Music struct {
	Title        string
	Description  string
	MusicUrl     string
	HQMusicUrl   string
	ThumbMediaId string
}

type ArticlesItem struct {
	Title       string
	Description string
	PicUrl      string
	Url         string
}
