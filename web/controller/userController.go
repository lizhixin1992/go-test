package controller

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/lizhixin1992/test/commons"
	"github.com/lizhixin1992/test/models"
	"github.com/lizhixin1992/test/services"
	"log"
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

//localhost:8080/save
func (c *UserController) PostSave() mvc.Result {
	info := models.User{}
	err := c.Ctx.ReadJSON(&info)
	if err != nil {
		log.Fatal(err)
		return setResponseFail()
	} else {
		date := int(commons.GetNowUnix())
		info.CreateTime = date
		info.UpdateTime = date
		c.Service.Save(&info)
		return setResponseSuccess()
	}
}

func setResponseSuccess() mvc.Response {
	return mvc.Response{
		Object: iris.Map{
			"errorCode":    "0",
			"errorMessage": "success",
		},
	}
}

func setResponseFail() mvc.Response {
	return mvc.Response{
		Object: iris.Map{
			"errorCode":    "-1",
			"errorMessage": "fail",
		},
	}
}

func (c *UserController) BeforeActivation(b mvc.BeforeActivation) {
	fmt.Println("******************* before *********************")
}

func (c *UserController) AfterActivation(b mvc.AfterActivation) {
	fmt.Println("******************* after *********************")
}
