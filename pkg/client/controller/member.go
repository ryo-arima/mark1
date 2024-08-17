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

func InitBootstrapMemberCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	bootstrapMemberCmd := &cobra.Command{
		Use:   "member",
		Short: "bootstrap the value of a key",
		Long:  "bootstrap the value of a key",
		Run: func(cmd *cobra.Command, args []string) {
			option, err := cmd.Flags().GetString("key")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(option)
			memberRepository := repository.NewMemberRepository(conf)
			memberUsecase := usecase.NewMemberUsecase(memberRepository)
			memberUsecase.BootstrapMemberForDB(request.MemberRequest{})
		},
	}
	bootstrapMemberCmd.Flags().StringP("key", "k", "", "cache key")
	return bootstrapMemberCmd
}

func InitCreateMemberCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	createMemberCmd := &cobra.Command{
		Use:   "member",
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
	createMemberCmd.Flags().StringP("key", "k", "", "cache key")
	return createMemberCmd
}

func InitCreateMemberCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	createMemberCmd := &cobra.Command{
		Use:   "member",
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
	createMemberCmd.Flags().StringP("key", "k", "", "cache key")
	return createMemberCmd
}

func InitCreateMemberCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	createMemberCmd := &cobra.Command{
		Use:   "member",
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
	createMemberCmd.Flags().StringP("key", "k", "", "cache key")
	return createMemberCmd
}

func InitGetMemberCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	getMemberCmd := &cobra.Command{
		Use:   "member",
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
	getMemberCmd.Flags().StringP("key", "k", "", "cache key")
	return getMemberCmd
}

func InitGetMemberCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	getMemberCmd := &cobra.Command{
		Use:   "member",
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
	getMemberCmd.Flags().StringP("key", "k", "", "cache key")
	return getMemberCmd
}

func InitGetMemberCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	getMemberCmd := &cobra.Command{
		Use:   "member",
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
	getMemberCmd.Flags().StringP("key", "k", "", "cache key")
	return getMemberCmd
}

func InitUpdateMemberCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	updateMemberCmd := &cobra.Command{
		Use:   "member",
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
	updateMemberCmd.Flags().StringP("key", "k", "", "cache key")
	return updateMemberCmd
}

func InitUpdateMemberCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	updateMemberCmd := &cobra.Command{
		Use:   "member",
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
	updateMemberCmd.Flags().StringP("key", "k", "", "cache key")
	return updateMemberCmd
}

func InitUpdateMemberCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	updateMemberCmd := &cobra.Command{
		Use:   "member",
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
	updateMemberCmd.Flags().StringP("key", "k", "", "cache key")
	return updateMemberCmd
}

func InitDeleteMemberCmdForAnonymousUser(conf config.BaseConfig) *cobra.Command {
	deleteMemberCmd := &cobra.Command{
		Use:   "member",
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
	deleteMemberCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteMemberCmd
}

func InitDeleteMemberCmdForAppUser(conf config.BaseConfig) *cobra.Command {
	deleteMemberCmd := &cobra.Command{
		Use:   "member",
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
	deleteMemberCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteMemberCmd
}

func InitDeleteMemberCmdForAdminUser(conf config.BaseConfig) *cobra.Command {
	deleteMemberCmd := &cobra.Command{
		Use:   "member",
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
	deleteMemberCmd.Flags().StringP("key", "k", "", "cache key")
	return deleteMemberCmd
}
