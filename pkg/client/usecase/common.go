package usecase

import (
	"github.com/ryo-arima/mark1/pkg/client/repository"
	"github.com/ryo-arima/mark1/pkg/entity/request"
)

type CommonUsecase interface {
	CreateEmailForPublic(request request.UserRequest)
	VerifyEmailForPublic(request request.UserRequest)
	CreateTokenForPublic()
}

type commonUsecase struct {
	CommonRepository repository.CommonRepository
}

func (commonUsecase commonUsecase) CreateEmailForPublic(request request.UserRequest) {
	commonUsecase.CommonRepository.CreateEmailForPublic(request)
}

func (commonUsecase commonUsecase) VerifyEmailForPublic(request request.UserRequest) {
	commonUsecase.CommonRepository.VerifyEmailForPublic(request)
}

func (commonUsecase commonUsecase) CreateTokenForPublic() {
	response := commonUsecase.CommonRepository.CreateTokenForPublic()
	//取得したトークンをHomeDir/tokenに保存
	commonUsecase.CommonRepository.SaveTokenAtHomeDir(response.Token)
}

func NewCommonUsecase(commonRepository repository.CommonRepository) CommonUsecase {
	return &commonUsecase{CommonRepository: commonRepository}
}
