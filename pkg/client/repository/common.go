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
	CreateTokenForPublic() (response response.TokenResponse)
	SaveTokenAtHomeDir(token string)
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

func (commonRepository commonRepository) CreateTokenForPublic() (response response.TokenResponse) {
	serverEndpoint := commonRepository.BaseConfig.YamlConfig.Application.Client.ServerEndpoint
	userEmail := commonRepository.BaseConfig.YamlConfig.Application.Client.UserEmail
	userPassword := commonRepository.BaseConfig.YamlConfig.Application.Client.UserPassword

	URL := serverEndpoint + "/api/public/token"

	requestData := request.UserRequest{
		User: request.User{
			Email:    userEmail,
			Password: userPassword,
		},
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error unmarshalling response:", err)
		return
	}

	return response
}

func (commonRepository commonRepository) SaveTokenAtHomeDir(token string) {
	// ファイルにトークンを保存
	err := ioutil.WriteFile(commonRepository.BaseConfig.YamlConfig.Application.Client.HomeDir+"/token", []byte(token), 0644)
	if err != nil {
		fmt.Println("Error writing token to file:", err)
	}
}

func NewCommonRepository(conf config.BaseConfig) CommonRepository {
	return &commonRepository{BaseConfig: conf}
}
