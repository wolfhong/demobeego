package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["demobeego/controllers:ArticleController"] = append(beego.GlobalControllerRouter["demobeego/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "Create",
			Router: `/create`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["demobeego/controllers:ArticleController"] = append(beego.GlobalControllerRouter["demobeego/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/delete`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["demobeego/controllers:ArticleController"] = append(beego.GlobalControllerRouter["demobeego/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/get`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["demobeego/controllers:ArticleController"] = append(beego.GlobalControllerRouter["demobeego/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/update`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["demobeego/controllers:UserController"] = append(beego.GlobalControllerRouter["demobeego/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["demobeego/controllers:UserController"] = append(beego.GlobalControllerRouter["demobeego/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get","post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["demobeego/controllers:UserController"] = append(beego.GlobalControllerRouter["demobeego/controllers:UserController"],
		beego.ControllerComments{
			Method: "UpdateProfile",
			Router: `/profile/update`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["demobeego/controllers:UserController"] = append(beego.GlobalControllerRouter["demobeego/controllers:UserController"],
		beego.ControllerComments{
			Method: "Signup",
			Router: `/signup`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
