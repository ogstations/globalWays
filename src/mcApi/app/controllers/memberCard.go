// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package controllers

import (
	"encoding/json"
	"github.com/revel/revel"
	"io/ioutil"
	"mcApi/app/models"
	"memberCard"
	"net/http"
	"regexp"
	"strconv"
	"utils/errors"
	"utils/page"
	"bytes"
	"fmt"
	"time"
)

type MemberCard struct {
	*revel.Controller
}

const (
	DigitalRegexp string = "^[0-9]*$"
)

// 批量生成会员卡
func (c *MemberCard) GenMemberCards() revel.Result {

	// unmarshal request body
	reqBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		revel.WARN.Printf("read http request body return error: %v", err)

		// fill response
		rspErr := &models.RspError{
			ErrorCode: errors.CODE_INVALID_PARAMS,
			Message:   errors.GlobalWaysErrors[errors.CODE_INVALID_PARAMS],
		}
		c.Response.Status = http.StatusBadRequest

		return c.RenderJson(rspErr)
	}
	defer c.Request.Body.Close()

	reqMsg := new(models.ReqNewMemberCard)
	if err := json.Unmarshal(reqBytes, reqMsg); err != nil {
		revel.WARN.Printf("unmarshal http request body return error: %v", err)

		// fill response
		rspErr := &models.RspError{
			ErrorCode: errors.CODE_INVALID_PARAMS,
			Message:   errors.GlobalWaysErrors[errors.CODE_INVALID_PARAMS],
		}
		c.Response.Status = http.StatusBadRequest

		return c.RenderJson(rspErr)
	}

	rspMsg := new(models.RspNewMemberCard)
	// error valid
	fieldErrors := make([]models.FieldError, 0)
	if reqMsg.CardCount < 0 {
		fieldErrors = append(fieldErrors, models.FieldError{Field: "CardCount", Message: "CardCount should not be empty"})
	}
	if len(fieldErrors) != 0 {
		rspMsg.FieldErrors = fieldErrors

		revel.WARN.Printf("request message has error.")
		return c.RenderJson(rspMsg)
	}

	// make new memberCards
	var gErr errors.GlobalWaysError
	rspMsg.CardNumbers, gErr = memberCard.GenMemberCards(reqMsg.MajorIndustryIdentifier, reqMsg.CompanyIdentifier,
		reqMsg.DomainIdentifier, reqMsg.CardCount, reqMsg.ChannelType, reqMsg.ChannelId, models.WriterEngine)
	if gErr.IsError() {
		revel.WARN.Printf("make new memberCards return error: %v", gErr.ErrorMessage())

		rspErr := &models.RspError{
			ErrorCode: gErr.GetCode(),
			Message:   gErr.GetMessage(),
		}
		c.Response.Status = http.StatusInternalServerError

		return c.RenderJson(rspErr)
	}

	return c.RenderJson(rspMsg)
}

// 会员卡列表
func (c *MemberCard) MemberCardList() revel.Result {

	//参数验证
	pageNum, pageSize := 1, 10

	pageStr := c.Params.Get("page")
	c.Validation.Match(pageStr, regexp.MustCompile(DigitalRegexp))
	if len(pageStr) != 0 && !c.Validation.HasErrors() {
		pageNum, _ = strconv.Atoi(pageStr)
	}

	pageSizeStr := c.Params.Get("size")
	c.Validation.Match(pageSizeStr, regexp.MustCompile(DigitalRegexp))
	if len(pageSizeStr) != 0 && !c.Validation.HasErrors() {
		pageSize, _ = strconv.Atoi(pageSizeStr)
	}

	//获取会员卡列表
	pager := &page.Page{
		Perpage:      int64(pageSize),
		Current_page: int64(pageNum),
	}

	cards, gErr := memberCard.FindMemberCard(pager, models.ReaderEngine)
	if gErr.IsError() {
		revel.WARN.Printf("make new memberCards return error: %v", gErr.ErrorMessage())

		rspErr := &models.RspError{
			ErrorCode: gErr.GetCode(),
			Message:   gErr.GetMessage(),
		}
		c.Response.Status = http.StatusInternalServerError

		return c.RenderJson(rspErr)
	}

	rspMsg := new(models.RspMemberCardList)
	rspMsg.Cards = cards
	rspMsg.Count = 100

	return c.RenderJson(rspMsg)
}

// 会员卡二维码二进制流
func (c *MemberCard) GenQrCode(id int64) revel.Result {

	card, gErr := memberCard.GetMemberCardById(id, models.ReaderEngine)
	if gErr.IsError() {
		revel.WARN.Printf("get memberCard by id return error: %v", gErr.ErrorMessage())

		rspErr := &models.RspError{
			ErrorCode: gErr.GetCode(),
			Message:   gErr.GetMessage(),
		}
		c.Response.Status = http.StatusInternalServerError

		return c.RenderJson(rspErr)
	}

	return c.RenderBinary(bytes.NewReader(card.GenQrStream()), fmt.Sprintf("%s.png", card.String()), revel.Inline, time.Now())
}
