package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)


type TotalMember struct {
	Totalitem  int64 `json:"totalitem"`
	Totalpages int64 `json:"totalpages"`
	List []Member 	 `json:"list"`
}

type FieldTotalMember struct {
	Totalitem  int64 `json:"totalitem"`
	Totalpages int64 `json:"totalpages"`
	List []interface{} `json:"list"`
}

// GetStaffByID retrieves member by Id. Returns error if
// Id doesn't exist
func GetStaffByID(id int) (v *Member, err error) {
	o := orm.NewOrm()
	v = &Member{ID: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetStaffUser retrieves member by username. Returns error if
func GetStaffUser(username string) (v *Member, err error) {
	o := orm.NewOrm()
	var user Member
	err = o.QueryTable("member").Filter("Username", username).One(&user)
	if err == orm.ErrMultiRows {
		// 多条的时候报错
		logs.Info("Returned Multi Rows Not One")
	}
	if err == orm.ErrNoRows {
		// 没有找到记录
		logs.Info("Not row found")
	}
	logs.Info("user:", &user)
	return &user, err
}

// GetAllUser retrieves all App matches certain condition. Returns empty list if
// no records exist
func GetAllUser(query map[string]string, fields []string, sortby []string, order []string, offset int64, limit int64) (ml TotalMember,mf FieldTotalMember, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Member))
	count, _ := qs.Count()

	// query k=v
	fmt.Printf("query:%v\n", query)
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return ml,mf, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return ml,mf, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return ml,mf, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return ml,mf, errors.New("Error: unused 'order' fields")
		}
	}

	var list []interface{}
	var l []Member
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(limit, offset).All(&l, fields...); err == nil {
		ml.Totalitem = count
		mf.Totalitem = count
		if (count % limit) == 0 {
			ml.Totalpages = count / limit
			mf.Totalpages = count / limit
		} else {
			ml.Totalpages = count/limit + 1
			mf.Totalpages = count/limit + 1
		}

		if len(fields) == 0 {
			for _, v := range l {
				ml.List = append(ml.List, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				fmt.Println("===",val)
				for _, fname := range fields {
					m[strings.ToLower(fname)] = val.FieldByName(fname).Interface()
				}
				list = append(list, m)
			}
			mf.List = list
		}
		return ml,mf, nil
	}
	fmt.Println(ml,mf)
	return ml,mf, err
}

// GetAllAboutUser retrieves all App matches certain condition. Returns empty list if
// no records exist
func GetAllAboutUser(username string, query map[string]string, fields []string, sortby []string, order []string, offset int64, limit int64) (ml TotalMember,mf FieldTotalMember, err error) {
	//var user []Member
	o := orm.NewOrm()
	qs := o.QueryTable(new(Member))
	if err == orm.ErrNoRows {
		// 没有找到记录
		fmt.Printf("Not row found")
	}
	count, _ := qs.Count()

	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return ml,mf, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return ml,mf, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return ml,mf, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return ml,mf, errors.New("Error: unused 'order' fields")
		}
	}

	var list []interface{}
	var l []Member
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(limit, offset).Filter("Username__icontains", username).All(&l, fields...); err == nil {
		ml.Totalitem = count
		mf.Totalitem = count
		if (count % limit) == 0 {
			ml.Totalpages = count / limit
			mf.Totalpages = count / limit
		} else {
			ml.Totalpages = count/limit + 1
			mf.Totalpages = count/limit + 1
		}

		if len(fields) == 0 {
			for _, v := range l {
				ml.List = append(ml.List, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[strings.ToLower(fname)] = val.FieldByName(fname).Interface()
				}
				list = append(list, m)
			}
			mf.List = list
		}
		return ml,mf, nil
	}
	fmt.Println(ml,mf)
	return ml,mf, err
}