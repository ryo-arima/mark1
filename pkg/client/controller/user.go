package controller

import (
	"fmt"
	"log"

	"github.com/ryo-arima/mark1/pkg/client/repository"
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
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
			userRepository := repository.NewUserRepository(conf)
			userUsecase := usecase.NewUserUsecase(userRepository)
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
