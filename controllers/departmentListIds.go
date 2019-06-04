package controllers

import (
	"dingtalk-staff-uc-api/models"

	"github.com/astaxie/beego"
)

// Operations about Users
type DepartmenListIdsController struct {
	beego.Controller
}

func (u *DepartmenListIdsController) Get() {
	uid := u.GetString(":uid")
	if uid != "" {
		user, err := models.GetUser(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}
