package book

import (
	"time"
	misc "github.com/ppastorf/psel-smartmei/internal/misc"
)

type Book struct { 
	ID misc.UniqueID `json:"id"`
	Title string `json:"title"`
	Pages int `json:"pages"`
	CreatedAt time.Time `json:"created_at"`
}

type AddBookRequest struct {
	User misc.UniqueID `json:"logged_user_id" validate:"required,uuid4_rfc4122"`
	Title string `json:"title" validate:"required"`
	Pages int32 `json:"pages" validate:"required"`
}

type BookLending struct {
	ID	misc.UniqueID `json:"book_lending_id"`
	BookID misc.UniqueID `json:"book_id"`
	OwnerID misc.UniqueID `json:"from_user"`
	ToUserID misc.UniqueID `json:"to_user"`
	LentAt time.Time `json:"lent_at"`
	ReturnedAt time.Time `json:"returned_at"`
}

type LendBookRequest struct {
	UserID misc.UniqueID `json:"logged_user_id" validate:"required,uuid4_rfc4122"`
	BookID misc.UniqueID `json:"book_id" validate:"required,uuid4_rfc4122"`
	ToUserID misc.UniqueID `json:"to_user_id" validate:"required,uuid4_rfc4122"`
}

type ReturnBookRequest struct {
	UserID misc.UniqueID `json:"logged_user_id" validate:"required,uuid4_rfc4122"`
	BookID misc.UniqueID `json:"book_id" validate:"required,uuid4_rfc4122"`
}
