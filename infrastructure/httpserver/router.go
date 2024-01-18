package httpserver

import (
	controller "datingapp/infrastructure/httpserver/controller"
	"datingapp/infrastructure/httpserver/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoute(dep Dependency, r *gin.Engine) {
	r.Use(middlewares.AuthMiddleware)

	v1 := r.Group("v1")
	{
		credController := controller.NewCredentialController(dep.CredentialUseCase, dep.CredentialRepo, dep.SessionUseCase)
		v1.POST("credentials/register", credController.RegisterCredential)

		sessionController := controller.NewSessionController(dep.SessionUseCase, dep.CredentialRepo, dep.SessionRepo, dep.SessionTokenProvider)
		v1.POST("auth/initiate-login", sessionController.InitiateLogin)
		v1.POST("auth/complete-login", sessionController.Login)

		profileController := controller.NewProfileController(dep.SessionTokenProvider)
		v1.GET("profile", profileController.GetProfile)
	}

}
