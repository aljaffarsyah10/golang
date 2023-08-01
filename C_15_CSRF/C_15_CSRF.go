package main

import (
    "fmt"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    "html/template"
    "net/http"
)

type M map[string]interface{}

func main() {
    tmpl := template.Must(template.ParseGlob("./*.html"))

    e := echo.New()

	const CSRFTokenHeader = "X-CSRF-Token"
	const CSRFKey = "csrf"
	
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "header:" + CSRFTokenHeader,
		ContextKey:  CSRFKey,
	}))
	
	e.GET("/index", func(c echo.Context) error {
		data := make(M)
		data[CSRFKey] = c.Get(CSRFKey)
		return tmpl.Execute(c.Response(), data)
	})

	e.POST("/sayhello", func(c echo.Context) error {
		data := make(M)
		if err := c.Bind(&data); err != nil {
			return err
		}
	
		message := fmt.Sprintf("hello %s", data["name"])
		return c.JSON(http.StatusOK, message)
	})

    e.Logger.Fatal(e.Start(":9000"))
}


// Testing pakai curl
// curl -X POST http://localhost:9000/sayhello -H 'Content-Type: application/json' -d '{"name":"noval","gender":"male"}'