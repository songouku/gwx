package wx

import (
	"encoding/json"
	"github.com/songouku/gwx/constant"
	"github.com/songouku/gwx/model"
	"github.com/songouku/gwx/util"
)

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
