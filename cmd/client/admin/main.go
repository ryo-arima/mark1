package main

import (
	"github.com/ryo-arima/mark1/pkg/client"
	"github.com/ryo-arima/mark1/pkg/config"
)

func main() {
	conf := config.NewBaseConfig()
	client.ClientForAdminUser(conf)
}