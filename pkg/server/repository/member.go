package repository

import (
	"time"

	"github.com/ryo-arima/mark1/pkg/config"
	"github.com/ryo-arima/mark1/pkg/entity/model"
	"gorm.io/gorm"
)

type MemberRepository interface {
	GetMembers() []model.Members
	CreateMember(member model.Members) *gorm.DB
	UpdateMember(member model.Members) *gorm.DB
	DeleteMember(uuid string) *gorm.DB
}

type memberRepository struct {
	BaseConfig config.BaseConfig
}


func (memberRepository memberRepository) GetMembers() []model.Members {
	var members []model.Members
	memberRepository.BaseConfig.DBConnection.Find(&members)
	return members
}

func (memberRepository memberRepository) CreateMember(member model.Members) *gorm.DB {
	results := memberRepository.BaseConfig.DBConnection.Create(&member)
	return results
}

func (memberRepository memberRepository) UpdateMember(member model.Members) *gorm.DB {
	results := memberRepository.BaseConfig.DBConnection.Model(model.Members{}).Where("uuid = ?", member.UUID).Updates(member)
	return results
}

func (memberRepository memberRepository) DeleteMember(uuid string) *gorm.DB {
	results := memberRepository.BaseConfig.DBConnection.Model(model.Members{}).Where("uuid = ?", uuid).Update("deleted_at", time.Now())
	return results
}

func NewMemberRepository(conf config.BaseConfig) MemberRepository {
	return &memberRepository{BaseConfig: conf}
}