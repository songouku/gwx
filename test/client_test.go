package test

import (
	"fmt"
	"gwx/client"
	"gwx/util"
	"testing"
)

var wx = client.Client{
	AppId:  "wxc5a30a91fda3c2d",
	Secret: "79fff5da1f931bccd6814578cc413fee",
}

func TestToken(t *testing.T) {
	token, err := wx.GetToken()
	if err != nil {
		fmt.Errorf("error is %v", err)
		return
	}
	fmt.Printf("res is %s", util.ConvertStr(token))
}
