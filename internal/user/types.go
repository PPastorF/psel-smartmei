package user

import (
	"time"
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	
	misc "github.com/ppastorf/psel-smartmei/internal/misc"
)

type User struct {
	tableName struct{} `pg:"users"`

	ID misc.UniqueID `json:"id" pg:",pk"`
	Name string `json:"name" pg:",notnull"`
	Email string `json:"email" pg:",unique,notnull"`
	
	CreatedAt time.Time `json:"created_at"`
	Collection []misc.UniqueID `json:"collection"`
	Lent []misc.UniqueID `json:"lent_books"`
	Borrowed []misc.UniqueID `json:"borrowed_books"`	
}

type CreateUserRequest struct {
	Name string `json:"name"`
	Email string `json:"email"`
}

func (r *CreateUserRequest) Sanitize() (*CreateUserRequest, error) {
	name, err := misc.SanitizeString(r.Name)
	if err != nil {
		return nil, err
	}
	email, err := misc.SanitizeString(r.Email)
	if err != nil {
		return nil, err
	}

	sr := &CreateUserRequest{
		Name: name,
		Email: email,
	}
	return sr, nil
}

func (r *CreateUserRequest) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Name, validation.Required, is.UTFLetter),
		validation.Field(&r.Email, validation.Required, is.EmailFormat),
	)
}

type GetUserRequest struct {
	UserID misc.UniqueID `json:"id"`
}

func (r *GetUserRequest) Sanitize() (*GetUserRequest, error) {
	uid, err := misc.SanitizeString(r.UserID.String())
	if err != nil {
		return nil, err
	}

	sr := &GetUserRequest{
		UserID: misc.UniqueID(uid),
	}
	return sr, nil
}

func (r *GetUserRequest) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.UserID, validation.Required, is.UUIDv4),
	)
}