package api

import (
	"github.com/kataras/iris/v12/context"
)

type Response struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func successData(context *context.Context, Data interface{}) {
	context.JSON(Response{
		Code: 200,
		Msg:  "请求成功",
		Data: Data,
	})
}

func successMsg(context *context.Context, Msg string) {
	context.JSON(Response{
		Code: 200,
		Msg:  Msg,
		Data: nil,
	})
}

func successDataMsg(context *context.Context, Data interface{}, Msg string) {
	context.JSON(Response{
		Code: 200,
		Msg:  Msg,
		Data: nil,
	})
}

func successNil(context *context.Context) {
	context.JSON(Response{
		Code: 200,
		Msg:  "请求成功",
		Data: nil,
	})
}

func errorCodeMsg(context *context.Context, Code int64, Msg string) {
	context.JSON(Response{
		Code: Code,
		Msg:  Msg,
		Data: nil,
	})
}

func errorMsg(context *context.Context, Msg string) {
	context.JSON(Response{
		Code: 400,
		Msg:  Msg,
		Data: nil,
	})
}

func errorNil(context *context.Context) {
	context.JSON(Response{
		Code: 400,
		Msg:  "请求失败",
		Data: nil,
	})
}
