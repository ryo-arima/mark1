package repository

import (
	"bytes"
	"fmt"
	"io"
	"net/smtp"
	"strings"
	"text/template"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"

	"github.com/ryo-arima/mark1/pkg/config"
	"github.com/ryo-arima/mark1/pkg/entity/model"
)

type CommonRepository interface {
	CreateEmail(emal model.Email)
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
	fmt.Println(msg)
	user := commonRepository.BaseConfig.YamlConfig.Application.Server.Mail.User
	pass := commonRepository.BaseConfig.YamlConfig.Application.Server.Mail.Pass
	host := commonRepository.BaseConfig.YamlConfig.Application.Server.Mail.Host
	port := commonRepository.BaseConfig.YamlConfig.Application.Server.Mail.Port
	fmt.Println(user, pass, host, port)
	auth := smtp.PlainAuth("", user, pass, host)
	fmt.Println("--------------")
	fmt.Println(auth)
	if err := smtp.SendMail(host+":"+port, auth, user, toMailAddress, msg); err != nil {
		fmt.Println(err)
	}
}

func NewCommonRepository(conf config.BaseConfig) CommonRepository {
	return &commonRepository{BaseConfig: conf}
}
