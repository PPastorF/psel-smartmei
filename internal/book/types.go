package book

import (
	"time"
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"

	misc "github.com/ppastorf/psel-smartmei/internal/misc"
)

type Book struct { 
	tableName struct{} `pg:"books"`

	ID misc.UniqueID `json:"id" pg:",pk"`
	Title string `json:"title" pg:",notnull"`
	Pages int `json:"pages" pg:",notnull"`

	CreatedAt time.Time `json:"created_at"`
}

type AddBookRequest struct {
	UserID misc.UniqueID `json:"logged_user_id"`
	Title string `json:"title" `
	Pages int32 `json:"pages" `
}

func (r *AddBookRequest) Sanitize() (*AddBookRequest, error) {
	title, err := misc.SanitizeString(r.Title)
	if err != nil {
		return nil, err
	}
	uid, err := misc.SanitizeString(r.UserID.String())
	if err != nil {
		return nil, err
	}

	sr := &AddBookRequest{
		UserID: misc.UniqueID(uid),
		Title: title,
		Pages: r.Pages,
	}
	return sr, nil
}

func (r *AddBookRequest) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.UserID, validation.Required, is.UUIDv4),
		validation.Field(&r.Title, validation.Required, is.UTFLetter),
		validation.Field(&r.Pages, validation.Min(1)),
	)
}

type BookLending struct {
	tableName struct{} `pg:"book_lendings"`

	ID	misc.UniqueID `json:"book_lending_id" pg:",pk"`
	BookID misc.UniqueID `json:"book_id" pg:",notnull"`
	OwnerID misc.UniqueID `json:"from_user" pg:",notnull"`
	ToUserID misc.UniqueID `json:"to_user" pg:",notnull"`
	
	LentAt time.Time `json:"lent_at"`
	ReturnedAt time.Time `json:"returned_at"`
}

type LendBookRequest struct {
	UserID misc.UniqueID `json:"logged_user_id"`
	BookID misc.UniqueID `json:"book_id"`
	ToUserID misc.UniqueID `json:"to_user_id"`
}

func (r *LendBookRequest) Sanitize() (*LendBookRequest, error) {
	uid, err := misc.SanitizeString(r.UserID.String())
	if err != nil {
		return nil, err
	}
	bid, err := misc.SanitizeString(r.BookID.String())
	if err != nil {
		return nil, err
	}
	touid, err := misc.SanitizeString(r.ToUserID.String())
	if err != nil {
		return nil, err
	}

	sr := &LendBookRequest{
		UserID: misc.UniqueID(uid),
		BookID: misc.UniqueID(bid),
		ToUserID: misc.UniqueID(touid),
	}
	return sr, nil
}

func (r *LendBookRequest) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.UserID, validation.Required, is.UUIDv4),
		validation.Field(&r.BookID, validation.Required, is.UUIDv4),
		validation.Field(&r.ToUserID, validation.Required, is.UUIDv4),
	)
}

type ReturnBookRequest struct {
	UserID misc.UniqueID `json:"logged_user_id"`
	BookID misc.UniqueID `json:"book_id"`
}

func (r *ReturnBookRequest) Sanitize() (*ReturnBookRequest, error) {
	uid, err := misc.SanitizeString(r.UserID.String())
	if err != nil {
		return nil, err
	}
	bid, err := misc.SanitizeString(r.BookID.String())
	if err != nil {
		return nil, err
	}

	sr := &ReturnBookRequest{
		UserID: misc.UniqueID(uid),
		BookID: misc.UniqueID(bid),
	}
	return sr, nil
}

func (r *ReturnBookRequest) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.UserID, validation.Required, is.UUIDv4),
		validation.Field(&r.BookID, validation.Required, is.UUIDv4),
	)
}
