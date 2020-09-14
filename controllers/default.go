package controllers

import "github.com/astaxie/beego"

type ApiController struct {
	beego.Controller
}

func (c *ApiController) GetTopPosts() {
	c.Data["json"] = "OK"
	c.ServeJSON()
}
