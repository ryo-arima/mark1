package controller

import (
	"fmt"
	"log"

	"github.com/ryo-arima/mark1/pkg/client/repository"
	"github.com/ryo-arima/mark1/pkg/client/repository/templates"
	"github.com/ryo-arima/mark1/pkg/client/usecase"
	"github.com/ryo-arima/mark1/pkg/config"
	"github.com/ryo-arima/mark1/pkg/entity/request"
	"github.com/spf13/cobra"
)

func InitBootstrapUserCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	bootstrapUserCmd := &cobra.Command{
		Use:   "user",
		Short: "bootstrap the value of a key",
		Long:  "bootstrap the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			userRepository := repository.NewUserRepository(conf)
			templateRepository := templates.NewTemplateRepository(conf)
			userUsecase := usecase.NewUserUsecase(userRepository, templateRepository)
			userUsecase.BootstrapUserForDB(request.UserRequest{})
		},
	}
	bootstrapUserCmd.Flags().StringP("key", "k", "", "cache key")
	return bootstrapUserCmd
}

func InitCreateUserCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	createUserCmd := &cobra.Command{
		Use:   "user",
		Short: "create the value of a key",
		Long:  "create the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			email, err := cmd.Flags().GetString("email")
			if err != nil {
				log.Fatal(err)
			}
			name, err := cmd.Flags().GetString("name")
			if err != nil {
				log.Fatal(err)
			}
			password, err := cmd.Flags().GetString("password")
			if err != nil {
				log.Fatal(err)
			}
			_request := request.UserRequest{
				User: request.User{
					Email:    email,
					Name:     name,
					Password: password,
				},
			}
			userRepository := repository.NewUserRepository(conf)
			templateRepository := templates.NewTemplateRepository(conf)
			userUsecase := usecase.NewUserUsecase(userRepository, templateRepository)
			userUsecase.CreateUserForPublic(_request)
		},
	}
	createUserCmd.Flags().StringP("email", "", "", "email")
	createUserCmd.Flags().StringP("name", "", "", "name")
	createUserCmd.Flags().StringP("password", "", "", "password")
	return createUserCmd
}

func InitCreateUserCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	createUserCmd := &cobra.Command{
		Use:   "user",
		Short: "create the value of a key",
		Long:  "create the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	createUserCmd.Flags().StringP("key", "k", "", "cache key")
	return createUserCmd
}

func InitGetUserCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	getUserCmd := &cobra.Command{
		Use:   "user",
		Short: "get the value of a key",
		Long:  "get the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	getUserCmd.Flags().StringP("key", "k", "", "cache key")
	return getUserCmd
}

func InitGetUserCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	getUserCmd := &cobra.Command{
		Use:   "users",
		Short: "get the value of a key",
		Long:  "get the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			_id, err := cmd.Flags().GetString("id")
			if err != nil {
				log.Fatal(err)
			}
			_uuid, err := cmd.Flags().GetString("uuid")
			if err != nil {
				log.Fatal(err)
			}
			_email, err := cmd.Flags().GetString("email")
			if err != nil {
				log.Fatal(err)
			}
			_name, err := cmd.Flags().GetString("name")
			if err != nil {
				log.Fatal(err)
			}
			_status, err := cmd.Flags().GetString("status")
			if err != nil {
				log.Fatal(err)
			}
			_isJson, err := cmd.Flags().GetBool("json")
			if err != nil {
				log.Fatal(err)
			}
			_request := request.UserRequest{
				User: request.User{
					ID:     _id,
					UUID:   _uuid,
					Email:  _email,
					Name:   _name,
					Status: _status,
				},
			}
			userRepository := repository.NewUserRepository(conf)
			templateRepository := templates.NewTemplateRepository(conf)
			userUsecase := usecase.NewUserUsecase(userRepository, templateRepository)
			userUsecase.GetUserForPrivate(_request, _isJson)
		},
	}
	getUserCmd.Flags().StringP("id", "", "", "id")
	getUserCmd.Flags().StringP("uuid", "", "", "uuid")
	getUserCmd.Flags().StringP("email", "", "", "email")
	getUserCmd.Flags().StringP("name", "", "", "name")
	getUserCmd.Flags().StringP("status", "", "", "status")
	getUserCmd.Flags().BoolP("json", "", false, "json")
	return getUserCmd
}

func InitUpdateUserCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	updateUserCmd := &cobra.Command{
		Use:   "user",
		Short: "update the value of a key",
		Long:  "udpate the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	updateUserCmd.Flags().StringP("key", "k", "", "cache key")
	return updateUserCmd
}

func InitUpdateUserCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	updateUserCmd := &cobra.Command{
		Use:   "user",
		Short: "update the value of a key",
		Long:  "update the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	updateUserCmd.Flags().StringP("key", "k", "", "cache key")
	return updateUserCmd
}

func InitDeleteUserCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	deleteUserCmd := &cobra.Command{
		Use:   "user",
		Short: "delete the value of a key",
		Long:  "delete the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	deleteUserCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteUserCmd
}

func InitDeleteUserCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	deleteUserCmd := &cobra.Command{
		Use:   "user",
		Short: "delete the value of a key",
		Long:  "delete the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
		},
	}
	deleteUserCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteUserCmd
}
