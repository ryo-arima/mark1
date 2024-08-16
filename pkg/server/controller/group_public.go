package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/mark1/pkg/entity/request"
	"github.com/ryo-arima/mark1/pkg/entity/response"
	"github.com/ryo-arima/mark1/pkg/server/repository"
)

type GroupControllerForPublic interface {
	GetGroups(c *gin.Context)
}

type groupControllerForPublic struct {
	GroupRepository repository.GroupRepository
}

func (groupController groupControllerForPublic) GetGroups(c *gin.Context) {
	var groupRequest request.GroupRequest
	if err := c.Bind(&groupRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.GroupResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Groups: []response.Group{}})
		return
	}
	res := groupController.GroupRepository.GetGroups()
	c.JSON(http.StatusOK, res)
	return
}


func NewGroupControllerForPublic(groupRepository repository.GroupRepository) GroupControllerForPublic {
	return &groupControllerForPublic{ GroupRepository: groupRepository}
}