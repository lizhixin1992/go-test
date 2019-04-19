package bootstrap

import (
	"github.com/gorilla/securecookie"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/sessions"
	"github.com/lizhixin1992/test/conf"
	"time"
)

type Bootstrapper struct {
	*iris.Application
	Appname      string
	AppOwner     string
	AppSpawnDate time.Time
	Sessions     *sessions.Sessions
}

type Configurator func(*Bootstrapper)

func New(appName, appOwner string, cfgs ...Configurator) *Bootstrapper {
	b := &Bootstrapper{
		Appname:      appName,
		AppOwner:     appOwner,
		AppSpawnDate: time.Now(),
		Application:  iris.New(),
	}

	for _, cfg := range cfgs {
		cfg(b)
	}
	return b
}

//设置视图层相关配置
func (b *Bootstrapper) SetUpViews() {
	//设置html的文件目录地址，后缀名，模版
	htmlEngine := iris.HTML(conf.ViewDir, conf.ViewExtension).Layout("shared/error.html")
	//设置是否每次都从新加载模版（线上环境关掉）
	htmlEngine.Reload(conf.HtmlReload)
	// 给模版内置各种定制的方法
	htmlEngine.AddFunc("FromUnixtimeShort", func(t int) string {
		dt := time.Unix(int64(t), int64(0))
		return dt.Format(conf.SysTimeFormShort)
	})
	htmlEngine.AddFunc("FromUnixtime", func(t int) string {
		dt := time.Unix(int64(t), int64(0))
		return dt.Format(conf.SysTimeForm)
	})
	b.RegisterView(htmlEngine)
}

//设置session
func (b *Bootstrapper) SetUpSession(expires time.Duration, cookieHashKey, cookieBlockKey []byte) {
	sessions.New(sessions.Config{
		Cookie:   "SECRET_SESS_COOKIE_" + b.Appname,
		Expires:  expires,
		Encoding: securecookie.New(cookieHashKey, cookieBlockKey),
	})
}

//设置当错误时的处理（< 200 || >= 400）
func (b *Bootstrapper) SetUpErrorHandlers() {
	b.OnAnyErrorCode(func(ctx iris.Context) {
		err := iris.Map{
			"app":     b.Appname,
			"status":  ctx.GetStatusCode(),
			"message": ctx.Values().Get("message"),
		}
		if jsonOutput := ctx.URLParamExists("json"); jsonOutput {
			ctx.JSON(err)
		}

		ctx.ViewData("Err", err)
		ctx.ViewData("Title", "Error")
		ctx.View("shared/error.html")
	})
}

func (b *Bootstrapper) Configure(cs ...Configurator) {
	for _, c := range cs {
		c(b)
	}
}

//初始化相关
func (b *Bootstrapper) Bootstrap() *Bootstrapper {
	b.SetUpViews()
	b.SetUpSession(conf.SessionExpires,
		[]byte("the-big-and-secret-fash-key-here"),
		[]byte("lot-secret-of-characters-big-too"),
	)
	b.SetUpErrorHandlers()

	// static files
	b.Favicon(conf.StaticAssets + conf.Favicon)
	b.StaticWeb(conf.StaticAssets[1:len(conf.StaticAssets)-1], conf.StaticAssets)

	// middleware, after static files
	b.Use(recover.New())
	b.Use(logger.New())

	return b
}

//启动iris
func (b *Bootstrapper) Listen(addr string, cfgs ...iris.Configurator) {
	b.Run(iris.Addr(addr), cfgs...)
}