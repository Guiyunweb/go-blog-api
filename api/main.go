package api

import (
	"blog-api/serializer"
)

func ErrorResponse(err error) serializer.Response {
	return serializer.Response{
		Success: false,
		Message: "服务器繁忙",
		Data:    err,
	}
}

func SuccessResponse() serializer.Response {
	return serializer.Response{
		Success: true,
		Message: "操作成功",
		Data:    nil,
	}
}
