package routers

import (
	"github.com/astaxie/beego"
	"simqo.com/mygospace/learngo/example/sample_im/controllers"
)

func init()  {
	beego.Router("/", &controllers.AppController{})
	beego.Router("/join", &controllers.AppController{}, "post:Join")
}


