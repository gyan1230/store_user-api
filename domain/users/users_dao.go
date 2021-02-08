package users

import (
	"fmt"

	"github.com/gyan1230/store_user-api/datasources/mysql/usersdb"
	"github.com/gyan1230/store_user-api/utils/date"
	"github.com/gyan1230/store_user-api/utils/errors"
)

const (
	insertQuery = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	findQuery   = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
	updateQuery = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
)

//Get :
func (u *User) Get() *errors.RestErr {
	client := usersdb.Connect()
	if client == nil {
		return errors.NewInternalServerError("Users DB connect error")
	}
	if err := client.Ping(); err != nil {
		return errors.NewInternalServerError("Ping error")
	}

	stmnt, err := client.Prepare(findQuery)
	if err != nil {
		errors.NewBadRequestError(err.Error())
	}
	defer stmnt.Close()
	r := stmnt.QueryRow(u.ID)
	if err := r.Scan(&u.ID, &u.FirstName, &u.Lastname, &u.Email, &u.DateCreated); err != nil {
		return errors.NewBadRequestError("No user found")
	}
	return nil
}

//Save :
func (u *User) Save() *errors.RestErr {
	client := usersdb.Connect()
	if client == nil {
		return errors.NewInternalServerError("Users DB connect error")
	}
	if err := client.Ping(); err != nil {
		return errors.NewInternalServerError("Ping error")
	}
	stmnt, err := client.Prepare(insertQuery)
	if err != nil {
		return errors.NewBadRequestError(err.Error())
	}
	defer stmnt.Close()
	u.DateCreated = date.GetNowString()
	insertResult, err := stmnt.Exec(u.FirstName, u.Lastname, u.Email, u.DateCreated)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error while saving user: %s", err.Error()))
	}
	uID, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError("Error while getting last insert id")
	}
	u.ID = uID
	return nil
}

//Update :
func (u *User) Update() *errors.RestErr {
	client := usersdb.Connect()
	if client == nil {
		return errors.NewInternalServerError("Users DB connect error")
	}
	if err := client.Ping(); err != nil {
		return errors.NewInternalServerError("Ping error")
	}
	stmnt, err := client.Prepare(updateQuery)
	if err != nil {
		return errors.NewBadRequestError(err.Error())
	}

	defer stmnt.Close()
	_, err = stmnt.Exec(u.FirstName, u.Lastname, u.Email, u.ID)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error while updating user: %s", err.Error()))
	}
	return nil
}
