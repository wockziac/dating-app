package usecase

import (
	"datingapp/application/core"
	"fmt"
)

type UserSessionUseCase struct {
	UserSessionRepo    core.IUserSessionRepository
	UserCredentialRepo core.IUserCredentialRepository
}

func (uc *UserSessionUseCase) InitiateLogin(request core.InitiateLoginRequest) error {

	_, err := uc.UserCredentialRepo.GetUserCredentials(request.PhoneNumber)
	if err != nil {
		return fmt.Errorf("error when trying to find user")
	}

	// Generate OTP

	// Initiate Login
	// uc.UserSessionRepo.InitiateLogin()

	// Initiate send OTP & store OTP

	return nil
}

func (uc *UserSessionUseCase) Login(request core.LoginRequest) (string, error) {

	return "", nil
}
