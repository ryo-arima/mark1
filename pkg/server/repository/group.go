package repository

import (
	"time"

	"github.com/ryo-arima/mark1/pkg/config"
	"github.com/ryo-arima/mark1/pkg/entity/model"
	"gorm.io/gorm"
)

type GroupRepository interface {
	GetGroups() []model.Groups
	CreateGroup(group model.Groups) *gorm.DB
	UpdateGroup(group model.Groups) *gorm.DB
	DeleteGroup(uuid string) *gorm.DB
}

type groupRepository struct {
	BaseConfig config.BaseConfig
}


func (groupRepository groupRepository) GetGroups() []model.Groups {
	var groups []model.Groups
	groupRepository.BaseConfig.DBConnection.Find(&groups)
	return groups
}

func (groupRepository groupRepository) CreateGroup(group model.Groups) *gorm.DB {
	results := groupRepository.BaseConfig.DBConnection.Create(&group)
	return results
}

func (groupRepository groupRepository) UpdateGroup(group model.Groups) *gorm.DB {
	results := groupRepository.BaseConfig.DBConnection.Model(model.Groups{}).Where("uuid = ?", group.UUID).Updates(group)
	return results
}

func (groupRepository groupRepository) DeleteGroup(uuid string) *gorm.DB {
	results := groupRepository.BaseConfig.DBConnection.Model(model.Groups{}).Where("uuid = ?", uuid).Update("deleted_at", time.Now())
	return results
}

func NewGroupRepository(conf config.BaseConfig) GroupRepository {
	return &groupRepository{BaseConfig: conf}
}