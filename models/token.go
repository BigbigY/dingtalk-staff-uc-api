package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func GetCompanyDingTalkClient() *DingTalkClient {
	// CorpID := os.Getenv("corpid")
	// CorpSecret := os.Getenv("corpsecret")
	CorpID := beego.AppConfig.String("corpid")
	CorpSecret := beego.AppConfig.String("corpsecret")
	logs.Info("token:***%v***,***%v***", CorpID, CorpSecret)
	config := &DTConfig{
		CorpID:     CorpID,
		CorpSecret: CorpSecret,
	}
	logs.Info("config", config)
	c := NewDingTalkCompanyClient(config)
	return c
}
