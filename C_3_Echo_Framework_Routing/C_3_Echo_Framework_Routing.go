package main

import (
    "fmt"
    "github.com/labstack/echo"
    "net/http"
    "strings"
)

type M map[string]interface{}

func main() {
    r := echo.New()

    r.GET("/", func(ctx echo.Context) error {
        data := "Hello from /index"
        return ctx.String(http.StatusOK, data)
    })

	r.GET("/index", func(ctx echo.Context) error{
		data := "Hello from /index"
		return ctx.String(http.StatusOK, data)
	})

	r.GET("/html", func(ctx echo.Context) error{
		data := "Hello from /html"
		return ctx.HTML(http.StatusOK, data)
	})

	r.GET("/index", func(ctx echo.Context) error{
		return ctx.Redirect(http.StatusTemporaryRedirect, "/")
	})

	r.GET("/json", func(ctx echo.Context) error{
		data := M{"Message": "Hello", "Counter":2}
		return ctx.JSON(http.StatusOK, data)
	})



	// Parsing Request
	// Parsing Query String
	r.GET("/page1", func(ctx echo.Context) error {
		name := ctx.QueryParam("name")
		data := fmt.Sprintf("Hello %s", name)
	
		return ctx.String(http.StatusOK, data)
	})
	// Test menggunakan curl:
	// curl -X GET http://localhost:9000/page1?name=grayson
	
	// Parsing URL Path Param
	r.GET("/page2/:name", func(ctx echo.Context) error {
		name := ctx.Param("name")
		data := fmt.Sprintf("Hello %s", name)
	
		return ctx.String(http.StatusOK, data)
	})	
	// Test menggunakan curl:
	// curl -X GET http://localhost:9000/page2/grayson	

	// Parsing URL Path Param dan Setelahnya
	r.GET("/page3/:name/*", func(ctx echo.Context) error {
		name := ctx.Param("name")
		message := ctx.Param("*")
	
		data := fmt.Sprintf("Hello %s, I have message for you: %s", name, message)
	
		return ctx.String(http.StatusOK, data)
	})
	// Test menggunakan curl:
	// curl -X GET http://localhost:9000/page3/tim/need/some/sleep	

	// Parsing Form Data
	r.POST("/page4", func(ctx echo.Context) error {
		name := ctx.FormValue("name")
		message := ctx.FormValue("message")
	
		data := fmt.Sprintf(
			"Hello %s, I have message for you: %s",
			name,
			strings.Replace(message, "/", "", 1),
		)
	
		return ctx.String(http.StatusOK, data)
	})

	// Test menggunakan curl:
	// curl -X POST -F name=damian -F message=angry http://localhost:9000/page4	



	// Penggunaan echo.WrapHandler Untuk Routing Handler Bertipe func(http.ResponseWriter,*http.Request) atau http.HandlerFunc
    r.GET("/index", echo.WrapHandler(http.HandlerFunc(ActionIndex)))
    r.GET("/home", echo.WrapHandler(ActionHome))
    r.GET("/about", ActionAbout)
	r.Static("/static", "assets")
	

	r.Start(":9000")
}

var ActionIndex = func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("from action index"))
}

var ActionHome = http.HandlerFunc(
    func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("from action home"))
    },
)

var ActionAbout = echo.WrapHandler(
    http.HandlerFunc(
        func(w http.ResponseWriter, r *http.Request) {
            w.Write([]byte("from action about"))
        },
    ),
)