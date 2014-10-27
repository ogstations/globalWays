// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/revel/revel"
	"io/ioutil"
	"mcApi/app/models"
	"memberCard"
	"net/http"
	"utils/errors"
)

type ChannelType struct {
	*revel.Controller
}

// 会员卡分发渠道列表
func (c *ChannelType) ChannelTypeList() revel.Result {

	rspMsg := new(models.RspMemberCardChannelList)

	var gErr errors.GlobalWaysError
	rspMsg.Channels, gErr = memberCard.FindMemberCardChannel(models.ReaderEngine)
	if gErr.IsError() {
		rspErr := &models.RspError{
			ErrorCode: gErr.GetCode(),
			Message:   gErr.GetMessage(),
		}
		c.Response.Status = http.StatusInternalServerError

		return c.RenderJson(rspErr)
	}

	return c.RenderJson(rspMsg)
}

// 新建会员卡分发渠道
func (c *ChannelType) CreateChannelType() revel.Result {

	// unmarshal request body
	reqBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		revel.WARN.Printf("read http request body return error: %v", err)

		// fill response
		rspErr := &models.RspError{
			ErrorCode: errors.CODE_HTTP_READ_REQ_BODY,
			Message:   fmt.Sprintf(errors.GlobalWaysErrors[errors.CODE_HTTP_READ_REQ_BODY], err),
		}
		c.Response.Status = http.StatusBadRequest

		return c.RenderJson(rspErr)
	}
	defer c.Request.Body.Close()

	reqMsg := new(models.ReqNewMemberCardChannel)
	if err := json.Unmarshal(reqBytes, reqMsg); err != nil {
		revel.WARN.Printf("unmarshal http request body return error: %v", err)

		// fill response
		rspErr := &models.RspError{
			ErrorCode: errors.CODE_JSON_ERR_UNMASHAL,
			Message:   fmt.Sprintf(errors.GlobalWaysErrors[errors.CODE_JSON_ERR_UNMASHAL], err),
		}
		c.Response.Status = http.StatusBadRequest

		return c.RenderJson(rspErr)
	}

	rspMsg := new(models.RspNewMemberCardChannel)
	// error valid
	fieldErrors := make([]models.FieldError, 0)
	if len(reqMsg.ChannelName) < 0 {
		fieldErrors = append(fieldErrors, models.FieldError{Field: "ChannelName", Message: "ChannelName should not be empty"})
	}
	if len(reqMsg.ChannelDesc) < 0 {
		fieldErrors = append(fieldErrors, models.FieldError{Field: "ChannelDesc", Message: "ChannelDesc should not be empty"})
	}
	if len(fieldErrors) != 0 {
		rspMsg.FieldErrors = fieldErrors

		return c.RenderJson(rspMsg)
	}

	// make new memberCard channel
	var gErr errors.GlobalWaysError
	rspMsg.Channel, _, gErr = memberCard.InsertMemberCardChannel(reqMsg.ChannelName, reqMsg.ChannelDesc, models.WriterEngine)
	if gErr.IsError() {
		rspErr := &models.RspError{
			ErrorCode: gErr.GetCode(),
			Message:   gErr.GetMessage(),
		}
		c.Response.Status = http.StatusInternalServerError

		return c.RenderJson(rspErr)
	}

	return c.RenderJson(rspMsg)
}

// 请求特定分发渠道信息
func (c *ChannelType) GetChannelById(id int64) revel.Result {

	rspMsg := new(models.RspSingleChannel)
	channel, gErr := memberCard.GetMemberCardChannel(id, models.ReaderEngine)
	if gErr.IsError() {
		rspErr := &models.RspError{
			ErrorCode: gErr.GetCode(),
			Message:   gErr.GetMessage(),
		}
		c.Response.Status = http.StatusInternalServerError

		return c.RenderJson(rspErr)
	}

	rspMsg.ChannelId = channel.Id
	rspMsg.ChannelName = channel.ChannelName
	rspMsg.ChannelDesc = channel.ChannelDesc

	return c.RenderJson(rspMsg)
}

func (c *ChannelType) UpdateChannel() revel.Result {

	// unmarshal request body
	reqBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		revel.WARN.Printf("read http request body return error: %v", err)

		// fill response
		rspErr := &models.RspError{
			ErrorCode: errors.CODE_HTTP_READ_REQ_BODY,
			Message:   fmt.Sprintf(errors.GlobalWaysErrors[errors.CODE_HTTP_READ_REQ_BODY], err),
		}
		c.Response.Status = http.StatusBadRequest

		return c.RenderJson(rspErr)
	}
	defer c.Request.Body.Close()

	reqMsg := new(models.ReqUpdChannel)
	if err := json.Unmarshal(reqBytes, reqMsg); err != nil {
		revel.WARN.Printf("unmarshal http request body return error: %v", err)

		// fill response
		rspErr := &models.RspError{
			ErrorCode: errors.CODE_JSON_ERR_UNMASHAL,
			Message:   fmt.Sprintf(errors.GlobalWaysErrors[errors.CODE_JSON_ERR_UNMASHAL], err),
		}
		c.Response.Status = http.StatusBadRequest

		return c.RenderJson(rspErr)
	}

	// update channel
	channel := &memberCard.MemberCardChannel{
		Id:          reqMsg.ChannelId,
		ChannelName: reqMsg.ChannelName,
		ChannelDesc: reqMsg.ChannelDesc,
	}
	_, gErr := memberCard.UpdMemberCardChannel(channel, models.WriterEngine)
	if gErr.IsError() {
		rspErr := &models.RspError{
			ErrorCode: gErr.GetCode(),
			Message:   gErr.GetMessage(),
		}
		c.Response.Status = http.StatusInternalServerError

		return c.RenderJson(rspErr)
	}

	rspMsg := &models.RspUpdChannel{
		ChannelId:   reqMsg.ChannelId,
		ChannelName: reqMsg.ChannelName,
		ChannelDesc: reqMsg.ChannelDesc,
	}

	return c.RenderJson(rspMsg)
}
