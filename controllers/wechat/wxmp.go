package wechat

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"gopkg.in/chanxuehong/wechat.v2/mp/core"
	"gopkg.in/chanxuehong/wechat.v2/mp/message/callback/request"
	"gopkg.in/chanxuehong/wechat.v2/mp/message/callback/response"
	"io/ioutil"
	"net/http"
	"wxTribe/controllers"
	"wxTribe/dto"
)

var (
	msgHandler core.Handler
	msgServer  *core.Server

	accessTokenServer core.AccessTokenServer = core.NewDefaultAccessTokenServer(viper.GetString("wx.appId"), viper.GetString("wx.appSecret"), nil)
	weChatClient      *core.Client           = core.NewClient(accessTokenServer, nil)
)

// 微信公众平台
type MainServiceController struct {
	controllers.BaseController
}

func init() {
	mux := core.NewServeMux()
	mux.EventHandleFunc(request.EventTypeSubscribe, requestSubscribeEventHandler)
	mux.MsgHandleFunc(request.MsgTypeText, textMsgHandler)
	msgHandler = mux
	msgServer = core.NewServer(viper.GetString("wx.oriId"), viper.GetString("wx.appId"), viper.GetString("wx.token"), viper.GetString("wx.encodedAESKey"), msgHandler, nil)
}

// 微信公众平台
func (c *MainServiceController) Get() {

	logs.Info("=======wxmp Get Start======>")

	msgServer.ServeHTTP(c.Ctx.ResponseWriter, c.Ctx.Request, nil)

	logs.Info("=======wxmp Get End======>")

	c.ServeXML()
}

// 微信公众平台
func (c *MainServiceController) Post() {

	logs.Info("=======wxmp Post Start======>")

	msgServer.ServeHTTP(c.Ctx.ResponseWriter, c.Ctx.Request, nil)

	logs.Info("=======wxmp Post End======>")
}

// textMsgHandler 收到文本消息
func textMsgHandler(ctx *core.Context) {
	logs.Info("收到文本消息:\n", getRequestXml(ctx))
	msg := request.GetText(ctx.MixedMsg)
	switch msg.Content {
	default:
		// https://openai.weixin.qq.com/openapi/message/TOKEN 普通消息接口 只签名不加密
		signedToken, _ := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), jwt.MapClaims{"username": msg.FromUserName, "msg": msg.Content}).SignedString([]byte(viper.GetString("airobot-benben.encodingAESKey")))
		resp, _ := http.Post(fmt.Sprintf("https://openai.weixin.qq.com/openapi/message/%s?query=%s", viper.GetString("airobot-benben.token"), signedToken), "", nil)
		defer resp.Body.Close()
		message, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("+++++++textMsgHandler message：%s+++++++++++\n", string(message))
		var bot dto.AiBot
		json.Unmarshal(message, &bot)
		if resp := response.NewText(msg.FromUserName, msg.ToUserName, msg.CreateTime, bot.Answer); bot.Answer != "" {
			ctx.AESResponse(resp, 0, "", nil)
			return
		}
	}
	ctx.NoneResponse()
}

// requestSubscribeEventHandler 新用户关注
func requestSubscribeEventHandler(ctx *core.Context) {

	logs.Info("收到服务号关注 事件:\n", getRequestXml(ctx))

	event := request.GetSubscribeEvent(ctx.MixedMsg)

	var articleList []response.Article
	article2 := response.Article{
		Title:       "GopherCon",
		Description: "🎉🎉🎉欢迎来到GopherCon！",
		PicURL:      "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1585741426158&di=0ee3f0387e2dd8c562bb1b9cea99866e&imgtype=0&src=http%3A%2F%2Fpic1.xtuan.com%2Fupload%2Fimage%2F20131128%2F09555565911_w.jpg",
		URL:         "https://github.com/gopher110",
	}
	articleList = append(articleList, article2)

	resp := response.NewNews(event.FromUserName, event.ToUserName, event.CreateTime, articleList)
	ctx.AESResponse(resp, 0, "", nil)
	return
}

// getRequestXml 获取请求 xml 信息
func getRequestXml(ctx *core.Context) (retVal string) {
	retVal = string(ctx.MsgPlaintext[:])
	return
}
