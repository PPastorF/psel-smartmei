package types

import (
)

type Book struct { 
	Id UniqueId `json:"id"`
	Title string `json:"title"`
	Pages int `json:"pages"`
	CreatedAt Timestamp `json:"created_at"`
}

type AddBookRequest struct {
	User string `json:"logged_user_id"`
	Title string `json:"title"`
	Pages int32 `json:"pages"`
}

type BookLending struct {
	BookId UniqueId `json:"book_id"`
	Owner UniqueId `json:"from_user"`
	To UniqueId `json:"to_user"`
	LentAt Timestamp `json:"lent_at"`
	ReturnedAt Timestamp `json:"returned_at"`
}

type LendBookRequest struct {
	User UniqueId `json:"logged_user_id"`
	BookId UniqueId `json:"book_id"`
	To UniqueId `json:"to_user_id"`
}

type ReturnBookRequest struct {
	User UniqueId `json:"logged_user_id"`
	BookId UniqueId `json:"book_id"`
}
