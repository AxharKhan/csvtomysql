package routers

import (
	"csvtomysql/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/store-files", &controllers.MainController{}, "post:StoreFiles")
	beego.Router("/read-db", &controllers.MainController{}, "get:ReadDB")

}
