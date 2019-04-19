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

//localhost:8080/
func (c *UserController) Get() mvc.Result {
	dataList := c.Service.GetAll()
	return mvc.Response{
		Object: iris.Map{
			"errorCode":    "0",
			"errorMessage": "success",
			"data":         dataList,
		},
	}

	//return mvc.Response{
	//	Text:"11111111",
	//}
}

//localhost:8080/{id}
func (c *UserController) GetBy(id int) mvc.Result {
	data := c.Service.GetById(id)
	return mvc.Response{
		Object: iris.Map{
			"errorCode":    "0",
			"errorMessage": "success",
			"data":         data,
		},
	}
}
