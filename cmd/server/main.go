package main

import (
	"os"

	"github.com/ryo-arima/mark1/pkg/config"
	"github.com/ryo-arima/mark1/pkg/server"
)

func main() {
	yamlConfig, err := config.LoadYaml(os.Getenv("MARK1_CONFIG_PATH"))
	if err != nil {
		panic(err)
	}

	dbConnector := config.NewDBConnector()
	redisConnector := config.NewRedisConnector()

	conf, err := config.NewBaseConfig(yamlConfig, dbConnector, redisConnector)
	if err != nil {
		panic(err)
	}
	server.Main(*conf)
}
