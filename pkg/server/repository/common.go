package repository

import (
	"bytes"
	"fmt"
	"io"
	"net/smtp"
	"strings"
	"text/template"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"

	"github.com/google/uuid"
	"github.com/ryo-arima/mark1/pkg/config"
	"github.com/ryo-arima/mark1/pkg/entity/model"
	"github.com/ryo-arima/mark1/pkg/server/middleware"
)

type CommonRepository interface {
	CreateEmail(email model.Email, tempCode string)
	CreateJwtToken(user model.Users) string
	SetTempCode(email model.Email) string
	GetTempCode(email model.Email) string
}

type commonRepository struct {
	BaseConfig config.BaseConfig
}

func (commonRepository commonRepository) CreateEmail(email model.Email, tempCode string) {
	tmpl, err := template.ParseFiles("pkg/server/template/verify_email.tpl")
	if err != nil {
		fmt.Println("Error parsing template file:", err)
		return
	}
	var userModel model.Users
	commonRepository.BaseConfig.DBConnection.Where("email = ?", email.To).First(&userModel)
	email.From = commonRepository.BaseConfig.YamlConfig.Application.Server.Mail.User
	email.VeryfyEmailURL = commonRepository.BaseConfig.YamlConfig.Application.Client.ServerEndpoint + "/api/public/email?code=" + tempCode + "&uuid=" + userModel.UUID

	var emailContent bytes.Buffer
	err = tmpl.Execute(&emailContent, email)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}

	toMailAddress := []string{email.To}
	reader := strings.NewReader(emailContent.String())
	transformer := japanese.ISO2022JP.NewEncoder()
	msgISO2022JP, err := io.ReadAll(transform.NewReader(reader, transformer))
	if err != nil {
		fmt.Println(err)
	}
	msg := []byte(msgISO2022JP)
	user := commonRepository.BaseConfig.YamlConfig.Application.Server.Mail.User
	pass := commonRepository.BaseConfig.YamlConfig.Application.Server.Mail.Pass
	host := commonRepository.BaseConfig.YamlConfig.Application.Server.Mail.Host
	port := commonRepository.BaseConfig.YamlConfig.Application.Server.Mail.Port
	auth := smtp.PlainAuth("", user, pass, host)
	if err := smtp.SendMail(host+":"+port, auth, user, toMailAddress, msg); err != nil {
		fmt.Println(err)
	}
}

func (commonRepository commonRepository) CreateJwtToken(user model.Users) string {
	var jwtKey = []byte(commonRepository.BaseConfig.YamlConfig.Application.Server.Jwt.Secret)

	expirationTime := time.Now().Add(1 * time.Minute)
	jti := uuid.New().String()
	claims := &middleware.Claims{
		UserEmail: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "mark1",
			Subject:   "login",
			Audience:  []string{"audience"},
			NotBefore: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        user.UUID + ":" + jti,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		fmt.Println(err)
	}
	return tokenString
}

func (commonRepository commonRepository) SetTempCode(email model.Email) string {
	length := commonRepository.BaseConfig.YamlConfig.Application.Server.Tmp.Length
	letters := commonRepository.BaseConfig.YamlConfig.Application.Server.Tmp.Letters
	tempCode := middleware.GenTempCode(length, letters)
	context := commonRepository.BaseConfig.RedisConf.Ctx
	err := commonRepository.BaseConfig.RedisConf.RedisConnection.Set(context, email.To, tempCode, time.Hour).Err()
	if err != nil {
		fmt.Println("Error setting key:", err)
	}
	return tempCode
}

func (commonRepository commonRepository) GetTempCode(email model.Email) string {
	context := commonRepository.BaseConfig.RedisConf.Ctx
	tempCode, err := commonRepository.BaseConfig.RedisConf.RedisConnection.Get(context, email.To).Result()
	if err != nil {
		fmt.Println(err)
	}
	return tempCode
}

func NewCommonRepository(conf config.BaseConfig) CommonRepository {
	return &commonRepository{BaseConfig: conf}
}
