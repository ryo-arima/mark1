package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/ryo-arima/mark1/pkg/entity/model"
	"github.com/ryo-arima/mark1/pkg/entity/request"
	"github.com/ryo-arima/mark1/pkg/entity/response"
	"github.com/ryo-arima/mark1/pkg/server/middleware"
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
	var res response.UserResponse
	users := userController.UserRepository.GetUsers()
	for _, user := range users {
		res.Users = append(res.Users, response.User{
			ID:        fmt.Sprintf("%v", user.ID),
			UUID:      user.UUID,
			Email:     user.Email,
			Name:      user.Name,
			Status:    user.Status,
			CreatedAt: middleware.GetStringTime(user.CreatedAt),
			UpdatedAt: middleware.GetStringTime(user.UpdatedAt),
			DeletedAt: middleware.GetStringTime(user.DeletedAt),
		})
	}
	c.JSON(http.StatusOK, response.UserResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: "success", Users: res.Users})
	return
}

func (userController userControllerForPrivate) CreateUser(c *gin.Context) {
	var userRequest request.UserRequest
	// Bind the request with the struct
	if err := c.Bind(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.UserResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Users: []response.User{}})
	}
	validate := validator.New()
	err := validate.Struct(userRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, &response.UserResponse{Code: "SERVER_CONTROLLER_GET__FOR__002", Message: err.Error(), Users: []response.User{}})
	} else {
		c.JSON(http.StatusOK, &response.UserResponse{Code: "SERVER_CONTROLLER_GET__FOR__003", Message: "Validation passed", Users: []response.User{}})
	}

	hash, err := middleware.GenHash(userRequest.User.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, &response.UserResponse{Code: "SERVER_CONTROLLER_GET__FOR__002", Message: err.Error(), Users: []response.User{}})
	}

	userModel := model.Users{
		UUID:      middleware.NewUUID(),
		Name:      userRequest.User.Name,
		Password:  hash,
		Email:     userRequest.User.Email,
		CreatedAt: middleware.GetNowTime(),
		UpdatedAt: middleware.GetNowTime(),
	}
	userController.UserRepository.CreateUser(userModel)
	createdUser := userController.UserRepository.FindUserByUUID(userModel.UUID)
	response := response.UserResponse{
		Code:    "SERVER_CONTROLLER_CREATE__FOR__001",
		Message: "User created successfully",
		Users:   []response.User{{UUID: createdUser.UUID, Name: createdUser.Name, Email: createdUser.Email, CreatedAt: *createdUser.CreatedAt, UpdatedAt: *createdUser.UpdatedAt}},
	}
	c.JSON(http.StatusOK, response)
	return
}

func (userController userControllerForPrivate) UpdateUser(c *gin.Context) {
	var userRequest request.UserRequest
	if err := c.Bind(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.UserResponse{Code: "SERVER_CONTROLLER_UPDATE__FOR__001", Message: err.Error(), Users: []response.User{}})
		return
	}

	// Find the existing user
	existingUser := userController.UserRepository.FindUserByUUID(userRequest.User.UUID)
	if existingUser.UUID == "" {
		c.JSON(http.StatusNotFound, &response.UserResponse{Code: "SERVER_CONTROLLER_UPDATE__FOR__002", Message: "User not found", Users: []response.User{}})
		return
	}

	// Update user fields
	existingUser.Name = userRequest.User.Name
	if existingUser.Email != userRequest.User.Email {
		existingUser.Email = userRequest.User.Email
		existingUser.Status = "EmailNotVerified"
	}
	existingUser.UpdatedAt = middleware.GetNowTime()

	// Save the updated user
	userController.UserRepository.UpdateUser(existingUser)

	// Prepare the response
	updatedUser := userController.UserRepository.FindUserByUUID(existingUser.UUID)
	response := response.UserResponse{
		Code:    "SERVER_CONTROLLER_UPDATE__FOR__003",
		Message: "User updated successfully",
		Users:   []response.User{{UUID: updatedUser.UUID, Name: updatedUser.Name, Email: updatedUser.Email, Status: updatedUser.Status, CreatedAt: *updatedUser.CreatedAt, UpdatedAt: *updatedUser.UpdatedAt}},
	}
	c.JSON(http.StatusOK, response)
	return
}

func (userController userControllerForPrivate) DeleteUser(c *gin.Context) {
	var userRequest request.UserRequest
	if err := c.Bind(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.UserResponse{Code: "SERVER_CONTROLLER_DELETE__FOR__001", Message: err.Error(), Users: []response.User{}})
		return
	}
	existingUser := userController.UserRepository.FindUserByUUID(userRequest.User.UUID)
	if existingUser.UUID == "" {
		c.JSON(http.StatusNotFound, &response.UserResponse{Code: "SERVER_CONTROLLER_UPDATE__FOR__002", Message: "User not found", Users: []response.User{}})
		return
	}
	res := userController.UserRepository.DeleteUser(existingUser.UUID)
	c.JSON(http.StatusOK, res)
	return
}

func NewUserControllerForPrivate(userRepository repository.UserRepository) UserControllerForPrivate {
	return &userControllerForPrivate{UserRepository: userRepository}
}
