package routers

import (
	"wxTribe/controllers/wechat"

	"github.com/astaxie/beego"
)

func init() {

	//微信公众号信息处理
	beego.Router("/wechat", &wechat.MainServiceController{})

}
