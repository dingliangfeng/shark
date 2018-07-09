package controllers

import (
	"github.com/astaxie/beego"
)

type BasePreAppend interface {
	PreAppend()
}

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare() {
	this.Data["name"] = "dingliangfeng"

	if app, ok := this.AppController.(BasePreAppend); ok {
		app.PreAppend()
	}
}
