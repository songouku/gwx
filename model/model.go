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
