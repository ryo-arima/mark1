package repository

import (
	"fmt"
	"io"
	"net/smtp"
	"strings"

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
	mailAddress := "dst mail addr"
	mailSubject := "<Subject>"

	toMailAddress := []string{mailAddress}
	// メールの内容を定義
	mailBody := "テストメールです。"
	msgStr :=
		"To:" + mailAddress + "\r\n" + //送信先
			"reply-to: " + "<src mail addr>" + "\r\n" + //送信元
			"Subject:" + mailSubject + "\r\n" +
			"\r\n" + mailBody
	reader := strings.NewReader(msgStr)
	transformer := japanese.ISO2022JP.NewEncoder()
	msgISO2022JP, err := io.ReadAll(transform.NewReader(reader, transformer))
	if err != nil {
		fmt.Println(err)
	}
	msg := []byte(msgISO2022JP)

	auth := smtp.PlainAuth("", "src mail addr", "src mail key", "smtp.gmail.com")
	if err := smtp.SendMail("smtp.gmail.com:587", auth, "src mail addr", toMailAddress, msg); err != nil {
		fmt.Println(err)
	}
}

func NewCommonRepository(conf config.BaseConfig) CommonRepository {
	return &commonRepository{BaseConfig: conf}
}
