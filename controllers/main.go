package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Prepare() {
	c.Layout = "layout/application.tpl"
}

func (c *MainController) Home() {
	c.TplName = "main/home.tpl"
}