package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/mark1/pkg/entity/request"
	"github.com/ryo-arima/mark1/pkg/entity/response"
	"github.com/ryo-arima/mark1/pkg/server/repository"
)

type MemberControllerForPublic interface {
	GetMembers(c *gin.Context)
}

type memberControllerForPublic struct {
	MemberRepository repository.MemberRepository
}

func (memberController memberControllerForPublic) GetMembers(c *gin.Context) {
	var memberRequest request.MemberRequest
	if err := c.Bind(&memberRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.MemberResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Members: []response.Member{}})
		return
	}
	res := memberController.MemberRepository.GetMembers()
	c.JSON(http.StatusOK, res)
	return
}


func NewMemberControllerForPublic(memberRepository repository.MemberRepository) MemberControllerForPublic {
	return &memberControllerForPublic{ MemberRepository: memberRepository}
}