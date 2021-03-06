// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package models

import "memberCard"

type RspNewMemberCard struct {
	CardNumbers []memberCard.CardNumber `json:"cardNumbers"`
	FieldErrors []FieldError            `json:"fieldErrors"`
}

type RspMemberCardChannelList struct {
	Channels []*memberCard.MemberCardChannel `json:"channels"`
}

type RspMemberCardList struct {
	Cards []*memberCard.MemberCard `json:"cards"`
	Count int64                    `json:"count`
}

type RspNewMemberCardChannel struct {
	Channel     *memberCard.MemberCardChannel `json:"channel"`
	FieldErrors []FieldError                  `json:"fieldErrors"`
}

type RspSingleChannel struct {
	ChannelId   int64  `json:"channelId"`
	ChannelName string `json:"channelName"`
	ChannelDesc string `json:"channelDesc"`
}

type RspUpdChannel struct {
	ChannelId   int64  `json:"channelId"`
	ChannelName string `json:"channelName"`
	ChannelDesc string `json:"channelDesc"`
}

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type RspError struct {
	ErrorCode int    `json:"errorCode"`
	Message   string `json:"message"`
}
