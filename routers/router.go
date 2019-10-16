// @APIVersion 1.0.0
// @Title API Tienda IA
// @Description Documentacion del Api para tiendas de inteligencia artificial; estos servicios se deben consumir desde la URL: http://52.229.9.122:8085/v1
// @Contact dellylucas@hotmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"marketIA/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/categories",
			beego.NSInclude(
				&controllers.CategoriesController{},
			),
		),
		beego.NSNamespace("/store",
			beego.NSInclude(
				&controllers.StoreController{},
			),
		),
		beego.NSNamespace("/payMet",
			beego.NSInclude(
				&controllers.PayMethodsController{},
			),
		),
		beego.NSNamespace("/service",
			beego.NSInclude(
				&controllers.ServiceController{},
			),
		),
		beego.NSNamespace("/shopping",
			beego.NSInclude(
				&controllers.ShoppingController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
