package usecase

import (
	"fmt"

	"github.com/ryo-arima/mark1/pkg/client/repository"
	"github.com/ryo-arima/mark1/pkg/entity/request"
)

type MemberUsecase interface {
	BootstrapMemberForDB(request request.MemberRequest)
	GetMemberForPublic(request request.MemberRequest)
	GetMemberForInternal(request request.MemberRequest)
	GetMemberForPrivate(request request.MemberRequest)
	CreateMemberForPublic(request request.MemberRequest)
	CreateMemberForInternal(request request.MemberRequest)
	CreateMemberForPrivate(request request.MemberRequest)
	UpdateMemberForPublic(request request.MemberRequest)
	UpdateMemberForInternal(request request.MemberRequest)
	UpdateMemberForPrivate(request request.MemberRequest)
	DeleteMemberForPublic(request request.MemberRequest)
	DeleteMemberForInternal(request request.MemberRequest)
	DeleteMemberForPrivate(request request.MemberRequest)
}

type memberUsecase struct {
	MemberRepository   repository.MemberRepository
}

//Bootstrap
func (memberUsecase memberUsecase) BootstrapMemberForDB(request request.MemberRequest) {
	members := memberUsecase.MemberRepository.BootstrapMemberForDB(request)
	fmt.Println(members)
}

//GET
func (memberUsecase memberUsecase) GetMemberForPublic(request request.MemberRequest) {
	members := memberUsecase.MemberRepository.GetMemberForPublic(request)
	fmt.Println(members)
}

func (memberUsecase memberUsecase) GetMemberForInternal(request request.MemberRequest) {
	members := memberUsecase.MemberRepository.GetMemberForInternal(request)
	fmt.Println(members)
}

func (memberUsecase memberUsecase) GetMemberForPrivate(request request.MemberRequest) {
	members := memberUsecase.MemberRepository.GetMemberForPrivate(request)
	fmt.Println(members)
}

//CREATE
func (memberUsecase memberUsecase) CreateMemberForPublic(request request.MemberRequest) {
	members := memberUsecase.MemberRepository.CreateMemberForPublic(request)
	fmt.Println(members)
}

func (memberUsecase memberUsecase) CreateMemberForInternal(request request.MemberRequest) {
	members := memberUsecase.MemberRepository.CreateMemberForInternal(request)
	fmt.Println(members)
}

func (memberUsecase memberUsecase) CreateMemberForPrivate(request request.MemberRequest) {
	members := memberUsecase.MemberRepository.CreateMemberForPrivate(request)
	fmt.Println(members)
}

//UPDATE
func (memberUsecase memberUsecase) UpdateMemberForPublic(request request.MemberRequest) {
	members := memberUsecase.MemberRepository.UpdateMemberForPublic(request)
	fmt.Println(members)
}

func (memberUsecase memberUsecase) UpdateMemberForInternal(request request.MemberRequest) {
	members := memberUsecase.MemberRepository.UpdateMemberForInternal(request)
	fmt.Println(members)
}

func (memberUsecase memberUsecase) UpdateMemberForPrivate(request request.MemberRequest) {
	members := memberUsecase.MemberRepository.UpdateMemberForPrivate(request)
	fmt.Println(members)
}

//DELETE
func (memberUsecase memberUsecase) DeleteMemberForPublic(request request.MemberRequest) {
	members := memberUsecase.MemberRepository.DeleteMemberForPublic(request)
	fmt.Println(members)
}

func (memberUsecase memberUsecase) DeleteMemberForInternal(request request.MemberRequest) {
	members := memberUsecase.MemberRepository.DeleteMemberForInternal(request)
	fmt.Println(members)
}

func (memberUsecase memberUsecase) DeleteMemberForPrivate(request request.MemberRequest) {
	members := memberUsecase.MemberRepository.DeleteMemberForPrivate(request)
	fmt.Println(members)
}

func NewMemberUsecase(memberRepository repository.MemberRepository) MemberUsecase {
	return &memberUsecase{ MemberRepository: memberRepository}
}