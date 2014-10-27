// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package controllers

import (
	"github.com/astaxie/beego"
	"gwsAdmin/models"
	"github.com/astaxie/beego/context"
)

var (
	_ beego.Controller
	_ context.Context
)

type AjaxController struct {
	BaseController
}

// @router /ajax/pos [post]
func (c *AjaxController) Pos() {
	c.Controller.EnableRender = false

	mid, err := c.GetInt("mid")
	if err == nil {
		pos := models.GetMenuPos(mid, models.Reader)
		c.Controller.Ctx.WriteString(pos)
	}
}
