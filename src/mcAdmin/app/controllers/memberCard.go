// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package controllers

import (
	"encoding/json"
	"github.com/revel/revel"
	"mcAdmin/app"
	"net/http"
	"strconv"
	"bytes"
)

type MemberCard struct {
	*revel.Controller
}

// api 请求新建会员卡 request
type ReqNewMemberCard struct {
	MajorIndustryIdentifier byte   `json:"majorIndustryIdentifier"`
	CompanyIdentifier       byte   `json:"companyIdentifier"`
	DomainIdentifier        uint16 `json:"domainIdentifier"`
	CardCount               int64  `json:"cardCount"`
	ChannelType             int64  `json:"channelType"`
	ChannelId               int64  `json:"channelId"`
}

// fill form
func (c *MemberCard) DisplayForm() revel.Result {

	// 分发渠道取得
	channels, err := getChannelTypeList()
	if err.IsError() {
		revel.WARN.Printf("get channelType return error: %v", err.ErrorMessage())
		return c.Render()
	}

	c.RenderArgs["channels"] = channels
	return c.RenderTemplate("MemberCard/genMemberCard.html")
}

// gen member card
func (c *MemberCard) GenMemberCard() revel.Result {

	//init
	mii := byte(6)
	cpi := byte(32)
	cdi := uint16(86)
	cardCount := int64(100)
	channelType := int64(1)
	channelId := int64(1)

	// 产业标识符
	if miiStr := c.Params.Get("mii"); len(miiStr) != 0 {
		tVal, _ := strconv.ParseInt(miiStr, 10, 8)
		mii = byte(tVal)
	}
	// 公司标识符
	if cpiStr := c.Params.Get("cpi"); len(cpiStr) != 0 {
		tVal, _ := strconv.ParseInt(cpiStr, 10, 8)
		cpi = byte(tVal)
	}
	// 地域标识符
	if cdiStr := c.Params.Get("cdi"); len(cdiStr) != 0 {
		tVal, _ := strconv.ParseUint(cdiStr, 10, 16)
		cdi = uint16(tVal)
	}
	// 会员卡数量
	if cardCntStr := c.Params.Get("cardCount"); len(cardCntStr) != 0 {
		cardCount, _ = strconv.ParseInt(cardCntStr, 10, 64)
	}
	// 分发渠道
	if channelTypeStr := c.Params.Get("channelType"); len(channelTypeStr) != 0 {
		channelType, _ = strconv.ParseInt(channelTypeStr, 10, 64)
	} else {

	}
	// 渠道ID
	if channelIDStr := c.Params.Get("channelID"); len(channelIDStr) != 0 {
		channelId, _ = strconv.ParseInt(channelIDStr, 10, 64)
	} else {

	}

	reqMsg := &ReqNewMemberCard{
		MajorIndustryIdentifier: mii,
		CompanyIdentifier:       cpi,
		DomainIdentifier:        cdi,
		CardCount:               cardCount,
		ChannelType:             channelType,
		ChannelId:               channelId,
	}

	reqBytes, err := json.Marshal(reqMsg)
	if err != nil {
		revel.WARN.Printf("Request marshal return error: %v", err)

		return c.Render()
	}

	rsp, err := http.Post(app.NewMemberCardUrl, "application/binary", bytes.NewReader(reqBytes))
	if rsp.StatusCode != http.StatusOK || err != nil {
		revel.WARN.Printf("http post request return error: %v, app.NewMemberCardUrl: %v", err, app.NewMemberCardUrl)
		return c.Render()
	}

	return c.Render()
}
