package constant

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
}

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
	ThumbMediaId string
	LocationX    string `json:"Location_X"`
	LocationY    string `json:"Location_Y"`
	Scale        string
	Label        string
	Title        string
	Description  string
	Url          string
}

type Menu struct {
}
