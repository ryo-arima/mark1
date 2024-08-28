package server

import (
	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/mark1/pkg/config"
	"github.com/ryo-arima/mark1/pkg/server/controller"
	"github.com/ryo-arima/mark1/pkg/server/middleware"
	"github.com/ryo-arima/mark1/pkg/server/repository"
)

func InitRouter(conf config.BaseConfig) *gin.Engine {

	userRepository := repository.NewUserRepository(conf)
	userControllerForPublic := controller.NewUserControllerForPublic(userRepository)
	userControllerForInternal := controller.NewUserControllerForInternal(userRepository)
	userControllerForPrivate := controller.NewUserControllerForPrivate(userRepository)

	groupRepository := repository.NewGroupRepository(conf)
	groupControllerForPublic := controller.NewGroupControllerForPublic(groupRepository)
	groupControllerForInternal := controller.NewGroupControllerForInternal(groupRepository)
	groupControllerForPrivate := controller.NewGroupControllerForPrivate(groupRepository)

	memberRepository := repository.NewMemberRepository(conf)
	memberControllerForPublic := controller.NewMemberControllerForPublic(memberRepository)
	memberControllerForInternal := controller.NewMemberControllerForInternal(memberRepository)
	memberControllerForPrivate := controller.NewMemberControllerForPrivate(memberRepository)

	commonRepository := repository.NewCommonRepository(conf)
	commonControllerForPublic := controller.NewCommonControllerForPublic(userRepository, commonRepository)

	router := gin.Default()
	privateAPI := router.Group("api/private")
	internalAPI := router.Group("api/internal")
	publicAPI := router.Group("api/public")

	publicAPI.Use(middleware.ForPublic(conf))
	internalAPI.Use(middleware.ForInternal(conf))
	privateAPI.Use(middleware.ForPrivate(conf))

	//common
	publicAPI.POST("/email", commonControllerForPublic.CreateEmail) //done
	publicAPI.GET("/email", commonControllerForPublic.VerifyEmail)  //done
	publicAPI.POST("/token", commonControllerForPublic.CreateToken) //done

	//user
	publicAPI.POST("/user", userControllerForPublic.CreateUser) //done
	internalAPI.GET("/users", userControllerForInternal.GetUsers)
	internalAPI.PUT("/user", userControllerForInternal.UpdateUser)
	internalAPI.DELETE("/user", userControllerForInternal.DeleteUser)
	privateAPI.GET("/users", userControllerForPrivate.GetUsers)
	privateAPI.POST("/user", userControllerForPrivate.CreateUser)
	privateAPI.PUT("/user", userControllerForPrivate.UpdateUser)
	privateAPI.DELETE("/user", userControllerForPrivate.DeleteUser)

	//group
	publicAPI.GET("/groups", groupControllerForPublic.GetGroups)
	internalAPI.GET("/groups", groupControllerForInternal.GetGroups)
	internalAPI.POST("/group", groupControllerForInternal.CreateGroup)
	internalAPI.PUT("/group", groupControllerForInternal.UpdateGroup)
	internalAPI.DELETE("/group", groupControllerForInternal.DeleteGroup)
	privateAPI.GET("/groups", groupControllerForPrivate.GetGroups)
	privateAPI.POST("/group", groupControllerForPrivate.CreateGroup)
	privateAPI.PUT("/group", groupControllerForPrivate.UpdateGroup)
	privateAPI.DELETE("/group", groupControllerForPrivate.DeleteGroup)

	//member
	publicAPI.GET("/members", memberControllerForPublic.GetMembers)
	internalAPI.GET("/members", memberControllerForInternal.GetMembers)
	internalAPI.POST("/member", memberControllerForInternal.CreateMember)
	internalAPI.PUT("/member", memberControllerForInternal.UpdateMember)
	internalAPI.DELETE("/member", memberControllerForInternal.DeleteMember)
	privateAPI.GET("/members", memberControllerForPrivate.GetMembers)
	privateAPI.POST("/member", memberControllerForPrivate.CreateMember)
	privateAPI.PUT("/member", memberControllerForPrivate.UpdateMember)
	privateAPI.DELETE("/member", memberControllerForPrivate.DeleteMember)

	return router
}
