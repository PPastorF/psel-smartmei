package types

import (
)

type User struct {
	Id UniqueId `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	CreatedAt Timestamp `json:"created_at"`
	Collection []Book `json:"collection"`
	Lent []BookLending `json:"lent_books"`
	Borrowed []BookLending `json:"borrowed_books"`
}

type CreateUserRequest struct {
	Name string `json:"name"`
	Email string `json:"email"`
}