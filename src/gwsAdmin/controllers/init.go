// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package controllers

import (
	"github.com/astaxie/beego"
	"strings"
	"github.com/beego/i18n"
	"github.com/howeyc/fsnotify"
	"path/filepath"
	"gwsAdmin/models"
	"gwsAdmin/utils"
)

// 初始化地理位置
func initLocales() {
	// Initialized language type list.
	langs := strings.Split(beego.AppConfig.String("lang::types"), "|")
	names := strings.Split(beego.AppConfig.String("lang::names"), "|")
	langTypes = make([]*langType, 0, len(langs))
	for i, v := range langs {
		langTypes = append(langTypes, &langType{
				Lang: v,
				Name: names[i],
			})
	}

	for _, lang := range langs {
		beego.Trace("Loading language: " + lang)
		if err := i18n.SetMessage(lang, "conf/"+"locale_"+lang+".ini"); err != nil {
			beego.Error("Fail to set message file: " + err.Error())
			return
		}
	}
}

// 初始化watcher
func initWatcher() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic("Failed start app watcher: " + err.Error())
	}

	go func() {
		for {
			select {
			case event := <-watcher.Event:
				switch filepath.Ext(event.Name) {
				case ".ini":
					beego.Info(event)

					if err := i18n.ReloadLangs(); err != nil {
						beego.Error("Conf Reload: ", err)
					}

					beego.Info("Config Reloaded")
				}
			}
		}
	}()

	if err := watcher.WatchFlags("conf", fsnotify.FSN_MODIFY); err != nil {
		beego.Error(err)
	}
}

func InitApp() {
	initLocales()
	initWatcher()
}

// 根据Session获取管理员信息
func getAdminInfoBySession(c beego.Controller) (*models.Admin, bool) {
	if admin := c.GetSession(utils.Session_Admin); admin != nil {
		return admin.(*models.Admin), true
	}

	return nil, false
}
