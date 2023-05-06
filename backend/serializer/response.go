package serializer

import (
	error "GoAsk/utils/error"
)

// Response 基础序列化器
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
}

func BuildResponse(status int, data interface{}) Response {
	return Response{
		Status: status,
		Data:   data,
		Msg:    error.GetMsg(status),
	}
}
