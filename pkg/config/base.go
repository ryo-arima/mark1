package config

import (
	"context"
	"io/ioutil"
	"strconv"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"fmt"

	yaml "gopkg.in/yaml.v3"
)

type BaseConfig struct {
	DBConnection *gorm.DB
	RedisConf    RedisConf
	YamlConfig   YamlConfig
}

type YamlConfig struct {
	Application Application `yaml:"Application"`
	MySQL       MySQL       `yaml:"MySQL"`
	Redis       Redis       `yaml:"Redis"`
}

type Server struct {
	Port      string `yaml:"port"`
	Admin     Admin  `yaml:"admin"`
	Jwt       Jwt    `yaml:"jwt"`
	Mail      Mail   `yaml:"mail"`
	TmpLength int    `yaml:"tmplength"`
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
	redisConf := NewRedisConf(d1)
	baseConfig := &BaseConfig{
		DBConnection: dbConnection,
		YamlConfig:   d1,
		RedisConf:    redisConf,
	}
	return *baseConfig
}

func NewDBConnection(conf YamlConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.MySQL.User, conf.MySQL.Pass, conf.MySQL.Host, conf.MySQL.Port, conf.MySQL.Db)
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db
}

func NewRedisConf(conf YamlConfig) RedisConf {
	ctx := context.Background()

	db, err := strconv.Atoi(conf.Redis.Db)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		return RedisConf{}
	}

	redisConnection := redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Host + ":" + conf.Redis.Port,
		Username: conf.Redis.User,
		Password: conf.Redis.Pass,
		DB:       db,
	})
	return RedisConf{
		RedisConnection: redisConnection,
		Ctx:             ctx,
	}
}
