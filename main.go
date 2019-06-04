package main

import (
	"dingtalk-staff-uc-api/models"
	_ "dingtalk-staff-uc-api/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/toolbox"
)

func init() {
	models.RegisterDB()
}

func main() {
	// Logs
    logs.SetLogger("file",`{"filename":"logs/dingtalk-staff-uc-api.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)

	// logs.Debug("my book is bought in the year of ", 2016)
    // logs.Info("this %s cat is %v years old", "yellow", 3)
    // logs.Warn("json is a type of kv like", map[string]int{"key": 2016})
    // logs.Error(1024, "is a very", "good game")
    // logs.Critical("oh,crash")

	// Crontab
	rsyncCronTime := beego.AppConfig.String("rsyncCronTime")
	tk := toolbox.NewTask("RsyncDepartment", rsyncCronTime, func() error { models.RsyncDepartment(); models.RsyncUser(); return nil })
	toolbox.AddTask("RsyncDepartment", tk)
	logs.Info("Start sync [%v] DingTalk to DB...",rsyncCronTime)
	toolbox.StartTask()
	defer toolbox.StopTask()

	// API Document
	if beego.BConfig.RunMode == "dev" || beego.BConfig.RunMode == "prod" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	// Debug model default = false
	debug, err := beego.AppConfig.Bool("debug")
	if err != nil {
		logs.Info("debug mode is false")
		debug = false
	}
	orm.Debug = debug

	beego.Run()
}
