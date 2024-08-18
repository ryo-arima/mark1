package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/ryo-arima/mark1/pkg/entity/model"
	"github.com/ryo-arima/mark1/pkg/entity/request"
	"github.com/ryo-arima/mark1/pkg/entity/response"
	"github.com/ryo-arima/mark1/pkg/server/repository"
)

type CommonControllerForPublic interface {
	CreateEmail(c *gin.Context)
	VerifyEmail(c *gin.Context)
}

type commonControllerForPublic struct {
	UserRepository   repository.UserRepository
	CommonRepository repository.CommonRepository
}

func (commonControllerForPublic commonControllerForPublic) CreateEmail(c *gin.Context) {
	var userRequest request.UserRequest
	if err := c.Bind(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.UserResponse{Code: "SERVER_CONTROLLER_CREATE__FOR__001", Message: err.Error(), Users: []response.User{}})
		return
	}
	validator := validator.New()
	err := validator.Struct(userRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, &response.UserResponse{Code: "SERVER_CONTROLLER_CREATE__FOR__002", Message: err.Error(), Users: []response.User{}})
		return
	}
	tempCode := commonControllerForPublic.CommonRepository.SetTempCode(model.Email{
		To: userRequest.User.Email,
	})
	url := fmt.Sprintf("http://localhost:8080/email?code=%s", tempCode)
	commonControllerForPublic.CommonRepository.CreateEmail(model.Email{
		To:             userRequest.User.Email,
		VeryfyEmailURL: url,
	})
}

func (commonControllerForPublic commonControllerForPublic) VerifyEmail(c *gin.Context) {
	var userRequest request.UserRequest
	fmt.Println(userRequest)
	c.Query("code")
}

func NewCommonControllerForPublic(userRepository repository.UserRepository, commonRepository repository.CommonRepository) CommonControllerForPublic {
	return &commonControllerForPublic{
		UserRepository:   userRepository,
		CommonRepository: commonRepository,
	}
}
