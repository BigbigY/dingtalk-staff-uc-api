// @APIVersion 1.0.0
// @Title DingTalk Staff UC API
// @Description Very cool tools to autogenerate documents for your API
// @Contact yy520it@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"dingtalk-staff-uc-api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/api/v1",
		// beego.NSNamespace("/object",
		// 	beego.NSInclude(
		// 		&controllers.ObjectController{},
		// 	),
		// ),
		// beego.NSNamespace("/user",
		// 	beego.NSInclude(
		// 		&controllers.UserController{},
		// 	),
		// ),
		// Rsync DingTalk to DB
		beego.NSNamespace("/sync",
			beego.NSInclude(
				&controllers.SyncController{},
			),
		),
		beego.NSNamespace("/staff/user",
			beego.NSInclude(
				&controllers.StaffUserController{},
			),
		),
		beego.NSNamespace("/staff/department",
			beego.NSInclude(
				&controllers.StaffDepartmentController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
