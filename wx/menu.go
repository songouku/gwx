package wx

import (
	"encoding/json"
	"fmt"
	"github.com/songouku/gwx/constant"
	"github.com/songouku/gwx/model"
	"github.com/songouku/gwx/util"
	"io/ioutil"
	"net/http"
)

type CreateMenuArgs struct {
	Button []Button `json:"button"`
}

type Button struct {
	Type      string   `json:"type"`
	Name      string   `json:"name"`
	Key       string   `json:"key"`
	Url       string   `json:"url"`
	AppId     string   `json:"appid"`
	PagePath  string   `json:"pagepath"`
	MediaId   string   `json:"media_id"`
	SubButton []Button `json:"sub_button"`
}

//创建菜单
func CreateMenu(token string, args CreateMenuArgs) (*model.Response, error) {
	content, err := util.Post(fmt.Sprintf(constant.CreateMenu, token), args)
	if err != nil {
		return nil, err
	}
	var result model.Response
	err = json.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	return &result, err
}

type MenuResponse struct {
	ErrCode      int          `json:"errcode"`
	ErrMsg       string       `json:"errmsg"`
	IsMenuOpen   int          `json:"is_menu_open"`
	SelfMenuInfo SelfMenuInfo `json:"selfmenu_info"`
}

type SelfMenuInfo struct {
	Button []ButtonResponse `json:"button"`
}

type ButtonResponse struct {
	Type      string            `json:"type"`
	Name      string            `json:"name"`
	Key       string            `json:"key"`
	Url       string            `json:"url"`
	AppId     string            `json:"appid"`
	PagePath  string            `json:"pagepath"`
	MediaId   string            `json:"media_id"`
	SubButton SubButtonResponse `json:"sub_button"`
}

type SubButtonResponse struct {
	List []ButtonResponse `json:"list"`
}

//查询菜单
func QueryMenu(token string) (*MenuResponse, error) {
	res, err := http.Get(fmt.Sprintf(constant.QueryMenu, token))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var result MenuResponse
	err = json.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func DelMenu(token string) (*model.Response, error) {
	res, err := http.Get(fmt.Sprintf(constant.DelMenu, token))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var result model.Response
	err = json.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
