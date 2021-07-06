package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

var (
	cookieNameForSessionID = "mycookiesessionnameid"
	sess                   = sessions.New(sessions.Config{Cookie: cookieNameForSessionID})
)

func main() {
	app := iris.Default()
	app.Use(myMiddleware)
	app.RegisterView(iris.HTML("./iris/views", ".html"))

	app.Get("/", func(ctx iris.Context) {
		// Bind: {{.message}} with "Hello world!"
		ctx.ViewData("message", "Hello world!")
		// Render template file: ./views/hello.html
		ctx.View("hello.html")
	})
	app.Get("/metrics", iris.FromStd(promhttp.Handler()))
	app.Get("/user/{id:uint64}", func(ctx iris.Context) {
		userID, _ := ctx.Params().GetUint64("id")
		ctx.Writef("User ID: %d", userID)
	})
	app.Handle("GET", "/ping/{id}", func(ctx iris.Context) {
		session := sess.Start(ctx)
		smid := session.GetString("serverMessageId")
		logrus.Infof("serverMessageId: %s", smid)
		session.Set("serverMessageId", "123456")
		cId := ctx.GetCookie("cookieId")
		logrus.Infof("cookieId: %s", cId)
		sId := ctx.GetCookie(cookieNameForSessionID)
		logrus.Infof("sessionId: %s", sId)
		ctx.SetCookieKV("cookieId", "123")

		p := ctx.Params().Get("id")
		logrus.Infof("params Id: %s", p)
		params := ctx.URLParams()
		logrus.Info(params)

		ctx.JSON(iris.Map{"message": "pong"})
	})

	// Listens and serves incoming http requests
	// on http://localhost:8080.
	app.Run(iris.Addr(":8080"))
}

func myMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
	ctx.Next()
}
