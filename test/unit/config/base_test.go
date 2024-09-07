package config

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ryo-arima/mark1/pkg/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MockDBConnector struct {
	mock.Mock
}

func (m *MockDBConnector) Open(conf config.YamlConfig) (*gorm.DB, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		fmt.Println("Error creating mock database")
		return nil, err
	}

	mock.ExpectQuery("SELECT VERSION()").WillReturnRows(sqlmock.NewRows([]string{"version"}).AddRow("8.0.23"))

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		fmt.Println("Error creating GORM database")
		return nil, err
	}

	return gormDB, nil
}

type MockRedisConnector struct {
	mock.Mock
}

func (m *MockRedisConnector) Open(conf config.YamlConfig) (config.RedisClient, error) {
	args := m.Called(conf)
	return args.Get(0).(config.RedisClient), args.Error(1)
}

func TestNewBaseConfig(t *testing.T) {
	mockDBConnector := new(MockDBConnector)
	mockRedisConnector := new(MockRedisConnector)
	conf, _ := config.LoadYaml("../../etc/test_main.yaml")

	mockDBConnector.On("Open", conf).Return(&gorm.DB{}, nil)
	mockRedisConnector.On("Open", conf).Return(config.RedisClient{}, nil)

	baseConfig, err := config.NewBaseConfig(conf, mockDBConnector, mockRedisConnector)

	assert.NoError(t, err)
	assert.NotNil(t, baseConfig)
	assert.NotNil(t, baseConfig.DBConnection)
	assert.NotNil(t, baseConfig.RedisClient)
}
