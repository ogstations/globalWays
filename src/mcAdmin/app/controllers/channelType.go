// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package controllers

import "github.com/revel/revel"

type ChannelType struct {
	*revel.Controller
}

func (c *ChannelType) ListChannelType() revel.Result {
	return c.Render()
}

func (c *ChannelType) GetChannelTypeById() revel.Result {
	return c.Render()
}

func (c *ChannelType) NewChannelType() revel.Result {
	return c.Render()
}
