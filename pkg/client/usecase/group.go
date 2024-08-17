package usecase

import (
	"fmt"

	"github.com/ryo-arima/mark1/pkg/client/repository"
	"github.com/ryo-arima/mark1/pkg/entity/request"
)

type GroupUsecase interface {
	BootstrapGroupForDB(request request.GroupRequest)
	GetGroupForPublic(request request.GroupRequest)
	GetGroupForInternal(request request.GroupRequest)
	GetGroupForPrivate(request request.GroupRequest)
	CreateGroupForInternal(request request.GroupRequest)
	CreateGroupForPrivate(request request.GroupRequest)
	UpdateGroupForInternal(request request.GroupRequest)
	UpdateGroupForPrivate(request request.GroupRequest)
	DeleteGroupForInternal(request request.GroupRequest)
	DeleteGroupForPrivate(request request.GroupRequest)
}

type groupUsecase struct {
	GroupRepository repository.GroupRepository
}

// Bootstrap
func (groupUsecase groupUsecase) BootstrapGroupForDB(request request.GroupRequest) {
	groups := groupUsecase.GroupRepository.BootstrapGroupForDB(request)
	fmt.Println(groups)
}

// GET
func (groupUsecase groupUsecase) GetGroupForPublic(request request.GroupRequest) {
	groups := groupUsecase.GroupRepository.GetGroupForPublic(request)
	fmt.Println(groups)
}

func (groupUsecase groupUsecase) GetGroupForInternal(request request.GroupRequest) {
	groups := groupUsecase.GroupRepository.GetGroupForInternal(request)
	fmt.Println(groups)
}

func (groupUsecase groupUsecase) GetGroupForPrivate(request request.GroupRequest) {
	groups := groupUsecase.GroupRepository.GetGroupForPrivate(request)
	fmt.Println(groups)
}

// CREATE
func (groupUsecase groupUsecase) CreateGroupForInternal(request request.GroupRequest) {
	groups := groupUsecase.GroupRepository.CreateGroupForInternal(request)
	fmt.Println(groups)
}

func (groupUsecase groupUsecase) CreateGroupForPrivate(request request.GroupRequest) {
	groups := groupUsecase.GroupRepository.CreateGroupForPrivate(request)
	fmt.Println(groups)
}

// UPDATE
func (groupUsecase groupUsecase) UpdateGroupForInternal(request request.GroupRequest) {
	groups := groupUsecase.GroupRepository.UpdateGroupForInternal(request)
	fmt.Println(groups)
}

func (groupUsecase groupUsecase) UpdateGroupForPrivate(request request.GroupRequest) {
	groups := groupUsecase.GroupRepository.UpdateGroupForPrivate(request)
	fmt.Println(groups)
}

// DELETE
func (groupUsecase groupUsecase) DeleteGroupForInternal(request request.GroupRequest) {
	groups := groupUsecase.GroupRepository.DeleteGroupForInternal(request)
	fmt.Println(groups)
}

func (groupUsecase groupUsecase) DeleteGroupForPrivate(request request.GroupRequest) {
	groups := groupUsecase.GroupRepository.DeleteGroupForPrivate(request)
	fmt.Println(groups)
}

func NewGroupUsecase(groupRepository repository.GroupRepository) GroupUsecase {
	return &groupUsecase{GroupRepository: groupRepository}
}
