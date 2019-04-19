package routes

import (
	"github.com/kataras/iris/mvc"
	"github.com/lizhixin1992/test/bootstrap"
	"github.com/lizhixin1992/test/services"
	"github.com/lizhixin1992/test/web/controller"
)

func Configure(b *bootstrap.Bootstrapper)  {
	userService := services.NewUserservice()

	userRoute := mvc.New(b.Party("/"))
	userRoute.Register(userService)
	userRoute.Handle(new(controller.UserController))
}
