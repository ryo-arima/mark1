package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/mark1/pkg/entity/model"
	"github.com/ryo-arima/mark1/pkg/entity/request"
	"github.com/ryo-arima/mark1/pkg/entity/response"
	"github.com/ryo-arima/mark1/pkg/server/repository"
)

type GroupControllerForInternal interface {
	GetGroups(c *gin.Context)
	CreateGroup(c *gin.Context)
	UpdateGroup(c *gin.Context)
	DeleteGroup(c *gin.Context)
}

type groupControllerForInternal struct {
	GroupRepository repository.GroupRepository
}

func (groupController groupControllerForInternal) GetGroups(c *gin.Context) {
	var groupRequest request.GroupRequest
	if err := c.Bind(&groupRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.GroupResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Groups: []response.Group{}})
		return
	}
	res := groupController.GroupRepository.GetGroups()
	c.JSON(http.StatusOK, res)
	return
}


func (groupController groupControllerForInternal) CreateGroup(c *gin.Context) {
	var groupRequest request.GroupRequest
	if err := c.Bind(&groupRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.GroupResponse{Code: "SERVER_CONTROLLER_CREATE__FOR__001", Message: err.Error(), Groups: []response.Group{}})
		return
	}
	var groupModel model.Groups
	res := groupController.GroupRepository.CreateGroup(groupModel)
	c.JSON(http.StatusOK, res)
	return
}


func (groupController groupControllerForInternal) UpdateGroup(c *gin.Context) {
	var groupRequest request.GroupRequest
	if err := c.Bind(&groupRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.GroupResponse{Code: "SERVER_CONTROLLER_UPDATE__FOR__001", Message: err.Error(), Groups: []response.Group{}})
		return
	}
	var groupModel model.Groups
	res := groupController.GroupRepository.UpdateGroup(groupModel)
	c.JSON(http.StatusOK, res)
	return
}


func (groupController groupControllerForInternal) DeleteGroup(c *gin.Context) {
	var groupRequest request.GroupRequest
	if err := c.Bind(&groupRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.GroupResponse{Code: "SERVER_CONTROLLER_DELETE__FOR__001", Message: err.Error(), Groups: []response.Group{}})
		return
	}
	var uuid string
	res := groupController.GroupRepository.DeleteGroup(uuid)
	c.JSON(http.StatusOK, res)
	return
}


func NewGroupControllerForInternal(groupRepository repository.GroupRepository) GroupControllerForInternal {
	return &groupControllerForInternal{ GroupRepository: groupRepository}
}