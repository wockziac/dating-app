package controller

import (
	"datingapp/infrastructure/httpserver/controller/models"
	"datingapp/infrastructure/httpserver/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	SessionTokenProvider jwt.JWTProvider
}

func NewProfileController(sessionTokenProvider jwt.JWTProvider) ProfileController {
	return ProfileController{
		SessionTokenProvider: sessionTokenProvider,
	}
}

func (p *ProfileController) GetProfile(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, models.BaseResponseModel{
		Message: "success",
	})
	return
}
