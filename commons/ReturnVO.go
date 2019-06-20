package commons

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type returnData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SetResponse(code int, msg string, data interface{}) mvc.Response {
	ret := returnData{Code: 0, Message: "success", Data: data}
	return mvc.Response{
		Object: ret,
	}
}

func SetResponseSuccessData(data interface{}) mvc.Response {
	ret := returnData{Code: 0, Message: "success", Data: data}
	return mvc.Response{
		Object: ret,
	}
}

func SetResponseSuccess() mvc.Response {
	return mvc.Response{
		Object: iris.Map{
			"errorCode":    0,
			"errorMessage": "success",
		},
	}
}

func SetResponseFail() mvc.Response {
	return mvc.Response{
		Object: iris.Map{
			"errorCode":    1,
			"errorMessage": "fail",
		},
	}
}
