package client

import (
	"github.com/ryo-arima/mark1/pkg/client/controller"
	"github.com/ryo-arima/mark1/pkg/config"
	"github.com/spf13/cobra"
)

type BaseCmdForAnonymousUser struct {
	Get    *cobra.Command
	Create *cobra.Command
	Verify *cobra.Command
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
	createCmd := &cobra.Command{
		Use:   "create",
		Short: "create the value of a key",
		Long:  "create the value of a key",
	}
	verifyCmd := &cobra.Command{
		Use:   "verify",
		Short: "verify the value of a key",
		Long:  "verify the value of a key",
	}
	baseCmdForAnonymousUser := BaseCmdForAnonymousUser{
		Get:    getCmd,
		Create: createCmd,
		Verify: verifyCmd,
	}
	return baseCmdForAnonymousUser
}

func ClientForAnonymousUser(conf config.BaseConfig) {
	rootCmdForAnonymousUser := InitRootCmdForAnonymousUser()
	rootCmdForAnonymousUser.CompletionOptions.HiddenDefaultCmd = true
	baseCmdForAnonymousUser := InitBaseCmdForAnonymousUser()

	//get
	getGroupCmdForAnonymousUser := controller.InitGetGroupCmdForAnonymousUser(conf)
	baseCmdForAnonymousUser.Get.AddCommand(getGroupCmdForAnonymousUser)
	getMemberCmdForAnonymousUser := controller.InitGetMemberCmdForAnonymousUser(conf)
	baseCmdForAnonymousUser.Get.AddCommand(getMemberCmdForAnonymousUser)
	rootCmdForAnonymousUser.AddCommand(baseCmdForAnonymousUser.Get)

	//create
	createTokenCmdForAnonymousUser := controller.InitCreateTokenCmdForAnonymousUser(conf)
	baseCmdForAnonymousUser.Create.AddCommand(createTokenCmdForAnonymousUser)
	createUserCmdForAnonymousUser := controller.InitCreateUserCmdForAnonymousUser(conf)
	baseCmdForAnonymousUser.Create.AddCommand(createUserCmdForAnonymousUser)
	createEmailCmdForAnonymousUser := controller.InitCreateEmailCmdForAnonymousUser(conf)
	baseCmdForAnonymousUser.Create.AddCommand(createEmailCmdForAnonymousUser)
	rootCmdForAnonymousUser.AddCommand(baseCmdForAnonymousUser.Create)
	rootCmdForAnonymousUser.Execute()
}
