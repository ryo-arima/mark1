package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/mark1/pkg/entity/model"
	"github.com/ryo-arima/mark1/pkg/entity/request"
	"github.com/ryo-arima/mark1/pkg/entity/response"
	"github.com/ryo-arima/mark1/pkg/server/repository"
)

type UserControllerForPrivate interface {
	GetUsers(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type userControllerForPrivate struct {
	UserRepository repository.UserRepository
}

func (userController userControllerForPrivate) GetUsers(c *gin.Context) {
	var userRequest request.UserRequest
	if err := c.Bind(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.UserResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Users: []response.User{}})
		return
	}
	res := userController.UserRepository.GetUsers()
	c.JSON(http.StatusOK, res)
	return
}


func (userController userControllerForPrivate) CreateUser(c *gin.Context) {
	var userRequest request.UserRequest
	if err := c.Bind(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.UserResponse{Code: "SERVER_CONTROLLER_CREATE__FOR__001", Message: err.Error(), Users: []response.User{}})
		return
	}
	var userModel model.Users
	res := userController.UserRepository.CreateUser(userModel)
	c.JSON(http.StatusOK, res)
	return
}


func (userController userControllerForPrivate) UpdateUser(c *gin.Context) {
	var userRequest request.UserRequest
	if err := c.Bind(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.UserResponse{Code: "SERVER_CONTROLLER_UPDATE__FOR__001", Message: err.Error(), Users: []response.User{}})
		return
	}
	var userModel model.Users
	res := userController.UserRepository.UpdateUser(userModel)
	c.JSON(http.StatusOK, res)
	return
}


func (userController userControllerForPrivate) DeleteUser(c *gin.Context) {
	var userRequest request.UserRequest
	if err := c.Bind(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.UserResponse{Code: "SERVER_CONTROLLER_DELETE__FOR__001", Message: err.Error(), Users: []response.User{}})
		return
	}
	var uuid string
	res := userController.UserRepository.DeleteUser(uuid)
	c.JSON(http.StatusOK, res)
	return
}


func NewUserControllerForPrivate(userRepository repository.UserRepository) UserControllerForPrivate {
	return &userControllerForPrivate{ UserRepository: userRepository}
}