package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/gyan1230/store_user-api/domain/users"
	"github.com/gyan1230/store_user-api/services"
	"github.com/gyan1230/store_user-api/utils/errors"
)

//GetAllUser :
func GetAllUser(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"Message": "Implement Me",
	})
}

//GetUser :
func GetUser(c *gin.Context) {
	id, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("Id invalid")
		c.JSON(err.Status, err)
		return
	}
	u, getErr := services.GetUser(id)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusFound, u)
}

//CreateUser :
func CreateUser(c *gin.Context) {
	var u users.User
	err := c.ShouldBindJSON(&u)
	if err != nil {
		resterr := errors.NewBadRequestError("Invalid JSON")
		c.JSON(resterr.Status, resterr)
		return
	}
	user, saveErr := services.CreateUser(u)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, user)
}
