package models

import (
	//使用beego orm 必备
	"fmt"
	"net/url"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	//使用的数据库 必备
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

type Total struct {
	total  int64 `json:"totle"`
	Member []Member
}

type Member struct {
	ID        int    `json:"id" orm:"pk;column(id);auto"`
	UserID    string `json:"userid" orm:"column(Userid)"`
	UnionID   string `json:"unionid" orm:"column(Unionid)"`
	UserName  string `json:"username" orm:"column(Username)"`
	Mobile    string `json:"mobile" orm:"column(Mobile)"`
	Tel       string `json:"tel" orm:"column(Tel)"`
	WorkPlace string `json:"workplace" orm:"column(Workplace)"`
	Remark    string `json:"remark" orm:"column(Remark)"`
	Order     int64  `json:"order" orm:"column(Order)"`
	IsAdmin   bool   `json:"isadmin" orm:"column(Isadmin)"`
	IsBoss    bool   `json:"isboss" orm:"column(Isboss)"`
	IsHide    bool   `json:"ishide" orm:"column(Ishide)"`
	IsLeader  bool   `json:"isleader" orm:"column(Isleader)"`
	Name      string `json:"name" orm:"column(Name)"`
	Active    bool   `json:"active" orm:"column(Active)"`
	//DepartmentID []*Department `orm:"reverse(many)"`
	DepartmentID int    `json:"departmentid" orm:"column(Departmentid)"`
	Position     string `json:"position" orm:"column(Position)"`
	Email        string `json:"email" orm:"column(Email)"`
	Avatar       string `json:"avatar" orm:"column(Avatar)"`
	Jobnumber    string `json:"jobnumber" orm:"column(Jobnumber)"`
	Extattr      string `json:"extattr" orm:"column(Extattr)"`
	//Extattr    interface{} `orm:"column(Extattr)"`
	CreatedTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedTime time.Time `orm:"auto_now;type(datetime)"`
}

func RegisterDB() {

	dbUser := beego.AppConfig.String("db_user")
	dbPwd := beego.AppConfig.String("db_passowrd")
	dbHost := beego.AppConfig.String("db_host")
	dbPort := beego.AppConfig.String("db_port")
	dbName := beego.AppConfig.String("db_name")
	maxIdle, err := beego.AppConfig.Int("db_maxidle")
	if err != nil {
		maxIdle = 30
	}
	maxConn, err := beego.AppConfig.Int("db_maxconn")
	if err != nil {
		maxConn = 30
	}

	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPwd, dbHost, dbPort, dbName) + "&loc=" + url.QueryEscape("Asia/Shanghai")

	//注册默认数据库
	orm.RegisterDataBase("default", "mysql", dbURL, maxIdle, maxConn)

	//注册 model
	orm.RegisterModel(new(Member))
	orm.RegisterModel(new(Department))

	// 同步(第2个参数代表创建表，第二个参数代表更新表)
	orm.RunSyncdb("default", false, true)
}
