package misc

import (
	"github.com/labstack/echo/v4"
	"github.com/go-pg/pg/v10"
)

type UniqueID string

type CustomContext struct {
	echo.Context
	DB *pg.DB
	Config *Config
}