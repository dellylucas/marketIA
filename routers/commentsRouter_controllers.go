package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["marketIA/controllers:CategoriesController"] = append(beego.GlobalControllerRouter["marketIA/controllers:CategoriesController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["marketIA/controllers:CategoriesController"] = append(beego.GlobalControllerRouter["marketIA/controllers:CategoriesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["marketIA/controllers:CategoriesController"] = append(beego.GlobalControllerRouter["marketIA/controllers:CategoriesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:catID`,
            AllowHTTPMethods: []string{"Delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["marketIA/controllers:CategoriesController"] = append(beego.GlobalControllerRouter["marketIA/controllers:CategoriesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["marketIA/controllers:PayMethodsController"] = append(beego.GlobalControllerRouter["marketIA/controllers:PayMethodsController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["marketIA/controllers:PayMethodsController"] = append(beego.GlobalControllerRouter["marketIA/controllers:PayMethodsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["marketIA/controllers:PayMethodsController"] = append(beego.GlobalControllerRouter["marketIA/controllers:PayMethodsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["marketIA/controllers:PayMethodsController"] = append(beego.GlobalControllerRouter["marketIA/controllers:PayMethodsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:payID`,
            AllowHTTPMethods: []string{"Delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["marketIA/controllers:ServiceController"] = append(beego.GlobalControllerRouter["marketIA/controllers:ServiceController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["marketIA/controllers:ServiceController"] = append(beego.GlobalControllerRouter["marketIA/controllers:ServiceController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["marketIA/controllers:ServiceController"] = append(beego.GlobalControllerRouter["marketIA/controllers:ServiceController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["marketIA/controllers:ServiceController"] = append(beego.GlobalControllerRouter["marketIA/controllers:ServiceController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:serviceID`,
            AllowHTTPMethods: []string{"Delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["marketIA/controllers:ShoppingController"] = append(beego.GlobalControllerRouter["marketIA/controllers:ShoppingController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["marketIA/controllers:ShoppingController"] = append(beego.GlobalControllerRouter["marketIA/controllers:ShoppingController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["marketIA/controllers:ShoppingController"] = append(beego.GlobalControllerRouter["marketIA/controllers:ShoppingController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["marketIA/controllers:ShoppingController"] = append(beego.GlobalControllerRouter["marketIA/controllers:ShoppingController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:serviceID`,
            AllowHTTPMethods: []string{"Delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["marketIA/controllers:StoreController"] = append(beego.GlobalControllerRouter["marketIA/controllers:StoreController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["marketIA/controllers:StoreController"] = append(beego.GlobalControllerRouter["marketIA/controllers:StoreController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["marketIA/controllers:StoreController"] = append(beego.GlobalControllerRouter["marketIA/controllers:StoreController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["marketIA/controllers:StoreController"] = append(beego.GlobalControllerRouter["marketIA/controllers:StoreController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:storeID`,
            AllowHTTPMethods: []string{"Delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["marketIA/controllers:UserController"] = append(beego.GlobalControllerRouter["marketIA/controllers:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["marketIA/controllers:UserController"] = append(beego.GlobalControllerRouter["marketIA/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["marketIA/controllers:UserController"] = append(beego.GlobalControllerRouter["marketIA/controllers:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["marketIA/controllers:UserController"] = append(beego.GlobalControllerRouter["marketIA/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:userID`,
            AllowHTTPMethods: []string{"Delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["marketIA/controllers:UserController"] = append(beego.GlobalControllerRouter["marketIA/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
