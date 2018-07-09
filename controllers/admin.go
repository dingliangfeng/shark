package controllers

type AdminController struct {
	BaseController
}

func (this *AdminController) PreAppend() {
	this.Data["id"] = 10001
}

func (this *AdminController) Index() {
	this.Data["view"] = "index.tpl"
	this.TplName = "admin/index.tpl"
}