package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func Get(url string, params map[string]interface{}) ([]byte, error) {
	client := &http.Client{}
	args := ""
	for key, value := range params {
		args = args + key + "=" + value.(string) + "&"
	}
	req, err := http.NewRequest("get", url, strings.NewReader(args[:len(args)-1]))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func Post(requestUrl string, args interface{}) ([]byte, error) {
	body, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	contentType := "application/x-www-form-urlencoded"
	res, err := http.Post(requestUrl, contentType, strings.NewReader(string(body)))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

func Upload(url, fileName, token, mediaType string) ([]byte, error) {
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
	req, err := http.NewRequest("POST", url, bodyBuf)
	req.Header.Add("Content-Type", bodyWriter.FormDataContentType())
	urlQuery := req.URL.Query()
	if err != nil {
		return nil, err
	}
	urlQuery.Add("access_token", token)
	urlQuery.Add("type", mediaType)
	req.URL.RawQuery = urlQuery.Encode()
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

func PostForm(requestUrl, fileName, token, mediaType string) ([]byte, error) {
	//打开文件
	fh, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer fh.Close()
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	fileWriter, err := bodyWriter.CreateFormFile("media", filepath.Base(fileName))
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return nil, err
	}
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()
	res, err := http.Post(fmt.Sprintf(requestUrl, token, mediaType), contentType, bodyBuf)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}
