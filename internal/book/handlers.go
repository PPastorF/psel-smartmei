package book

import (
	// "net/http"
	// "github.com/rs/zerolog/log"
	"github.com/labstack/echo/v4"
	// "github.com/go-pg/pg/v10"

	// misc "github.com/ppastorf/psel-smartmei/internal/misc"
)

// Req:
// {
// 	"logged_user_id": UserID,
// 	"title": string,
// 	"pages": int
// }

// Res:
// { ...Book }
func AddBook(c echo.Context) error {
	return nil
}

// Req:
// {
// 	"logged_user_id": UserID,
// 	"book_id": BookID,
// 	"to_user_id": UserID
// }

// Res:
// { ...BookLending }
func LendBook(c echo.Context) error {
	return nil
}

// Req:
// {
// 	"logged_user_id": UserID,
// 	"book_id": BookID
// }

// Res:
// { ...BookLending }
func ReturnBook(c echo.Context) error {
	return nil
}