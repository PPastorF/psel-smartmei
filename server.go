package main

import (
	"os"
	"fmt"
    "strings"
	"context"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
    "github.com/go-pg/pg/v10"
    "github.com/go-pg/pg/v10/orm"
	misc "github.com/ppastorf/psel-smartmei/internal/misc"
	types "github.com/ppastorf/psel-smartmei/internal/types"
	user "github.com/ppastorf/psel-smartmei/internal/user"
	book "github.com/ppastorf/psel-smartmei/internal/book"
)

func main() {
	config := &misc.Config{}
	if err := config.ReadFromFile("config/config.yaml"); err != nil {
		log.Fatal().Err(err).Msg("Erro ao ler o arquivo de configurações.")
		panic(err)
    }

	if inProd() {
		log.Logger = misc.ProdLogger()
	} else {
		log.Logger = misc.DevLogger()
	}

	db := pg.Connect(&pg.Options{
		Addr: serverAddress(config.Db.Endpoint, config.Db.Port),
		User: config.Db.Auth.User,
		Password: config.Db.Auth.Pass,
		Database: config.Db.DbName,
		ApplicationName: "psel-smartmei-api",
    })
    defer db.Close()

	ctx := context.Background()
	
	if err := db.Ping(ctx); err != nil {
		panic(err)
	}

    err := createSchema(db)
    if err != nil {
		log.Fatal().Err(err).Msgf("Não foi possivel se conectar ao banco.")
		panic(err)
    }

	e := echo.New()
	v1 := e.Group(apiRoutePrefix(config.Api.ApiVersion))

	// Rotas
	v1.GET("/user/:id", user.GetUser)
	v1.POST("/user", user.CreateUser)
	v1.POST("/book", book.CreateBook)
	v1.PUT("/book/lend", book.LendBook)
	v1.PUT("/book/return", book.ReturnBook)

	// Servidor
	e.Logger.Fatal(e.Start(serverAddress(config.Server.Address, config.Server.HttpPort)))
}

func inProd() bool {
    if strings.ToUpper(os.Getenv("DEPLOY_ENV")) == "PROD" {
        return true
    } else {
        return false
    }
}

func apiRoutePrefix(version string) string {
	return fmt.Sprintf("/api/%s", version)
}

func serverAddress(addr, port string) string {
	return fmt.Sprintf("%s:%s", addr, port)
}

func createSchema(db *pg.DB) error {
    models := []interface{}{
        (*types.User)(nil),
        (*types.Book)(nil),
        (*types.BookLending)(nil),
    }

    for _, model := range models {
        err := db.Model(model).CreateTable(&orm.CreateTableOptions{
            Temp: true,
        })
        if err != nil {
            return err
        }
    }
    return nil
}