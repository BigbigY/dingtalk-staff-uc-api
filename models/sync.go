package models

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

// DepartmentResp from the DingTalk get department id List
func DepartmentResp() (departmentListResponse DepartmentListResponse, departmentList []int) {
	c := GetCompanyDingTalkClient()
	c.RefreshCompanyAccessToken()

	// 获取所有部门ID列表 https://open-doc.dingtalk.com/microapp/serverapi2/dubakq#-1 {department:{name:SRE部门,deparid:123}:2},{name:XXXX部门,deparid:122}}
	departmentListResponse, err := c.DepartmentList(1, "zh_CN")
	if err != nil {
		logs.Info("Get list department,err:", err)
		return
	}

	// 生成所有部门列表[1,2,3,4]
	departmentList = make([]int, 300, 500)
	for i := 0; i < len(departmentListResponse.Department); i++ {
		departmentList[i] = departmentListResponse.Department[i].Id
	}

	return
}

// RsyncDepartment from the DingTalk sync.
func RsyncDepartment() {
	/*
		1、查询部门列表，获取ID
		2、根据ID循环获取每个用户信息
		3、新增到数据库
	*/

	c := GetCompanyDingTalkClient()
	c.RefreshCompanyAccessToken()
	logs.Info("开始同步Department")

	// 获取所有部门ID列表 https://open-doc.dingtalk.com/microapp/serverapi2/dubakq#-1 {department:{name:SRE部门,deparid:123}:2},{name:XXXX部门,deparid:122}}
	// departmentListResponse, err := c.DepartmentList(1, "zh_CN")
	// if err != nil {
	// 	fmt.Println("获取部门列表", err)
	// }

	departmentListResponse, _ := DepartmentResp()

	//写入Department表
	for i := 0; i < len(departmentListResponse.Department); i++ {
		if departmentListResponse.Department[i].Id == 1 {
			continue
		}
		// 插入数据库
		o := orm.NewOrm()
		departmentInfo := Department{
			Id:              departmentListResponse.Department[i].Id,
			Name:            departmentListResponse.Department[i].Name,
			ParentId:        departmentListResponse.Department[i].ParentId,
			CreateDeptGroup: departmentListResponse.Department[i].CreateDeptGroup,
			AutoAddUser:     departmentListResponse.Department[i].AutoAddUser,
		}
		if created, id, err := o.ReadOrCreate(&departmentInfo, "Id"); err == nil {
			if created {
				logs.Info("New Insert an object. Id:", id)
			} else {
				// Update the existing departmentID field information
				departmentInfo.Name = departmentListResponse.Department[i].Name
				departmentInfo.ParentId = departmentListResponse.Department[i].ParentId

				o.Update(&departmentInfo, "Name", "Parentid")
				logs.Info("Get an department object. Id:", id)
			}
		}

	}

	// // 生成所有部门列表[1,2,3,4]
	// departmentList := make([]int, 300, 500)
	// for i := 0; i < len(departmentListResponse.Department); i++ {
	// 	departmentList[i] = departmentListResponse.Department[i].Id
	// }

}

// RsyncUser from the DingTalk sync.
func RsyncUser() {

	c := GetCompanyDingTalkClient()
	c.RefreshCompanyAccessToken()
	logs.Info("开始同步User")
	_, departmentList := DepartmentResp()
	// step2 : 插入用户数据
	// 部门循环详情(did部门ID),取部门下的成员
	for _, did := range departmentList {

		if did == 0 {
			continue
		}
		userListResponse, err := c.UserList(did)
		//fmt.Println(resp)
		if err != nil {
			logs.Info("获取部门成员", err)
		}

		// 插入数据库
		o := orm.NewOrm()
		fmt.Println(o)

		for i := 0; i < len(userListResponse.UserList); i++ {
			// 计划转字符串
			// var toDepartment string
			// for j := 0; j < len(resp.UserList[i].Department); j++ {
			// 	toDepartment += strconv.Itoa(resp.UserList[i].Department[j]) + " "
			// 	fmt.Println(toDepartment)
			// }

			// get username
			username := strings.Split(userListResponse.UserList[i].Email, "@")
			
			// create people
			people := Member{
				UserID:       userListResponse.UserList[i].UserID,
				UnionID:      userListResponse.UserList[i].UnionID,
				UserName:     username[0],
				Mobile:       userListResponse.UserList[i].Mobile,
				Tel:          userListResponse.UserList[i].Tel,
				WorkPlace:    userListResponse.UserList[i].WorkPlace,
				Remark:       userListResponse.UserList[i].Remark,
				Order:        userListResponse.UserList[i].Order,
				IsAdmin:      userListResponse.UserList[i].IsAdmin,
				IsBoss:       userListResponse.UserList[i].IsBoss,
				IsHide:       userListResponse.UserList[i].IsHide,
				IsLeader:     userListResponse.UserList[i].IsLeader,
				Name:         userListResponse.UserList[i].Name,
				Active:       userListResponse.UserList[i].Active,
				DepartmentID: userListResponse.UserList[i].Department[0],
				Position:     userListResponse.UserList[i].Position,
				Email:        userListResponse.UserList[i].Email,
				Avatar:       userListResponse.UserList[i].Avatar,
				Jobnumber:    userListResponse.UserList[i].Jobnumber,
				Extattr:      userListResponse.UserList[i].Extattr,
			}

			//insert
			if created, id, err := o.ReadOrCreate(&people, "Userid"); err == nil {
				if created {
					// add new pople info
					logs.Info("New Insert an object. Id:", id)
				} else {
					// Update the existing user field information
					people.Mobile = userListResponse.UserList[i].Mobile
					people.WorkPlace = userListResponse.UserList[i].WorkPlace
					people.Avatar = userListResponse.UserList[i].Avatar
					people.DepartmentID = userListResponse.UserList[i].Department[0]
					people.Position = userListResponse.UserList[i].Position

					o.Update(&people, "Mobile", "Workplace", "Avatar", "Position","Departmentid")
					logs.Info("Get an user object. Id:", id)
				}
			}
		}

	}
}
