package controller

import (
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
	commonControllerForPublic.CommonRepository.CreateEmail(model.Email{
		To: userRequest.User.Email,
	}, tempCode)
}

func (commonControllerForPublic commonControllerForPublic) VerifyEmail(c *gin.Context) {
	code := c.Query("code")
	uuid := c.Query("uuid")
	if code == "" {
		c.JSON(http.StatusBadRequest, &response.UserResponse{Code: "SERVER_CONTROLLER_VERIFY__FOR__001", Message: "code is required", Users: []response.User{}})
		return
	} else {
		user := commonControllerForPublic.UserRepository.FindUserByUUID(uuid)
		tempCode := commonControllerForPublic.CommonRepository.GetTempCode(model.Email{
			To: user.Email,
		})
		if code != tempCode {
			c.JSON(http.StatusBadRequest, &response.UserResponse{Code: "SERVER_CONTROLLER_VERIFY__FOR__002", Message: "code is invalid", Users: []response.User{}})
			return
		} else {
			commonControllerForPublic.UserRepository.UpdateUser(model.Users{
				UUID:   uuid,
				Status: "EmailVerified",
			})
			c.JSON(http.StatusOK, &response.UserResponse{Code: "SERVER_CONTROLLER_VERIFY__FOR__003", Message: "ok", Users: []response.User{}})
			return
		}
	}
}

func NewCommonControllerForPublic(userRepository repository.UserRepository, commonRepository repository.CommonRepository) CommonControllerForPublic {
	return &commonControllerForPublic{
		UserRepository:   userRepository,
		CommonRepository: commonRepository,
	}
}
