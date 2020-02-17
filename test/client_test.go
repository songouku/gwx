package test

import (
	"encoding/xml"
	"fmt"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
	"gwx/client"
	"gwx/util"
	"io/ioutil"
	"testing"
)

var wx = client.Client{
	AppId:  "wxc5a30a91fda3c2de",
	Secret: "79fff5da1f931bccd6814578cc413fee",
}

func TestToken(t *testing.T) {
	token, err := wx.GetToken()
	if err != nil {
		fmt.Errorf("error is %v", err)
		return
	}
	fmt.Printf("res is %s", gconv.String(token))
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
	fmt.Printf("result is %s", gconv.String(result))
}

func TestXml(t *testing.T) {
	content := gfile.GetBytes("../test.xml")
	var result Content
	util.ConvertXmlToStruct(content, &result)
	glog.Infof("json is %v", result)

}

type Model struct {
	Xml Content
}

type Content struct {
	ToUserName   string
	FromUserName string
	CreateTime   string
	MsgType      string
	Content      string
	MsgId        string
}
