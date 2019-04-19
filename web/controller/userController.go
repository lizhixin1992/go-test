package controller

import (
	"github.com/kataras/iris"
	"github.com/lizhixin1992/test/services"
)

type UserController struct {
	Ctx     iris.Context
	Service services.UserService
}
