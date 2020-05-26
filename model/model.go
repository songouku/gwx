package model

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
