package misc

import (
	"github.com/labstack/echo/v4"
	"github.com/go-pg/pg/v10"

	conf "github.com/ppastorf/psel-smartmei/internal/conf"
)

type UniqueID string

func (id UniqueID) String() string {
	return string(id)
}

type CustomContext struct {
	echo.Context
	DB *pg.DB
	Config *conf.AppConfig
}