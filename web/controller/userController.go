package controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/lizhixin1992/test/services"
)

type UserController struct {
	Ctx     iris.Context
	Service services.UserService
}

func (c *UserController) Get() mvc.Result  {
	dataList := c.Service.GetAll()
	return mvc.Response{
		Object:iris.Map{
			"errorCode":"0",
			"errorMessage":"success",
			"data":dataList,
		},
	}
}