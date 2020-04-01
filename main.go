package main

import (
	"github.com/spf13/viper"
	_ "wxTribe/initial"
	_ "wxTribe/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run(viper.GetString("addr"))
}
