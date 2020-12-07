package api

import "github.com/kataras/iris/v12/context"

func HelloWorld(context *context.Context) {
	successData(context, "hello world")
}
