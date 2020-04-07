package wx

import (
	"encoding/json"
	"errors"
	"github.com/songouku/gwx/constant"
	"github.com/songouku/gwx/model"
	"github.com/songouku/gwx/util"
)

func (c *Config) CreateMenu(param []model.MenuItem) error {
	args := map[string]interface{}{
		"button": param,
	}
	content, err := util.Post(constant.CreateMenu, args)
	if err != nil {
		return errors.New("create menu failed")
	}
	var result model.Response
	err = json.Unmarshal([]byte(content), &result)
	if err != nil {
		return errors.New("create menu failed")
	}
	if result.ErrCode != 0 {
		return errors.New(result.ErrMsg)
	}
	return nil
}

func (c *Config) QueryMenu() {

}
