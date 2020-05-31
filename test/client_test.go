package test

import (
	"fmt"
	"github.com/songouku/gwx/constant"
	"github.com/songouku/gwx/wx"
	"testing"
)

var (
	AppId  = "wxc5a30a91fda3c2de"
	Secret = "79fff5da1f931bccd6814578cc413fee"
	Token  = "zifeiyu"
)

var config *wx.Config
var token = "33_7idsR4DaDFyZZ1sgo2zxNsA2sgTzgFElRSt-uXkjaoHVCIZXtqptwrCU27Cp7qKWWrPkDNPMa2nfdmum6efx917XWxnNk0qPwTmVunvsW5VddpSnc5vYJsynruVw4j541Hie4aMqPxQqosliIUGdAIATFR"

func init() {
	config = wx.NewConfig(AppId, Secret, Token)
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
	res, err := wx.UploadTmpMedia(token, filmName)
	if err != nil {
		fmt.Errorf("error is %v\n", err)
		return
	}
	fmt.Printf("msgId is %v\n", res)
}

func TestMedia(t *testing.T) {
	filmName := "/Users/allen/Pictures/tt.jpeg"
	res, err := wx.UploadImg(token, filmName)
	if err != nil {
		fmt.Errorf("error is %v\n", err)
		return
	}
	fmt.Printf("msgId is %v\n", res)
}

func TestMaterial(t *testing.T) {
	filmName := "/Users/allen/Pictures/tt.jpeg"
	res, err := wx.AddMaterial(token, filmName, constant.Thumb)
	if err != nil {
		fmt.Errorf("error is %v\n", err)
		return
	}
	fmt.Printf("msgId is %v\n", res)
}

func TestGetTmpMaterial(t *testing.T) {
	mediaId := "lEnSJW54UWeWNBPFTqpkFc95U0DbMoKBgXyyDdnsWlfFkOLdLc-32Nc0ZJle9EWH"
	res, err := wx.GetTmpMaterial(token, mediaId)
	if err != nil {
		fmt.Errorf("error is %v\n", err)
		return
	}
	fmt.Printf("msgId is %v\n", res)
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
	res, err := wx.AddNews(token, args)
	if err != nil {
		fmt.Errorf("error is %v\n", err)
		return
	}
	fmt.Printf("res is %v", res)
}

func TestGetMaterial(t *testing.T) {
	mediaId := "0GPeqeABYec913q2Ho6p6S27QW8Bgb4DDXhx3cmoeeE"
	res, err := wx.GetMaterial(token, mediaId)
	if err != nil {
		fmt.Errorf("error is %v\n", err.Error())
		return
	}
	fmt.Printf("res is %v\n", res.NewsItem)
}

func TestBatchMaterial(t *testing.T) {
	res, err := wx.BatchGetMaterial(token, wx.BatchGetMaterialArgs{
		Type:   constant.News.Type,
		Offset: 0,
		Count:  20,
	})
	if err != nil {
		fmt.Printf("error is %v\n", err)
		return
	}
	fmt.Printf("res is %v\n", res)
}

func TestMaterialCount(t *testing.T) {
	res, err := wx.GetMaterialCount(token)
	if err != nil {
		fmt.Printf("error is %v\n", err)
		return
	}
	fmt.Printf("res is %v\n", res)
}

func TestDelMaterial(t *testing.T) {
	mediaId := "0GPeqeABYec913q2Ho6p6S27QW8Bgb4DDXhx3cmoeeE"
	res, err := wx.DelMaterial(token, mediaId)
	if err != nil {
		fmt.Printf("error is %v\n", err)
		return
	}
	fmt.Printf("res is %v\n", res)
}
