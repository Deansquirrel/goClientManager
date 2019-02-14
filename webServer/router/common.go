package router

import (
	"encoding/json"
	"fmt"
)

const (
	FormatErr = "{errcode:-1,errmsg:\"%s\"}"
	FormatMsg = "{errcode:0,errmsg:\"%s\"}"
)

func GetMsgReturn(msg string) string {
	rd := responseDao{
		ErrCode: 0,
		ErrMsg:  msg,
	}
	rb, err := json.Marshal(rd)
	if err != nil {
		return fmt.Sprintf(FormatMsg, msg)
	}
	return string(rb)
}

func GetErrReturn(msg string) string {
	rd := responseDao{
		ErrCode: -1,
		ErrMsg:  msg,
	}
	rb, err := json.Marshal(rd)
	if err != nil {
		return fmt.Sprintf(FormatErr, msg)
	}
	return string(rb)
}

type responseDao struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}
