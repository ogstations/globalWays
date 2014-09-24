// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package controllers

import (
	"encoding/json"
	"github.com/revel/revel"
	"io/ioutil"
	"mcAdmin/app"
	"memberCard"
	"net/http"
	"utils/errors"
	"strconv"
	"mcAdmin/app/routes"
	"bytes"
)

type ChannelType struct {
	*revel.Controller
}

// 分发渠道列表
func (c *ChannelType) ListChannelType() revel.Result {

	// 分发渠道取得
	channels, err := getChannelTypeList()
	if err.IsError() {
		revel.WARN.Printf("get channelType return error: %v", err.ErrorMessage())
		return c.Render()
	}

	c.RenderArgs["channels"] = channels
	c.RenderArgs["channelCnt"] = len(channels)

	return c.Render()
}

// 单个分发渠道
func (c *ChannelType) GetSingleChannelType(id int64) revel.Result {

	channel, err := getChannelTypeById(id)
	if err.IsError() {
		revel.WARN.Printf("get single channelType return error: %v", err.ErrorMessage())
		return c.Render()
	}

	c.RenderArgs["channel"] = channel

	return c.Render()
}

// 新建分发渠道
func (c *ChannelType) NewChannelType() revel.Result {

	if c.Request.Method == "GET" {
		return c.Render()
	} else if c.Request.Method == "POST" {
		reqMsg := &ReqNewMemberCardChannel{
			ChannelName: c.Params.Get("channelName"),
			ChannelDesc: c.Params.Get("channelDesc"),
		}

		reqBytes, err := json.Marshal(reqMsg)
		if err != nil {

		}

		_, err = http.Post(app.NewChannelTypeUrl, "application/binary", bytes.NewReader(reqBytes))
		if err != nil {

		}

		return c.Redirect(routes.ChannelType.ListChannelType())
	}

	return c.Render()
}

// 更新分发渠道
func (c *ChannelType) UpdateChannelType() revel.Result {
	return c.Redirect(routes.ChannelType.ListChannelType())
}

// api 请求渠道 response
type RspMemberCardChannelList struct {
	Channels []*memberCard.MemberCardChannel `json:"channels"`
}

// api 请求单个渠道 request
type ReqSingleChannel struct {
	ChannelId int64 `json:"channelId"`
}

// api 请求单个渠道 response
type RspSingleChannel struct {
	ChannelId   int64  `json:"channelId"`
	ChannelName string `json:"channelName"`
	ChannelDesc string `json:"channelDesc"`
}

// api 新建分发渠道 request
type ReqNewMemberCardChannel struct {
	ChannelName string `json:"channelName"`
	ChannelDesc string `json:"channelDesc"`
}

// 分发渠道列表
func getChannelTypeList() ([]*memberCard.MemberCardChannel, errors.GlobalWaysError) {

	channels := make([]*memberCard.MemberCardChannel, 0)

	rsp, err := http.Get(app.ListChannelTypeUrl)
	if rsp.StatusCode != http.StatusOK || err != nil {
		return channels, errors.Newf(errors.CODE_HTTP_ERR_GET, errors.GlobalWaysErrors[errors.CODE_HTTP_ERR_GET], err)
	}

	rspBytes, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return channels, errors.Newf(errors.CODE_HTTP_READ_REQ_BODY, errors.GlobalWaysErrors[errors.CODE_HTTP_READ_REQ_BODY], err)
	}
	defer rsp.Body.Close()

	rspMsg := new(RspMemberCardChannelList)
	if err := json.Unmarshal(rspBytes, rspMsg); err != nil {
		return channels, errors.Newf(errors.CODE_JSON_ERR_UNMASHAL, errors.GlobalWaysErrors[errors.CODE_JSON_ERR_UNMASHAL], err)
	}
	channels = append(channels, rspMsg.Channels...)

	return channels, errors.ErrorOK()
}

// 单个分发渠道
func getChannelTypeById(id int64) (*memberCard.MemberCardChannel, errors.GlobalWaysError) {
	channel := new(memberCard.MemberCardChannel)

	rsp, err := http.Get(app.SingleChannelTypeUrl + strconv.Itoa(int(id)))
	if rsp.StatusCode != http.StatusOK || err != nil {
		return channel, errors.Newf(errors.CODE_HTTP_ERR_GET, errors.GlobalWaysErrors[errors.CODE_HTTP_ERR_GET], err)
	}

	rspBytes, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return channel, errors.Newf(errors.CODE_HTTP_READ_REQ_BODY, errors.GlobalWaysErrors[errors.CODE_HTTP_READ_REQ_BODY], err)
	}
	defer rsp.Body.Close()

	rspMsg := new(RspSingleChannel)
	if err := json.Unmarshal(rspBytes, rspMsg); err != nil {
		return channel, errors.Newf(errors.CODE_JSON_ERR_UNMASHAL, errors.GlobalWaysErrors[errors.CODE_JSON_ERR_UNMASHAL], err)
	}
	channel.Id = rspMsg.ChannelId
	channel.ChannelName = rspMsg.ChannelName
	channel.ChannelDesc = rspMsg.ChannelDesc

	return channel, errors.ErrorOK()
}

