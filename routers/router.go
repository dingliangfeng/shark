package routers

import (
	"github.com/dingliangfeng/shark/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
