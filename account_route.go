package v1

import (
	"github.com/Elsystm-Inc/social-medium-accounts/controllers"
	"github.com/gin-gonic/gin"
)

var userController = controllers.NewUserController()
var accountMediaController = controllers.NewAccountMediaController()
var topicsController = controllers.NewTopicsController()

func SetupAcountRoute(route *gin.Engine) {
	accounts := route.Group("/accounts/v1")
	{
		api := accounts.Group("auth")
		{
			api.GET("/user", userController.User)
			api.POST("/register", userController.Register)
			api.POST("/login", userController.Login)
			api.POST("/forgot-password", userController.ForgotPassword)
			api.POST("/forgot-password-check", userController.ForgotPasswordCheck)
			api.PUT("/reset-password", userController.ResetPassword)
			api.POST("/account-media", accountMediaController.Store)
			api.DELETE("/delete-avatar", accountMediaController.DeleteAvatar)
			api.DELETE("/delete-cover", accountMediaController.DeleteCover)
			api.POST("/topic", topicsController.Store)
			api.GET("/topics", topicsController.Topics)
			api.DELETE("/topic/:id/delete", topicsController.Delete)
			api.POST("/topic-user", topicsController.TopicUser)

		}
	}
}
