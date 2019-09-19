// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"marketIA/controllers"
	"marketIA/utils"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace(utils.UserPath,
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
