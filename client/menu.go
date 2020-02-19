package client

import (
	"errors"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"gwx/constant"
	"gwx/model"
)

func (c *Client) CreateMenu(param []model.MenuItem) error {
	args := g.Map{"button": param}
	content := ghttp.PostContent(constant.CreateMenu, args)
	var result model.Response
	err := gjson.DecodeTo(content, &result)
	if err != nil {
		glog.Errorf("create menu failed , error is %v", err)
		return err
	}
	if result.ErrCode != 0 {
		return errors.New(result.ErrMsg)
	}
	return nil
}

func (c *Client) QueryMenu() {

}
