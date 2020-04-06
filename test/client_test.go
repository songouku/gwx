package test

import (
	"fmt"
	"gwx/client"
	"gwx/util"
	"io/ioutil"
	"testing"
)

var (
	AppId  = "wxc5a30a91fda3c2de"
	Secret = "79fff5da1f931bccd6814578cc413fee"
)

var wx = client.NewClient(AppId, Secret)

func TestToken(t *testing.T) {
	token, err := wx.GetToken()
	if err != nil {
		fmt.Errorf("error is %v", err)
		return
	}
	fmt.Printf("res is %v", token)
}

func TestXmlParse(t *testing.T) {
	str, err := ioutil.ReadFile("person.xml")
	if err != nil {
		t.Errorf("error is %v", err)
		return
	}
	var result Company
	err = util.ConvertXmlToStruct(str, &result)
	if err != nil {
		t.Errorf("error is %v", err)
		return
	}
	fmt.Printf("result is %v", result)
}

type Company struct {
	Staff []Staff `xml:"staff"`
}

type Staff struct {
	Id        string `xml:"id" json:"id"`
	FirstName string `xml:"firstname" json:"firstName"`
	LastName  string `xml:"lastname" json:"lastName"`
	UserName  string `xml:"username" json:"userName"`
}
