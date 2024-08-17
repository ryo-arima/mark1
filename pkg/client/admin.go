package client

import (
	"github.com/ryo-arima/mark1/pkg/client/controller"
	"github.com/ryo-arima/mark1/pkg/config"
	"github.com/spf13/cobra"
)

type BaseCmdForAdminUser struct {
	Bootstrap *cobra.Command
	Create    *cobra.Command
	Get       *cobra.Command
	Update    *cobra.Command
	Delete    *cobra.Command
}

func InitRootCmdForAdminUser() *cobra.Command {
	var rootCmdForAdminUser = &cobra.Command{
		Use:   "mark1-admin",
		Short: "'mark1' is a CLI tool to manage anniversaries",
		Long:  `''`,
	}
	return rootCmdForAdminUser
}

func InitBaseCmdForAdminUser() BaseCmdForAdminUser {
	bootstrapCmd := &cobra.Command{
		Use:   "bootstrap",
		Short: "bootstrap the value of a key",
		Long:  "bootstrap the value of a key",
	}
	createCmd := &cobra.Command{
		Use:   "create",
		Short: "create the value of a key",
		Long:  "create the value of a key",
	}
	getCmd := &cobra.Command{
		Use:   "get",
		Short: "get the value of a key",
		Long:  "get the value of a key",
	}
	updateCmd := &cobra.Command{
		Use:   "update",
		Short: "update the value of a key",
		Long:  "update the value of a key",
	}
	deleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "delete the value of a key",
		Long:  "delete the value of a key",
	}
	baseCmdForAdminUser := BaseCmdForAdminUser{
		Bootstrap: bootstrapCmd,
		Create:    createCmd,
		Get:       getCmd,
		Update:    updateCmd,
		Delete:    deleteCmd,
	}
	return baseCmdForAdminUser
}

func ClientForAdminUser(conf config.BaseConfig) {
	rootCmdForAdminUser := InitRootCmdForAdminUser()
	rootCmdForAdminUser.CompletionOptions.HiddenDefaultCmd = true
	baseCmdForAdminUser := InitBaseCmdForAdminUser()

	//bootstrap
	bootstrapUserCmdForAdminUser := controller.InitBootstrapUserCmdForAdminUser(conf)
	baseCmdForAdminUser.Bootstrap.AddCommand(bootstrapUserCmdForAdminUser)
	bootstrapGroupCmdForAdminUser := controller.InitBootstrapGroupCmdForAdminUser(conf)
	baseCmdForAdminUser.Bootstrap.AddCommand(bootstrapGroupCmdForAdminUser)
	bootstrapMemberCmdForAdminUser := controller.InitBootstrapMemberCmdForAdminUser(conf)
	baseCmdForAdminUser.Bootstrap.AddCommand(bootstrapMemberCmdForAdminUser)
	rootCmdForAdminUser.AddCommand(baseCmdForAdminUser.Bootstrap)

	//create
	createUserCmdForAdminUser := controller.InitCreateUserCmdForAdminUser(conf)
	baseCmdForAdminUser.Create.AddCommand(createUserCmdForAdminUser)
	createGroupCmdForAdminUser := controller.InitCreateGroupCmdForAdminUser(conf)
	baseCmdForAdminUser.Create.AddCommand(createGroupCmdForAdminUser)
	createMemberCmdForAdminUser := controller.InitCreateMemberCmdForAdminUser(conf)
	baseCmdForAdminUser.Create.AddCommand(createMemberCmdForAdminUser)
	rootCmdForAdminUser.AddCommand(baseCmdForAdminUser.Create)

	//get
	getUserCmdForAdminUser := controller.InitGetUserCmdForAdminUser(conf)
	baseCmdForAdminUser.Get.AddCommand(getUserCmdForAdminUser)
	getGroupCmdForAdminUser := controller.InitGetGroupCmdForAdminUser(conf)
	baseCmdForAdminUser.Get.AddCommand(getGroupCmdForAdminUser)
	getMemberCmdForAdminUser := controller.InitGetMemberCmdForAdminUser(conf)
	baseCmdForAdminUser.Get.AddCommand(getMemberCmdForAdminUser)
	rootCmdForAdminUser.AddCommand(baseCmdForAdminUser.Get)

	//update
	updateUserCmdForAdminUser := controller.InitUpdateUserCmdForAdminUser(conf)
	baseCmdForAdminUser.Update.AddCommand(updateUserCmdForAdminUser)
	updateGroupCmdForAdminUser := controller.InitUpdateGroupCmdForAdminUser(conf)
	baseCmdForAdminUser.Update.AddCommand(updateGroupCmdForAdminUser)
	updateMemberCmdForAdminUser := controller.InitUpdateMemberCmdForAdminUser(conf)
	baseCmdForAdminUser.Update.AddCommand(updateMemberCmdForAdminUser)
	rootCmdForAdminUser.AddCommand(baseCmdForAdminUser.Update)

	//delete
	deleteUserCmdForAdminUser := controller.InitDeleteUserCmdForAdminUser(conf)
	baseCmdForAdminUser.Delete.AddCommand(deleteUserCmdForAdminUser)
	deleteGroupCmdForAdminUser := controller.InitDeleteGroupCmdForAdminUser(conf)
	baseCmdForAdminUser.Delete.AddCommand(deleteGroupCmdForAdminUser)
	deleteMemberCmdForAdminUser := controller.InitDeleteMemberCmdForAdminUser(conf)
	baseCmdForAdminUser.Delete.AddCommand(deleteMemberCmdForAdminUser)
	rootCmdForAdminUser.AddCommand(baseCmdForAdminUser.Delete)

	rootCmdForAdminUser.Execute()
}
