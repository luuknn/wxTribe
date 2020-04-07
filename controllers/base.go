package controllers

import (
	"github.com/astaxie/beego"
)

// BaseController Base Controller
type BaseController struct {
	beego.Controller
	IsLogin      bool
	UserOpenid   string
	UserUsername string
	UserAvatar   string
}

func (c *BaseController) ResponseMessage(code int, data interface{}, message string) {
	out := make(map[string]interface{})
	out["code"] = code
	out["data"] = data
	out["message"] = message
	c.Data["json"] = out
	c.ServeJSON()
}

func (c *BaseController) Output(bodyContent []byte, contentType string) {
	c.Ctx.Output.ContentType(contentType)
	c.Ctx.Output.Body(bodyContent)
}
