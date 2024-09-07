package config

import (
	"context"
	"io/ioutil"

	"github.com/go-redis/redis/v8"
	yaml "gopkg.in/yaml.v3"
)

type YamlConfig struct {
	Application Application `yaml:"Application"`
	MySQL       MySQL       `yaml:"MySQL"`
	Redis       Redis       `yaml:"Redis"`
}

type Server struct {
	Port  string `yaml:"port"`
	Admin Admin  `yaml:"admin"`
	Jwt   Jwt    `yaml:"jwt"`
	Mail  Mail   `yaml:"mail"`
	Tmp   Tmp    `yaml:"tmp"`
}

type Tmp struct {
	Length  int    `yaml:"length"`
	Letters string `yaml:"letters"`
}

type Mail struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
}

type Jwt struct {
	Secret string `yaml:"Secret"`
}

type Client struct {
	ServerEndpoint string `yaml:"ServerEndpoint"`
	UserEmail      string `yaml:"UserEmail"`
	UserPassword   string `yaml:"UserPassword"`
	HomeDir        string `yaml:"HomeDir"`
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

type RedisConf struct {
	RedisConnection *redis.Client
	Ctx             context.Context
}

type Redis struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
	Db   string `yaml:"db"`
}

type YamlLoader interface {
	LoadYaml(filePath string) (YamlConfig, error)
}

func LoadYaml(filePath string) (YamlConfig, error) {
	var config YamlConfig
	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(buf, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}
