package users

import (
	"strings"

	"github.com/gyan1230/store_user-api/utils/errors"
)

//User :
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Email       string `json:"email"`
	DateCreated string `json:"date"`
}

//Validate :
func (u *User) Validate() *errors.RestErr {
	u.Email = strings.TrimSpace(strings.ToLower(u.Email))
	if u.Email == "" {
		return errors.NewInternalServerError("Invalid email address")
	}
	return nil
}
