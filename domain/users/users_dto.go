package users

import (
	"fmt"
	"time"

	"github.com/gyan1230/store_user-api/utils/errors"
)

var (
	userDb = make(map[int64]*User)
)

//Get :
func (u *User) Get() *errors.RestErr {
	result := userDb[u.ID]
	if result == nil {
		return errors.NewNotFoundErr(fmt.Sprintf("User %d not found\n", u.ID))
	}
	u.ID = result.ID
	u.FirstName = result.FirstName
	u.Lastname = result.Lastname
	u.Email = result.Email
	u.DateCreated = result.DateCreated
	return nil
}

//Save :
func (u *User) Save() *errors.RestErr {
	result := userDb[u.ID]
	if result != nil {
		if result.Email == u.Email {
			return errors.NewNotFoundErr(fmt.Sprintf("User email %s already taken\n", u.Email))
		}
		return errors.NewNotFoundErr(fmt.Sprintf("User %d already exists.\n", u.ID))
	}
	now := time.Now()
	u.DateCreated = now.Format("19-02-2006T15:04:00Z")
	userDb[u.ID] = u
	return nil
}
