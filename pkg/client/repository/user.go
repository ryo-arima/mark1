package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ryo-arima/mark1/pkg/config"
	"github.com/ryo-arima/mark1/pkg/entity/model"
	"github.com/ryo-arima/mark1/pkg/entity/request"
	"github.com/ryo-arima/mark1/pkg/entity/response"
)

type UserRepository interface {
	BootstrapUserForDB(request request.UserRequest) (response response.UserResponse)
	GetUserForPublic(request request.UserRequest) (response response.UserResponse)
	GetUserForInternal(request request.UserRequest) (response response.UserResponse)
	GetUserForPrivate(request request.UserRequest) (response response.UserResponse)
	CreateUserForPublic(request request.UserRequest) (response response.UserResponse)
	CreateUserForInternal(request request.UserRequest) (response response.UserResponse)
	CreateUserForPrivate(request request.UserRequest) (response response.UserResponse)
	UpdateUserForInternal(request request.UserRequest) (response response.UserResponse)
	UpdateUserForPrivate(request request.UserRequest) (response response.UserResponse)
	DeleteUserForInternal(request request.UserRequest) (response response.UserResponse)
	DeleteUserForPrivate(request request.UserRequest) (response response.UserResponse)
}

type userRepository struct {
	BaseConfig config.BaseConfig
}

// Bootstrap
func (userRepository userRepository) BootstrapUserForDB(request request.UserRequest) (response response.UserResponse) {
	userRepository.BaseConfig.DBConnection.Migrator().DropTable(&model.Users{})
	userRepository.BaseConfig.DBConnection.Migrator().CreateTable(&model.Users{})
	return response
}

// GET
func (userRepository userRepository) GetUserForPublic(request request.UserRequest) (response response.UserResponse) {
	return response
}

func (userRepository userRepository) GetUserForInternal(request request.UserRequest) (response response.UserResponse) {
	fmt.Println("GetUserForInternal")
	return response
}

func (userRepository userRepository) GetUserForPrivate(request request.UserRequest) (response response.UserResponse) {
	fmt.Println("GetUserForPrivate")
	return response
}

// CREATE
func (userRepository userRepository) CreateUserForPublic(request request.UserRequest) (response response.UserResponse) {
	URL := userRepository.BaseConfig.YamlConfig.Application.Client.ServerEndpoint + "/api/public/user"

	jsonData, err := json.Marshal(request)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	resp, err := http.Post(URL, "application/json", bytes.NewBuffer([]byte(jsonData)))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Status Code:", resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Response Body:", string(body))
	return response
}

func (userRepository userRepository) CreateUserForInternal(request request.UserRequest) (response response.UserResponse) {
	fmt.Println("CreateUserForInternal()")
	return response
}

func (userRepository userRepository) CreateUserForPrivate(request request.UserRequest) (response response.UserResponse) {
	fmt.Println("CreateUserForPrivate()")
	return response
}

// UPDATE
func (userRepository userRepository) UpdateUserForInternal(request request.UserRequest) (response response.UserResponse) {
	fmt.Println("UpdateUserForInternal")
	return response
}

func (userRepository userRepository) UpdateUserForPrivate(request request.UserRequest) (response response.UserResponse) {
	fmt.Println("UpdateUserForPrivate")
	return response
}

// DELETE
func (userRepository userRepository) DeleteUserForInternal(request request.UserRequest) (response response.UserResponse) {
	fmt.Println("DeleteUserForInternal")
	return response
}

func (userRepository userRepository) DeleteUserForPrivate(request request.UserRequest) (response response.UserResponse) {
	fmt.Println("DeleteUserForPrivate")
	return response
}

func NewUserRepository(conf config.BaseConfig) UserRepository {
	return &userRepository{BaseConfig: conf}
}
