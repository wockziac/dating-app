package usecase

import (
	"datingapp/application/core"
	"datingapp/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegisterUserCredential(t *testing.T) {

	t.Run("register user credentials", func(t *testing.T) {
		req := core.UserCredentialRequest{
			PhoneNumber:  "081234567891",
			EmailAddress: "banshee@gmng.com",
		}

		usrCredRepo := mocks.IUserCredentialRepository{}
		usrCredRepo.On("GetUserCredentials", mock.Anything).Return(nil, nil)
		usrCredRepo.On("InsertUserCredentials", mock.Anything).Return(core.UserCredentials{
			UserID:       1,
			PhoneNumber:  req.PhoneNumber,
			EmailAddress: req.EmailAddress,
		}, nil)

		usrSessUseCase := mocks.IUserSessionUseCase{}
		usrSessUseCase.On("InitiateLogin", mock.Anything).Return(nil)

		usrCredUseCase := NewUserCredentialsUsecase(&usrCredRepo, &usrSessUseCase)
		err := usrCredUseCase.RegisterUserCredential(req)

		assert.Nil(t, err)
	})
}
