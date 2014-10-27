// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package controllers

import (
	"gwsAdmin/models"
	"utils/ipv4"
)

// 插入操作日志
func addLogs(uid int64, controller, action, desc string) {
	//******************************************
	//管理员日志
	ip, _ := ipv4.GetClientIP()
	logs := &models.Logs{
		Uid:        uid,
		Controller: controller,
		Action:     action,
		Ip:         ip,
		Desc:       desc,
	}
	models.AddLogs(logs, models.Writter)
	//*****************************************
}
