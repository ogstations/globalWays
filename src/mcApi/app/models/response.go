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

type RspNewMemberCardChannel struct {
	Channel     *memberCard.MemberCardChannel `json:"channel"`
	FieldErrors []FieldError                  `json:"fieldErrors"`
}

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type RspError struct {
	ErrorCode int    `json:"errorCode"`
	Message   string `json:"message"`
}
