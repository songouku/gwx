package util

import (
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/encoding/gxml"
	"github.com/gogf/gf/os/glog"
	"io/ioutil"
	"net/http"
	"strings"
)

func Get(params map[string]interface{}, url string) (string, error) {
	client := &http.Client{}
	args := ""
	for key, value := range params {
		args = args + key + "=" + value.(string) + "&"
	}
	req, err := http.NewRequest("get", url, strings.NewReader(args[:len(args)-1]))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", nil
	}

	return string(body), nil
}

//xml数据转为实体
// result 为指针类型
func ConvertXmlToStruct(data []byte, result interface{}) {
	json, err := gxml.ToJson(data)
	if err != nil {
		glog.Errorf("convert to json failed , error is %v", err)
		return
	}
	tmp, err := gjson.DecodeToJson(json)
	if err != nil {
		glog.Errorf("convert to struct failed , error is %v", err)
		return
	}
	encode, err := gjson.Encode(tmp.GetJson("xml"))
	if err != nil {
		glog.Errorf("encode failed , error is %v", err)
		return
	}
	err = gjson.DecodeTo(encode, result)
	if err != nil {
		glog.Errorf("convert to struct failed , error is %v", err)
		return
	}
}
