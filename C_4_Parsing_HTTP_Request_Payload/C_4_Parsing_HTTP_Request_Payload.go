package main

import (
    "fmt"
    "github.com/labstack/echo"
    "net/http"
)

type User struct {
    Name  string `json:"name" form:"name" query:"name"`
    Email string `json:"email" form:"email" query:"email"`
}

func main() {
    r := echo.New()

	r.Any("/user", func(c echo.Context) (err error) {
		u := new(User)
		if err = c.Bind(u); err != nil {
			return
		}
	
		return c.JSON(http.StatusOK, u)
	})

    fmt.Println("server started at :9000")
    r.Start(":9000")
}

// Testing
// Form Data
// curl -X POST http://localhost:9000/user -d "name=Okta" -d "email=okta@google.com"
// # output => {"name":"Okta","email":"okta@google.com"}

// JSON Payload
// curl -X POST http://localhost:9000/user -H "Content-Type: application/json" -d "{\"name\":\"Okta\",\"email\":\"okta@google.com\"}"
// # output => {"name":"Okta","email":"okta@google.com"}

// XML Payload
// curl -X POST http://localhost:9000/user -H "Content-Type: application/xml" -d "<?xml version="1.0"?> <Data> <Name>Okta</Name> <Email>okta@google.com</Email> </Data>"
// # output => {"name":"Okta","email":"okta@google.com"}

// Query String
// curl -X GET "http://localhost:9000/user?name=Okta&email=okta@google.com"
// # output => {"name":"Okta","email":"okta@google.com"}