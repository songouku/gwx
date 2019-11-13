package test

import (
	"encoding/xml"
	"fmt"
	"gwx/client"
	"gwx/util"
	"io/ioutil"
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

type DataModel struct {
	Persons Data
}

type Data struct {
	Person []Person
}

type Person struct {
	Name      string
	Age       int
	Career    string
	Interests Interests
}

type Interests struct {
	Interest []string
}

type Result struct {
	Person []Person
}

func TestXmlParse(t *testing.T) {
	str, err := ioutil.ReadFile("person.xml")
	if err != nil {
		t.Errorf("error is %v", err)
		return
	}
	var result Result
	err = xml.Unmarshal(str, &result)
	if err != nil {
		t.Errorf("error is %v", err)
		return
	}
	fmt.Printf("result is %s", util.ConvertStr(result))
}
