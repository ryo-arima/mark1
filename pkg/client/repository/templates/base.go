package templates

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/ryo-arima/mark1/pkg/config"
	"github.com/ryo-arima/mark1/pkg/entity/response"
)

type TemplateRepository interface {
	OutputUsers(users []response.User)
	OutputUsersJson(users []response.User)
}

type templateRepository struct {
	BaseConfig config.BaseConfig
}

func (templateRepository templateRepository) OutputUsers(users []response.User) {
	templateDir := templateRepository.BaseConfig.YamlConfig.Application.Client.TemplatesDir
	t, err := template.ParseFiles(templateDir + "users.tpl")
	if err != nil {
		log.Fatalf("Error parsing template file: %v", err)
	}
	data := map[string]interface{}{
		"Users": users,
	}
	err = t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}

func (templateRepository templateRepository) OutputUsersJson(users []response.User) {
	jsonData, err := json.MarshalIndent(users, "", "  ") // インデント付きのJSON出力
	if err != nil {
		log.Fatalf("Error marshalling to JSON: %v", err)
	}
	fmt.Println(string(jsonData))
}

func NewTemplateRepository(baseConfig config.BaseConfig) TemplateRepository {
	return templateRepository{BaseConfig: baseConfig}
}
