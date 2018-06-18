package main

import (
	"fmt"
	"net/http"

	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	//adicionando middleware
	irisMiddleware := iris.FromStd(middleLog)
	app.Use(irisMiddleware)

	app.RegisterView(iris.HTML("./views", ".html"))

	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("message", "Golang")
		ctx.View("hello.html")
	})

	app.Get("/log", func(ctx iris.Context) {
		ctx.ViewData("message", "Log!")
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

func middleLog(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	app := iris.New()
	if r.URL.Path == "/log" && r.Method == "GET" {
		//w.Write([]byte("Log Ok. "))
		app.Logger().Info("Log Ok. ")
		next(w, r)
		return
	}
	//w.WriteHeader(iris.StatusBadRequest)
	//w.Write([]byte("Bad Request"))
	fmt.Sprint("fmt ok")
	app.Logger().Info("Not Log endpoint! ")
	next(w, r)
	return
}
