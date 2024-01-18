package usecase

import (
	"datingapp/application/core"
	"fmt"
)

type UserCredentialsUseCase struct {
	UserCredentialsRepo core.IUserCredentialRepository
	UserSessionUseCase  core.IUserSessionUseCase
}

func NewUserCredentialsUsecase(
	credRepo core.IUserCredentialRepository,
	userSessUseCase core.IUserSessionUseCase,
) UserCredentialsUseCase {
	return UserCredentialsUseCase{
		UserCredentialsRepo: credRepo,
		UserSessionUseCase:  userSessUseCase,
	}
}

func (uc UserCredentialsUseCase) RegisterUserCredential(request core.UserCredentialRequest) error {
	userCreds, err := uc.UserCredentialsRepo.GetUserCredentials(request.PhoneNumber)
	if err != nil {
		return fmt.Errorf("error when process user credential")
	}

	if userCreds != nil {
		return fmt.Errorf("credentials already exists")
	}

	usrCrd, err := core.CreateNewUserCredential(request)
	if err != nil {
		return err
	}

	usrCrd, err = uc.UserCredentialsRepo.InsertUserCredentials(usrCrd)
	if err != nil {
		return fmt.Errorf("error register user credentials")
	}

	err = uc.UserSessionUseCase.InitiateLogin(core.InitiateLoginRequest{PhoneNumber: request.PhoneNumber})
	return err
}
