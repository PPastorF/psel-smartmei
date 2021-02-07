package user

import (
	"github.com/labstack/echo/v4"
)

// Req:
// {
// 	"logged_user_id": UserID,
// 	"title": string,
// 	"pages": int
// }

// Res:
// { ...Book }
func CreateBook(c echo.Context) error {
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