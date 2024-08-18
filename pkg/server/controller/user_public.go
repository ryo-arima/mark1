package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/ryo-arima/mark1/pkg/entity/model"
	"github.com/ryo-arima/mark1/pkg/entity/request"
	"github.com/ryo-arima/mark1/pkg/entity/response"
	"github.com/ryo-arima/mark1/pkg/server/middleware"
	"github.com/ryo-arima/mark1/pkg/server/repository"
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
}

func NewUserControllerForPublic(userRepository repository.UserRepository) UserControllerForPublic {
	return &userControllerForPublic{UserRepository: userRepository}
}
