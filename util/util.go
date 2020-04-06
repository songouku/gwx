package util

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/url"
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

func Post(requestUrl string, data map[string]interface{}) (string, error) {
	client := &http.Client{}
	param := url.Values{}
	for k, v := range data {
		param[k] = []string{v.(string)}
	}
	res, err := client.PostForm(requestUrl, param)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

//xml数据转为实体
// result 为指针类型
func ConvertXmlToStruct(data []byte, result interface{}) error {
	err := xml.Unmarshal(data, result)
	if err != nil {
		return err
	}
	return nil
}
