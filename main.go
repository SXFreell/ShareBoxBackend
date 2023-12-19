package main

import (
	_ "sharebox/dao"
	"sharebox/service"
	"sharebox/utils"

	"github.com/kataras/iris/v12"
)

func main() {
	// Web server
	webServer := iris.New()
	webServer.Use(iris.Compression)
	webServer.Use(corsMiddleware())
	webServer.HandleDir("/", "./static/dist")
	webServer.OnAnyErrorCode(func(ctx iris.Context) {
		if ctx.GetStatusCode() == iris.StatusNotFound {
			ctx.ServeFile("./static/dist/index.html")
		} else {
			ctx.WriteString(iris.StatusText(ctx.GetStatusCode()))
		}
	})

	go func() {
		webServer.Listen(":41251")
	}()

	// booksAPI := app.Party("/books")
	// {
	// 	booksAPI.Use(iris.Compression)

	// 	// GET: http://localhost:8080/books
	// 	booksAPI.Get("/", list)
	// 	// POST: http://localhost:8080/books
	// 	booksAPI.Post("/", create)
	// }

	// API
	app := iris.New()

	authAPI := app.Party("/auth")
	{
		authAPI.Post("/login", service.Login)
		// authAPI.Post("/register", register)
		authAPI.Post("/logout", service.Logout)
	}

	getAPI := app.Party("/get")
	{
		getAPI.Get("/", service.GetSomething)
	}

	setAPI := app.Party("/set")
	{
		setAPI.Post("/", service.SetSomething)
	}

	app.Get("/api", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello Iris!"})
		utils.Log.Info("Hello Iris!")
	})

	app.Listen(":41250")
}

// CORS Middleware
func corsMiddleware() iris.Handler {
	return func(ctx iris.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin")
		ctx.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		if ctx.Method() == iris.MethodOptions {
			ctx.StatusCode(iris.StatusNoContent)
			return
		}
		ctx.Next()
	}
}

// // Book example.
// type Book struct {
// 	Title string `json:"title"`
// }

// func list(ctx iris.Context) {
// 	books := []Book{
// 		{"Mastering Concurrency in Go"},
// 		{"Go Design Patterns"},
// 		{"Black Hat Go"},
// 	}

// 	ctx.JSON(books)
// 	// TIP: negotiate the response between server's prioritizes
// 	// and client's requirements, instead of ctx.JSON:
// 	// ctx.Negotiation().JSON().MsgPack().Protobuf()
// 	// ctx.Negotiate(books)
// }

// func create(ctx iris.Context) {
// 	var b Book
// 	err := ctx.ReadJSON(&b)
// 	// TIP: use ctx.ReadBody(&b) to bind
// 	// any type of incoming data instead.
// 	if err != nil {
// 		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
// 			Title("Book creation failure").DetailErr(err))
// 		// TIP: use ctx.StopWithError(code, err) when only
// 		// plain text responses are expected on errors.
// 		return
// 	}

// 	println("Received Book: " + b.Title)

// 	ctx.StatusCode(iris.StatusCreated)
// }
