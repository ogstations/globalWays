// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"html/template"
	"time"
	"utils/page"
	"utils/ipv4"
)

type Admin struct {
	Id            int64
	UserName      string    `orm:"unique;column(user_name)" valid:"Required;Match(^\\w*$)"`
	Password      string    `orm:"column(password)" valid:"Required"`
	Email         string    `orm:"size(32)" valid:"Email;MaxSize(32)"`
	Mobile        string    `orm:"size(11) valid:"Phone""`
	RealName      string    `orm:"column(real_name);size(32)"`
	RoleId        int64     `orm:"column(role_id);index"`
	Role          *Role     `orm:"-"`
	LastLoginIp   string    `orm:"column(lastLogin_ip);size(32)" valid:"IP"`
	LastLoginTime time.Time `orm:"column(lastLogin_time)"`
	Lang          string    `orm:"column(lang);type(char);size(5)"`
	Status        bool
	Created       time.Time `orm:"auto_now_add"`
	Updated       time.Time `orm:"auto_now"`
}

func (m *Admin) TableName() string {
	return "admin"
}

// 验证
func (m *Admin) Valid(v *validation.Validation) {
	//	v.SetError("error", "error message")
}

// 管理员列表
func ListAdmin(pager *page.Page, roleId int64, ormer orm.Ormer) ([]*Admin, template.HTML) {

	adminList := make([]*Admin, 0)
	adminCnt := int64(0)
	querySeter := ormer.QueryTable(new(Admin))

	if roleId > 0 { // 特定角色管理员列表
		adminCnt, _ = querySeter.Filter("role_id", roleId).Count()
	} else { // 全部管理员列表
		adminCnt, _ = querySeter.Count()
	}

	pager.Nums = adminCnt
	pager.Sub_pages = 2
	pageHtml := pager.Show()

	// 如果管理员数量 > 0
	if adminCnt > 0 {
		querySeter.Limit(pager.Perpage, (pager.Current_page-1)*pager.Perpage).All(&adminList)
	}

	// fill role
	for _, admin := range adminList {
		admin.Role = GetRoleById(admin.RoleId, ormer)
	}

	return adminList, pageHtml
}

// 检查管理员是否存在
func isAdminExist(name string, ormer orm.Ormer) bool {
	cnt, _ := ormer.QueryTable(new(Admin)).Filter("user_name", name).Count()
	return cnt != 0
}

// 获取管理员信息by name
func GetAdminByName(name string, ormer orm.Ormer) *Admin {
	admin := new(Admin)
	ormer.QueryTable(admin).Filter("user_name", name).One(admin)
	admin.Role = GetRoleById(admin.RoleId, ormer)

	return admin
}

// 获取管理员信息by id
func GetAdminById(id int64, ormer orm.Ormer) *Admin {
	admin := new(Admin)
	ormer.QueryTable(admin).Filter("id", id).One(admin)
	admin.Role = GetRoleById(admin.RoleId, ormer)

	return admin
}

//更新登录IP&时间
func UpdateAdminLoginInfo(id int64, ormer orm.Ormer) bool {
	ip, _ := ipv4.GetClientIP()
	if _, err := ormer.QueryTable(new(Admin)).Filter("id", id).Update(orm.Params{"lastLogin_ip":ip, "lastLogin_time":time.Now().UTC()}); err != nil {
		return false
	}

	return true
}

// 更新管理员信息
func UpdateAdminInfo(adminInfo *Admin, ormer orm.Ormer) bool {
	if _, err := ormer.QueryTable(new(Admin)).Filter("id", adminInfo.Id).Update(orm.Params{

}); err != nil {
		return false
	}

	return true
}

// 更新管理员密码
func UpdateAdminPwd(adminInfo *Admin, ormer orm.Ormer) bool {
	return true
}
