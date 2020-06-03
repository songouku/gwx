package test

import (
	"fmt"
	"github.com/songouku/gwx/constant"
	"github.com/songouku/gwx/wx"
	"testing"
)

var (
	AppId       = "wxc5a30a91fda3c2de"
	Secret      = "79fff5da1f931bccd6814578cc413fee"
	Token       = "zifeiyu"
	EncodingKey = "h3gJrFDBcvJL29iJOl9ITNH1L0mH05WqW1haZTZbShr"
)

var config *wx.Config
var token = "33_d3PylGkPFkYrMsEcbLiJq9YLJjS2Hp4ETbYh4KR9EJU8Xw9i8oD5Py3z6VKMLWVnfwzZ1PaZZtNXet1oBzHElrf6tC1tSf5NJ8CTzLMJMLJ91bFzExfFhoCdnuiKr4QU1SNvuOQRSxjUIHOpHYLeAIABSF"

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
	content := "BYFSjVEEMS/U6ZzOETIxTlx+7fZVwkmOX95CoLsPd/t3p/GS9pDuJ3pfuHHASv7a4FHlr3Ix5NHZ9mefCFzWCF2TabS9EqSl1v1RxyB04o4bkkp0MAufLTofvAIijTOGeIzQwEI95SImWT1hk6tdCfy3K1eMjyQrVl22cxSEIwF4xEZOS+SnhORFi/5J+qEagX9ZZTkBBX8EI1l7ZZ/KVhPam1WQQ4r+3VnzcHj0y5w/EQhLGt9FlglHMXnhBxgwCAJ062JESRSy8aqEDI/RMbPaXAxbkvAc9og0HOGw9Z8Cyp2PuldUmuHQWXogLD5M24R0/FZah1qpfIpxSoNJDOBREjwB9OVfCxBXlC6KsqeNv4xHzBK8xLI94QUBg+w/Xlmw/eXMmhijz29lC1xpJCCEaKo1gyKyNZxflaouTeE="
	console(config.Decrypt(content))
}
