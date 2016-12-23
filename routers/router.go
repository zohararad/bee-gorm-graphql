package routers

import (
	"github.com/zohararad/bee-gorm-graphql/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
