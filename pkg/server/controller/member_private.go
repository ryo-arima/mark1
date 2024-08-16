package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/mark1/pkg/entity/model"
	"github.com/ryo-arima/mark1/pkg/entity/request"
	"github.com/ryo-arima/mark1/pkg/entity/response"
	"github.com/ryo-arima/mark1/pkg/server/repository"
)

type MemberControllerForPrivate interface {
	GetMembers(c *gin.Context)
	CreateMember(c *gin.Context)
	UpdateMember(c *gin.Context)
	DeleteMember(c *gin.Context)
}

type memberControllerForPrivate struct {
	MemberRepository repository.MemberRepository
}

func (memberController memberControllerForPrivate) GetMembers(c *gin.Context) {
	var memberRequest request.MemberRequest
	if err := c.Bind(&memberRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.MemberResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Members: []response.Member{}})
		return
	}
	res := memberController.MemberRepository.GetMembers()
	c.JSON(http.StatusOK, res)
	return
}


func (memberController memberControllerForPrivate) CreateMember(c *gin.Context) {
	var memberRequest request.MemberRequest
	if err := c.Bind(&memberRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.MemberResponse{Code: "SERVER_CONTROLLER_CREATE__FOR__001", Message: err.Error(), Members: []response.Member{}})
		return
	}
	var memberModel model.Members
	res := memberController.MemberRepository.CreateMember(memberModel)
	c.JSON(http.StatusOK, res)
	return
}


func (memberController memberControllerForPrivate) UpdateMember(c *gin.Context) {
	var memberRequest request.MemberRequest
	if err := c.Bind(&memberRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.MemberResponse{Code: "SERVER_CONTROLLER_UPDATE__FOR__001", Message: err.Error(), Members: []response.Member{}})
		return
	}
	var memberModel model.Members
	res := memberController.MemberRepository.UpdateMember(memberModel)
	c.JSON(http.StatusOK, res)
	return
}


func (memberController memberControllerForPrivate) DeleteMember(c *gin.Context) {
	var memberRequest request.MemberRequest
	if err := c.Bind(&memberRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.MemberResponse{Code: "SERVER_CONTROLLER_DELETE__FOR__001", Message: err.Error(), Members: []response.Member{}})
		return
	}
	var uuid string
	res := memberController.MemberRepository.DeleteMember(uuid)
	c.JSON(http.StatusOK, res)
	return
}


func NewMemberControllerForPrivate(memberRepository repository.MemberRepository) MemberControllerForPrivate {
	return &memberControllerForPrivate{ MemberRepository: memberRepository}
}