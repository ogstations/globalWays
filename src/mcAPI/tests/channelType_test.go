// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package tests

import (
	"testing"
	"github.com/astaxie/beego"
	"net/http/httptest"
	"net/http"
	"mcAPI/models"
	"bytes"
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"runtime"
	"path/filepath"
	_ "mcAPI/routers"
	"github.com/astaxie/beego/orm"
)

var (
	_ models.ChannelType
	_ json.Decoder
	_ orm.Ormer
	_ bytes.Buffer
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".." + string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

func TestChannelPost(t *testing.T) {

	channelType := &models.ChannelType{
		ChannelName: "test",
		ChannelDesc: "test channel",
	}

	bodyBytes, _ := json.Marshal(channelType)
	r, _ := http.NewRequest("POST", "/v1/channelType/", bytes.NewReader(bodyBytes))
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestPost", "Code[%v]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
			Convey("Status Code Should Be 203", func() {
					So(w.Code, ShouldEqual, 203)
				})
			Convey("The Result Should Not Be Empty", func() {
					So(w.Body.Len(), ShouldBeGreaterThan, 0)
				})
		})
}

func TestChannelGetAll(t *testing.T) {
	r, _ := http.NewRequest("GET", "/v1/channelType/", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestGet", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
			Convey("Status Code Should Be 200", func() {
					So(w.Code, ShouldEqual, 200)
				})
			Convey("The Result Should Not Be Empty", func() {
					So(w.Body.Len(), ShouldBeGreaterThan, 0)
				})
		})
}

func TestChannelGetOne(t *testing.T) {
	r, _ := http.NewRequest("GET", "/v1/channelType/1", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestGet", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
			Convey("Status Code Should Be 200", func() {
					So(w.Code, ShouldEqual, 200)
				})
			Convey("The Result Should Not Be Empty", func() {
					So(w.Body.Len(), ShouldBeGreaterThan, 0)
				})
		})
}

func TestChannelUpdOne(t *testing.T) {

	channelType := &models.ChannelType{
		ChannelName: "test",
		ChannelDesc: "test channel for upd",
	}

	bodyBytes, _ := json.Marshal(channelType)
	r, _ := http.NewRequest("PUT", "/v1/channelType/1", bytes.NewReader(bodyBytes))
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestGet", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
			Convey("Status Code Should Be 200", func() {
					So(w.Code, ShouldEqual, 200)
				})
			Convey("The Result Should Not Be Empty", func() {
					So(w.Body.Len(), ShouldBeGreaterThan, 0)
				})
		})
}
