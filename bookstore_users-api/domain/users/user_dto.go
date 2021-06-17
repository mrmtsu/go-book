package users

import (
	"strings"

	"github.com/mrmtsu/go-book/bookstore_users-api/utils/errors"
)

const (
	StatusActive = "active"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Status      string `json:"status"`
	Password    string `json:"password"`
	DateCreated string `json:"date_created"`
}

func (u *User) Validate() *errors.RestErr {
	u.FirstName = strings.TrimSpace(u.FirstName)
	u.LastName = strings.TrimSpace(u.LastName)

	u.Email = strings.TrimSpace(strings.ToLower(u.Email))
	if u.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}

	u.Password = strings.TrimSpace(u.Password)
	if u.Password == "" {
		return errors.NewBadRequestError("invalid email password")
	}
	return nil
}
