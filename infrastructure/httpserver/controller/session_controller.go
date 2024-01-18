package controller

import (
	"datingapp/application/core"
	"datingapp/infrastructure/httpserver/controller/models"
	"datingapp/infrastructure/httpserver/jwt"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SessionController struct {
	SessionUseCase       core.ISessionUseCase
	CredentialRepo       core.ICredentialRepository
	SessionRepo          core.ISessionRepository
	SessionTokenProvider jwt.JWTProvider
}

func NewSessionController(
	sessionUseCase core.ISessionUseCase,
	credRepo core.ICredentialRepository,
	sessionRepo core.ISessionRepository,
	tokenProvider jwt.JWTProvider,
) SessionController {
	return SessionController{
		SessionUseCase:       sessionUseCase,
		CredentialRepo:       credRepo,
		SessionRepo:          sessionRepo,
		SessionTokenProvider: tokenProvider,
	}
}

func (s *SessionController) InitiateLogin(ctx *gin.Context) {
	var request models.InitiateLoginRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		fmt.Println(fmt.Sprintf("error: %v", err))
		ctx.JSON(http.StatusBadRequest, models.BaseResponseModel{
			Message: "unable to convert body",
		})
		return
	}

	err = s.SessionUseCase.InitiateLogin(core.InitiateLoginRequest{PhoneNumber: request.PhoneNumber})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.BaseResponseModel{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.BaseResponseModel{
		Message: "success",
	})
	return
}

func (s *SessionController) Login(ctx *gin.Context) {
	var request models.CompleteLoginRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		fmt.Println(fmt.Sprintf("error: %v", err))
		ctx.JSON(http.StatusBadRequest, models.CompleteLoginResponse{
			Message: "unable to convert body",
		})
		return
	}

	token, err := s.SessionUseCase.Login(request.ParseToCoreModel())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.CompleteLoginResponse{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.CompleteLoginResponse{
		Message: "success",
		Token:   token,
	})
	return
}
