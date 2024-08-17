package repository

import (
	"fmt"

	"github.com/ryo-arima/mark1/pkg/config"
	"github.com/ryo-arima/mark1/pkg/entity/request"
	"github.com/ryo-arima/mark1/pkg/entity/response"
)

type MemberRepository interface {
	BootstrapMemberForDB(request request.MemberRequest) (response response.MemberResponse)
	GetMemberForPublic(request request.MemberRequest) (response response.MemberResponse)
	GetMemberForInternal(request request.MemberRequest) (response response.MemberResponse)
	GetMemberForPrivate(request request.MemberRequest) (response response.MemberResponse)
	CreateMemberForInternal(request request.MemberRequest) (response response.MemberResponse)
	CreateMemberForPrivate(request request.MemberRequest) (response response.MemberResponse)
	UpdateMemberForInternal(request request.MemberRequest) (response response.MemberResponse)
	UpdateMemberForPrivate(request request.MemberRequest) (response response.MemberResponse)
	DeleteMemberForInternal(request request.MemberRequest) (response response.MemberResponse)
	DeleteMemberForPrivate(request request.MemberRequest) (response response.MemberResponse)
}

type memberRepository struct {
	BaseConfig config.BaseConfig
}

// Bootstrap
func (memberRepository memberRepository) BootstrapMemberForDB(request request.MemberRequest) (response response.MemberResponse) {
	fmt.Println("BootstrapMemberForDB")
	return response
}

// GET
func (memberRepository memberRepository) GetMemberForPublic(request request.MemberRequest) (response response.MemberResponse) {
	fmt.Println("GetMemberForPublic")
	return response
}

func (memberRepository memberRepository) GetMemberForInternal(request request.MemberRequest) (response response.MemberResponse) {
	fmt.Println("GetMemberForInternal")
	return response
}

func (memberRepository memberRepository) GetMemberForPrivate(request request.MemberRequest) (response response.MemberResponse) {
	fmt.Println("GetMemberForPrivate")
	return response
}

// CREATE
func (memberRepository memberRepository) CreateMemberForInternal(request request.MemberRequest) (response response.MemberResponse) {
	fmt.Println("CreateMemberForInternal()")
	return response
}

func (memberRepository memberRepository) CreateMemberForPrivate(request request.MemberRequest) (response response.MemberResponse) {
	fmt.Println("CreateMemberForPrivate()")
	return response
}

// UPDATE
func (memberRepository memberRepository) UpdateMemberForInternal(request request.MemberRequest) (response response.MemberResponse) {
	fmt.Println("UpdateMemberForInternal")
	return response
}

func (memberRepository memberRepository) UpdateMemberForPrivate(request request.MemberRequest) (response response.MemberResponse) {
	fmt.Println("UpdateMemberForPrivate")
	return response
}

// DELETE
func (memberRepository memberRepository) DeleteMemberForInternal(request request.MemberRequest) (response response.MemberResponse) {
	fmt.Println("DeleteMemberForInternal")
	return response
}

func (memberRepository memberRepository) DeleteMemberForPrivate(request request.MemberRequest) (response response.MemberResponse) {
	fmt.Println("DeleteMemberForPrivate")
	return response
}

func NewMemberRepository(conf config.BaseConfig) MemberRepository {
	return &memberRepository{BaseConfig: conf}
}
