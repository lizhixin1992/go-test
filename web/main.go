package main

import (
	"github.com/lizhixin1992/test/bootstrap"
	"github.com/lizhixin1992/test/web/routes"
)

func newApp() (b *bootstrap.Bootstrapper) {
	app := bootstrap.New("test-go", "lizhixin")
	app.Bootstrap()
	app.Configure(routes.Configure)
	return app
}

func main() {
	app := newApp()
	app.Listen(":8080")
}
