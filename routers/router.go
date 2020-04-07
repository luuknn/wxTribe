package routers

import (
	"wxTribe/controllers/user"
	"wxTribe/controllers/wechat"

	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &user.ResumeController{})

	//微信公众号信息处理
	beego.Router("/wechat", &wechat.MainServiceController{})

}
