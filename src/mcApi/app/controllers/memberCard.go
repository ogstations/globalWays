// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package controllers

import (
	"github.com/revel/revel"
	"io/ioutil"
	"mcApi/app/models"
	"utils/errors"
	"encoding/json"
	"memberCard"
)

type MemberCard struct {
	*revel.Controller
}

func (c *MemberCard) GenMemberCards() revel.Result {

	// unmarshal request body
	reqBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		revel.WARN.Printf("read http request body return error: %v", err)

		// fill response
		rspErr := &models.RspError{
			ErrorCode: errors.CODE_INVALID_PARAMS,
			Message: errors.GlobalWaysErrors[errors.CODE_INVALID_PARAMS],
		}

		return c.RenderJson(rspErr)
	}
	defer c.Request.Body.Close()

	reqMsg := new(models.ReqNewMemberCard)
	if err := json.Unmarshal(reqBytes, reqMsg); err != nil {
		revel.WARN.Printf("unmarshal http request body return error: %v", err)

		// fill response
		rspErr := &models.RspError{
			ErrorCode: errors.CODE_INVALID_PARAMS,
			Message: errors.GlobalWaysErrors[errors.CODE_INVALID_PARAMS],
		}

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

		return c.RenderJson(rspMsg)
	}

	// make new memberCards
	var gErr errors.GlobalWaysError
	rspMsg.CardNumbers, gErr = memberCard.GenMemberCards(reqMsg.MajorIndustryIdentifier, reqMsg.CompanyIdentifier,
		reqMsg.DomainIdentifier,reqMsg.CardCount,reqMsg.ChannelType, reqMsg.ChannelId, models.WriterEngine)
	if gErr.IsError() {
		rspErr := &models.RspError{
			ErrorCode: gErr.GetCode(),
			Message: gErr.GetMessage(),
		}

		return c.RenderJson(rspErr)
	}

	return c.RenderJson(rspMsg)
}
