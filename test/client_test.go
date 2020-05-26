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
var token = "33_0cnu0HWyu2FJvoqTLzwqwfAch21qjZnbNqxmF2asMIXFNLEvR5VEGX-gN2EMMmohZnWrkXWp9nygW3G71UPJbAu2AQm28RJF8o33clTSVDxIjnY1V2nCUm7yhhY4UeXmkxqJMTKHlzAFdaWBNSMdAFARWQ"

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
	fmt.Printf("msgId is %s", config.UploadTmpMedia(token, filmName))
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
