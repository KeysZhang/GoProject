/*
Create by zhangzemian on 2017/11/03
*/

package controllers

import (
	"web/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type LoginController struct {
	beego.Controller
}

type RegisterController struct {
	beego.Controller
}

//controler of homePage
func (c *MainController) Get() {
	c.TplName = "login.html"
	c.Data["IsError"] = false
}

//controler of login
func (login *LoginController) Post() {

	//get the user's message
	Inputs := login.Input()
	username := Inputs.Get("username")
	password := Inputs.Get("password")

	var user models.User = models.User{username, password}

	if models.IsAllRight(user) != ""{
		login.TplName = "login.html"
		login.Data["IsError"] = true
		login.Data["Message"] = models.IsAllRight(user)
	} else {
		login.Data["username"] = user.Username
		login.TplName = "success.html"
	}

}

//controler of registerPage
func (register *RegisterController) Get(){
	register.TplName = "register.html"
}

//controler of register
func (register *RegisterController) Post(){

	//get the user's message
	Inputs := register.Input()
	username := Inputs.Get("username")
	password := Inputs.Get("password")

	var user models.User = models.User{username, password}


	if models.RegisterUser(user) != "" {
		register.TplName = "register.html"
		register.Data["IsError"] = true
		register.Data["Message"] = models.RegisterUser(user)
	} else {
		register.TplName = "login.html"
	}
}