package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ryo-arima/mark1/pkg/config"
	"github.com/ryo-arima/mark1/pkg/entity/request"
	"github.com/ryo-arima/mark1/pkg/entity/response"
)

type CommonRepository interface {
	CreateEmailForPublic(request request.UserRequest) (response response.EmailResponse)
	VerifyEmailForPublic(request request.UserRequest) (response response.EmailResponse)
}

type commonRepository struct {
	BaseConfig config.BaseConfig
}

func (commonRepository commonRepository) CreateEmailForPublic(request request.UserRequest) (response response.EmailResponse) {
	URL := commonRepository.BaseConfig.YamlConfig.Application.Client.ServerEndpoint + "/api/public/email"

	jsonData, err := json.Marshal(request)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	resp, err := http.Post(URL, "application/json", bytes.NewBuffer([]byte(jsonData)))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Status Code:", resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Response Body:", string(body))
	return response
}

func (commonRepository commonRepository) VerifyEmailForPublic(request request.UserRequest) (response response.EmailResponse) {
	// GETリクエストを送信するURLを指定
	URL := commonRepository.BaseConfig.YamlConfig.Application.Client.ServerEndpoint + "/api/public/email"

	// GETリクエストを送信
	resp, err := http.Get(URL)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}
	defer resp.Body.Close() // リソースの解放

	// レスポンスボディを読み込む
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// ステータスコードとレスポンスボディを表示
	fmt.Println("Status Code:", resp.StatusCode)
	fmt.Println("Response Body:", string(body))
	return response
}

func NewCommonRepository(conf config.BaseConfig) CommonRepository {
	return &commonRepository{BaseConfig: conf}
}
