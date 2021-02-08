package types

import (
)

type User struct {
	tableName struct{} `pg:"users"`

	Id UniqueId `json:"id" pg:",pk"`
	Name string `json:"name"`
	Email string `json:"email" pg:",unique,notnull"`
	CreatedAt Timestamp `json:"created_at"`
	Collection []Book `json:"collection"`
	Lent []BookLending `json:"lent_books"`
	Borrowed []BookLending `json:"borrowed_books"`
	
}

type CreateUserRequest struct {
	Name string `json:"name"`
	Email string `json:"email"`
}