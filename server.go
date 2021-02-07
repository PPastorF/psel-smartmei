package main

import (
	"github.com/rs/zerolog/log"
	"github.com/labstack/echo/v4"
	misc "github.com/ppastorf/psel-smartmei/internal/misc"
	user "github.com/ppastorf/psel-smartmei/internal/user"
	book "github.com/ppastorf/psel-smartmei/internal/book"
)

func main() {
	c, err := misc.ReadConfig("config/config.yaml")
    if err != nil {
        log.Fatal()
    }

	if misc.InProduction() {
		log.Logger = misc.ProdLogger()
	} else {
		log.Logger = misc.DevLogger()
	}

	e := echo.New()

	v1 := e.Group(c.Api.ApiVersion)

	// Rotas
	v1.GET("/user/:id", user.GetUser)
	v1.POST("/user", user.CreateUser)
	v1.POST("/book", book.CreateBook)
	v1.PUT("/book/lend", book.LendBook)
	v1.PUT("/book/return", book.ReturnBook)

	// Servidor
	e.Logger.Fatal(e.Start(":8080"))
}
