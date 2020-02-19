package constant

type MenuType struct {
	Key  string
	Desc string
}

var (
	Click           = MenuType{Key: "click", Desc: "点击事件"}
	View            = MenuType{Key: "view", Desc: "视图事件"}
	ScanCodePush    = MenuType{Key: "scancode_push", Desc: "扫码推送事件"}
	ScanCodeWaitMsg = MenuType{Key: "scancode_waitmsg", Desc: "扫码等待信息"}
	PicSysPhoto     = MenuType{Key: "pic_sysphoto", Desc: ""}
	PicPhotoOrAlbum = MenuType{Key: "pic_photo_or_album", Desc: ""}
	PicWx           = MenuType{Key: "pic_weixin", Desc: ""}
	LocationSelect  = MenuType{Key: "location_select", Desc: ""}
	MediaId         = MenuType{Key: "media_id", Desc: ""}
	ViewLimited     = MenuType{Key: "view_limited", Desc: ""}
)
