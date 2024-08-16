package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/mark1/pkg/entity/request"
	"github.com/ryo-arima/mark1/pkg/entity/response"
	"github.com/ryo-arima/mark1/pkg/server/repository"
)

type UserControllerForPublic interface {
	GetUsers(c *gin.Context)
}

type userControllerForPublic struct {
	UserRepository repository.UserRepository
}

func (userController userControllerForPublic) GetUsers(c *gin.Context) {
	var userRequest request.UserRequest
	if err := c.Bind(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.UserResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Users: []response.User{}})
		return
	}
	res := userController.UserRepository.GetUsers()
	c.JSON(http.StatusOK, res)
	return
}


func NewUserControllerForPublic(userRepository repository.UserRepository) UserControllerForPublic {
	return &userControllerForPublic{ UserRepository: userRepository}
}