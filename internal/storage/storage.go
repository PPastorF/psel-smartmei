package storage

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"

	user "github.com/ppastorf/psel-smartmei/internal/user"
	book "github.com/ppastorf/psel-smartmei/internal/book"
	misc "github.com/ppastorf/psel-smartmei/internal/misc"
)

func DatabaseConnection(config *misc.Config) *pg.DB {
	db := pg.Connect(&pg.Options{
		Addr: misc.ConnectionURL(config.Db.Endpoint, config.Db.Port),
		User: config.Db.Auth.User,
		Password: config.Db.Auth.Pass,
		Database: config.Db.DbName,
		ApplicationName: "psel-smartmei-api",
    })
	return db
}

func CreateSchema(db *pg.DB) error {
    models := []interface{}{
        (*user.User)(nil),
        (*book.Book)(nil),
        (*book.BookLending)(nil),
    }

    for _, model := range models {
        err := db.Model(model).CreateTable(&orm.CreateTableOptions{
            IfNotExists: true,
        })
        if err != nil {
            return err
        }
    }
    return nil
}