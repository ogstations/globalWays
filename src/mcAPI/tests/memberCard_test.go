// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package tests

import (
	"bytes"
	"encoding/json"
	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
	"mcAPI/models"
	_ "mcAPI/routers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMemberCardsPost(t *testing.T) {

	reqMsg := &models.ReqNewMemberCards{
		MII:         6,
		CPI:         32,
		CDI:         86,
		ChannelType: 1,
		ChannelId:   2,
		CardCnt:     100,
	}

	bodyBytes, _ := json.Marshal(reqMsg)
	r, _ := http.NewRequest("POST", "/v1/memberCard/", bytes.NewReader(bodyBytes))
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestPost", "Code[%v]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

func TestMemberCardGetAll(t *testing.T) {
	r, _ := http.NewRequest("GET", "/v1/memberCard?page=1&size=20", nil)
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

func TestMemberCardGetOne(t *testing.T) {
	r, _ := http.NewRequest("GET", "/v1/memberCard/1", nil)
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
