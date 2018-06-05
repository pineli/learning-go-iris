package main

import "github.com/kataras/iris"

func main() {
	app := iris.New()

	app.RegisterView(iris.HTML("./views", ".html"))

	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("message", "Golang")
		ctx.View("hello.html")
	})

	app.Get("/user/{id:long}", func(ctx iris.Context) {
		userID, _ := ctx.Params().GetInt64("id")
		ctx.Writef("User ID: %d", userID)
	})

	app.Get("/json", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "JSON woth Iris web framework. "})
	})
	app.Run(iris.Addr(":8080"))

}
