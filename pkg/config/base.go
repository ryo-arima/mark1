package config

import (
	"io/ioutil"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"fmt"

	yaml "gopkg.in/yaml.v3"
)

type BaseConfig struct {
	DBConnection *gorm.DB
	YamlConfig   YamlConfig
}

type YamlConfig struct {
	Application Application `yaml:"Application"`
	MySQL       MySQL       `yaml:"MySQL"`
}

type Server struct {
	Admin Admin `yaml:"admin"`
}

type Client struct {
	ServerEndpoint string `yaml:"ServerEndpoint"`
	UserEmail      string `yaml:"UserEmail"`
	UserPassword   string `yaml:"UserPassword"`
}

type Application struct {
	Server Server `yaml:"Server"`
	Client Client `yaml:"Client"`
}

type Admin struct {
	Emails []string `yaml:"emails"`
}

type MySQL struct {
	Host string `yaml:"host"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
	Port string `yaml:"port"`
	Db   string `yaml:"db"`
}

func NewBaseConfig() BaseConfig {
	buf1, err := ioutil.ReadFile("etc/main.yaml")
	if err != nil {
		panic(err)
	}

	var d1 YamlConfig
	err = yaml.Unmarshal(buf1, &d1)
	if err != nil {
		panic(err)
	}

	dbConnection := NewDBConnection(d1)
	baseConfig := &BaseConfig{
		DBConnection: dbConnection,
		YamlConfig:   d1,
	}
	return *baseConfig
}

func NewDBConnection(conf YamlConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.MySQL.User, conf.MySQL.Pass, conf.MySQL.Host, conf.MySQL.Port, conf.MySQL.Db)
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db
}
