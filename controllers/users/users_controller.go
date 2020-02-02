package users

import (
	"github.com/bookstore_users-api/domain/users"
	"github.com/bookstore_users-api/services"
	"github.com/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context){
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil{
		restErr := errors.NewBadRequest("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveError := services.CreateUser(user)
	if saveError != nil{
		c.JSON(saveError.Status, saveError)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context){
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil{
		err := errors.NewBadRequest("user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	user, getError := services.GetUser(userId)
	if getError != nil{
		c.JSON(getError.Status, getError)
		return
	}
	c.JSON(http.StatusOK, user)
}
