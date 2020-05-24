package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/songouku/gwx/constant"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
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

func Upload(url, fileName, token, mediaType string) (*constant.WxImage, error) {
	//打开文件
	fh, err := os.Open(fileName)
	if err != nil {
		fmt.Errorf("error is %v", err.Error())
		return nil, err
	}
	defer fh.Close()
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	fileWriter, err := bodyWriter.CreateFormFile(mediaType, filepath.Base(fileName))
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return nil, err
	}
	bodyWriter.Close()

	//upload
	req, err := http.NewRequest("POST", fmt.Sprintf(url, token, mediaType), bodyBuf)
	req.Header.Add("Content-Type", bodyWriter.FormDataContentType())
	urlQuery := req.URL.Query()
	if err != nil {
		return nil, err
	}
	urlQuery.Add("access_token", token)
	urlQuery.Add("type", "image")

	req.URL.RawQuery = urlQuery.Encode()
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	content, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}
	var result constant.WxImage
	err = json.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
