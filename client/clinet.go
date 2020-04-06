package client

type client struct {
	AppId  string `json:"appId"`
	Secret string `json:"secret"`
}

func NewClient(appId string, secret string) *client {
	return &client{
		AppId:  appId,
		Secret: secret,
	}
}
