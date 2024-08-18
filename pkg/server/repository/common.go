package repository

import (
	"bytes"
	"fmt"
	"io"
	"net/smtp"
	"strings"
	"text/template"
	"time"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"

	"github.com/ryo-arima/mark1/pkg/config"
	"github.com/ryo-arima/mark1/pkg/entity/model"
	"github.com/ryo-arima/mark1/pkg/server/middleware"
)

type CommonRepository interface {
	CreateEmail(email model.Email)
	SetTempCode(email model.Email) string
	GetTempCode(email model.Email) string
}

type commonRepository struct {
	BaseConfig config.BaseConfig
}

func (commonRepository commonRepository) CreateEmail(email model.Email) {
	tmpl, err := template.ParseFiles("pkg/server/template/verify_email.tpl")
	if err != nil {
		fmt.Println("Error parsing template file:", err)
		return
	}
	email.From = commonRepository.BaseConfig.YamlConfig.Application.Server.Mail.User

	var emailContent bytes.Buffer
	err = tmpl.Execute(&emailContent, email)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}

	toMailAddress := []string{email.To}
	fmt.Println(emailContent.String())
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

func (commonRepository commonRepository) SetTempCode(email model.Email) string {
	length := commonRepository.BaseConfig.YamlConfig.Application.Server.TmpLength
	tempCode := middleware.GenTempCode(length)
	context := commonRepository.BaseConfig.RedisConf.Ctx
	commonRepository.BaseConfig.RedisConf.RedisConnection.Set(context, email.To, tempCode, time.Hour)
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
