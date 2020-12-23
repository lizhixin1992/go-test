package routes

import (
	"github.com/kataras/iris/v12/mvc"
	"github.com/lizhixin1992/go-test/bootstrap"
	"github.com/lizhixin1992/go-test/services"
	"github.com/lizhixin1992/go-test/web/controller"
)

func Configure(b *bootstrap.Bootstrapper) {
	userService := services.NewUserservice()

	userRoute := mvc.New(b.Party("/"))
	userRoute.Register(userService)
	userRoute.Handle(new(controller.UserController))
}
