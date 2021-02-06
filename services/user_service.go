package services

import (
	"github.com/gyan1230/store_user-api/domain/users"
	"github.com/gyan1230/store_user-api/utils/errors"
)

//CreateUser :
func CreateUser(u users.User) (*users.User, *errors.RestErr) {
	if err := u.Validate(); err != nil {
		return nil, err
	}
	if err := u.Save(); err != nil {
		return nil, err
	}
	return &u, nil
}

//GetUser :
func GetUser(id int64) (*users.User, *errors.RestErr) {
	result := &users.User{
		ID: id,
	}
	err := result.Get()
	if err != nil {
		return nil, err
	}
	return result, nil
}
