// Gorilla Session
package main

import (
    "fmt"
    // "github.com/gorilla/context"
    // "github.com/gorilla/sessions"
    "github.com/labstack/echo"
    "net/http"
)

const SESSION_ID = "id"

func main() {
    store := newMongoStore()

    e := echo.New()

    e.Use(echo.WrapMiddleware(context.ClearHandler))

	e.GET("/set", func(c echo.Context) error {
		session, _ := store.Get(c.Request(), SESSION_ID)
		session.Values["message1"] = "hello"
		session.Values["message2"] = "world"
		session.Save(c.Request(), c.Response())
	
		return c.Redirect(http.StatusTemporaryRedirect, "/get")
	})
	
	e.GET("/get", func(c echo.Context) error {
		session, _ := store.Get(c.Request(), SESSION_ID)
	
		if len(session.Values) == 0 {
			return c.String(http.StatusOK, "empty result")
		}
	
		return c.String(http.StatusOK, fmt.Sprintf(
			"%s %s",
			session.Values["message1"],
			session.Values["message2"],
		))
	})

	e.GET("/delete", func(c echo.Context) error {
		session, _ := store.Get(c.Request(), SESSION_ID)
		session.Options.MaxAge = -1
		session.Save(c.Request(), c.Response())
	
		return c.Redirect(http.StatusTemporaryRedirect, "/get")
	})
}
