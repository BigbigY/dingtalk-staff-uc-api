package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["dingtalk-staff-uc-api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["dingtalk-staff-uc-api/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dingtalk-staff-uc-api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["dingtalk-staff-uc-api/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dingtalk-staff-uc-api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["dingtalk-staff-uc-api/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dingtalk-staff-uc-api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["dingtalk-staff-uc-api/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dingtalk-staff-uc-api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["dingtalk-staff-uc-api/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dingtalk-staff-uc-api/controllers:StaffDepartmentController"] = append(beego.GlobalControllerRouter["dingtalk-staff-uc-api/controllers:StaffDepartmentController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dingtalk-staff-uc-api/controllers:StaffDepartmentController"] = append(beego.GlobalControllerRouter["dingtalk-staff-uc-api/controllers:StaffDepartmentController"],
		beego.ControllerComments{
			Method: "GetAboutDepartment",
			Router: `/about/querydepartment`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dingtalk-staff-uc-api/controllers:StaffDepartmentController"] = append(beego.GlobalControllerRouter["dingtalk-staff-uc-api/controllers:StaffDepartmentController"],
		beego.ControllerComments{
			Method: "GetDepartment",
			Router: `/querydepartment`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dingtalk-staff-uc-api/controllers:StaffDepartmentController"] = append(beego.GlobalControllerRouter["dingtalk-staff-uc-api/controllers:StaffDepartmentController"],
		beego.ControllerComments{
			Method: "GetID",
			Router: `/querydepartmentid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dingtalk-staff-uc-api/controllers:StaffUserController"] = append(beego.GlobalControllerRouter["dingtalk-staff-uc-api/controllers:StaffUserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dingtalk-staff-uc-api/controllers:StaffUserController"] = append(beego.GlobalControllerRouter["dingtalk-staff-uc-api/controllers:StaffUserController"],
		beego.ControllerComments{
			Method: "GetAboutUser",
			Router: `/about/queryuser`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dingtalk-staff-uc-api/controllers:StaffUserController"] = append(beego.GlobalControllerRouter["dingtalk-staff-uc-api/controllers:StaffUserController"],
		beego.ControllerComments{
			Method: "GetID",
			Router: `/queryid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dingtalk-staff-uc-api/controllers:StaffUserController"] = append(beego.GlobalControllerRouter["dingtalk-staff-uc-api/controllers:StaffUserController"],
		beego.ControllerComments{
			Method: "GetUser",
			Router: `/queryuser`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dingtalk-staff-uc-api/controllers:SyncController"] = append(beego.GlobalControllerRouter["dingtalk-staff-uc-api/controllers:SyncController"],
		beego.ControllerComments{
			Method: "PostSync",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dingtalk-staff-uc-api/controllers:UserController"] = append(beego.GlobalControllerRouter["dingtalk-staff-uc-api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dingtalk-staff-uc-api/controllers:UserController"] = append(beego.GlobalControllerRouter["dingtalk-staff-uc-api/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dingtalk-staff-uc-api/controllers:UserController"] = append(beego.GlobalControllerRouter["dingtalk-staff-uc-api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dingtalk-staff-uc-api/controllers:UserController"] = append(beego.GlobalControllerRouter["dingtalk-staff-uc-api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dingtalk-staff-uc-api/controllers:UserController"] = append(beego.GlobalControllerRouter["dingtalk-staff-uc-api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dingtalk-staff-uc-api/controllers:UserController"] = append(beego.GlobalControllerRouter["dingtalk-staff-uc-api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dingtalk-staff-uc-api/controllers:UserController"] = append(beego.GlobalControllerRouter["dingtalk-staff-uc-api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
