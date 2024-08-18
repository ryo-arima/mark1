package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ryo-arima/mark1/pkg/entity/request"
	"github.com/ryo-arima/mark1/pkg/entity/response"
	"github.com/ryo-arima/mark1/pkg/server/middleware"
	"github.com/ryo-arima/mark1/pkg/server/repository"
)

type CommonControllerForInternal interface {
	CreateToken(c *gin.Context)
}

type commonControllerForInternal struct {
	UserRepository   repository.UserRepository
	CommonRepository repository.CommonRepository
}

func (commonControllerForInternal commonControllerForInternal) CreateToken(c *gin.Context) {
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
	user := commonControllerForInternal.UserRepository.FindUserByEmail(userRequest.User.Email)
	isAuthenticated, err := middleware.CheckHash(userRequest.User.Password, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, &response.UserResponse{Code: "SERVER_CONTROLLER_CREATE__FOR__003", Message: err.Error(), Users: []response.User{}})
		return
	}
	if isAuthenticated {
		var jwtKey = []byte("my_secret_key")

		type Claims struct {
			Username string `json:"username"`
			jwt.RegisteredClaims
		}
		expirationTime := time.Now().Add(5 * time.Minute)

		claims := &Claims{
			Username: user.Name,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(expirationTime),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			c.JSON(http.StatusBadRequest, &response.TokenResponse{Code: "SERVER_CONTROLLER_CREATE__FOR__004", Message: err.Error(), Token: tokenString})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, &response.UserResponse{Code: "SERVER_CONTROLLER_CREATE__FOR__005", Message: "Invalid", Users: []response.User{}})
		return
	}
}

func NewCommonControllerForInternal(userRepository repository.UserRepository, commonRpository repository.CommonRepository) CommonControllerForInternal {
	return &commonControllerForInternal{
		UserRepository:   userRepository,
		CommonRepository: commonRpository,
	}
}
