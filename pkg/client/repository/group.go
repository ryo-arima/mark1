package repository

import (
	"fmt"

	"github.com/ryo-arima/mark1/pkg/config"
	"github.com/ryo-arima/mark1/pkg/entity/model"
	"github.com/ryo-arima/mark1/pkg/entity/request"
	"github.com/ryo-arima/mark1/pkg/entity/response"
)

type GroupRepository interface {
	BootstrapGroupForDB(request request.GroupRequest) (response response.GroupResponse)
	GetGroupForPublic(request request.GroupRequest) (response response.GroupResponse)
	GetGroupForInternal(request request.GroupRequest) (response response.GroupResponse)
	GetGroupForPrivate(request request.GroupRequest) (response response.GroupResponse)
	CreateGroupForInternal(request request.GroupRequest) (response response.GroupResponse)
	CreateGroupForPrivate(request request.GroupRequest) (response response.GroupResponse)
	UpdateGroupForInternal(request request.GroupRequest) (response response.GroupResponse)
	UpdateGroupForPrivate(request request.GroupRequest) (response response.GroupResponse)
	DeleteGroupForInternal(request request.GroupRequest) (response response.GroupResponse)
	DeleteGroupForPrivate(request request.GroupRequest) (response response.GroupResponse)
}

type groupRepository struct {
	BaseConfig config.BaseConfig
}

// Bootstrap
func (groupRepository groupRepository) BootstrapGroupForDB(request request.GroupRequest) (response response.GroupResponse) {
	fmt.Println("BootstrapGroupForDB")
	groupRepository.BaseConfig.DBConnection.Migrator().DropTable(&model.Groups{})
	groupRepository.BaseConfig.DBConnection.Migrator().CreateTable(&model.Groups{})
	return response
}

// GET
func (groupRepository groupRepository) GetGroupForPublic(request request.GroupRequest) (response response.GroupResponse) {
	fmt.Println("GetGroupForPublic")
	return response
}

func (groupRepository groupRepository) GetGroupForInternal(request request.GroupRequest) (response response.GroupResponse) {
	fmt.Println("GetGroupForInternal")
	return response
}

func (groupRepository groupRepository) GetGroupForPrivate(request request.GroupRequest) (response response.GroupResponse) {
	fmt.Println("GetGroupForPrivate")
	return response
}

// CREATE
func (groupRepository groupRepository) CreateGroupForPublic(request request.GroupRequest) (response response.GroupResponse) {
	fmt.Println("CreateGroupForPublic")
	return response
}

func (groupRepository groupRepository) CreateGroupForInternal(request request.GroupRequest) (response response.GroupResponse) {
	fmt.Println("CreateGroupForInternal()")
	return response
}

func (groupRepository groupRepository) CreateGroupForPrivate(request request.GroupRequest) (response response.GroupResponse) {
	fmt.Println("CreateGroupForPrivate()")
	return response
}

// UPDATE
func (groupRepository groupRepository) UpdateGroupForPublic(request request.GroupRequest) (response response.GroupResponse) {
	fmt.Println("UpdateGroupForPublic()")
	return response
}

func (groupRepository groupRepository) UpdateGroupForInternal(request request.GroupRequest) (response response.GroupResponse) {
	fmt.Println("UpdateGroupForInternal")
	return response
}

func (groupRepository groupRepository) UpdateGroupForPrivate(request request.GroupRequest) (response response.GroupResponse) {
	fmt.Println("UpdateGroupForPrivate")
	return response
}

// DELETE
func (groupRepository groupRepository) DeleteGroupForPublic(request request.GroupRequest) (response response.GroupResponse) {
	fmt.Println("DeleteGroupForPublic")
	return response
}

func (groupRepository groupRepository) DeleteGroupForInternal(request request.GroupRequest) (response response.GroupResponse) {
	fmt.Println("DeleteGroupForInternal")
	return response
}

func (groupRepository groupRepository) DeleteGroupForPrivate(request request.GroupRequest) (response response.GroupResponse) {
	fmt.Println("DeleteGroupForPrivate")
	return response
}

func NewGroupRepository(conf config.BaseConfig) GroupRepository {
	return &groupRepository{BaseConfig: conf}
}
