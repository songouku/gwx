package constant

var (
	//获取token
	Token = "https://api.weixin.qq.com/cgi-bin/token"

	//创建菜单 POST
	// https://api.weixin.qq.com/cgi-bin/menu/create?access_token=ACCESS_TOKEN
	CreateMenu = "https://api.weixin.qq.com/cgi-bin/menu/create?access_token=%s"

	//获取菜单 POST
	//https://api.weixin.qq.com/cgi-bin/get_current_selfmenu_info?access_token=ACCESS_TOKEN
	QueryMenu = "https://api.weixin.qq.com/cgi-bin/get_current_selfmenu_info?access_token=%s"

	//删除菜单 GET
	//https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=ACCESS_TOKEN
	DelMenu = "https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=%s"

	//创建个性化菜单 POST
	//https://api.weixin.qq.com/cgi-bin/menu/addconditional?access_token=ACCESS_TOKEN
	CreateConditionalMenu = "https://api.weixin.qq.com/cgi-bin/menu/addconditional"

	//删除个性化菜单 POST
	//https://api.weixin.qq.com/cgi-bin/menu/delconditional?access_token=ACCESS_TOKEN
	DelConditionalMenu = "https://api.weixin.qq.com/cgi-bin/menu/delconditional"

	//测试个性化菜单匹配结果 POST
	//https://api.weixin.qq.com/cgi-bin/menu/trymatch?access_token=ACCESS_TOKEN
	TryMatch = "https://api.weixin.qq.com/cgi-bin/menu/trymatch"

	//获取自定义菜单配置 GET
	//https://api.weixin.qq.com/cgi-bin/menu/get?access_token=ACCESS_TOKEN
	QueryConditionalMenu = "https://api.weixin.qq.com/cgi-bin/menu/get"

	//上传临时素材
	//https://api.weixin.qq.com/cgi-bin/media/upload?access_token=ACCESS_TOKEN&type=TYPE
	CreateMedia = "https://api.weixin.qq.com/cgi-bin/media/upload?access_token=%s&type=%s"

	//上传图文消息内的图片获取URL
	UploadImg = "https://api.weixin.qq.com/cgi-bin/media/uploadimg"

	//上传其他永久素材
	AddMaterial = "https://api.weixin.qq.com/cgi-bin/material/add_material?access_token=%s&type%s"

	//获取临时素材
	GetTmpMaterial = "https://api.weixin.qq.com/cgi-bin/media/get?access_token=%s&media_id=%s"

	//获取永久素材
	GetMaterial = "https://api.weixin.qq.com/cgi-bin/material/get_material?access_token=%s"

	//添加图文素材
	AddNews = "https://api.weixin.qq.com/cgi-bin/material/add_news?access_token=%s"

	//获取素材列表
	BatchGetMaterial = "https://api.weixin.qq.com/cgi-bin/material/batchget_material?access_token=%s"

	//获取素材数量
	GetMaterialCount = "https://api.weixin.qq.com/cgi-bin/material/get_materialcount?access_token=%s"

	//删除素材
	DelMaterial = "https://api.weixin.qq.com/cgi-bin/material/del_material?access_token=%s"
)
