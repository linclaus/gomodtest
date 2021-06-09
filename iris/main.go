package main

import (
	"github.com/kataras/iris/v12"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
	app.Handle("GET", "/ping", func(ctx iris.Context) {
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
