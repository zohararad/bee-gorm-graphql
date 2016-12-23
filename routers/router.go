package routers

import (
	"github.com/zohararad/bee-gorm-graphql/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Home")
	beego.Router("/session/login", &controllers.SessionController{}, "get:New")
	beego.Router("/session/login", &controllers.SessionController{}, "post:Create")
	beego.Router("/session/logout", &controllers.SessionController{}, "get:Destroy")
	beego.DelStaticPath("/static")
	beego.SetStaticPath("/assets", "public/assets")
	addGraphQLHandler()
	beego.Router("*", &controllers.MainController{}, "get:Home")
	addFilters()
}

func addGraphQLHandler() {
	beego.Router("/query", &controllers.GraphqlController{}, "post:Query")
	beego.Router("/graphql", &controllers.GraphqlController{}, "get:Index")
}

func addFilters()  {
	beego.InsertFilter("/*", beego.BeforeRouter, filterLoggedInUser)
}
