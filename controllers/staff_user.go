package controllers

import (
	"errors"
	"fmt"
	"dingtalk-staff-uc-api/models"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

	//使用beego orm 必备

	//使用的数据库 必备
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

// Operations about Users
type StaffUserController struct {
	beego.Controller
}

// GetID ...
// @Title Get ID
// @Description get user by id
// @Param	id	query	int	true	"Fields returned. e.g. col1,col2 ..."
// @Success 200 {object} models.Member
// @Failure 403 :id is empty
// @router /queryid [get]
func (c *StaffUserController) GetID() {
	var id int

	// id: 1 (default is 1)
	if v, err := c.GetInt("id"); err == nil {
		id = v
	}

	l, err := models.GetStaffByID(id)
	logs.Debug("GetID &v", &l)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// @Title GetUser
// @Description get user by Username
// @Param	username   query 	string	true	"The key for staticblock"
// @Success 200 {object} models.Member
// @Failure 403 :username is empty
// @router /queryuser [get]
func (c *StaffUserController) GetUser() {
	var username string

	if v := c.GetString("username"); v != "" {
		username = v
	}

	v, err := models.GetStaffUser(username)
	logs.Debug("GetUser &v", &v)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}

	c.ServeJSON()
}

// @Title GetAboutUser
// @Description get user by Username
// @Param	username	query	string	true	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ...ID、UserID、UnionID、UserName、Mobile、Tel、 WorkPlace、Remark、Order、IsAdmin、IsBoss、IsHide、IsLeader、Name、Active、DepartmentID、Position、Email、Avatar、Jobnumber、Extattr、CreatedTime、Updated、UpdatedTime"
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Member
// @Failure 403 :username is empty
// @router /about/queryuser [get]
func (c *StaffUserController) GetAboutUser() {
	var fields []string
	var sortby []string
	var order []string
	var query map[string]string = make(map[string]string)
	var limit int64 = 10
	var offset int64 = 0

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.Split(cond, ":")
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	var username string

	if v := c.GetString("username"); v != "" {
		username = v
	}

	fmt.Println(username)
	ml, mf, err := models.GetAllAboutUser(username, query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		// return value type
		if len(fields) == 0 {
			c.Data["json"] = ml
			logs.Debug("GetAboutUser 不带fields搜索 &v", &ml)
		} else {
			c.Data["json"] = mf
			logs.Debug("GetAboutUser 带fields搜索 &v", &mf)
		}
	}
	c.ServeJSON()
}

// @Title Get All User
// @Description get App
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ...ID、UserID、UnionID、UserName、Mobile、Tel、 WorkPlace、Remark、Order、IsAdmin、IsBoss、IsHide、IsLeader、Name、Active、DepartmentID、Position、Email、Avatar、Jobnumber、Extattr、CreatedTime、Updated、UpdatedTime"
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Member
// @Failure 403
// @router  / [get]
func (c *StaffUserController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query map[string]string = make(map[string]string)
	var limit int64 = 10
	var offset int64 = 0

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.Split(cond, ":")
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	ml, mf, err := models.GetAllUser(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		// return value type
		if len(fields) == 0 {
			c.Data["json"] = ml
			logs.Debug("GetAll 不带fields搜索")
		} else {
			c.Data["json"] = mf
			logs.Debug("GetAll 带fields搜索")
		}

	}
	c.ServeJSON()
}
