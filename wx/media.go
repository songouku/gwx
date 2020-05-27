package wx

import (
	"encoding/json"
	"github.com/songouku/gwx/constant"
	"github.com/songouku/gwx/model"
	"github.com/songouku/gwx/util"
)

//上传临时素材
func (c *Config) UploadTmpMedia(token, fileName string) (*model.WxImage, error) {
	body, err := util.Upload(constant.CreateMedia, fileName, token, constant.Image.Type)
	if err != nil {
		return nil, err
	}
	var result model.WxImage
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

//上传图文素材中的图片
func (c *Config) UploadImg(token, fileName string) (*model.WxMedia, error) {
	body, err := util.Upload(constant.UploadImg, fileName, token, constant.Image.Type)
	if err != nil {
		return nil, err
	}
	var result model.WxMedia
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, err
}

//添加其他永久素材
func (c *Config) AddMaterial(token, fileName string, mediaType constant.MediaType) (*model.WxMedia, error) {
	body, err := util.Upload(constant.AddMaterial, fileName, token, mediaType.Type)
	if err != nil {
		return nil, err
	}
	var result model.WxMedia
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

//获取临时素材
func (c *Config) GetTmpMaterial(token, mediaId string) (interface{}, error) {
	param := make(map[string]interface{})
	param["access_token"] = token
	param["media_id"] = mediaId

	res, err := util.Get(param, constant.GetTmpMaterial)
	if err != nil {
		return nil, err
	}
	return res, nil
}
