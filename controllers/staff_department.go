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
type StaffDepartmentController struct {
	beego.Controller
}

// GetDepartmentID ...
// @Title Get DepartmentID
// @Description get Department by departmentid
// @Param	departmentid	query	int	true	"Fields returned. e.g. col1,col2 ..."
// @Success 200 {object} models.Department
// @Failure 403 :departmentid is empty
// @router /querydepartmentid [get]
func (c *StaffDepartmentController) GetID() {
	var departmentid int

	// id: 1 (default is 1)
	if v, err := c.GetInt("departmentid"); err == nil {
		departmentid = v
	}

	l, err := models.GetStaffByDepartmentID(departmentid)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// @Title GetDepartment
// @Description get Department by Departmentid
// @Param	name   query 	string	true	"The key for staticblock"
// @Success 200 {object} models.Department
// @Failure 403 :name is empty
// @router /querydepartment [get]
func (c *StaffDepartmentController) GetDepartment() {
	var name string

	if v := c.GetString("name"); v != "" {
		name = v
	}

	v, err := models.GetStaffDepartment(name)

	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}

	c.ServeJSON()
}

// @Title GetAboutDepartment
// @Description get department by departmentid
// @Param	name	query	string	true	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ...Id、Name、ParentId、CreateDeptGroup、AutoAddUser、CreatedTime、UpdatedTime"
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Department
// @Failure 403 :name is empty
// @router /about/querydepartment [get]
func (c *StaffDepartmentController) GetAboutDepartment() {
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

	var name string

	if v := c.GetString("name"); v != "" {
		name = v
	}

	fmt.Println(name)
	ml, mf,err := models.GetAllAboutDepartment(name, query, fields, sortby, order, offset, limit)
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

// @Title Get All Department
// @Description get App
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ...Id、Name、ParentId、CreateDeptGroup、AutoAddUser、CreatedTime、UpdatedTime"
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Department
// @Failure 403
// @router  / [get]
func (c *StaffDepartmentController) GetAll() {
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

	ml, mf, err := models.GetAllDepartment(query, fields, sortby, order, offset, limit)
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
