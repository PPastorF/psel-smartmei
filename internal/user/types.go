package user

import (
	"time"
	"github.com/go-playground/validator/v10"

	misc "github.com/ppastorf/psel-smartmei/internal/misc"
)

type User struct {
	tableName struct{} `pg:"users"`

	ID misc.UniqueID `json:"id" pg:",pk"`
	Name string `json:"name"`
	Email string `json:"email" pg:",unique,notnull"`
	CreatedAt time.Time `json:"created_at"`
	Collection []misc.UniqueID `json:"collection"`
	Lent []misc.UniqueID `json:"lent_books"`
	Borrowed []misc.UniqueID `json:"borrowed_books"`	
}

type CreateUserRequest struct {
	Name string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

func (r *CreateUserRequest) SanitizeAndValidate() (*CreateUserRequest, error) {
	sanitized := &CreateUserRequest{
		misc.SanitizeString(r.Name),
		misc.SanitizeString(r.Email),
	}

	validate := validator.New()

	err := validate.Struct(sanitized)
	if err != nil {
		return nil, err
	}

	return sanitized, nil
}

type GetUserRequest struct {
	ID misc.UniqueID `json:"id" validate:"required,uuid4_rfc4122"`
}