// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package models

// 请求新建会员卡
type ReqNewMemberCard struct {
	MajorIndustryIdentifier byte   `json:"majorIndustryIdentifier"`
	CompanyIdentifier       byte   `json:"companyIdentifier"`
	DomainIdentifier        uint16 `json:"domainIdentifier"`
	CardCount               int64  `json:"cardCount"`
	ChannelType             int64  `json:"channelType"`
	ChannelId               int64  `json:"channelId"`
}

type ReqNewMemberCardChannel struct {
	ChannelName string `json:"channelName"`
	ChannelDesc string `json:"channelDesc"`
}

type ReqUpdChannel struct {
	ChannelId   int64  `json:"channelId"`
	ChannelName string `json:"channelName"`
	ChannelDesc string `json:"channelDesc"`
}
