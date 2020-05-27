package test

import (
	"fmt"
	"github.com/songouku/gwx/wx"
	"testing"
)

var (
	AppId  = "wxc5a30a91fda3c2de"
	Secret = "79fff5da1f931bccd6814578cc413fee"
	Token  = "zifeiyu"
)

var config *wx.Config
var token = "33_MPO9XV6KhSIPvlInF4eTI8pqxNMggqh1J99nzpba3fKsd-ZBMhi2iv71G5juOZuGa6CqB4T7DbzPlEmybvTby1bJxcM6QXFnIpJi1S2fwL5-McXXjmmgLFw10ybLT1TniCKXtNnkbj8yyrI8CDHfAFAKLP"

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
	res, err := config.UploadTmpMedia(token, filmName)
	if err != nil {
		fmt.Errorf("error is %v\n", err)
		return
	}
	fmt.Printf("msgId is %v\n", res)
}

func TestMedia(t *testing.T) {
	filmName := "/Users/allen/Pictures/tt.jpeg"
	res, err := config.UploadImg(token, filmName)
	if err != nil {
		fmt.Errorf("error is %v\n", err)
		return
	}
	fmt.Printf("msgId is %v\n", res)
}

func TestMaterial(t *testing.T) {
	filmName := "/Users/allen/Pictures/tt.jpeg"
	res, err := config.UploadImg(token, filmName)
	if err != nil {
		fmt.Errorf("error is %v\n", err)
		return
	}
	fmt.Printf("msgId is %v\n", res)
}

func TestGetTmpMaterial(t *testing.T) {
	mediaId := "lEnSJW54UWeWNBPFTqpkFc95U0DbMoKBgXyyDdnsWlfFkOLdLc-32Nc0ZJle9EWH"
	res, err := config.GetTmpMaterial(token, mediaId)
	if err != nil {
		fmt.Errorf("error is %v\n", err)
		return
	}
	fmt.Printf("msgId is %v\n", res)
}
