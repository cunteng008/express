package main

import (
	_ "express/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/ttf", "static/ttf")
	beego.SetStaticPath("/images", "static/images")
	beego.SetStaticPath("/css", "static/css")
	beego.SetStaticPath("/js", "static/js")
	beego.SetStaticPath("/music", "static/music")

	beego.Run()
}
