package wx

import (
	"encoding/json"
	"fmt"
	"github.com/songouku/gwx/constant"
	util "github.com/songouku/gwx/util"
)

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
	QrSceneStr     string      `json:"qr_scene_str"`
}

//用户信息
func UserInfo(token, openId string) (*UserInfoResponse, error) {
	param := make(map[string]interface{})
	param["access_token"] = token
	param["openid"] = openId
	param["lang"] = "zh_CN"
	tmp, err := util.Get(constant.UserInfo, param)
	if err != nil {
		return nil, err
	}
	var result UserInfoResponse
	err = json.Unmarshal(tmp, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type Info struct {
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
	QrSceneStr     string      `json:"qr_scene_str"`
}

type BatchUserInfoResponse struct {
	ErrCode      int    `json:"errcode"`
	ErrMsg       string `json:"errmsg"`
	UserInfoList []Info `json:"user_info_list"`
}

//批量获取用户信息
func BatchUserInfo(token string, openId ...string) (*BatchUserInfoResponse, error) {
	param := make(map[string]interface{})
	var list []map[string]interface{}
	for index := range openId {
		tmp := make(map[string]interface{})
		tmp["openid"] = openId[index]
		tmp["lang"] = "zh_CN"
		list = append(list, tmp)
	}
	param["user_list"] = list
	res, err := util.Post(fmt.Sprintf(constant.BatchUserInfo, token), param)
	if err != nil {
		return nil, err
	}
	var result BatchUserInfoResponse
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type UserTokenResponse struct {
	ErrCode      int    `json:"errcode"`
	ErrMsg       string `json:"errmsg"`
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenId       string `json:"openid"`
	Scope        string `json:"scope"`
}

//用户授权，获取token
func GetUserToken(appId, secret, code string) (*UserTokenResponse, error) {
	params := make(map[string]interface{})
	params["appid"] = appId
	params["secret"] = secret
	params["code"] = code
	params["grant_type"] = "authorization_code"
	res, err := util.Get(constant.GetUserToken, params)
	if err != nil {
		return nil, err
	}
	var result UserTokenResponse
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

//用户授权后获取用户信息
//此token为用户授权后获取的token
func GetUserInfo(token, openId string) (*Info, error) {
	params := make(map[string]interface{})
	params["access_token"] = token
	params["openid"] = openId
	params["lang"] = "zh_CN"
	res, err := util.Get(constant.GetUserInfo, params)
	if err != nil {
		return nil, err
	}
	var result Info
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
