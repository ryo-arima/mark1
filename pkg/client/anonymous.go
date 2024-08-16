package client

import (
	"github.com/ryo-arima/mark1/pkg/client/controller"
	"github.com/ryo-arima/mark1/pkg/config"
	"github.com/spf13/cobra"
)

type BaseCmdForAnonymousUser struct {
	Get       *cobra.Command
}

func InitRootCmdForAnonymousUser() *cobra.Command {
	var rootCmdForAnonymousUser = &cobra.Command{
		Use:   "mark1-anonymous",
		Short: "'mark1' is a CLI tool to manage anniversaries",
		Long:  `''`,
	}
	return rootCmdForAnonymousUser
}

func InitBaseCmdForAnonymousUser() BaseCmdForAnonymousUser {
	getCmd := &cobra.Command{
		Use:   "get",
		Short: "get the value of a key",
		Long:  "get the value of a key",
	}
	baseCmdForAnonymousUser := BaseCmdForAnonymousUser{
		Get:       getCmd,
	}
	return baseCmdForAnonymousUser
}

func ClientForAnonymousUser(conf config.BaseConfig) {
	rootCmdForAnonymousUser := InitRootCmdForAnonymousUser()
	rootCmdForAnonymousUser.CompletionOptions.HiddenDefaultCmd = true
	baseCmdForAnonymousUser := InitBaseCmdForAnonymousUser()

	//get
	getUserCmdForAnonymousUser := controller.InitGetUserCmdForAnonymousUser(conf)
	baseCmdForAnonymousUser.Get.AddCommand(getUserCmdForAnonymousUser)
	getGroupCmdForAnonymousUser := controller.InitGetGroupCmdForAnonymousUser(conf)
	baseCmdForAnonymousUser.Get.AddCommand(getGroupCmdForAnonymousUser)
	getMemberCmdForAnonymousUser := controller.InitGetMemberCmdForAnonymousUser(conf)
	baseCmdForAnonymousUser.Get.AddCommand(getMemberCmdForAnonymousUser)
	rootCmdForAnonymousUser.AddCommand(baseCmdForAnonymousUser.Get)
	
	rootCmdForAnonymousUser.Execute()
}