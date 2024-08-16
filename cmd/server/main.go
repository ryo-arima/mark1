package main

import (
	"github.com/ryo-arima/mark1/pkg/config"
	"github.com/ryo-arima/mark1/pkg/server"
)

func main() {
	conf := config.NewBaseConfig()
	server.Main(conf)
}