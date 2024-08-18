package controller

import (
	"log"

	"github.com/ryo-arima/mark1/pkg/client/repository"
	"github.com/ryo-arima/mark1/pkg/client/usecase"
	"github.com/ryo-arima/mark1/pkg/config"
	"github.com/ryo-arima/mark1/pkg/entity/request"
	"github.com/spf13/cobra"
)

func InitCreateEmailCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	createTokenCmd := &cobra.Command{
		Use:   "email",
		Short: "create the value of a key",
		Long:  "create the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			email, err := cmd.Flags().GetString("email")
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
					Password: password,
				},
			}
			commonRepository := repository.NewCommonRepository(conf)
			commonUsecase := usecase.NewCommonUsecase(commonRepository)
			commonUsecase.CreateEmailForPublic(_request)
		},
	}
	createTokenCmd.Flags().StringP("email", "", "", "email")
	createTokenCmd.Flags().StringP("password", "", "", "password")
	return createTokenCmd
}

func InitCreateTokenCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	createTokenCmd := &cobra.Command{
		Use:   "token",
		Short: "create the value of a key",
		Long:  "create the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	createTokenCmd.Flags().StringP("email", "", "", "email")
	createTokenCmd.Flags().StringP("password", "", "", "password")
	return createTokenCmd
}

func InitVerifyEmailCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	verifyEmailCmd := &cobra.Command{
		Use:   "email",
		Short: "verify the value of a key",
		Long:  "verify the value of a key",
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
			commonRepository := repository.NewCommonRepository(conf)
			commonUsecase := usecase.NewCommonUsecase(commonRepository)
			commonUsecase.VerifyEmailForPublic(_request)
		},
	}
	verifyEmailCmd.Flags().StringP("email", "", "", "email")
	verifyEmailCmd.Flags().StringP("token", "", "", "token")
	return verifyEmailCmd
}
