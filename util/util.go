package util

import (
	"gwx/constant"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetForJson(params map[string]interface{}) (string, error) {
	res, err := get(params, constant.Token)
	if err != nil {

	}
	return res, nil
}

func get(params map[string]interface{}, url string) (string, error) {
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
