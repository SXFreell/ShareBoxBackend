package main

import (
	_ "sharebox/dao"
	"sharebox/service"

	"github.com/kataras/iris/v12"
	"github.com/robfig/cron/v3"
)

func main() {
	// Web server
	// webServer := iris.New()
	// webServer.Use(iris.Compression)
	// webServer.Use(corsMiddleware())
	// webServer.HandleDir("/", "./static/dist")
	// webServer.OnAnyErrorCode(func(ctx iris.Context) {
	// 	if ctx.GetStatusCode() == iris.StatusNotFound {
	// 		ctx.ServeFile("./static/dist/index.html")
	// 	} else {
	// 		ctx.WriteString(iris.StatusText(ctx.GetStatusCode()))
	// 	}
	// })

	// go func() {
	// 	webServer.Listen(":41251")
	// }()

	// API
	app := iris.New()
	app.Use(iris.Compression)
	app.Use(corsMiddleware())

	api := app.Party("/api")
	{
		authAPI := api.Party("/auth")
		{
			authAPI.Get("/", returnIndex)
			authAPI.Post("/login", service.Login)
			// authAPI.Post("/register", register)
			authAPI.Post("/logout", service.Logout)
		}

		getAPI := api.Party("/get")
		{
			getAPI.Get("/", returnIndex)
			getAPI.Post("/", service.GetSomething)
		}

		setAPI := api.Party("/set")
		{
			setAPI.Get("/", returnIndex)
			setAPI.Post("/", service.SetSomething)
		}
	}

	app.HandleDir("/", "./static/dist")
	app.OnAnyErrorCode(func(ctx iris.Context) {
		if ctx.GetStatusCode() == iris.StatusNotFound {
			ctx.ServeFile("./static/dist/index.html")
		} else {
			ctx.WriteString(iris.StatusText(ctx.GetStatusCode()))
		}
	})

	c := cron.New(cron.WithSeconds())
	c.AddFunc("45 23 1 * * *", service.ClearExpiredFile)
	go c.Start()

	app.Listen(":41250")
}

// CORS Middleware
func corsMiddleware() iris.Handler {
	return func(ctx iris.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		// ctx.Header("Access-Control-Allow-Credentials", "true")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin")
		ctx.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		if ctx.Method() == iris.MethodOptions {
			ctx.StatusCode(iris.StatusNoContent)
			return
		}
		ctx.Next()
	}
}

// 返回index.html
func returnIndex(ctx iris.Context) {
	ctx.ServeFile("./static/dist/index.html")
}
