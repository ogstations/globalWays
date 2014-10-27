// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package controllers

import "github.com/astaxie/beego"

type PublicController struct {
	BaseController
}

// @router /message [get]
func (c *PublicController) Message() {
	beego.ReadFromRequest(&c.Controller)
	c.TplNames = "public/message.html"
}
