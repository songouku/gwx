package test

import (
	"fmt"
	"github.com/songouku/gwx/constant"
	"github.com/songouku/gwx/util"
	"testing"
)

var (
	AppId  = "wxc5a30a91fda3c2de"
	Secret = "79fff5da1f931bccd6814578cc413fee"
	Token  = "zifeiyu"
)

func TestMedia(t *testing.T) {
	//config := wx.NewConfig(AppId, Secret, Token)
	filmName := "/Users/allen/Pictures/tt.jpeg"
	token := "32_4LccVh6Il0SmYYKiMF0dU9JMT-eOdKSq-2PddQBwcxPIQ8YnZieu4KSa4oGUg4u9et3NYWwpd0CrO_fDGvpBJ3Dl4jmLcEETwMUZO3xpyhi2dJmQcjB2PMdrLILfViStfHXKtFNoEkjuKxt3BAQbAAAVUK"
	//acc, err := config.GetToken()
	//if err != nil {
	//	fmt.Errorf("error is %v", err)
	//	return
	//}
	//token := acc.AccessToken
	//fmt.Printf("token is %s", token)
	result, err := util.Upload(constant.CreateMedia, filmName, token, "image")
	if err != nil {
		fmt.Errorf("upload image failed , error is %v", err)
		return
	}
	fmt.Printf(result.MediaId)
}

type User struct {
}

func (u *User) login() {
	fmt.Printf("111")
}

func (u *User) loginOut() {
	fmt.Printf("222")
}

func loginwx(service UserService) {
	service.login()
}

func TestInterface(t *testing.T) {
	user := new(User)
	loginwx(user)
}
