package middleware

import (
	"crypto/rand"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/mark1/pkg/config"
	"golang.org/x/crypto/bcrypt"
)

func ForPublic(conf config.BaseConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func ForInternal(conf config.BaseConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func ForPrivate(conf config.BaseConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func GenHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CheckHash(password string, expectedHash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(expectedHash), []byte(password))
	if err != nil {
		return false, nil
	}
	return true, nil
}

func GenTempCode(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return ""
	}

	var result string
	for _, v := range b {
		result += string(letters[int(v)%len(letters)])
	}
	return result
}
