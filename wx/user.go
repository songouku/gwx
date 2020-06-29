package wx

import (
	"encoding/json"
	"fmt"
	"github.com/songouku/gwx/constant"
	"github.com/songouku/gwx/model"
	"github.com/songouku/gwx/util"
)

//用户信息
func UserInfo(token, openId string) (*model.UserInfoResponse, error) {
	param := make(map[string]interface{})
	param["access_token"] = token
	param["openid"] = openId
	param["lang"] = "zh_CN"
	tmp, err := util.Get(constant.UserInfo, param)
	if err != nil {
		return nil, err
	}
	var result model.UserInfoResponse
	err = json.Unmarshal(tmp, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

//批量获取用户信息
func BatchUserInfo(token string, openId ...string) (*model.BatchUserInfoResponse, error) {
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
	var result model.BatchUserInfoResponse
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
