package usecase

import "datingapp/application/core"

type UserCredentialsUseCase struct {
	UserCredentialsRepo core.IUserCredentialRepository
}

func NewUserCredentialsUsecase(credRepo core.IUserCredentialRepository) UserCredentialsUseCase {
	return UserCredentialsUseCase{
		UserCredentialsRepo: credRepo,
	}
}

func (uc UserCredentialsUseCase) RegisterUserCredential() {

}

func (uc UserCredentialsUseCase) Login() {

}
