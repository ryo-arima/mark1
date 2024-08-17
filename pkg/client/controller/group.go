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

func InitBootstrapGroupCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	bootstrapGroupCmd := &cobra.Command{
		Use:   "group",
		Short: "bootstrap the value of a key",
		Long:  "bootstrap the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
			groupRepository := repository.NewGroupRepository(conf)
			groupUsecase := usecase.NewGroupUsecase(groupRepository)
			groupUsecase.BootstrapGroupForDB(request.GroupRequest{})
		},
	}
	bootstrapGroupCmd.Flags().StringP("key", "k", "", "cache key")
	return bootstrapGroupCmd
}

func InitCreateGroupCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	createGroupCmd := &cobra.Command{
		Use:   "group",
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
	createGroupCmd.Flags().StringP("key", "k", "", "cache key")
	return createGroupCmd
}

func InitCreateGroupCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	createGroupCmd := &cobra.Command{
		Use:   "group",
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
	createGroupCmd.Flags().StringP("key", "k", "", "cache key")
	return createGroupCmd
}

func InitCreateGroupCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	createGroupCmd := &cobra.Command{
		Use:   "group",
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
	createGroupCmd.Flags().StringP("key", "k", "", "cache key")
	return createGroupCmd
}

func InitGetGroupCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	getGroupCmd := &cobra.Command{
		Use:   "group",
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
	getGroupCmd.Flags().StringP("key", "k", "", "cache key")
	return getGroupCmd
}

func InitGetGroupCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	getGroupCmd := &cobra.Command{
		Use:   "group",
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
	getGroupCmd.Flags().StringP("key", "k", "", "cache key")
	return getGroupCmd
}

func InitGetGroupCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	getGroupCmd := &cobra.Command{
		Use:   "group",
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
	getGroupCmd.Flags().StringP("key", "k", "", "cache key")
	return getGroupCmd
}

func InitUpdateGroupCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	updateGroupCmd := &cobra.Command{
		Use:   "group",
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
	updateGroupCmd.Flags().StringP("key", "k", "", "cache key")
	return updateGroupCmd
}

func InitUpdateGroupCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	updateGroupCmd := &cobra.Command{
		Use:   "group",
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
	updateGroupCmd.Flags().StringP("key", "k", "", "cache key")
	return updateGroupCmd
}

func InitUpdateGroupCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	updateGroupCmd := &cobra.Command{
		Use:   "group",
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
	updateGroupCmd.Flags().StringP("key", "k", "", "cache key")
	return updateGroupCmd
}

func InitDeleteGroupCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	deleteGroupCmd := &cobra.Command{
		Use:   "group",
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
	deleteGroupCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteGroupCmd
}

func InitDeleteGroupCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	deleteGroupCmd := &cobra.Command{
		Use:   "group",
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
	deleteGroupCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteGroupCmd
}

func InitDeleteGroupCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	deleteGroupCmd := &cobra.Command{
		Use:   "group",
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
	deleteGroupCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteGroupCmd
}
