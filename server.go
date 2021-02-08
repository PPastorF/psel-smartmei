package main

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	misc "github.com/ppastorf/psel-smartmei/internal/misc"
	user "github.com/ppastorf/psel-smartmei/internal/user"
	book "github.com/ppastorf/psel-smartmei/internal/book"
	storage "github.com/ppastorf/psel-smartmei/internal/storage"
)

func main() {
	config := &misc.Config{}
	if err := config.ReadFromFile("config/config.yaml"); err != nil {
		log.Fatal().Err(err).Msg("Erro ao ler o arquivo de configuracoes")
		panic(err)
    }

	if misc.InProductionEnv() {
		log.Logger = misc.ProdLogger()
	} else {
		log.Logger = misc.DevLogger()
	}

	db := storage.DatabaseConnection(config)
    defer db.Close()

    err := storage.CreateSchema(db)
    if err != nil {
		log.Fatal().Err(err).Msgf("Não foi possivel se conectar ao banco")
		panic(err)
    }

	e := echo.New()
	v1 := e.Group(misc.ApiRoutePrefix(config.Api.ApiVersion))
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &misc.CustomContext{c, db, config}
			return next(cc)
		}
	})

	// Rotas
	v1.GET("/user/:id", user.GetUser)
	v1.POST("/user", user.CreateUser)
	v1.POST("/book", book.CreateBook)
	v1.PUT("/book/lend", book.LendBook)
	v1.PUT("/book/return", book.ReturnBook)

	// Servidor
	e.Logger.Fatal(e.Start(misc.ConnectionURL(config.Server.Address, config.Server.HttpPort)))
}