package controllers

import (
	"dingtalk-staff-uc-api/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

	//使用的数据库 必备
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

// Operations about Users
type SyncController struct {
	beego.Controller
}

// @Title Post
// @Description DingTalk sync user to DB
// @Success 200 {object} models.DingTalk
// @router / [post]
func (u *SyncController) PostSync() {
	logs.Info("run RsyncDepartment....")
	models.RsyncDepartment()
	logs.Info("run RsyncUser....")
	models.RsyncUser()
	u.Data["json"] = "ok"
	u.ServeJSON()
}
