package main

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"

	_ "simqo.com/mygospace/learngo/example/sample_im/routers"
)

const (
	VERSION = "0.0.1"
)

func main() {
	beego.Info(beego.BConfig.AppName, VERSION)
	beego.AddFuncMap("i18n", i18n.Tr)
	beego.Run()
}
