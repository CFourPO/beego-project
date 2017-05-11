package routers

import (
	"first-project/controllers"
	"github.com/astaxie/beego"
)

func init() {
    
    beego.Router("/", &controllers.MainController{})
    beego.SetStaticPath("/down1", "down1")
    
}
