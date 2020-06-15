package controllers

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"strings"
)

var langTypes []string

func init()  {
	langTypes = strings.Split(beego.AppConfig.String("lang_typess"), "|")
	for _, lang := range langTypes {
		err := i18n.SetMessage(lang, "conf/"+"locale_"+lang+".ini")
		if err != nil {
			beego.Error(err)
			return
		}
	}
}

type baseController struct {
	beego.Controller
	i18n.Locale
}

func (this *baseController) Prepare()  {
	this.Lang = ""
	// 1. 获取lang 通过 Accept-language
	a1 := this.Ctx.Request.Header.Get("Accept-language")
	if len(a1) > 5 {
		a1 = a1[:5]
		if i18n.IsExist(a1) {
			this.Lang = a1
		}
	}

	// 2. 默认语言
	if len(this.Lang) == 0 {
		this.Lang = "en-US"
	}

	// 3. 设置当前使用语言
	this.Data["Lang"] = this.Lang
}

type AppController struct {
	baseController
}

func (this *AppController) Get()  {
	this.TplName = "wecome.tpl"
}

func (this *AppController) Join()  {
	uname := this.GetString("uname")
	tech := this.GetString("tech")

	if len(uname) == 0 {
		this.Redirect("/", 302)
		return
	}

	switch tech {
	case "longpolling":
		this.Redirect("/lp?uname="+uname, 302)
	case "websocket":
		this.Redirect("/ws?uname="+uname, 302)
	default:
		this.Redirect("/", 302)
	}
	return
}
