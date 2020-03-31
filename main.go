package main

import (
	_ "wxTribe/initial"
	_ "wxTribe/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run(":9000")
}
