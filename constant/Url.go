package constant

var (
	//获取token
	Token = "https://api.weixin.qq.com/cgi-bin/token"

	//创建菜单 POST
	// https://api.weixin.qq.com/cgi-bin/menu/create?access_token=ACCESS_TOKEN
	CreateMenu = "https://api.weixin.qq.com/cgi-bin/menu/create"

	//获取菜单 POST
	//https://api.weixin.qq.com/cgi-bin/get_current_selfmenu_info?access_token=ACCESS_TOKEN
	QueryMenu = "https://api.weixin.qq.com/cgi-bin/get_current_selfmenu_info"

	//删除菜单 GET
	//https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=ACCESS_TOKEN
	DelMenu = "https://api.weixin.qq.com/cgi-bin/menu/delete"

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
)
