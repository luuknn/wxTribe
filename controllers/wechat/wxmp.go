package wechat

import (
	"github.com/astaxie/beego/logs"
	"github.com/spf13/viper"
	"gopkg.in/chanxuehong/wechat.v2/mp/core"
	"gopkg.in/chanxuehong/wechat.v2/mp/message/callback/request"
	"gopkg.in/chanxuehong/wechat.v2/mp/message/callback/response"
	"wxTribe/controllers"
)

var (
	msgHandler core.Handler
	msgServer  *core.Server

	accessTokenServer core.AccessTokenServer = core.NewDefaultAccessTokenServer(viper.GetString("wx.appId"), viper.GetString("wx.appSecret"), nil)
	weChatClient      *core.Client           = core.NewClient(accessTokenServer, nil)
)

//微信公众平台
type MainServiceController struct {
	controllers.BaseController
}

func init() {
	mux := core.NewServeMux()
	mux.EventHandleFunc(request.EventTypeSubscribe, requestSubscribeEventHandler)
	msgHandler = mux
	msgServer = core.NewServer(viper.GetString("wx.oriId"), viper.GetString("wx.appId"), viper.GetString("wx.token"), viper.GetString("wx.encodedAESKey"), msgHandler, nil)
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

func requestSubscribeEventHandler(ctx *core.Context) {

	msg := string(ctx.MsgPlaintext[:])

	logs.Info("收到服务号关注 事件:\n", msg)

	event := request.GetSubscribeEvent(ctx.MixedMsg)

	var articleList []response.Article
	article2 := response.Article{
		Title:       "test",
		Description: "test110",
		PicURL:      "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1585741426158&di=0ee3f0387e2dd8c562bb1b9cea99866e&imgtype=0&src=http%3A%2F%2Fpic1.xtuan.com%2Fupload%2Fimage%2F20131128%2F09555565911_w.jpg",
		URL:         "https://www.baidu.com",
	}
	articleList = append(articleList, article2)

	resp := response.NewNews(event.FromUserName, event.ToUserName, event.CreateTime, articleList)
	ctx.AESResponse(resp, 0, "", nil)
	return
}
