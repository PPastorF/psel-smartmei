package storage

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"

	user "github.com/ppastorf/psel-smartmei/internal/user"
	book "github.com/ppastorf/psel-smartmei/internal/book"
	conf "github.com/ppastorf/psel-smartmei/internal/conf"
)

func NewDBConnection(dbConf conf.DBConfig) *pg.DB {
	db := pg.Connect(&pg.Options{
		Addr: dbConf.DBConnectionURL(),
		User: dbConf.Auth.User,
		Password: dbConf.Auth.Pass,
		Database: dbConf.Name,
		ApplicationName: "psel-smartmei-api",
    })
	return db
}

func CreateDBSchema(db *pg.DB) error {
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