package usecase

import (
	"datingapp/application/core"
	"datingapp/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegisterCredential(t *testing.T) {

	t.Run("register credentials", func(t *testing.T) {
		req := core.CredentialRequest{
			PhoneNumber:  "081234567891",
			EmailAddress: "banshee@gmng.com",
		}

		usrCredRepo := mocks.ICredentialRepository{}
		usrCredRepo.On("GetCredentials", mock.Anything).Return(nil, nil)
		usrCredRepo.On("InsertCredentials", mock.Anything).Return(core.Credentials{
			UserID:       1,
			PhoneNumber:  req.PhoneNumber,
			EmailAddress: req.EmailAddress,
		}, nil)

		usrSessUseCase := mocks.ISessionUseCase{}
		usrSessUseCase.On("InitiateLogin", mock.Anything).Return(nil)

		usrCredUseCase := NewCredentialsUsecase(&usrCredRepo, &usrSessUseCase)
		err := usrCredUseCase.RegisterCredential(req)

		assert.Nil(t, err)
	})
}
