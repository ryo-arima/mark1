package config

import (
	"gorm.io/gorm"
)

type BaseConfig struct {
	DBConnection *gorm.DB
	RedisClient  RedisClient
	YamlConfig   YamlConfig
	Token        string
}

func NewBaseConfig(conf YamlConfig, dbConnector DBConnector, redisConnector RedisConnector) (*BaseConfig, error) {
	dbConnection, err := dbConnector.Open(conf)
	if err != nil {
		return nil, err
	}

	redisClient, err := redisConnector.Open(conf)
	if err != nil {
		return nil, err
	}

	token := LoadToken(conf.Application.Client.HomeDir)
	if err != nil {
		return nil, err
	}

	return &BaseConfig{
		DBConnection: dbConnection,
		YamlConfig:   conf,
		RedisClient:  redisClient,
		Token:        token,
	}, nil
}
