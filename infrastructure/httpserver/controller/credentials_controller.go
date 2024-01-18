package controller

import (
	"datingapp/application/core"
	"datingapp/infrastructure/httpserver/controller/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CredentialController struct {
	CredentialUseCase core.ICredentialUseCase
	CredentialRepo    core.ICredentialRepository
	SessionUseCase    core.ISessionUseCase
}

func NewCredentialController(
	credUseCase core.ICredentialUseCase,
	credRepo core.ICredentialRepository,
	sessionUseCase core.ISessionUseCase,
) CredentialController {

	return CredentialController{
		CredentialUseCase: credUseCase,
		CredentialRepo:    credRepo,
		SessionUseCase:    sessionUseCase,
	}
}

func (c *CredentialController) RegisterCredential(ctx *gin.Context) {
	var request models.RegisterCredentialsRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		fmt.Println(fmt.Sprintf("error: %v", err))
		ctx.JSON(http.StatusBadRequest, models.BaseResponseModel{
			Message: "unable to convert body",
		})
		return
	}

	err = c.CredentialUseCase.RegisterCredential(request.ParseToCoreModel())
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
