package usecase

import (
	"datingapp/application/core"
	"fmt"
)

type CredentialsUseCase struct {
	CredentialsRepo core.ICredentialRepository
	SessionUseCase  core.ISessionUseCase
}

func NewCredentialsUsecase(
	credRepo core.ICredentialRepository,
	sessionUseCase core.ISessionUseCase,
) CredentialsUseCase {
	return CredentialsUseCase{
		CredentialsRepo: credRepo,
		SessionUseCase:  sessionUseCase,
	}
}

func (uc CredentialsUseCase) RegisterCredential(request core.CredentialRequest) error {
	Creds, err := uc.CredentialsRepo.GetCredentials(request.PhoneNumber)
	if err != nil {
		return fmt.Errorf("error when process  credential")
	}

	if Creds != nil {
		return fmt.Errorf("credentials already exists")
	}

	usrCrd, err := core.CreateNewCredential(request)
	if err != nil {
		return err
	}

	usrCrd, err = uc.CredentialsRepo.InsertCredentials(usrCrd)
	if err != nil {
		return fmt.Errorf("error register  credentials")
	}

	err = uc.SessionUseCase.InitiateLogin(core.InitiateLoginRequest{PhoneNumber: request.PhoneNumber})
	return err
}
