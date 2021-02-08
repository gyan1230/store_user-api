package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gyan1230/store_user-api/controllers/users"
)

func mapUrls() {
	api := router.Group("/api/v1")
	api.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Health": "OK",
		})
	})
	// api.GET("/user", users.GetAllUser)
	api.POST("/create", users.CreateUser)
	api.GET("/user/:user_id", users.GetUser)
	api.PUT("/user/:user_id", users.UpdateUser)

}
