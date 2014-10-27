// Package: models
// File: logs.go
// Created by mint
// Useage: 管理员后台操作记录
// DATE: 14-6-27 22:58
package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Logs struct {
	Id         int64
	Uid        int64     `orm:"column(uid)"`
	Admin      *Admin    `orm:"-"`
	Controller string    `orm:"size(50);column(module)"`
	Action     string    `orm:"size(100);column(action)"`
	Ip         string    `orm:"size(15);column(ip)"`
	Desc       string    `orm:"type(text)"`
	Created    time.Time `orm:"auto_now_add"`
}

func (m *Logs) TableName() string {
	return "logs"
}

//添加日志记录
func AddLogs(log *Logs, ormer orm.Ormer) bool {
	if _, err := ormer.Insert(log); err != nil {
		return false
	}

	return true
}

//清空日志
//TODO 清空日志是将所有后台操作日志清空么？还是清空当前管理员日志？谁有清空的权限？
func DelLogs() bool {
	return false
}

/*
func (L *Logs) DelAll() bool {
	sql := "TRUNCATE Logs;"
	_, err := DB_Write.Exec(sql)
	if err != nil {
		revel.WARN.Printf("清空日志错误: %v", err)
		return false
	}

	return true
}

//获取日志列表
func (L *Logs) GetByAll(search string, Page int64, Perpage int64) (logs_arr []*Logs, html template.HTML, where map[string]string) {

	logs_list := []*Logs{}

	//查询条件
	var WhereStr string = " 1 AND "

	if len(search) > 0 {
		//解码
		where = utils.DecodeSegment(search)

		revel.WARN.Println(where)

		if len(where["module"]) > 0 {
			WhereStr += " `module`='" + where["module"] + "' AND "
		}

		if len(where["username"]) > 0 {
			admin := new(Admin)
			AdminInfo := admin.GetByName(where["username"])
			WhereStr += " `uid`=" + strconv.Itoa(int(AdminInfo.Id)) + " AND "
		}

		if len(where["realname"]) > 0 {
			admin := new(Admin)
			AdminInfo := admin.GetByRealName(where["realname"])
			WhereStr += " `uid`='" + strconv.Itoa(int(AdminInfo.Id)) + "' AND "
		}

		if len(where["start_time"]) > 0 {
			WhereStr += " `createtime` >='" + where["start_time"] + " 00:00:00' AND "
		}

		if len(where["end_time"]) > 0 {
			WhereStr += " `createtime` <='" + where["end_time"] + " 23:59:59' AND "
		}
	}

	WhereStr += " 1 "

	//查询总数
	logs := new(Logs)
	Total, err := DB_Read.Where(WhereStr).Count(logs)
	if err != nil {
		revel.WARN.Printf("错误: %v", err)
	}

	//分页
	Pager := new(page.Page)
	if len(search) > 0 {
		Pager.SubPage_link = "/Logs/" + search + "/"
	} else {
		Pager.SubPage_link = "/Logs/"
	}

	Pager.Nums = Total
	Pager.Perpage = Perpage
	Pager.Current_page = Page
	Pager.SubPage_type = 2
	pages := Pager.Show()

	//查询数据
	DB_Read.Where(WhereStr).Limit(int(Perpage), int((Page-1)*Pager.Perpage)).Desc("id").Find(&logs_list)

	if len(logs_list) > 0 {
		admin := new(Admin)

		for i, v := range logs_list {
			logs_list[i].Admin = admin.GetById(v.Uid)
		}
	}

	return logs_list, pages, where
}
*/
