package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/mark1/pkg/entity/model"
	"github.com/ryo-arima/mark1/pkg/entity/request"
	"github.com/ryo-arima/mark1/pkg/entity/response"
	"github.com/ryo-arima/mark1/pkg/server/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserControllerForPublic interface {
	CreateUser(c *gin.Context)
}

type userControllerForPublic struct {
	UserRepository repository.UserRepository
}

func (userController userControllerForPublic) CreateUser(c *gin.Context) {
	var userRequest request.UserRequest
	// Bind the request with the struct
	if err := c.Bind(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.UserResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Users: []response.User{}})
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(userRequest.User.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, &response.UserResponse{Code: "SERVER_CONTROLLER_GET__FOR__002", Message: err.Error(), Users: []response.User{}})
	}

	// Create the user
	userModel := model.Users{
		Name:     userRequest.User.Name,
		Password: string(hash),
		Email:    userRequest.User.Email,
	}
	userController.UserRepository.CreateUser(userModel)

	// Return the response
	response := response.UserResponse{
		Code:    "SERVER_CONTROLLER_CREATE__FOR__001",
		Message: "User created successfully",
		Users:   []response.User{},
	}
	c.JSON(http.StatusOK, response)
}

func NewUserControllerForPublic(userRepository repository.UserRepository) UserControllerForPublic {
	return &userControllerForPublic{UserRepository: userRepository}
}
