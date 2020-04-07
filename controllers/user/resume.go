package user

import (
	"fmt"
	"io/ioutil"
	"os"
	"wxTribe/controllers"
	"wxTribe/utils"
)

type ResumeController struct {
	controllers.BaseController
	lang string
}

func (c *ResumeController) Get() {
	switch c.GetString("lang") {
	case "en":
		c.lang = "en"
	default:
		c.lang = "cn"
	}
	file, _ := os.Open(fmt.Sprintf("%s/../../views/static/resume-%s.pdf", utils.GetCurrentPath(), c.lang))
	buf, _ := ioutil.ReadAll(file)
	c.Output(buf, "application/pdf")
}
