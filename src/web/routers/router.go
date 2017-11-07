/*
Create by zhangzemian on 2017/11/03
*/

package routers

import (
	"web/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/regist", &controllers.RegisterController{})
}
