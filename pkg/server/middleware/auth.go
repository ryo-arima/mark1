package middleware

import (
	"crypto/rand"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ryo-arima/mark1/pkg/config"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	UserEmail string `json:"user_email"`
	jwt.RegisteredClaims
}

func ForPublic(conf config.BaseConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func ForInternal(conf config.BaseConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(c.GetHeader("Authorization"), claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(conf.YamlConfig.Application.Server.Jwt.Secret), nil
		})
		if err != nil {
			fmt.Println(err)
		}

		if !token.Valid {
			fmt.Println("invalid token")
			c.JSON(401, gin.H{"message": "invalid token"})
		}
	}
}

func ForPrivate(conf config.BaseConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(c.GetHeader("Authorization"), claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(conf.YamlConfig.Application.Server.Jwt.Secret), nil
		})
		if err != nil {
			fmt.Println(err)
		}

		if !token.Valid {
			fmt.Println("invalid token")
			c.JSON(401, gin.H{"message": "invalid token"})
		}
		isAdmin := false
		for _, email := range conf.YamlConfig.Application.Server.Admin.Emails {
			if email == claims.UserEmail {
				isAdmin = true
			}
		}
		if !isAdmin {
			c.JSON(401, gin.H{"message": "invalid token"})
		}
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

func GenTempCode(n int, letters string) string {
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
