package test

import (
	"encoding/xml"
	"fmt"
	"github.com/songouku/gwx/constant"
	"github.com/songouku/gwx/model"
	"github.com/songouku/gwx/wx"
	"math/rand"
	"testing"
	"time"
)

var (
	AppId       = "wxc5a30a91fda3c2de"
	Secret      = "79fff5da1f931bccd6814578cc413fee"
	Token       = "zifeiyu"
	EncodingKey = "h3gJrFDBcvJL29iJOl9ITNH1L0mH05WqW1haZTZbShr"
)

var config *wx.Config
var token = "34_QaUy_WqBnQjjLP3CqRsiGtPP8cBTFSiKOIcfEs6AeUk5MlodmPFXWNujF48tBvy6uOdJRjrfRMd7f6ctcP3jJd9hzhioc3Pid9GRYnOCO73tbj6uurjqChcBJDyJY6O7082haXzKgiB9BAkYJOZaAGADIT"

func init() {
	config = wx.NewConfig(AppId, Secret, Token, EncodingKey)
}

func console(data interface{}, err error) {
	if err != nil {
		fmt.Errorf("error is %v\n", err)
		return
	}
	fmt.Printf("result is %v\n", data)
}

func TestToken(t *testing.T) {
	acc, err := config.GetToken()
	if err != nil {
		fmt.Errorf("error is %v\n", err)
		return
	}
	token := acc.AccessToken
	fmt.Printf("token is %s\n", token)
}

func TestTmpMedia(t *testing.T) {
	filmName := "/Users/allen/Pictures/tt.jpeg"
	console(wx.UploadTmpMedia(token, filmName))
}

func TestMedia(t *testing.T) {
	filmName := "/Users/allen/Pictures/tt.jpeg"
	console(wx.UploadImg(token, filmName))
}

func TestMaterial(t *testing.T) {
	filmName := "/Users/allen/Pictures/tt.jpeg"
	console(wx.AddMaterial(token, filmName, constant.Thumb))
}

func TestGetTmpMaterial(t *testing.T) {
	mediaId := "lEnSJW54UWeWNBPFTqpkFc95U0DbMoKBgXyyDdnsWlfFkOLdLc-32Nc0ZJle9EWH"
	console(wx.GetTmpMaterial(token, mediaId))
}

func TestAddNews(t *testing.T) {
	news := []wx.Item{
		{
			Title:              "测试",
			ThumbMediaId:       "0GPeqeABYec913q2Ho6p6ci3fMzyf5UExbLQh1uLbq4",
			Author:             "子非鱼",
			Digest:             "测试图文",
			ShowCoverPic:       1,
			Content:            "测试图文消息",
			ContentSourceUrl:   "http://www.baidu.com",
			NeedOpenComment:    1,
			OnlyFansCanComment: 1,
		},
	}
	args := wx.New{Articles: news}
	console(wx.AddNews(token, args))
}

func TestGetMaterial(t *testing.T) {
	mediaId := "0GPeqeABYec913q2Ho6p6S27QW8Bgb4DDXhx3cmoeeE"
	console(wx.GetMaterial(token, mediaId))
}

func TestBatchMaterial(t *testing.T) {
	console(wx.BatchGetMaterial(token, wx.BatchGetMaterialArgs{
		Type:   constant.News.Type,
		Offset: 0,
		Count:  20,
	}))
}

func TestMaterialCount(t *testing.T) {
	console(wx.GetMaterialCount(token))
}

func TestDelMaterial(t *testing.T) {
	mediaId := "0GPeqeABYec913q2Ho6p6S27QW8Bgb4DDXhx3cmoeeE"
	console(wx.DelMaterial(token, mediaId))
}

func TestCreateMenu(t *testing.T) {
	button := wx.Button{
		Type: constant.Click.Key,
		Name: "百度",
		Key:  "baidu",
	}
	args := wx.CreateMenuArgs{
		Button: []wx.Button{button},
	}
	console(wx.CreateMenu(token, args))
}

func TestQueryMenu(t *testing.T) {
	console(wx.QueryMenu(token))
}

func TestDelMenu(t *testing.T) {
	console(wx.DelMenu(token))
}

func TestAes(t *testing.T) {
	str := " <xml>\n    <ToUserName><![CDATA[gh_b2eeb48362f3]]></ToUserName>\n    <Encrypt><![CDATA[37+xerWBnoHHcv0hPcZaJfhwLiFcc2LgJPnxdwITdgXx54y30v5jN5iz4b4fLMS84oLFqJ0TYoK6yK1i4ulsx4N15tWnDcqOWZNn+dYk8eTs6V1hw9df6EpBemJSQ/f0WgxsJnE3kmc0Fpm72l+zS4Th9QwxsJTgEpnMjAddHWMYkgrfNm8n9SFMXd/1m8Y5rEm4osYW3u21+2h0gao/h5jk3BpchMVcBMu1TgHPjYNenWA1f7tLxYJ2V4/C5iQk9WIBnRA0VvPxByyBRX3DX9We7KriPJHEntRZPqtjbYXBr001V1sPd1IMPdS78uv0symDrsFmDPXvRXF3Wgx4/rjv7q8AIz+TAD/sFI3MajVN7MYaP1kLZ6DyDtU1pO3HVhs9q544izTYfO4Ai3vSfiS20EKf5GZ2sPlwAKJSZZY=]]></Encrypt>\n</xml>"
	var res model.Message
	err := xml.Unmarshal([]byte(str), &res)
	if err != nil {
		fmt.Errorf("error is %v\n", res)
		return
	}
	console(config.Decrypt(res.Encrypt))
}

func TestTime(t *testing.T) {
	now := time.Now().Unix()
	fmt.Printf("now is %v\n", now)
}

func TestRandom(t *testing.T) {
	str := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	b := make([]rune, 16)
	for i := range b {
		b[i] = str[rand.Intn(len(str))]
	}
	fmt.Printf("str is %s\n", string(b))
}

func TestUserInfo(t *testing.T) {
	openId := "o9h_PvhcxthiebXImiHCyJkVp-Eg"
	res, err := wx.UserInfo(token, openId)
	console(res, err)
}

func TestBatchUserInfo(t *testing.T) {
	openId := "o9h_PvhcxthiebXImiHCyJkVp-Eg"
	res, err := wx.BatchUserInfo(token, openId)
	console(res, err)
}
