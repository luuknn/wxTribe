package wechat

import (
	"github.com/astaxie/beego/logs"
	"wxTribe/controllers"
	"wxTribe/utils"

	"gopkg.in/chanxuehong/wechat.v2/mp/core"
)

var (
	msgHandler core.Handler
	msgServer  *core.Server
)

func init() {
	mux := core.NewServeMux()
	msgHandler = mux
	msgServer = core.NewServer(utils.WxOriId, utils.WxAppId, utils.WxToken, utils.WxEncodedAESKey, msgHandler, nil)
}

//微信公众平台
type MainServiceController struct {
	controllers.BaseController
}

//微信公众平台
func (c *MainServiceController) Get() {

	logs.Info("=======wxmp Get Start======>")

	msgServer.ServeHTTP(c.Ctx.ResponseWriter, c.Ctx.Request, nil)

	logs.Info("=======wxmp Get End======>")

	c.ServeXML()
}

//微信公众平台
func (c *MainServiceController) Post() {

	logs.Info("=======wxmp Post Start======>")

	msgServer.ServeHTTP(c.Ctx.ResponseWriter, c.Ctx.Request, nil)

	logs.Info("=======wxmp Post End======>")
}
