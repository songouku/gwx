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

type MediaResponse struct {
	MediaId      string `json:"media_id"`
	Url          string `json:"url"`
	Title        string `json:"title"`
	Introduction string `json:"introduction"`
	ErrCode      int    `json:"errcode"`
	ErrMsg       string `json:"errmsg"`
}

type ImageMedia struct {
	Type      string `json:"type"`
	MediaId   string `json:"media_id"`
	CreatedAt int    `json:"created_at"`
	ErrCode   int    `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
}

//上传临时素材
func UploadTmpMedia(token, fileName string) (*ImageMedia, error) {
	body, err := util.Upload(constant.CreateMedia, fileName, token, constant.Image.Type)
	if err != nil {
		return nil, err
	}
	var result ImageMedia
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

//上传图文素材中的图片
func UploadImg(token, fileName string) (*MediaResponse, error) {
	body, err := util.Upload(constant.UploadImg, fileName, token, constant.Image.Type)
	if err != nil {
		return nil, err
	}
	var result MediaResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, err
}

//添加其他永久素材
func AddMaterial(token, fileName string, mediaType constant.MediaType) (*MediaResponse, error) {
	body, err := util.PostForm(constant.AddMaterial, fileName, token, mediaType.Type)
	if err != nil {
		return nil, err
	}
	var result MediaResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

//获取临时素材
func GetTmpMaterial(token, mediaId string) (interface{}, error) {
	param := make(map[string]interface{})
	param["access_token"] = token
	param["media_id"] = mediaId
	res, err := util.Get(constant.GetTmpMaterial, param)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type Material struct {
	ErrCode     int        `json:"errcode"`
	ErrMsg      string     `json:"errmsg"`
	NewsItem    []NewsItem `json:"news_item"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	DownUrl     string     `json:"down_url"`
}

type NewsItem struct {
	Title            string `json:"title"`
	ThumbMediaId     string `json:"thumb_media_id"`
	ShowCoverPic     int    `json:"show_cover_pic"`
	Author           string `json:"author"`
	Digest           string `json:"digest"`
	Content          string `json:"content"`
	Url              string `json:"url"`
	ContentSourceUrl string `json:"content_source_url"`
}

//获取永久素材
func GetMaterial(token, mediaId string) (*Material, error) {
	param := make(map[string]interface{})
	param["media_id"] = mediaId
	res, err := util.Post(fmt.Sprintf(constant.GetMaterial, token), param)
	if err != nil {
		return nil, err
	}
	var result Material
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type Item struct {
	Title              string `json:"title" content:"标题"`
	ThumbMediaId       string `json:"thumb_media_id" content:"图文消息的封面图片素材id（必须是永久mediaID）"`
	Author             string `json:"author" content:"作者"`
	Digest             string `json:"digest" content:"图文消息的摘要，仅有单图文消息才有摘要，多图文此处为空。如果本字段为没有填写，则默认抓取正文前64个字。"`
	ShowCoverPic       int    `json:"show_cover_pic" content:"是否显示封面，0为false，即不显示，1为true，即显示"`
	Content            string `json:"content" content:"图文消息的具体内容，支持HTML标签，必须少于2万字符，小于1M，且此处会去除JS,涉及图片url必须来源 '上传图文消息内的图片获取URL'接口获取。外部图片url将被过滤。"`
	ContentSourceUrl   string `json:"content_source_url" content:"图文消息的原文地址，即点击“阅读原文”后的URL"`
	NeedOpenComment    int    `json:"need_open_comment" content:"Uint32 是否打开评论，0不打开，1打开"`
	OnlyFansCanComment int    `json:"only_fans_can_comment" content:"Uint32 是否粉丝才可评论，0所有人可评论，1粉丝才可评论"`
}

type New struct {
	Articles []Item `json:"articles"`
}

type NewsResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	MediaId string `json:"media_id"`
}

//上传图文素材
func AddNews(token string, args New) (*NewsResponse, error) {
	content, err := util.Post(fmt.Sprintf(constant.AddNews, token), args)
	var result NewsResponse
	err = json.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type BatchGetMaterialArgs struct {
	Type   string `json:"type" content:"素材的类型，图片（image）、视频（video）、语音 （voice）、图文（news）"`
	Offset int    `json:"offset" content:"从全部素材的该偏移位置开始返回，0表示从第一个素材 返回"`
	Count  int    `json:"count" content:"返回素材的数量，取值在1到20之间"`
}

type BatchMaterial struct {
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
	TotalCount int    `json:"total_count"`
	ItemCount  int    `json:"item_count"`
	Item       []interface{}
}

type BatchMaterialItem struct {
	MediaId    string               `json:"media_id"`
	Content    BatchMaterialContent `json:"content"`
	UpdateTime int64                `json:"update_time"`
	Name       string               `json:"name"`
	Url        string               `json:"url"`
}

type BatchMaterialContent struct {
	NewsItem []NewsItem `json:"news_item"`
}

//获取素材列表
func BatchGetMaterial(token string, args BatchGetMaterialArgs) (*BatchMaterial, error) {
	res, err := util.Post(fmt.Sprintf(constant.BatchGetMaterial, token), args)
	if err != nil {
		return nil, err
	}

	var result BatchMaterial
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

type MaterialCount struct {
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
	VoiceCount int    `json:"voice_count"`
	VideoCount int    `json:"video_count"`
	ImageCount int    `json:"image_count"`
	NewsCount  int    `json:"news_count"`
}

//素材数量
func GetMaterialCount(token string) (*MaterialCount, error) {
	res, err := http.Get(fmt.Sprintf(constant.GetMaterialCount, token))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var result MaterialCount
	err = json.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

//删除素材
func DelMaterial(token, mediaId string) (*model.Response, error) {
	param := make(map[string]interface{})
	param["media_id"] = mediaId
	res, err := util.Post(fmt.Sprintf(constant.DelMaterial, token), param)
	if err != nil {
		return nil, err
	}
	var result model.Response
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
