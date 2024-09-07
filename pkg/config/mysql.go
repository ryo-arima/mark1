package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConnector interface {
	Open(conf YamlConfig) (*gorm.DB, error)
}
type dbConnector struct{}

func (dbConnector dbConnector) Open(conf YamlConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.MySQL.User, conf.MySQL.Pass, conf.MySQL.Host, conf.MySQL.Port, conf.MySQL.Db)
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, nil
}

func NewDBConnector() dbConnector {
	return dbConnector{}
}
