package user

import (
	"net/http"
	"github.com/labstack/echo/v4"
	types "github.com/ppastorf/psel-smartmei/internal/types"
	misc "github.com/ppastorf/psel-smartmei/internal/misc"
)

// Res:
// { ...User }
func GetUser(c echo.Context) error {
	id := c.Param("id")

	// validar id

	// puxar usuario do banco

	// retornar

	return c.String(http.StatusOK, id)
}

// Req:
// {
// 	"name": string,
// 	"email": string
// }
// Res:
// { ...User }
func CreateUser(c echo.Context) error {
	params := new(types.CreateUserRequest)
	if err := c.Bind(params); err != nil {
		return err
	}

	// validate name
	name := params.Name

	// validate email
	email := params.Email
	
	id := misc.GenerateUniqueId()
	
	ts, err := misc.GenerateTimestamp()
	if err != nil {
		return err
	}
	
	u := types.User{
		Id: id,
		Name: name,
		Email: email,
		CreatedAt: ts,
		Collection: make([]types.Book, 0),
		Lent: make([]types.BookLending, 0),
		Borrowed: make([]types.BookLending, 0),
	}

	// checar se usuario ja esta cadastrado

	// salvar no banco

	err = misc.LogJsonData("New user created", u)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, u)
}
